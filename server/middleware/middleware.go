package middleware 

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"../models"
	"github.com/gorilla/mux"

	// Mongodb
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connectionString for MongoDB
const connectionString = "Connection String"

// Name of the Database
const dbName = "test"

// Name of the collection
const collectionName = "todolist"

// collection object
var collection *mongo.Collection

func connect(){
	clientOptions
}