package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "A brief description of your application",
		Long:  `A longer description that spans multiple lines and likely contains examples and usage.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Cobra!")
		},
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "subcommand",
		Short: "A brief description of the subcommand",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running subcommand")
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
