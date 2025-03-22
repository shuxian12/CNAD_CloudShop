package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database
func InitDB(dbPath string) (*sql.DB, error) {
	// 若檔案已存在，先刪除
	if _, err := os.Stat(dbPath); err == nil {
		os.Remove(dbPath)
	}
	// Ensure directory exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath+"?parseTime=true&loc=Local")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Create tables if they don't exist
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

// createTables creates necessary database tables if they don't exist
func createTables(db *sql.DB) error {
	// Create users table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			username TEXT PRIMARY KEY
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Create listings table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS listings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			price REAL NOT NULL,
			username TEXT NOT NULL,
			category TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			FOREIGN KEY (username) REFERENCES users(username)
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create listings table: %w", err)
	}

	// Create categories table for faster category lookups and heavy read operations
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			category TEXT PRIMARY KEY,
			count INTEGER NOT NULL
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create categories table: %w", err)
	}

	// Create index for faster category lookups
	_, err = db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_listings_category ON listings(category);
	`)
	if err != nil {
		return fmt.Errorf("failed to create category index: %w", err)
	}

	// Create index for faster username lookups
	_, err = db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_listings_username ON listings(username);
	`)
	if err != nil {
		return fmt.Errorf("failed to create username index: %w", err)
	}

	// Create index for faster category lookups
	_, err = db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_categories_count ON categories(count);
	`)
	if err != nil {
		return fmt.Errorf("failed to create category count index: %w", err)
	}

	log.Println("Database tables created successfully")
	return nil
}
