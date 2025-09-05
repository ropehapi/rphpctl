package cmd

import (
	"fmt"
	"github.com/ropehapi/rphpctl/cmd/idp"

	"github.com/spf13/cobra"
)

// idpCmd represents the idp command
var idpCmd = &cobra.Command{
	Use:   "idp",
	Short: "Acessa as funções do IDP",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("idp called")
	},
}

func init() {
	rootCmd.AddCommand(idpCmd)
	idpCmd.AddCommand(idp.LoginCmd)
}
