# Go-todolist

# Go To Do App

Make sure you have both Docker and Docker installed before hand

# Getting Started

## Configruation

- generate .env file `cp .env.demo .env` or just copy the `.env.demo` variable to a local `.env` file

- run `docker-compose up -d`, then navigate to http://localhost:8081 or interact directly with API with the localhost:8080 endpoint

`Routes` 


## GET `/api/task`

Get all tasks currently exists 


## POST `/api/task`

Create Task

## PUT `/api/task/{id}`

Complete a task

## PUT `/api/undoTask/{id}`

Undo a task

## DELETE `/api/deleteTask/{id}`

Delete A task
