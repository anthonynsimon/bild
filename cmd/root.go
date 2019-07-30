package cmd

import (
	"github.com/spf13/cobra"
)

// Version of bild's CLI, set by the compiler on release
var Version string

var rootCmd = &cobra.Command{
	Use:     "bild",
	Short:   "A collection of parallel image processing algorithms in pure Go",
	Version: Version,
}

func init() {
	rootCmd.AddCommand(createAdjust())
	rootCmd.AddCommand(createBlend())
	rootCmd.AddCommand(createBlur())
	rootCmd.AddCommand(createImgio())
	rootCmd.AddCommand(createNoise())
	rootCmd.AddCommand(createSegment())
	rootCmd.AddCommand(createHistogram())
	rootCmd.AddCommand(createChannel())
	rootCmd.AddCommand(createEffect())
}

// Execute starts the cli's root command
func Execute() {
	err := rootCmd.Execute()
	exitIfNotNil(err)
}
