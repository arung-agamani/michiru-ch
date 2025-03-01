package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "michiru",
	Short: "Michiru is how you would imagine handling CI/CD through Discord",
	Long: "Michiru is a web app that acts as agent to make you able to do CI/CD management through Discord",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}