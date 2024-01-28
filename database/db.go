package database

import (
    "database/sql"
    "fmt"
	"github.com/joho/godotenv"
    _ "github.com/denisenkom/go-mssqldb"
	"os"
    "golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func init() {

	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
        return
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    db, err = sql.Open("sqlserver", fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", dbUser, dbPassword, dbHost, dbName))
    if err != nil {
        fmt.Println("Error al abrir la conexión a la base de datos:", err)
        return
    }

    if err = db.Ping(); err != nil {
        fmt.Println("Error al hacer ping a la base de datos:", err)
        return
    }
}

func GetUser(username string) (string, error) {
    var passwordHash string
    err := db.QueryRow("SELECT password FROM users WHERE username = @p1", username).Scan(&passwordHash)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", fmt.Errorf("no se encontró ningún usuario con el nombre de usuario %s", username)
        } else {
            return "", fmt.Errorf("ocurrió un error al obtener el usuario de la base de datos: %v", err)
        }
    }

    return passwordHash, nil
}

func AddUser(username, password string) error {
    passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("ocurrió un error al generar el hash de la contraseña: %v", err)
    }

    stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(@p1, @p2)")
    if err != nil {
        return fmt.Errorf("ocurrió un error al preparar la consulta SQL: %v", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(username, string(passwordHash))
    if err != nil {
        return fmt.Errorf("ocurrió un error al agregar el usuario: %v", err)
    }

    return nil
}