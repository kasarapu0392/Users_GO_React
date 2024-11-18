package handlers

import (
	"backend/db"
	"backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := models.GetUsers(db.DB)
	if err != nil {
		log.Printf("[GetUsers] Error fetching users: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	var user models.User

	// Parse the request body
	if err := c.Bind(&user); err != nil {
		log.Printf("[CreateUser] Invalid input: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Validate required fields
	if user.UserName == "" || user.Email == "" {
		log.Printf("[CreateUser] Missing required fields")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User name and email are required"})
	}

	// Check if the user already exists
	exists, err := models.CheckUserExists(db.DB, user.UserName)
	if err != nil {
		log.Printf("[CreateUser] Error checking for existing user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error checking existing user"})
	}
	if exists {
		log.Printf("[CreateUser] User already exists: %s", user.UserName)
		return c.JSON(http.StatusConflict, map[string]string{"error": "User already exists"})
	}

	// Create user in the database
	err = models.CreateUser(db.DB, user)
	if err != nil {
		log.Printf("[CreateUser] Database error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	log.Printf("[CreateUser] User created successfully: %+v", user)
	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	user.ID = id

	if err := models.UpdateUser(db.DB, user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	if err := models.DeleteUser(db.DB, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}

	return c.NoContent(http.StatusNoContent)
}
