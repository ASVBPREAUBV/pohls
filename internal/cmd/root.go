package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: "Pohls is a fast and simple media file manager",
	Long: `A Fast and Flexible File Manager built with
                love by ASVBPREAUBV in Go.
                Complete documentation is available at http://NaN`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var InputDir string
var OutputDir string

func Execute() {
	rootCmd.Flags().StringVarP(&InputDir, "input", "i", "", "Source directory to read from")
	rootCmd.Flags().StringVarP(&OutputDir, "output", "o", "", "Target directory to write to")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	} else {
	}
}
