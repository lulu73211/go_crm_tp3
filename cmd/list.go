package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List contacts (memory store)",
	RunE: func(cmd *cobra.Command, args []string) error {
		rows, err := app.List()
		if err != nil {
			return err
		}
		if len(rows) == 0 {
			fmt.Println("(no contacts)")
			return nil
		}
		fmt.Printf("%-4s  %-20s  %-25s  %-15s\n", "ID", "Name", "Email", "Phone")
		for _, r := range rows {
			fmt.Printf("%-4d  %-20s  %-25s  %-15s\n", r.ID, r.Name, r.Email, r.Phone)
		}
		return nil
	},
}

func init() { rootCmd.AddCommand(listCmd) }
