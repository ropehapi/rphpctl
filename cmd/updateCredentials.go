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

// updateCredentialsCmd represents the login command
var updateCredentialsCmd = &cobra.Command{
	Use:   "update-credentials",
	Short: "Atualiza um par login/senha",
	Long:  `Faz uma requisição de atualizar conta no password vault`,
	Run: func(cmd *cobra.Command, args []string) {
		payload := map[string]string{"name": name, "login": login, "password": password}
		body, _ := json.Marshal(payload)

		url := os.Getenv("PASSWORD_VAULT_HOST") + ":" + os.Getenv("PASSWORD_VAULT_PORT") + "/account/" + id
		fmt.Println(url)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
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
	updateCredentialsCmd.Flags().StringVarP(&id, "id", "i", "", "Id da conta")
	updateCredentialsCmd.Flags().StringVarP(&name, "name", "n", "", "Nome da conta")
	updateCredentialsCmd.Flags().StringVarP(&login, "login", "l", "", "Login")
	updateCredentialsCmd.Flags().StringVarP(&password, "password", "p", "", "Senha")
}
