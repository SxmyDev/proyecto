package routes

import (
    "net/http"
    "github.com/Sxmmy2030/proyecto/controllers"
    "github.com/Sxmmy2030/proyecto/auth"
    "github.com/Sxmmy2030/proyecto/middleware"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func RoutesConfig() http.Handler {
    r := mux.NewRouter()

    // Public routes
    r.HandleFunc("/auth/login", auth.LoginHandler).Methods("POST")
    r.HandleFunc("/auth/register", auth.RegisterHandler).Methods("POST")

    // Private routes
    r.Handle("/proveedores", middleware.AuthenticationMiddleware(http.HandlerFunc(controllers.ProveedorHandler))).Methods("GET", "PUT", "DELETE")

    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:4200"},
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
    })

    return c.Handler(r)
}
