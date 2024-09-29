package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Project represents a project entity
type Project struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// DB is a global variable to hold the database connection
var DB *sql.DB

// InitDB initializes the database connection and creates the projects table
func InitDB(filepath string) {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS projects (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "name" TEXT NOT NULL,
        "description" TEXT,
        "created_at" TEXT,
        "updated_at" TEXT
    );`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Projects table created")
}

// InsertProject inserts a new project into the database
func InsertProject(name, description string) error {
	insertProjectSQL := `INSERT INTO projects (name, description, created_at, updated_at) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(insertProjectSQL, name, description, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	return err
}

// ListProjects lists recent projects from the database
func ListProjects(limit int) ([]Project, error) {
	rows, err := DB.Query(`SELECT id, name, description, created_at, updated_at FROM projects ORDER BY created_at DESC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		var createdAt, updatedAt string
		err = rows.Scan(&project.ID, &project.Name, &project.Description, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		project.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		project.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
		projects = append(projects, project)
	}

	return projects, nil
}
