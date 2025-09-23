package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var currency, name string
var balance int

// CreateAccountCmd represents the login command
var CreateAccountCmd = &cobra.Command{
	Use:   "create-account",
	Short: "Faz uma transferência de entrada na conta",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		payload := map[string]interface{}{
			"currencyCode": currency,
			"balance":      balance,
			"name":         name,
		}
		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", os.Getenv("FINANCE_MANAGER_HOST")+":"+os.Getenv("FINANCE_MANAGER_PORT")+"/accounts", bytes.NewBuffer(body))
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
	CreateAccountCmd.Flags().StringVarP(&currency, "currency", "c", "", "Moeda")
	CreateAccountCmd.Flags().IntVarP(&balance, "balance", "b", 0, "Valor")
	CreateAccountCmd.Flags().StringVarP(&name, "name", "n", "", "Descrição")
}
