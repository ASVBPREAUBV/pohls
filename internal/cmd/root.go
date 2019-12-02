package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "pohls",
	Short: "Pohls is a fast and simple media file manager",
	Long: `A Fast and Flexible File Manager built with
                love by ASVBPREAUBV in Go.
                Complete documentation is available at http://NaN`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
