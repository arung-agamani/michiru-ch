package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "Greets you",
	Long: "Greets you back",
	Run: helloHandler,
}

func helloHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Konnichiwa, sekai!")
}

func init() {
	rootCmd.AddCommand(helloCmd)
}