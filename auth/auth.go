package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "github.com/Sxmmy2030/proyecto/database"
	"github.com/Sxmmy2030/proyecto/jwt"
	"github.com/Sxmmy2030/proyecto/models"
	"github.com/Sxmmy2030/proyecto/repositories"
	"golang.org/x/crypto/bcrypt"
)

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

    var credentials models.Proveedor
    err = json.Unmarshal(body, &credentials)
    if err != nil {
        http.Error(w, "Error parsing request body", http.StatusBadRequest)
        return
    }

    proveedor, err := repositories.GetProveedorByEmail(credentials.Email)
    if err != nil {
        http.Error(w, "Error: Failed to get provider from database", http.StatusInternalServerError)
        return
    }

    passwordHash := []byte(proveedor.Password)

    err = bcrypt.CompareHashAndPassword(passwordHash, []byte(credentials.Password))
    if err != nil {
        http.Error(w, "Error: Invalid credentials", http.StatusUnauthorized)
        return
    }

    tokenString, err := jwt.GenerateJWT(proveedor.RazonSocial, proveedor.RUC, proveedor.IDProveedor)
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

	var proveedor models.Proveedor
	err = json.Unmarshal(body, &proveedor)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	err = repositories.AddProveedor(proveedor)
    if err != nil {
        // Devuelve el error real
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	tokenString, err := jwt.GenerateJWT(proveedor.RazonSocial, proveedor.RUC, proveedor.IDProveedor)
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
