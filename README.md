# TaskManagerAPI

## Overview
TaskManagerAPI is a backend application built using Go (Golang) to manage users, projects, and tasks. The API is designed to handle typical operations like creating, updating, deleting, and retrieving data for users, projects, and tasks. It uses PostgreSQL as the database and GORM as the ORM.

## Features
- **User Management**: Create, retrieve, update, and delete users.
- **Project Management**: Create, retrieve, update, and delete projects.
- **Task Management**: Create, retrieve, update, and delete tasks.
- **Relational Data**: Projects belong to users, and tasks belong to projects.

## Requirements
- **Go**: Version 1.20 or higher
- **Docker**: For containerized PostgreSQL and application
- **Make**: To streamline project commands

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/talyx/TaskManagerApi.git
   cd TaskManagerApi
   ```

2. Start PostgreSQL using Docker:
   ```bash
   make up
   ```

   This will run PostgreSQL in a Docker container.

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Start the application in debug mode:
   ```bash
   make run-debug
   ```
5. Start the application in production mode:
   ```bash
   make run-production
   ```
   The server will run on `http://localhost:8005` by default.

## Makefile Commands

- `make up`: Start PostgreSQL in a Docker container
- `make down`: Stop and remove PostgreSQL container
- `make psql`: To connect to a container with PostgreSQL and start an interactive psql session
- `make status`: To check the status of a database container
- `make reset-db:`: To reset database container
- `make run-debug`: Run the application in debug mode
- `make run-production`: Run the application in production mode

## API Endpoints

### User Endpoints

#### Create a User
**POST** `/register`
```json
{
  "names": "John Doe",
  "email": "john.doe@example.com",
  "passwordhash": "securepassword"
}
```

#### Get All Users
**GET** `/users`

#### Get User by ID
**GET** `/user/{id}`

#### Update User
**PUT** `/user/{id}`
```json
{
  "names": "Jane Doe",
  "email": "jane.doe@example.com",
  "passwordhash": "newpassword"
}
```

#### Delete User
**DELETE** `/user/{id}`

### Project Endpoints

#### Create a Project
**POST** `/project`
```json
{
  "name": "Project Alpha",
  "description": "This is the first project."
}
```
userID  is taken automatically from a cookie or token when creating a project

#### Get All Projects
**GET** `/projects`

#### Get Project by ID
**GET** `/project/{id}`

#### Update Project
**PUT** `/project/{id}`
```json
{
  "name": "Project Beta",
  "description": "Updated project description."
}
```

#### Delete Project
**DELETE** `/project/{id}`

### Task Endpoints

#### Create a Task
**POST** `/task`
```json
{
  "title": "Task One",
  "description": "This is the first task.",
  "ProjectID": 1
}
```

#### Get All Tasks
**GET** `/tasks`

#### Get Task by ID
**GET** `/task/{id}`

#### Update Task
**PUT** `/task/{id}`
```json
{
  "name": "Task Two",
  "description": "Updated task description."
}
```

#### Delete Task
**DELETE** `/task/{id}`

## Example curl Requests

### Create a User
```bash
curl -X POST http://localhost:8005/register \
-H "Content-Type: application/json" \
-d '{
  "names": "John Doe",
  "email": "john.doe@example.com",
  "passwordhash": "securepassword"
}'
```

### Authorization by a created user
```bash
curl -v -X POST http://localhost:8005/login \
-H  "Content-type: application/json" \
-d '{"login":"john.doe@example.com", "password":"securepassword"}' \
-c cookies.txt
```

### Get All Users
```bash
curl -X GET http://localhost:8005/users -b cookies.txt
```

### Create a Project
```bash
curl -X POST http://localhost:8005/project \
-H "Content-Type: application/json" \
-d '{
  "name": "Project Alpha",
  "description": "This is the first project."
}' \
-b cookies.txt

```

### Get All Projects
```bash
curl -X GET http://localhost:8005/projects -b cookies.txt
```

### Create a Task
```bash
curl -X POST http://localhost:8005/task \
-H "Content-Type: application/json" \
-d '{
  "title": "Task One",
  "description": "This is the first task.",
  "ProjectID": 1
}' \
-b cookies.txt
```

### Get All Tasks by project
```bash
curl -X GET http://localhost:8005/tasks/1 -b cookies.txt
```


