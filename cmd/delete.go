package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a contact by id (memory store)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := app.Delete(id); err != nil {
			return err
		}
		fmt.Printf("[OK] Deleted contact #%d\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().UintVar(&id, "id", 0, "contact id (required)")
	_ = deleteCmd.MarkFlagRequired("id")
}
