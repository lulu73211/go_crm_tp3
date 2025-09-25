package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name  string
	email string
	phone string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact (memory store)",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := app.Add(name, email, phone)
		if err != nil {
			return err
		}
		fmt.Printf("[OK] Created contact #%d: %s (%s, %s)\n", c.ID, c.Name, c.Email, c.Phone)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "name", "n", "", "name (required)")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "email")
	addCmd.Flags().StringVarP(&phone, "phone", "p", "", "phone")
	_ = addCmd.MarkFlagRequired("name")
}
