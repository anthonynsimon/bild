package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "bild",
	Short:   "A collection of parallel image processing algorithms in pure Go",
	Version: "v0.1.0",
}

func init() {
	adjust := createAdjust()
	rootCmd.AddCommand(adjust)

	blend := createBlend()
	rootCmd.AddCommand(blend)
}

// Execute starts the cli's root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
