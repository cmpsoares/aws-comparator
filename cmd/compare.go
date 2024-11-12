package cmd

import (
	"fmt"

	"github.com/cmpsoares/aws-comparator/awsutils"
	"github.com/spf13/cobra"
)

var accountA string
var accountB string

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare resources between two AWS accounts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Comparing AWS accounts: %s and %s\n", accountA, accountB)
		services := []string{"ec2", "s3", "rds", "lambda", "cloudformation", "dynamodb", "iam"}
		awsutils.FetchResources(accountA, services, "json")
		awsutils.FetchResources(accountB, services, "json")
	},
}

func init() {
	compareCmd.Flags().StringVarP(&accountA, "accountA", "a", "", "First AWS account (required)")
	compareCmd.Flags().StringVarP(&accountB, "accountB", "b", "", "Second AWS account (required)")
	compareCmd.MarkFlagRequired("accountA")
	compareCmd.MarkFlagRequired("accountB")
	rootCmd.AddCommand(compareCmd)
}
