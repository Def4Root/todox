package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add your tasks.",
	Long: `add your tasks.

tx add task [--tag <TAG>]`,
	Run: func(cmd *cobra.Command, args []string) {
		listfile := GetListFile()
		file, err := os.OpenFile(listfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		FailOnError(err)
		defer file.Close()

		task := strings.Join(os.Args[2:], " ")
		_, err = fmt.Fprintln(file, task)
		fmt.Printf("Task added: %s\n", task)
		FailOnError(err)

		// tag, _ := cmd.Flags().GetString("tag") 
		// fmt.Println("tags:", tag)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// addCmd.Flags().StringP("tag", "t", "", "adding a tag to tasks")
}
