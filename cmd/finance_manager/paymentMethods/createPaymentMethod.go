package payment_methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var name, paymentType, accountId string

// CreatePaymentMethodCmd represents the login command
var CreatePaymentMethodCmd = &cobra.Command{
	Use:   "create-payment-method",
	Short: "Faz uma transferência de entrada na conta",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		payload := map[string]interface{}{
			"accountId": accountId,
			"type":      paymentType,
			"name":      name,
		}
		body, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", os.Getenv("FINANCE_MANAGER_HOST")+":"+os.Getenv("FINANCE_MANAGER_PORT")+"/payment-methods", bytes.NewBuffer(body))
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
	CreatePaymentMethodCmd.Flags().StringVarP(&accountId, "accountId", "a", "", "Moeda")
	CreatePaymentMethodCmd.Flags().StringVarP(&paymentType, "type", "t", "", "Valor")
	CreatePaymentMethodCmd.Flags().StringVarP(&name, "name", "n", "", "Descrição")
}
