package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Module struct {
	Name         string        `json:"Name"`
	Type         string        `json:"Type"`
	ImageURL     string        `bson:"ImgUrl"`
	Environments []string      `json:"Environments"`
	Resources    []string      `json:"Resources"`
	Result       template.HTML `json:"Result"`
	Special      string        `json:"Special"`
	Prix         int           `json:"Prix"`
	Rarity       string        `json:"Rareté"`
	Description  template.HTML `json:"Description"`
	Note         template.HTML `json:"Note"`
}

func ImportModulesHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	module := Module{
		Name:         "___",
		Type:         "Targeting",
		ImageURL:     "ImageURL",
		Environments: []string{"Atmosphere", "Space"},
		Resources:    []string{"None"},
		Result:       template.HTML("Allow missiles to follow target"),
		Special:      "Advanced tracking algorithms",
		Prix:         1000,
		Rarity:       "Rare",
		Description:  template.HTML("<p>An advanced targeting module that enhances missile guidance systems.</p>"),
		Note:         template.HTML("<p>Requires skilled technicians for installation.</p>"),
	}

	err = insertModule(collection, module)
	if err != nil {
		log.Printf("Erreur lors de l'insertion du module : %v\n", err)
		http.Error(w, "Erreur lors de l'importation du module", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Module importé avec succès")
}

func insertModule(collection *mongo.Collection, module Module) error {
	_, err := collection.InsertOne(context.Background(), module)
	return err
}
