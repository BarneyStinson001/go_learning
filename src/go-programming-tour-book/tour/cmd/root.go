package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "",
	Short: "",
	Long: "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}
