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

	fmt.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
