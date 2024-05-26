package service

import (
	"fmt"
	"log"

	"github.com/blackflame007/nicklesseos.com/models"
)

type UserService struct {
	dbService *DatabaseService
}

func NewUserService(dbService *DatabaseService) *UserService {
	return &UserService{dbService: dbService}
}

func (us *UserService) CreateUserTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		high_score INT NOT NULL
	);`

	_, err := us.dbService.db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %s", err)
	}
}

func (us *UserService) InsertUser(user models.User) {
	var exists bool
	fmt.Println("User: ", user)
	err := us.dbService.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", user.UserInfo.Email).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if user exists: %s", err)
	}

	fmt.Println("Exists: ", exists)

	if exists {
		fmt.Println("User with this email already exists")
		return
	}

	insertUserSQL := `INSERT INTO users (name, email, high_score) VALUES (?, ?, ?);`
	_, err = us.dbService.db.Exec(insertUserSQL, user.UserInfo.FullName, user.UserInfo.Email, user.HighScore)
	if err != nil {
		log.Fatalf("Failed to insert user: %s", err)
	}
}

func (us *UserService) GetLeaderboard() ([]models.User, error) {
	query := `SELECT name, email, high_score FROM users ORDER BY high_score DESC;`
	fmt.Println("Query: ", query)
	rows, err := us.dbService.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to get leaderboard: %s", err)
	}
	defer rows.Close()

	var users []models.User
	fmt.Println("Rows: ", rows)
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserInfo.FullName, &u.UserInfo.Email, &u.HighScore); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	fmt.Println("Users: ", users)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
