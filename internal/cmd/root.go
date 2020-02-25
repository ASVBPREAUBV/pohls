package cmd

import (
	"fmt"
	"github.com/ASVBPREAUBV/pohls/internal/config"
	"github.com/ASVBPREAUBV/pohls/internal/detox"
	"github.com/ASVBPREAUBV/pohls/internal/extractor"
	"github.com/ASVBPREAUBV/pohls/internal/filePathToMedia"
	"github.com/ASVBPREAUBV/pohls/internal/files"
	"github.com/ASVBPREAUBV/pohls/internal/mediawriter"
	"github.com/spf13/cobra"
	"path"
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
var DryRun bool

func Execute() {
	rootCmd.Flags().StringVarP(&InputDir, "input", "i", "", "Source directory to read from")
	rootCmd.Flags().StringVarP(&OutputDir, "output", "o", "", "Target directory to write to")
	rootCmd.Flags().BoolVarP(&DryRun, "dry", "d", false, "dry run without writing file")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	} else {
		config.ReadConfig()
		filepathList := mediawriter.Walk(InputDir)
		fmt.Println(filepathList)
		mediaList := extractor.FilenameCleaner(filepathList)
		fmt.Println(mediaList)

		/*cleanFileNameList :=
		for _, m := range media {
			m.DestinationFilePath = path.Join(OutputDir, string(m.MediaType), m.Title)
			if DryRun {
				fmt.Println(m)
			} else {
				mediawriter.WriteFile(m.SourceFilePath,m.DestinationFilePath)
			}
		}*/
	}
}
