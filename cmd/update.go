package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	id uint
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing contact (memory store)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := app.Update(id, name, email, phone); err != nil {
			return err
		}
		fmt.Printf("[OK] Updated contact #%d\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().UintVar(&id, "id", 0, "contact id (required)")
	updateCmd.Flags().StringVarP(&name, "name", "n", "", "new name")
	updateCmd.Flags().StringVarP(&email, "email", "e", "", "new email")
	updateCmd.Flags().StringVarP(&phone, "phone", "p", "", "new phone")
	_ = updateCmd.MarkFlagRequired("id")
}
