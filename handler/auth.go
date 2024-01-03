package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"todo-cognixus/config"
	"todo-cognixus/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var oauthConfig *oauth2.Config

type Provider string

const (
	GOOGLE   Provider = "google"
	GITHUB   Provider = "github"
	FACEBOOK Provider = "facebook"
)

// @Summary Login
// @Description Social media login
// @Description Provider options: google, facebook, github
// @Description Use this API on browser
// @Tags Auth
// @Produce json
// @Param provider path string true "Provider"
// @Success 200 {object} map[string]interface{} "Successful login"
// @Router /auth/{provider} [get]
func Login(ctx *fiber.Ctx) error {
	provider := ctx.Params("provider")

	switch Provider(provider) {
	case GOOGLE:
		oauthConfig = &oauth2.Config{
			ClientID:     config.GOOGLE_KEY_ID,
			ClientSecret: config.GOOGLE_SECRET_KEY,
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint: google.Endpoint,
		}

	case FACEBOOK:
		oauthConfig = &oauth2.Config{
			ClientID:     config.FACEBOOK_KEY_ID,
			ClientSecret: config.FACEBOOK_SECRET_KEY,
			Scopes:       []string{"public_profile", "email"},
			Endpoint:     facebook.Endpoint,
		}

	case GITHUB:
		oauthConfig = &oauth2.Config{
			ClientID:     config.GITHUB_KEY_ID,
			ClientSecret: config.GITHUB_SECRET_KEY,
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		}

	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid provider"})
	}

	oauthConfig.RedirectURL = fmt.Sprintf("http://localhost:%s/auth/%s/callback", config.PORT, provider)

	// Redirect to the OAuth provider's login page
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

	return ctx.Redirect(url)
}

func LoginCallback(ctx *fiber.Ctx) error {
	provider := ctx.Params("provider")

	if ctx.Query("state") != "state" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "State does not match"})
	}

	tokenOauth, err := oauthConfig.Exchange(context.Background(), ctx.Query("code"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "Unable to extract token from Code-Exchange"})
	}

	var resp *http.Response

	user := model.User{
		Provider: provider,
	}

	// Get id, name and email from provider API
	switch Provider(provider) {
	case GOOGLE:
		resp, err = http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tokenOauth.AccessToken)

	case FACEBOOK:
		resp, err = http.Get("https://graph.facebook.com/v15.0/me?fields=id,name,email&access_token=" + tokenOauth.AccessToken)

	case GITHUB:
		err = GetGithubUserInfo(tokenOauth, &user)

	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid provider"})
	}

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Unable to fetch user data"})
	}

	// If user.UID is an empty string, unmarshal the response body in a single json
	if len(user.UID) == 0 {
		// Unmarshal body for Google and Facebook response in a json array
		var userDataJson map[string]interface{}

		err = UnmarshalJSONResponse(resp.Body, &userDataJson)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": err.Error()})
		}

		user.UID = userDataJson["id"].(string)
		user.Name = userDataJson["name"].(string)
		user.Email = userDataJson["email"].(string)
	}

	// If the user didn't exist, create the user
	isExisting, err := CreateUser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": err})
	}

	// If the user existed, save the information
	if isExisting {
		err = UpdateUser(&user)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": err})
		}
	}

	// Create JWT
	tokenJWT := jwt.New(jwt.SigningMethodHS256)
	claims := tokenJWT.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 1).Unix() // 1 day

	// Generate token (JWT) with a secret key
	token, err := tokenJWT.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "Failed to log in"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Successfully logged in",
		"token":  token,
		"name":   user.Name,
	})
}

func GetGithubUserInfo(tokenOauth *oauth2.Token, user *model.User) error {
	githubUserAPI := "https://api.github.com/user"

	// Get basic info from github API
	respInfo, err := GetGithubUserRequest(githubUserAPI, tokenOauth.AccessToken)

	// Unmarshal basic info response
	var userGithubJson map[string]interface{}
	err = UnmarshalJSONResponse(respInfo.Body, &userGithubJson)
	if err != nil {
		return err
	}

	// Get email from github API
	respEmail, err := GetGithubUserRequest(githubUserAPI+"/emails", tokenOauth.AccessToken)

	// Unmarshal email response in a json array
	var emailGithubJson []map[string]interface{}
	err = UnmarshalJSONResponse(respEmail.Body, &emailGithubJson)
	if err != nil {
		return err
	}

	// Assign into model user
	user.UID = strconv.FormatFloat(userGithubJson["id"].(float64), 'f', -1, 64)
	user.Name = userGithubJson["name"].(string)
	user.Email = emailGithubJson[0]["email"].(string)

	return nil
}

func GetGithubUserRequest(githubGetAPI string, accessToken string) (*http.Response, error) {
	var resp *http.Response

	req, err := http.NewRequest(http.MethodGet, githubGetAPI, nil)
	if err != nil {
		return resp, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err = http.DefaultClient.Do(req)

	return resp, err
}

func UnmarshalJSONResponse(body io.ReadCloser, dataJson interface{}) error {
	dataBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataBytes, &dataJson)
	if err != nil {
		return err
	}

	return nil
}
