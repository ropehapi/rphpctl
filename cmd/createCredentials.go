package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var name, login string

// createCredentialsCmd represents the login command
var createCredentialsCmd = &cobra.Command{
	Use:   "create-credentials",
	Short: "Cria um par login/senha",
	Long:  `Faz uma requisição de criar conta no password vault`,
	Run: func(cmd *cobra.Command, args []string) {
		payload := map[string]string{"name": name, "login": login, "password": password}
		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", os.Getenv("PASSWORD_VAULT_HOST")+":"+os.Getenv("PASSWORD_VAULT_PORT")+"/account", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+os.Getenv("BEARER_TOKEN"))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erro ao enviar requisição:", err)
			return
		}
		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		fmt.Println(string(respBody))
	},
}

func init() {
	createCredentialsCmd.Flags().StringVarP(&name, "name", "n", "", "Nome da conta")
	createCredentialsCmd.Flags().StringVarP(&login, "login", "l", "", "Login")
	createCredentialsCmd.Flags().StringVarP(&password, "password", "p", "", "Senha")
}
