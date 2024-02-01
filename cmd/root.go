/*
Copyright © 2024 Piotr Rzepkowski piotr.rzepkowski98@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	tfExecutable    string
	actionFilter    []string
	regexFilter     string
	interactiveMode bool
	rootCmd         = &cobra.Command{
		Use:   "tffilter",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tffilter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&regexFilter, "regex", "r", ".*", "Filter out resources to include in plan/apply. Defaults to all resources.")
	rootCmd.PersistentFlags().StringSliceVarP(&actionFilter, "action", "a", []string{}, "Filter out actions to include in plan/apply. Possible values: update, delete, create")
	rootCmd.PersistentFlags().StringVar(&tfExecutable, "executable", "terraform", "path to terraform/terragrunt executable to use for underlying commands (defaults to terraform)")
	//TODO: Add interactive mode using huh & bubbles
	rootCmd.PersistentFlags().BoolVarP(&interactiveMode, "interactive", "i", false, "use interactive mode when running plan/apply command (defaults to false)")
}
