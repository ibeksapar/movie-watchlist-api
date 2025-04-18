package handlers

import (
	"encoding/json"
	"movie-watchlist-api/db"
	"movie-watchlist-api/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	if err := db.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		http.Error(w, "Username already taken", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil || user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var existingUser models.User
	if err := db.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&existingUser).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
