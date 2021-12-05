/*
Copyright © 2021 Martin Bjeldbak Madsen <me@martinbjeldbak.com>

*/
package cmd

import (
	"bufio"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cfgFile string
var inputFile string
var inputData []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "advent-of-code",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		currentDir, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		inputFile, err := os.Open(filepath.Join(currentDir, inputFile))

		if err != nil {
			log.Fatal(err)
		}
		defer inputFile.Close()

		scanner := bufio.NewScanner(inputFile)

		for scanner.Scan() {
			inputData = append(inputData, i)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent-of-code.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&inputFile, "inputFile", "F", "", "input data file")
	rootCmd.MarkPersistentFlagRequired("inputFile")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
