package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func buildBlurCommand(name string) *cobra.Command {
	var strength float64

	var cmd = &cobra.Command{
		Use:     name,
		Args:    cobra.ExactArgs(1),
		Example: fmt.Sprintf("%s --strength 0.5 image.jpg", name),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.Flags().Float64VarP(&strength, "strength", "s", 1.0, "blend mode strength")

	return cmd
}

func createBlur() *cobra.Command {
	var blurCmd = &cobra.Command{
		Use:   "blur",
		Short: "blur an image using the specified algorithm",
	}

	boxCmd := buildBlurCommand("box")
	blurCmd.AddCommand(boxCmd)

	gaussianCmd := buildBlurCommand("gaussian")
	blurCmd.AddCommand(gaussianCmd)

	return blurCmd
}
