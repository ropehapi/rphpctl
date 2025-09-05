package idp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var username, password string

// LoginCmd represents the login command
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Gera um token JWT no IDP",
	Long:  `Faz uma requisição de login no IDP, passando login e senha. Se informados corretamente, a API deve retornar um token JWT`,
	Run: func(cmd *cobra.Command, args []string) {
		payload := map[string]string{"username": username, "password": password}
		body, _ := json.Marshal(payload)

		resp, err := http.Post(os.Getenv("KAIZEN_AUTH_SERVICE_HOST")+":"+os.Getenv("KAIZEN_AUTH_SERVICE_PORT")+"/login", "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Println(string(respBody))
	},
}

func init() {
	LoginCmd.Flags().StringVarP(&username, "username", "u", "", "Nome de usuário")
	LoginCmd.Flags().StringVarP(&password, "password", "p", "", "Senha")
}
