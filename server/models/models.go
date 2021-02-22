package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Creating a struct for ToDoList
type ToDoList struct {
	//Object Id that identify the task 
	ID primitive.ObjectID `json:"_id,omitempty"
	bson:"_id,omitempty"`

	// Task name in string
	Task string `json:"task,omitempty"`
	// Completed status whether true or false
	Completed bool `json:"completed,omitempty"`
}

