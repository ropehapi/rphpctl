package password_vault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// UpdateAccountCodesCmd represents the login command
var UpdateAccountCodesCmd = &cobra.Command{
	Use:   "update-codes",
	Short: "Atualiza um par conta/códigos",
	Long:  `Faz uma requisição de atualizar códigos no password vault`,
	Run: func(cmd *cobra.Command, args []string) {
		payload := map[string]string{"name": name, "codes": codes}
		body, _ := json.Marshal(payload)

		url := os.Getenv("PASSWORD_VAULT_HOST") + ":" + os.Getenv("PASSWORD_VAULT_PORT") + "/account-codes/" + id
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
	UpdateAccountCodesCmd.Flags().StringVarP(&id, "id", "i", "", "Id da conta")
	UpdateAccountCodesCmd.Flags().StringVarP(&name, "name", "n", "", "Nome da conta")
	UpdateAccountCodesCmd.Flags().StringVarP(&codes, "codes", "c", "", "Códigos")
}
