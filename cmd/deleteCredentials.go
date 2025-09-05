package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var id string

// deleteCredentialsCmd represents the login command
var deleteCredentialsCmd = &cobra.Command{
	Use:   "delete-credentials",
	Short: "Deleta um par de login/senha",
	Long:  `Faz uma requisição para deleter logins/senhas`,
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("DELETE", os.Getenv("PASSWORD_VAULT_HOST")+":"+os.Getenv("PASSWORD_VAULT_PORT")+"/account/"+id, nil)
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
	deleteCredentialsCmd.Flags().StringVarP(&id, "id", "i", "", "Id da conta")
}
