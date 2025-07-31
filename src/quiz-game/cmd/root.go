/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type problem struct {
	q string
	a string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quiz-game",
	Short: "A brief description of quiz game",
	Long:  `Usage of ./quiz-game:`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implement the main logic for the quiz game here
		cfgFile, _ := cmd.Flags().GetString("csv")
		//timeLimit, _ := cmd.Flags().GetInt("limit")
		file, err := os.Open(cfgFile)
		if err != nil {
			cmd.PrintErrf("Error opening file %s: %v\n", cfgFile, err)
			return
		}
		r := csv.NewReader(file)
		lines, err := r.ReadAll()
		if err != nil {
			cmd.PrintErrf("Error reading CSV file %s: %v\n", cfgFile, err)
			return
		}
		problems := parseLines(lines)
		
		correct := 0
		for i, p := range problems {
			fmt.Printf("Problem #%d: %s = ?\n", i+1, p.q)
			var answer string
			fmt.Scanln(&answer)
			if strings.TrimSpace(answer) == p.a {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Printf("Wrong! The correct answer is %s\n", p.a)
			}
		}
		fmt.Printf("You got %d out of %d correct!\n", correct, len(problems))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, 0, len(lines))
	for _, line := range lines {
		if len(line) != 2 {
			fmt.Printf("Invalid line: %v\n", line)
			continue
		}
		ret = append(ret, problem{q: line[0], a: strings.TrimSpace(line[1])})
	}
	return ret
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.Flags().StringP("csv", "", "problems.csv", "a csv file in the format of 'question,answer' for the quiz game")
	//rootCmd.PersistentFlags().IntVar(&timeLimit, "limit", 30, "the time limit for each question in seconds")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


