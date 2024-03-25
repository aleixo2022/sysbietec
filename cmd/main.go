package main

import (
	"fmt"
	"net/http"
	"os"
	"sysbietec/internal/api"
)

func main() {
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectUri := os.Getenv("REDIRECT_URI")

	// Handler para a página inicial
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<h1>Bem-vindo à página inicial do Sysbietec!</h1>")
	})

	// Handler para o callback de autenticação
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			fmt.Fprintln(w, "Código de autorização não encontrado")
			return
		}

		client := api.NewClient(clientId, clientSecret)
		accessToken, err := client.Authenticate(code, redirectUri)
		if err != nil {
			fmt.Fprintf(w, "Erro ao autenticar: %v", err)
			return
		}

		fmt.Fprintf(w, "Access Token: %s", accessToken)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão se não estiver definida
	}
	fmt.Printf("Servidor iniciado na porta %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
