package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	inputFile string
	inputData []string
	rootCmd   = &cobra.Command{
		Use:   "2023",
		Short: "Run 2023 questions",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			if inputFile == "" {
				inputFile = filepath.Join("cmd", fmt.Sprintf("%s.in", cmd.Name()))
			}


			inputFile, err := os.Open(filepath.Join(currentDir, inputFile))
			if err != nil {
				log.Fatal(err)
			}
			defer inputFile.Close()

			scanner := bufio.NewScanner(inputFile)

			for scanner.Scan() {
				inputData = append(inputData, scanner.Text())
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputFile, "inputFile", "F", "", "input data file")
}
