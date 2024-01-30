package data_import

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Brand struct {
	Name        string        `bson:"name"`
	Type        string        `bson:"type"`
	Logo        string        `bson:"logo"`
	Font        string        `bson:"font"`
	Description template.HTML `bson:"description"`
	Note        template.HTML `bson:"note"`
}

func ImportBrandsHandler(w http.ResponseWriter, r *http.Request) {
	collection, err := setupMongoDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
		return
	}

	brand := Brand{
		Name:        "Name",
		Type:        "Type",
		Logo:        "Logo",
		Font:        "Font",
		Description: "<p>Description...</p>",
		Note:        "<p>__</p>",
	}

	err = insertBrand(collection, brand)
	if err != nil {
		log.Printf("Erreur lors de l'insertion de la brand : %v\n", err)
		http.Error(w, "Erreur lors de l'importation de la brand", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Brand importée avec succès")
}

func insertBrand(collection *mongo.Collection, brand Brand) error {
	_, err := collection.InsertOne(context.Background(), brand)
	return err
}
