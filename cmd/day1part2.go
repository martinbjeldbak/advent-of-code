/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
)

func day2part2(measurements []string) int {
	return 0
}

// day1part2Cmd represents the day1part2 command
var day1part2Cmd = &cobra.Command{
	Use: "day1part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("day1part2 called")
		day2part2(inputData)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(day1part2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day1part2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1part2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
