package debt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var payerAccountId string

// PayDebtsCmd represents the login command
var PayDebtsCmd = &cobra.Command{
	Use:   "pay-debt",
	Short: "Paga uma conta a pagar",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("POST", os.Getenv("FINANCE_MANAGER_HOST")+":"+os.Getenv("FINANCE_MANAGER_PORT")+"/debts/pay/"+id+"/"+payerAccountId, nil)
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
	PayDebtsCmd.Flags().StringVarP(&id, "id", "i", "", "Id da conta a pagar")
	PayDebtsCmd.Flags().StringVarP(&payerAccountId, "payer-account-id", "p", "", "Id da conta que irá pagar")
}
