package auth

import (
    "encoding/json"
    "fmt"
    "github.com/Sxmmy2030/proyecto/database"
    "github.com/dgrijalva/jwt-go"
    "github.com/joho/godotenv"
    "golang.org/x/crypto/bcrypt"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func generateJWT(username string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("Failed to load .env file: %v", err)
	}

	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return "", fmt.Errorf("Failed to get JWT_SECRET from environment variables")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Issuer:    username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", fmt.Errorf("Failed to generate JWT token: %v", err)
	}
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    var user User
    err = json.Unmarshal(body, &user)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }

    passwordHash, err := database.GetUser(user.Username)
    if err != nil {
        http.Error(w, "Error: Failed to get user from database", http.StatusInternalServerError)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
    if err != nil {
        http.Error(w, "Error: Invalid credentials", http.StatusUnauthorized)
        return
    }

    tokenString, err := generateJWT(user.Username)
    if err != nil {
        http.Error(w, "Error: Failed to generate token", http.StatusInternalServerError)
        return
    }

    response := map[string]string{
        "token": tokenString,
    }

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error: Failed to generate JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	err = database.AddUser(user.Username, user.Password)
    if err != nil {
        // Devuelve el error real
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	tokenString, err := generateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error: Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token": tokenString,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error: Failed to generate JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
