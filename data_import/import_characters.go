package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// Character est une structure pour représenter un personnage
type Character struct {
	Name         string        `bson:"name"`
	ImageURL     string        `bson:"image_url"`
	Age          string        `bson:"age"`
	Sex          string        `bson:"sex"`
	Home         string        `bson:"home"`
	Job          string        `bson:"job"`
	Traits       string        `bson:"traits"`
	League       string        `bson:"league"`
	FriendFoe    string        `bson:"friend_foe"`
	MissionGiver string        `bson:"mission_giver"`
	Recruitable  string        `bson:"recruitable"`
	Romanceable  string        `bson:"romanceable"`
	Special      string        `bson:"special"`
	Description  template.HTML `bson:"description"`
	Note         template.HTML `bson:"note"`
}

// ImportCharactersHandler gère l'importation des personnages
func ImportCharactersHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB() // Établir la connexion à la base de données
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	// Créez une instance de Character avec les données d'un personnage
	character := Character{
		Name:         "Drayne",
		ImageURL:     "ffh-drayne.png",
		Age:          "Unknown",
		Sex:          "Player choice",
		Home:         "Unknown",
		Job:          "Voidrunner pilot",
		Traits:       "Mask",
		League:       "H.Z",
		FriendFoe:    "__",
		MissionGiver: "Yes",
		Recruitable:  "__",
		Romanceable:  "__",
		Special:      "__",
		Description:  "Pragmatique et réaliste!!!...",
		Note:         "La fibrose pulmonaire idiopathique (FPI)...",
	}

	// Insérez le personnage dans la base de données
	err = insertCharacter(collection, character)
	if err != nil {
		log.Printf("Erreur lors de l'insertion du personnage : %v\n", err)
		http.Error(w, "Erreur lors de l'importation du personnage", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Personnage importé avec succès")
}

func insertCharacter(collection *mongo.Collection, character Character) error {
	// Insérez le personnage dans la collection
	_, err := collection.InsertOne(context.Background(), character)
	return err
}
