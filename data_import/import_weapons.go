package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Weapon struct {
	Name         string        `bson:"Name"`
	Type         string        `bson:"Type"`
	Abbr         string        `bson:"Abbr"`
	ImageURL     string        `bson:"ImageUrl"`
	Environments []string      `bson:"Environments"`
	Resources    []string      `bson:"Resources"`
	Tool         string        `bson:"Tool"`
	Damages      string        `bson:"Damages"`
	Description  template.HTML `bson:"Description"`
	Note         template.HTML `bson:"Note"`
}

func ImportWeaponsHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	weapon := Weapon{
		Name:         "Weapon",
		Type:         "Type",
		Abbr:         "Abbr",
		ImageURL:     "ImageURL",
		Environments: []string{"Environments"},
		Resources:    []string{"None"},
		Tool:         "Tool",
		Damages:      "Damages",
		Description:  template.HTML("<p>___</p>"),
		Note:         template.HTML("<p>___</p>"),
	}

	err = insertWeapon(collection, weapon)
	if err != nil {
		log.Printf("Erreur lors de l'insertion de la weapon : %v\n", err)
		http.Error(w, "Erreur lors de l'importation de la weapon", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Weapon importée avec succès")
}

func insertWeapon(collection *mongo.Collection, weapon Weapon) error {
	_, err := collection.InsertOne(context.Background(), weapon)
	return err
}
