package account

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var id string

// GetAccountsCmd represents the login command
var GetAccountsCmd = &cobra.Command{
	Use:   "get-accounts",
	Short: "Lista as transferências",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		uri := "/accounts"

		if id != "" {
			uri += "/" + id
		} else if name != "" && currency != "" {
			uri += "?name=" + name + "&currency_code=" + currency
		} else if name != "" {
			uri += "?name=" + name
		} else if currency != "" {
			uri += "?currency_code=" + currency
		}

		req, err := http.NewRequest("GET",
			os.Getenv("FINANCE_MANAGER_HOST")+":"+os.Getenv("FINANCE_MANAGER_PORT")+uri, nil)
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

		var prettyJSON interface{} // pode ser []interface{} ou map[string]interface{}
		if err := json.Unmarshal(respBody, &prettyJSON); err == nil {
			formatted, _ := json.MarshalIndent(prettyJSON, "", "  ")
			fmt.Println("HTTP Status:", resp.Status)
			fmt.Println(string(formatted))
		} else {
			fmt.Println("HTTP Status:", resp.Status)
			fmt.Println(string(respBody))
		}
	},
}

func init() {
	GetAccountsCmd.Flags().StringVarP(&name, "name", "n", "", "Tipo (cashin, cashout ou debt_payment)")
	GetAccountsCmd.Flags().StringVarP(&currency, "currency", "c", "", "Categoria")
	GetAccountsCmd.Flags().StringVarP(&id, "id", "i", "", "Id")
}
