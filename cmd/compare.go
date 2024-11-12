package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/cmpsoares/aws-comparator/awsutils"
)

var accountA string
var accountB string

var compareCmd = &cobra.Command{
    Use:   "compare",
    Short: "Compare resources between two AWS accounts",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Comparing AWS accounts: %s and %s\n", accountA, accountB)
        awsutils.FetchResources(accountA)
        awsutils.FetchResources(accountB)
    },
}

func init() {
    compareCmd.Flags().StringVarP(&accountA, "accountA", "a", "", "First AWS account (required)")
    compareCmd.Flags().StringVarP(&accountB, "accountB", "b", "", "Second AWS account (required)")
    compareCmd.MarkFlagRequired("accountA")
    compareCmd.MarkFlagRequired("accountB")
    rootCmd.AddCommand(compareCmd)
}
