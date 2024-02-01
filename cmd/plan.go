/*
Copyright © 2024 Piotr Rzepkowski piotr.rzepkowski98@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	tfj "github.com/hashicorp/terraform-json"
	"github.com/piotr-rzepa/tffilter/pkg/utils"
	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var (
	planOutputName = "plan.json"
	planCmd        = &cobra.Command{
		Use:   "plan",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("plan called")
			log.Println("Filter targets based on these action:", actionFilter)
			log.Println("Filter targets based on this regex:", regexFilter)
			// filters := utils.Filters{actionFilters: actionFilter, regexFilter: regexFilter}
			filters := &utils.Filters{ActionFilters: actionFilter, RegexFilter: regexFilter}
			wp := new(utils.Wrapper)
			wp.SearchBinary(tfExecutable)
			wp.ExecuteCommand("plan", "-no-color", fmt.Sprintf("-out=%s", planOutputName))
			output := wp.ExecuteCommand("show", "-no-color", "-json", "plan.json")
			plan := new(tfj.Plan)
			plan.UnmarshalJSON([]byte(output))
			bf, err := utils.ProcessPlanChanges(plan.ResourceChanges, *filters)
			if err != nil {
				log.Fatal(err)
			}
			planBuffer := bf.String()
			bf.Reset()
			// Space is important to distinguish subcommand from arguments
			bf.WriteString("plan ")
			bf.WriteString(planBuffer)
			fmt.Println(strings.Fields(bf.String()))
			wp.ExecuteCommandWithOutput(strings.Fields(bf.String())...)
		},
	}
)

func init() {
	rootCmd.AddCommand(planCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// planCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// planCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
