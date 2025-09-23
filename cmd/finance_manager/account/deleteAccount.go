package account

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// DeleteAccountCmd represents the login command
var DeleteAccountCmd = &cobra.Command{
	Use:   "delete-account",
	Short: "Deleta uma transferência",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		uri := "/accounts"
		if id != "" {
			uri += "/" + id
		}

		req, err := http.NewRequest("DELETE",
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
	DeleteAccountCmd.Flags().StringVarP(&id, "id", "i", "", "Id")
}
