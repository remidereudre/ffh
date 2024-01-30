package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Faction struct {
	Name             string        `json:"name"`
	Abbreviation     string        `json:"abbreviation"`
	Logo             string        `json:"logo"`
	Headquarter      string        `json:"headquarter"`
	Leaders          []string      `json:"leaders"`
	ImportantFigures []string      `json:"importantFigures"`
	Description      template.HTML `json:"description"`
	Note             template.HTML `json:"note"`
}

func ImportFactionsHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	faction := Faction{
		Name:             "Liberty League",
		Abbreviation:     "L.L",
		Logo:             "ffh-liberty-league-faction-logo.png",
		Headquarter:      "Liberty Tower, City Center",
		Leaders:          []string{"Commander Maria Rodriguez", "Senator Jonathan Turner"},
		ImportantFigures: []string{"Dr. Amelia Grant", "Captain Mark Steele"},
		Description:      "<p></p>",
		Note:             "<p></p>",
	}

	err = insertFaction(collection, faction)
	if err != nil {
		log.Printf("Erreur lors de l'insertion de la faction : %v\n", err)
		http.Error(w, "Erreur lors de l'importation de la faction", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Faction importée avec succès")
}

func insertFaction(collection *mongo.Collection, faction Faction) error {
	_, err := collection.InsertOne(context.Background(), faction)
	return err
}
