package jwt

import (
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/joho/godotenv"
    "os"
    "time"
)

type CustomClaims struct {
    RazonSocial string `json:"razon_social"`
    RUC         string `json:"ruc"`
    IDProveedor int    `json:"id_proveedor"`
    jwt.StandardClaims
}

func GenerateJWT(razonSocial, ruc string, idProveedor int) (string, error) {
    err := godotenv.Load()
    if err != nil {
        return "", fmt.Errorf("Failed to load .env file: %v", err)
    }

    jwtKey := os.Getenv("JWT_SECRET")
    if jwtKey == "" {
        return "", fmt.Errorf("Failed to get JWT_SECRET from environment variables")
    }

    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &CustomClaims{
        RazonSocial: razonSocial,
        RUC:         ruc,
        IDProveedor: idProveedor,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(jwtKey))
    if err != nil {
        return "", fmt.Errorf("Failed to generate JWT token: %v", err)
    }

    return tokenString, nil
}
