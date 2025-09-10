package password_vault

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// GetAccountCodesCmd represents the login command
var GetAccountCodesCmd = &cobra.Command{
	Use:   "get-codes",
	Short: "Lista palavras chave de uma conta",
	Long:  `Faz uma requisição para listar as palavras chave de uma conta`,
	Run: func(cmd *cobra.Command, args []string) {
		uri := "/account-codes"
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
	GetAccountCodesCmd.Flags().StringVarP(&name, "name", "n", "", "Nome da conta")
}
