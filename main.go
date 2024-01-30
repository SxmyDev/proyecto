package main

import (
    "log"
    "net/http"
    "github.com/Sxmmy2030/proyecto/routes"
)

func main() {
    router := routes.RoutesConfig()

    log.Fatal(http.ListenAndServe(":8000", router))
}
