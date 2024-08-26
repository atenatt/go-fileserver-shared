package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "root" {
		// Password: devops1
		return "$1$8UCs.Hxd$MY6u51wFLuAuget/DZ72j."
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuservidor.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))

	fmt.Printf("Subindo servidor na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
