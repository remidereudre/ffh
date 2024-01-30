package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Voidrunner struct {
	Family      string        `bson:"family"`
	Model       string        `bson:"model"`
	ImageURL    string        `bson:"image_url"`
	Firepower   int           `bson:"firepower"`
	Shield      int           `bson:"shield"`
	Speed       int           `bson:"speed"`
	Agility     int           `bson:"agility"`
	Cargo       float64       `bson:"cargo"`
	Special     string        `bson:"special"`
	Rarity      string        `bson:"rarity"`
	Cost        int           `bson:"cost"`
	Description template.HTML `bson:"description"`
	Note        template.HTML `bson:"note"`
}

func ImportVoidrunnersHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	voidrunner := Voidrunner{
		Family:      "Family",
		Model:       "Model",
		ImageURL:    "ImageURL",
		Firepower:   1,
		Shield:      1,
		Speed:       1,
		Agility:     1,
		Cargo:       1.0,
		Special:     "Special",
		Rarity:      "Rarity",
		Cost:        1,
		Description: "<p>Description...</p>",
		Note:        "<p>__</p>",
	}

	// Insérez le personnage dans la base de données
	err = insertVoidrunner(collection, voidrunner)
	if err != nil {
		log.Printf("Erreur lors de l'insertion du voidrunner : %v\n", err)
		http.Error(w, "Erreur lors de l'importation du voidrunner", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Voidrunner importé avec succès")
}

func insertVoidrunner(collection *mongo.Collection, voidrunner Voidrunner) error {
	// Insérez le personnage dans la collection
	_, err := collection.InsertOne(context.Background(), voidrunner)
	return err
}
