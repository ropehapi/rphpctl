package cmd

import (
	"fmt"
	"github.com/ropehapi/rphpctl/cmd/finance_manager/account"
	payment_methods "github.com/ropehapi/rphpctl/cmd/finance_manager/paymentMethods"
	"github.com/ropehapi/rphpctl/cmd/finance_manager/transfer"

	"github.com/spf13/cobra"
)

// financeManagerCmd represents the vault command
var financeManagerCmd = &cobra.Command{
	Use:   "finance-manager",
	Short: "Acessa as funções do finance manager",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("finance manager called")
	},
}

func init() {
	rootCmd.AddCommand(financeManagerCmd)
	financeManagerCmd.AddCommand(transfer.GetTransfersCmd)
	financeManagerCmd.AddCommand(transfer.CreateCashinCmd)
	financeManagerCmd.AddCommand(transfer.CreateCashoutCmd)
	financeManagerCmd.AddCommand(transfer.DeleteTransfersCmd)

	financeManagerCmd.AddCommand(account.CreateAccountCmd)
	financeManagerCmd.AddCommand(account.GetAccountsCmd)
	financeManagerCmd.AddCommand(account.DeleteAccountCmd)
	financeManagerCmd.AddCommand(account.UpdateAccountCmd)

	financeManagerCmd.AddCommand(payment_methods.CreatePaymentMethodCmd)
	financeManagerCmd.AddCommand(payment_methods.GetPaymentMethodCmd)
	financeManagerCmd.AddCommand(payment_methods.DeletePaymentMethodCmd)
	financeManagerCmd.AddCommand(payment_methods.UpdatePaymentMethodCmd)
}
