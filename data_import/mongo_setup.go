package data_import

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database

// setupMongoDB retourne la collection pour permettre une utilisation flexible
func setupMongoDB() (*mongo.Collection, error) {
	// URL de connexion à MongoDB
	mongoURI := "mongodb://localhost:27017"

	// Options de configuration de la base de données
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Établissement de la connexion
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Vérification de la connexion
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Sélection de la base de données et de la collection
	database = client.Database("farfromhumans")
	return database.Collection("characters"), nil
}
