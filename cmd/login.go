/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var username, password string

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Faz login no IDP",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
        payload := map[string]string{"username": username, "password": password}
        body, _ := json.Marshal(payload)

        resp, err := http.Post("http://127.0.0.1:8080/login", "application/json", bytes.NewBuffer(body))
        if err != nil {
            fmt.Println("Erro:", err)
            return
        }
        defer resp.Body.Close()
        respBody, _ := io.ReadAll(resp.Body)
        fmt.Println(string(respBody))
	},
}

func init() {
	// idpCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Nome de usuário")
    loginCmd.Flags().StringVarP(&password, "password", "p", "", "Senha")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
