package controller

import (
	"fmt"
	"net/http"
)

// TestHandler es una función del controlador que maneja la ruta "/test".
func GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundoss")
}
