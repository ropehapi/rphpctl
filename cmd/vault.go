package cmd

import (
	"fmt"
	"github.com/ropehapi/rphpctl/cmd/password_vault"

	"github.com/spf13/cobra"
)

// passwordVaultCmd represents the vault command
var passwordVaultCmd = &cobra.Command{
	Use:   "password-vault",
	Short: "Acessa as funções do Password Vault",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vault called")
	},
}

func init() {
	rootCmd.AddCommand(passwordVaultCmd)
	passwordVaultCmd.AddCommand(password_vault.CreateAccountCmd)
	passwordVaultCmd.AddCommand(password_vault.GetAccountCmd)
	passwordVaultCmd.AddCommand(password_vault.UpdateAccountCmd)
	passwordVaultCmd.AddCommand(password_vault.DeleteAccountCmd)

	passwordVaultCmd.AddCommand(password_vault.GetAccountCodesCmd)
	passwordVaultCmd.AddCommand(password_vault.CreateAccountCodesCmd)
	passwordVaultCmd.AddCommand(password_vault.UpdateAccountCodesCmd)
	passwordVaultCmd.AddCommand(password_vault.DeleteAccountCodesCmd)
}
