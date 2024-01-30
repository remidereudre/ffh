package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Civida struct {
	Name                 string        `bson:"name"`
	ImageURL             template.HTML `bson:"image_url"`
	System               template.HTML `bson:"system"`
	Localization         template.HTML `bson:"localization"`
	League               template.HTML `bson:"league"`
	Population           int           `bson:"population"`
	PopulationType       string        `bson:"populationType"`
	Economy              string        `bson:"economy"`
	MainMissions         template.HTML `bson:"mainMissions"`
	SideMissions         template.HTML `bson:"sideMissions"`
	TimeBasedMission     template.HTML `bson:"timeBasedMission"`
	Artifacts            template.HTML `bson:"artifacts"`
	Scanner              template.HTML `bson:"scanner"`
	EnvironmentalHazards template.HTML `bson:"environmentalHazards"`
	Ships                template.HTML `bson:"ships"`
	MainCharacters       template.HTML `bson:"mainCharacters"`
	Description          template.HTML `bson:"description"`
	Note                 template.HTML `bson:"note"`
}

func ImportCividasHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	civida := Civida{
		Name:                 "Name",
		ImageURL:             "ImageURL",
		System:               "System",
		Localization:         "Localization",
		League:               "League",
		Population:           1,
		PopulationType:       "PopulationType",
		Economy:              "Economy",
		MainMissions:         "MainMissions",
		SideMissions:         "SideMissions",
		TimeBasedMission:     "TimeBasedMission",
		Artifacts:            "Artifacts",
		Scanner:              "Scanner",
		EnvironmentalHazards: "EnvironmentalHazards",
		Ships:                "Ships",
		MainCharacters:       "MainCharacters",
		Description:          "Pragmatique et réaliste!!!...",
		Note:                 "La fibrose pulmonaire idiopathique (FPI)...",
	}

	err = insertCivida(collection, civida)
	if err != nil {
		log.Printf("Erreur lors de l'insertion d'une civida : %v\n", err)
		http.Error(w, "Erreur lors de l'importation de la civida", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Civida importé avec succès")
}

func insertCivida(collection *mongo.Collection, civida Civida) error {
	_, err := collection.InsertOne(context.Background(), civida)
	return err
}
