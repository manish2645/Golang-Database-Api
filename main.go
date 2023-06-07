package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	State   string `json:"state"`
	City    string `json:"city"`
	Zip     string `json:"zipcode"`
	Country string `json:"country"`
}

const (
	DBHost = "localhost"
	DBUser = "postgres"
	DBPort = 5432
	DBPass = "manish"
	DBName = "postgres"
)

func saveUser(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBPass, DBName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (id, name, email, age, phone, state, city, zipcode,country) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		user.ID, user.Name, user.Email, user.Age, user.Phone, user.State, user.City, user.Zip, user.Country)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save the data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User saved successfully", "user": user})
}

func getUsers(c *gin.Context) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBPass, DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, email, age, phone, state, city, zipcode,country FROM users")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch user data"})
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Phone, &user.State, &user.City, &user.Zip, &user.Country)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User data not found!"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", saveUser)
	router.Run(":8080")
}
