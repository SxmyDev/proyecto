package repositories

import (
    "github.com/Sxmmy2030/proyecto/models"
    "github.com/Sxmmy2030/proyecto/database"
    "golang.org/x/crypto/bcrypt"
)

func GetProveedor(ruc string) (models.Proveedor, error) {
    var proveedor models.Proveedor
    result := database.DB.Where("RUC = ?", ruc).First(&proveedor)
    if result.Error != nil {
        return models.Proveedor{}, result.Error
    }
    return proveedor, nil
}

func GetProveedorByEmail(email string) (models.Proveedor, error) {
    var proveedor models.Proveedor
    result := database.DB.Where("EMAIL = ?", email).First(&proveedor)
    if result.Error != nil {
        return models.Proveedor{}, result.Error
    }
    return proveedor, nil
}

func GetAllProveedores() ([]models.Proveedor, error) {
	var proveedores []models.Proveedor
	result := database.DB.Find(&proveedores)
	if result.Error != nil {
		return nil, result.Error
	}
	return proveedores, nil
}

func AddProveedor(proveedor models.Proveedor) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(proveedor.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    proveedor.Password = string(hashedPassword)
    result := database.DB.Create(&proveedor)
    return result.Error
}
