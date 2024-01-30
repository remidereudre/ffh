package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"farfromhumans/data_import" // Importation correcte avec le chemin du package

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Brand struct {
	Name        string        `bson:"name"`
	Type        string        `bson:"type"`
	Logo        string        `bson:"logo"`
	Font        string        `bson:"font"`
	Description template.HTML `bson:"description"`
	Note        template.HTML `bson:"note"`
}

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

type GalaxyData struct {
	CelestialBodies []CelestialBody
	Systems         []System
}

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

var client *mongo.Client
var database *mongo.Database
var brandsCollection *mongo.Collection
var celestialBodiesCollection *mongo.Collection
var charactersCollection *mongo.Collection
var cividasCollection *mongo.Collection
var factionsCollection *mongo.Collection
var modulesCollection *mongo.Collection
var systemsCollection *mongo.Collection
var voidrunnersCollection *mongo.Collection
var weaponsCollection *mongo.Collection

func connectToAtlas() (*mongo.Client, error) {
	// Remplacez "votre_chaine_de_connexion" par la chaîne de connexion fournie par MongoDB Atlas
	//mongoURI := "mongodb+srv://nosdren:ZzQf9TODr0aEXrKL@farfromhumans.9qivxkl.mongodb.net/"
	mongoURI := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func setupMongoDB() {
	client, err := connectToAtlas()
	if err != nil {
		log.Fatal(err)
	}

	database = client.Database("farfromhumans")
	brandsCollection = database.Collection("brands")
	celestialBodiesCollection = database.Collection("celestialbodies")
	charactersCollection = database.Collection("characters")
	cividasCollection = database.Collection("cividas")
	factionsCollection = database.Collection("factions")
	modulesCollection = database.Collection("modules")
	systemsCollection = database.Collection("systems")
	voidrunnersCollection = database.Collection("voidrunners")
	weaponsCollection = database.Collection("weapons")
}

func setupHandlers() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/brands", brandsHandler)
	http.HandleFunc("/characters", charactersHandler)
	http.HandleFunc("/cividas", cividasHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/factions", factionsHandler)
	http.HandleFunc("/galaxy", galaxyHandler)
	http.HandleFunc("/logs", logsHandler)
	http.HandleFunc("/modules", modulesHandler)
	http.HandleFunc("/shields", shieldsHandler)
	http.HandleFunc("/shops", shopsHandler)
	http.HandleFunc("/todo", toDoHandler)
	http.HandleFunc("/updates", updatesHandler)
	http.HandleFunc("/voidrunners", voidrunnersHandler)
	http.HandleFunc("/weapons", weaponsHandler)

	http.HandleFunc("/import_brands", data_import.ImportBrandsHandler)
	http.HandleFunc("/import_celestialbodies", data_import.ImportCelestialBodiesHandler)
	http.HandleFunc("/import_characters", data_import.ImportCharactersHandler)
	http.HandleFunc("/import_cividas", data_import.ImportCividasHandler)
	http.HandleFunc("/import_factions", data_import.ImportFactionsHandler)
	http.HandleFunc("/import_modules", data_import.ImportModulesHandler)
	http.HandleFunc("/import_systems", data_import.ImportSystemsHandler)
	http.HandleFunc("/import_voidrunners", data_import.ImportVoidrunnersHandler)
	http.HandleFunc("/import_weapons", data_import.ImportWeaponsHandler)
}

func main() {
	setupMongoDB()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Erreur lors de la déconnexion de MongoDB: %v", err)
		}
	}()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	setupHandlers()
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, _ := template.ParseFiles("templates/base.html", tmpl)
	t.Execute(w, data)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/home.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/about.html", nil)
}

/* BRANDS */
func brandsHandler(w http.ResponseWriter, r *http.Request) {
	brands, err := getAllBrands()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des brands : %v", err), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "templates/brands.html", brands)
}
func getAllBrands() ([]Brand, error) {
	cursor, err := brandsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var brands []Brand
	for cursor.Next(context.Background()) {
		var brand Brand
		if err := cursor.Decode(&brand); err != nil {
			return nil, err
		}
		brands = append(brands, brand)
	}

	return brands, nil
}

/* CHARACTERS */
func charactersHandler(w http.ResponseWriter, r *http.Request) {
	characters, err := getAllCharacters()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des personnages", http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "templates/characters.html", characters)
}

