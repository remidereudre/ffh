package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type System struct {
	System        string        `bson:"System"`
	Name          string        `bson:"Name"`
	ImageURL      string        `bson:"ImageURL"`
	Type          string        `bson:"Type"`
	Stars         []string      `bson:"Stars"`
	Planets       []string      `bson:"Planets"`
	AsteroidBelts []string      `bson:"AsteroidBelts"`
	Description   template.HTML `bson:"Description"`
	Note          template.HTML `bson:"Note"`
}

func ImportSystemsHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	system := System{
		System:   "Eorana 8000-0-0",
		Name:     "Eorana 8000-0-0",
		ImageURL: "ffh-eorana-8000-0-0.png",
		Type:     "System",
		Stars:    []string{"Solis"},
		Planets: []string{
			"Crimsonia",
			"Vermillion",
			"Umbra",
			"Verdantia",
			"Chénmò",
			"Celestia",
			"Lumina",
			"Chromis",
			"Obsidian",
			"Solara",
		},
		AsteroidBelts: []string{
			"Inner belt 8000-0-0",
			"Outer belt 8000-0-0",
		},
		Description: "<p>___</p>",
		Note:        "<p>___</p>",
	}

	err = insertSystem(collection, system)
	if err != nil {
		log.Printf("Erreur lors de l'insertion du system : %v\n", err)
		http.Error(w, "Erreur lors de l'importation du system", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "System importée avec succès")
}

func insertSystem(collection *mongo.Collection, system System) error {
	_, err := collection.InsertOne(context.Background(), system)
	return err
}
