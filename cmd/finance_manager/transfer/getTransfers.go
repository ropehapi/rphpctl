package transfer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var transferType, category, id string

// GetTransfersCmd represents the login command
var GetTransfersCmd = &cobra.Command{
	Use:   "get-transfers",
	Short: "Lista as transferências",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		uri := "/transfers"

		if id != "" {
			uri += "/" + id
		} else if transferType != "" && category != "" {
			uri += "?type=" + transferType + "&category=" + category
		} else if transferType != "" {
			uri += "?type=" + transferType
		} else if category != "" {
			uri += "?category=" + category
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
	GetTransfersCmd.Flags().StringVarP(&transferType, "type", "t", "", "Tipo (cashin, cashout ou debt_payment)")
	GetTransfersCmd.Flags().StringVarP(&category, "category", "c", "", "Categoria")
	GetTransfersCmd.Flags().StringVarP(&id, "id", "i", "", "Id")
}
