// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Amira",
            "email": "syahidatulamira06@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todo/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all todo items by user id from login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Get all todo items by user id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add a new todo item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Create new todo item",
                "parameters": [
                    {
                        "description": "Todo object",
                        "name": "IncomingTodo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.IncomingTodo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/todo/complete/{id}": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a todo item status to complete by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Update a todo item status to complete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/todo/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a todo item by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo"
                ],
                "summary": "Delete todo item by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.IncomingTodo": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Todo",
	Description:      "This is a swagger for Todo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
