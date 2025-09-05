package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCredentialsCmd represents the login command
var getCredentialsCmd = &cobra.Command{
	Use:   "get-credentials",
	Short: "Lista pares de login/senha",
	Long:  `Faz uma requisição para listar os logins/senhas`,
	Run: func(cmd *cobra.Command, args []string) {
		uri := "/account"
		if name != "" {
			uri += "/" + name
		}

		req, err := http.NewRequest("GET", os.Getenv("PASSWORD_VAULT_HOST")+":"+os.Getenv("PASSWORD_VAULT_PORT")+uri, nil)
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
	getCredentialsCmd.Flags().StringVarP(&name, "name", "n", "", "Nome da conta")
}