func getAllCharacters() ([]Character, error) {
	cursor, err := charactersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var characters []Character
	for cursor.Next(context.Background()) {
		var character Character
		if err := cursor.Decode(&character); err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return characters, nil
}

/* CIVIDAS */
func cividasHandler(w http.ResponseWriter, r *http.Request) {
	cividas, err := getAllCividas()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des cividas : %v", err), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "templates/cividas.html", cividas)
}
func getAllCividas() ([]Civida, error) {
	cursor, err := cividasCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var cividas []Civida
	for cursor.Next(context.Background()) {
		var civida Civida
		if err := cursor.Decode(&civida); err != nil {
			return nil, err
		}
		cividas = append(cividas, civida)
	}

	return cividas, nil
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/contact.html", nil)
}

/* FACTIONS */
func factionsHandler(w http.ResponseWriter, r *http.Request) {
	factions, err := getAllFactions()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des factions : %v", err), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "templates/factions.html", factions)
}
func getAllFactions() ([]Faction, error) {
	cursor, err := factionsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var factions []Faction
	for cursor.Next(context.Background()) {
		var faction Faction
		if err := cursor.Decode(&faction); err != nil {
			return nil, err
		}
		factions = append(factions, faction)
	}

	return factions, nil
}

/*
GALAXY
CELSTIAL BODIES
*/
func galaxyHandler(w http.ResponseWriter, r *http.Request) {
	celestialBodies, err := getAllCelestialBodies()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des celestial bodies : %v", err), http.StatusInternalServerError)
		return
	}

	systems, err := getAllSystems()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des systèmes : %v", err), http.StatusInternalServerError)
		return
	}

	galaxyData := GalaxyData{
		CelestialBodies: celestialBodies,
		Systems:         systems,
	}

	renderTemplate(w, "templates/galaxy.html", galaxyData)
}

func getAllSystems() ([]System, error) {
	cursor, err := systemsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var systems []System
	for cursor.Next(context.Background()) {
		var system System
		if err := cursor.Decode(&system); err != nil {
			return nil, err
		}
		systems = append(systems, system)
	}

	return systems, nil
}
func getAllCelestialBodies() ([]CelestialBody, error) {
	cursor, err := celestialBodiesCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var celestialBodies []CelestialBody
	for cursor.Next(context.Background()) {
		var celestialBody CelestialBody
		if err := cursor.Decode(&celestialBody); err != nil {
			return nil, err
		}
		celestialBodies = append(celestialBodies, celestialBody)
	}

	return celestialBodies, nil
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/logs.html", nil)
}

func shopsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/shops.html", nil)
}

func toDoHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/todo.html", nil)
}

func updatesHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/updates.html", nil)
}

/* MODULES */
func modulesHandler(w http.ResponseWriter, r *http.Request) {
	modules, err := getAllModules()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des modules : %v", err), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "templates/modules.html", modules)
}
func getAllModules() ([]Module, error) {
	cursor, err := modulesCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var modules []Module
	for cursor.Next(context.Background()) {
		var module Module
		if err := cursor.Decode(&module); err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}

	return modules, nil
}

func shieldsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/shields.html", nil)
}

/*
VOIDRUNNERS
*/
func voidrunnersHandler(w http.ResponseWriter, r *http.Request) {
	voidrunners, err := getAllVoidrunners()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des voidrunners : %v", err), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "templates/voidrunners.html", voidrunners)
}

func getAllVoidrunners() ([]Voidrunner, error) {
	cursor, err := voidrunnersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var voidrunners []Voidrunner
	for cursor.Next(context.Background()) {
		var voidrunner Voidrunner
		if err := cursor.Decode(&voidrunner); err != nil {
			return nil, err
		}
		voidrunners = append(voidrunners, voidrunner)
	}

	return voidrunners, nil
}

/*
WEAPONS
*/
func weaponsHandler(w http.ResponseWriter, r *http.Request) {
	weapons, err := getAllWeapons()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des weapons : %v", err), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "templates/weapons.html", weapons)
}

func getAllWeapons() ([]Weapon, error) {
	cursor, err := weaponsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var weapons []Weapon
	for cursor.Next(context.Background()) {
		var weapon Weapon
		if err := cursor.Decode(&weapon); err != nil {
			return nil, err
		}
		weapons = append(weapons, weapon)
	}

	return weapons, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "templates/home.html", nil)
}
