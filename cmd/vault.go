package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vaultCmd represents the vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Acessa as funções do Password Vault",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vault called")
	},
}

func init() {
	rootCmd.AddCommand(vaultCmd)
	vaultCmd.AddCommand(createCredentialsCmd)
}
