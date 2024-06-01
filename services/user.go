package service

import (
	"database/sql"
	"log"

	"github.com/blackflame007/nicklesseos.com/models"
)

type UserService struct {
	dbService *DatabaseService
}

func NewUserService(dbService *DatabaseService) *UserService {
	return &UserService{dbService: dbService}
}

// CreateUserTable creates the user table
func (us *UserService) CreateUserTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		rank INT NOT NULL
	);`

	_, err := us.dbService.db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %s", err)
	}
}

// InsertUser inserts a user into the database
func (us *UserService) InsertUser(user models.User) {
	var exists bool
	err := us.dbService.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", user.UserInfo.Email).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if user exists: %s", err)
	}

	if exists {
		return
	}

	insertUserSQL := `INSERT INTO users (name, email, rank) VALUES (?, ?, ?);`
	_, err = us.dbService.db.Exec(insertUserSQL, user.UserInfo.FullName, user.UserInfo.Email, user.Rank)
	if err != nil {
		log.Fatalf("Failed to insert user: %s", err)
	}
}

// FindUserByEmail finds a user by email
func (us *UserService) FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := us.dbService.db.QueryRow("SELECT id, name, email, rank FROM users WHERE email = ?", email).Scan(&user.ID, &user.UserInfo.FullName, &user.UserInfo.Email, &user.Rank)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Check Rank checks the rank of a user
func (us *UserService) CheckRank(email string) int {
	var rank int
	err := us.dbService.db.QueryRow("SELECT rank FROM users WHERE email = ?", email).Scan(&rank)
	if err != nil {
		log.Fatalf("Failed to check rank: %s", err)
	}

	return rank
}

func (us *UserService) CreateTokenTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tokens (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INT REFERENCES users(id),
		name VARCHAR(255) UNIQUE NOT NULL,
		token VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		revoked_at TIMESTAMP
	);
	`
	_, err := us.dbService.db.Exec(query)
	return err
}

func (us *UserService) SaveToken(userID int, token string, name string, revoked_at sql.NullTime) error {
	query := `INSERT INTO tokens (user_id, token, name, revoked_at) VALUES (?, ?, ?, ?)`
	_, err := us.dbService.db.Exec(query, userID, token, name, revoked_at)
	return err
}

func (us *UserService) GetTokensByUserID(userID int) ([]models.Token, error) {
	var tokens []models.Token
	query := `SELECT id, user_id, name, token, created_at, revoked_at FROM tokens WHERE user_id = ?`
	rows, err := us.dbService.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var token models.Token
		if err := rows.Scan(&token.ID, &token.UserID, &token.Name, &token.Token, &token.CreatedAt, &token.RevokedAt); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

// RevokeToken removes a token
func (us *UserService) RevokeToken(tokenID int) error {
	query := `DELETE FROM tokens WHERE id = ?`
	_, err := us.dbService.db.Exec(query, tokenID)
	return err
}
