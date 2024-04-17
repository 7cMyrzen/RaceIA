package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type RobotData struct {
	IdRace   int     `json:"idrace"`
	Vitesse  float64 `json:"vitesse"`
	Distance float64 `json:"distance"`
	Tri_x    float64 `json:"tri_x"`
	Tri_y    float64 `json:"tri_y"`
	Tri_z    float64 `json:"tri_z"`
}

var db *sql.DB
var err error

func Init() {
	// Paramètres de connexion à la base de données
	username := "root"
	password := "w3gAqA888A37"
	hostname := "127.0.0.1"
	dbname := "raceia"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/raceia", username, password, hostname)

	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("Erreur lors de la connexion à MySQL: %v\n", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Erreur lors de la vérification de la connexion à MySQL: %v\n", err)
		return
	}

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbname)
	_, err = db.Exec(query)
	if err != nil {
		return
	}

	// Sélectionner la base de données
	_, err = db.Exec("USE raceia")
	if err != nil {
		fmt.Printf("Erreur lors de la sélection de la base de données: %v\n", err)
		return
	}

	// Créer les tables
	createTables(db)
}

func createTables(db *sql.DB) {
	// Requête SQL pour créer la table race
	createRaceTable := `
        CREATE TABLE IF NOT EXISTS race (
            id INT NOT NULL AUTO_INCREMENT,
            total_time DATETIME DEFAULT NULL,
            avg_speed FLOAT DEFAULT NULL,
            PRIMARY KEY (id)
        )
    `

	// Exécuter la requête pour créer la table race
	_, err := db.Exec(createRaceTable)
	if err != nil {
		fmt.Printf("Erreur lors de la création de la table race: %v\n", err)
		return
	}
	fmt.Println("Table race créée avec succès")

	// Requête SQL pour créer la table distance
	createDistanceTable := `
        CREATE TABLE IF NOT EXISTS distance (
            id INT NOT NULL AUTO_INCREMENT,
            id_race INT DEFAULT NULL,
            time DATETIME DEFAULT CURRENT_TIMESTAMP,
            distance FLOAT DEFAULT NULL,
            PRIMARY KEY (id),
            KEY id_race (id_race),
            CONSTRAINT distance_ibfk_1 FOREIGN KEY (id_race) REFERENCES race (id)
        )
    `

	// Exécuter la requête pour créer la table distance
	_, err = db.Exec(createDistanceTable)
	if err != nil {
		fmt.Printf("Erreur lors de la création de la table distance: %v\n", err)
		return
	}
	fmt.Println("Table distance créée avec succès")

	// Requête SQL pour créer la table speed
	createSpeedTable := `
        CREATE TABLE IF NOT EXISTS speed (
            id INT NOT NULL AUTO_INCREMENT,
            id_race INT DEFAULT NULL,
            time DATETIME DEFAULT CURRENT_TIMESTAMP,
            speed FLOAT DEFAULT NULL,
            PRIMARY KEY (id),
            KEY id_race (id_race),
            CONSTRAINT speed_ibfk_1 FOREIGN KEY (id_race) REFERENCES race (id)
        )
    `

	// Exécuter la requête pour créer la table speed
	_, err = db.Exec(createSpeedTable)
	if err != nil {
		fmt.Printf("Erreur lors de la création de la table speed: %v\n", err)
		return
	}
	fmt.Println("Table speed créée avec succès")

	// Requête SQL pour créer la table tri_dim
	createTriDimTable := `
        CREATE TABLE IF NOT EXISTS tri_dim (
            id INT NOT NULL AUTO_INCREMENT,
            id_race INT DEFAULT NULL,
            time DATETIME DEFAULT CURRENT_TIMESTAMP,
            td_x FLOAT DEFAULT NULL,
            td_y FLOAT DEFAULT NULL,
            td_z FLOAT DEFAULT NULL,
            PRIMARY KEY (id),
            KEY id_race (id_race),
            CONSTRAINT tri_dim_ibfk_1 FOREIGN KEY (id_race) REFERENCES race (id)
        )
    `

	// Exécuter la requête pour créer la table tri_dim
	_, err = db.Exec(createTriDimTable)
	if err != nil {
		fmt.Printf("Erreur lors de la création de la table tri_dim: %v\n", err)
		return
	}
	fmt.Println("Table tri_dim créée avec succès")
}

func InsertData(data RobotData) {
	// Récupérer les valeurs de la structure RobotData
	idRace := data.IdRace
	vitesse := data.Vitesse
	distance := data.Distance
	tri_x := data.Tri_x
	tri_y := data.Tri_y
	tri_z := data.Tri_z

	InsertRaceIfNotExists(idRace)

	// Requête SQL pour insérer les données dans la table distance
	insert, err := db.Prepare("INSERT INTO distance(id_race, distance) VALUES(?, ?)")
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête pour la table distance: %v\n", err)
		return
	}
	defer insert.Close()

	_, err = insert.Exec(idRace, distance)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion des données dans la table distance: %v\n", err)
		return
	}

	// Requête SQL pour insérer les données dans la table speed
	insert, err = db.Prepare("INSERT INTO speed(id_race, speed) VALUES(?, ?)")
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête pour la table speed: %v\n", err)
		return
	}
	defer insert.Close()

	_, err = insert.Exec(idRace, vitesse)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion des données dans la table speed: %v\n", err)
		return
	}

	// Requête SQL pour insérer les données dans la table tri_dim
	insert, err = db.Prepare("INSERT INTO tri_dim(id_race, td_x, td_y, td_z) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête pour la table tri_dim: %v\n", err)
		return
	}
	defer insert.Close()

	_, err = insert.Exec(idRace, tri_x, tri_y, tri_z)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion des données dans la table tri_dim: %v\n", err)
		return
	}

	fmt.Println("Données insérées avec succès")
}

func InsertRaceIfNotExists(id int) error {
	// Vérifier si l'ID de la course existe déjà
	var existingID int
	err := db.QueryRow("SELECT id FROM race WHERE id = ?", id).Scan(&existingID)
	switch {
	case err == sql.ErrNoRows:
		// L'ID de la course n'existe pas encore, insérer une nouvelle entrée
		_, err := db.Exec("INSERT INTO race (id) VALUES (?)", id)
		if err != nil {
			return fmt.Errorf("erreur lors de l'insertion de la course: %v", err)
		}
		fmt.Println("Course insérée avec succès, ID de la course:", id)
	case err != nil:
		// Une erreur s'est produite lors de la requête
		return fmt.Errorf("erreur lors de la vérification de l'existence de la course: %v", err)
	default:
		// L'ID de la course existe déjà, ne rien faire
		fmt.Println("L'ID de la course existe déjà:", existingID)
	}
	return nil
}
