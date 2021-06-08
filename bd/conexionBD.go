package bd

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://<username>:<password>@twider.dddfe.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
}
