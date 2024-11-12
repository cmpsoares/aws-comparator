package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Export resources of an AWS account",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Exporting AWS account resources...")
        // Implement export logic
    },
}

func init() {
    rootCmd.AddCommand(exportCmd)
}
