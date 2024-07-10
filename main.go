package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Definindo uma estrutura para a resposta
type Response struct {
	Message string `json:"message"`
}

// Função que será chamada quando o endpoint for acessado
func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Roteamento da função para o endpoint "/hello"
	http.HandleFunc("/hello", helloHandler)

	// Iniciando o servidor na porta 8080
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
