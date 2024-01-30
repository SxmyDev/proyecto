package controllers

import (
	"encoding/json"
	// "github.com/Sxmmy2030/proyecto/models"
	"github.com/Sxmmy2030/proyecto/repositories"
	"net/http"
)

// ProveedorHandler maneja las solicitudes HTTP relacionadas con los proveedores
func ProveedorHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		proveedores, err := repositories.GetAllProveedores()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(proveedores)
	case "PUT":
		// Aquí iría el código para actualizar un proveedor existente
	case "DELETE":
		// Aquí iría el código para eliminar un proveedor
	default:
		http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
	}
}
