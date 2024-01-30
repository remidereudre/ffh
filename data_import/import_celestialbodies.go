package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type CelestialBody struct {
	Name                 string        `json:"name"`
	System               string        `json:"system"`
	ImageURL             string        `json:"imageURL"`
	Type                 string        `json:"type"`
	Composition          []string      `json:"composition"`
	Gravity              float64       `json:"gravity"`
	SurfaceTemperature   int           `json:"surfaceTemperature"`
	Resources            []string      `json:"resources"`
	MainMissions         []string      `json:"mainMissions"`
	SideMissions         []string      `json:"sideMissions"`
	TimeBasedMissions    []string      `json:"timeBasedMissions"`
	Artifacts            []string      `json:"artifacts"`
	Scanner              []string      `json:"scanner"`
	EnvironmentalHazards []string      `json:"environmentalHazards"`
	MainCharacters       []string      `json:"mainCharacters"`
	Ships                []string      `json:"ships"`
	Description          template.HTML `bson:"Description"`
	Note                 template.HTML `bson:"Note"`
}

func ImportCelestialBodiesHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	celestialBody := CelestialBody{
		Name:                 "Solis",
		System:               "Eorana 8000-0-0",
		ImageURL:             "ffh-solis.png",
		Type:                 "G2V Main Sequence Star",
		Composition:          []string{"Hydrogen (74%)", "Helium (24%)", "Heavier elements (2%)"},
		Gravity:              28,
		SurfaceTemperature:   5500,
		Resources:            []string{"Heat", "Light"},
		MainMissions:         []string{"Build a Dyson Sphere (resources) -> Solis"},
		SideMissions:         []string{},
		TimeBasedMissions:    []string{"Watch out the watchers (kill) -> Solis"},
		Artifacts:            []string{"Scientific data (1)", "Scientific data (2)"},
		Scanner:              []string{"Eruption", "Surface stains"},
		EnvironmentalHazards: []string{"Heat", "Radiation", "Gravity"},
		MainCharacters:       []string{"Scientist", "Droid"},
		Ships:                []string{"Watchers (some)", "Small scientist ship"},
		Description:          "___",
		Note:                 "___",
	}

	err = insertCelestialBody(collection, celestialBody)
	if err != nil {
		log.Printf("Erreur lors de l'insertion du celestial body : %v\n", err)
		http.Error(w, "Erreur lors de l'importation du celestial body", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Celestial Body importé avec succès")
}

func insertCelestialBody(collection *mongo.Collection, celestialBody CelestialBody) error {
	_, err := collection.InsertOne(context.Background(), celestialBody)
	return err
}
