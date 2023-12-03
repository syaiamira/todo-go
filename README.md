# todo-go

## Run

To run the application, follow these steps:

1. Make sure you have Docker installed on your machine.
2. Clone this repository to your local machine.
3. Open a terminal and navigate to the project directory.
4. Run the following command to start the application:

   ```bash
   docker-compose up

5. Alternatively, pull the image from Docker Hub:

    Run:
    ```
    docker pull syaiamira/todo-go
    docker run -p 8000:8000 syaiamira/todo-go

    ```

## Test

### Add a new todo
curl -X POST -H "Content-Type: application/json" -d '{"title":"New Task"}' http://localhost:8000/todo/

### Get the list of todos
curl -X 'GET' 'http://127.0.0.1:8000/todo/'

### Mark a todo item to complete
curl -X 'PATCH' 'http://localhost:8000/todo/complete/{id}'

### Delete a todo item
curl -X 'DELETE' 'http://localhost:8000/todo/{id}'

## Interface documentation

Go to: [http://localhost:8000/swagger/index.html#]