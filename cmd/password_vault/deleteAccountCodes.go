package password_vault

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// DeleteAccountCodesCmd represents the login command
var DeleteAccountCodesCmd = &cobra.Command{
	Use:   "delete-codes",
	Short: "Deleta um par de conta/códigos",
	Long:  `Faz uma requisição para deletar códigos de conta`,
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("DELETE", os.Getenv("PASSWORD_VAULT_HOST")+":"+os.Getenv("PASSWORD_VAULT_PORT")+"/account-codes/"+id, nil)
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
	DeleteAccountCodesCmd.Flags().StringVarP(&id, "id", "i", "", "Id da conta")
}
