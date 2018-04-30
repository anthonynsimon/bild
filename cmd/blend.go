package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func buildBlendModeCommand(name string) *cobra.Command {
	var strength float64

	var cmd = &cobra.Command{
		Use:     name,
		Args:    cobra.ExactArgs(2),
		Example: fmt.Sprintf("%s --strength 0.5 a.jpg b.jpg", name),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.Flags().Float64VarP(&strength, "strength", "s", 1.0, "blend mode strength")

	return cmd
}

func createBlend() *cobra.Command {
	blendCmd := &cobra.Command{
		Use:   "blend",
		Short: "blend two images together",
	}

	addCmd := buildBlendModeCommand("add")
	blendCmd.AddCommand(addCmd)

	colorburnCmd := buildBlendModeCommand("colorburn")
	blendCmd.AddCommand(colorburnCmd)

	colordodgeCmd := buildBlendModeCommand("colordodge")
	blendCmd.AddCommand(colordodgeCmd)

	darkenCmd := buildBlendModeCommand("darken")
	blendCmd.AddCommand(darkenCmd)

	differenceCmd := buildBlendModeCommand("difference")
	blendCmd.AddCommand(differenceCmd)

	divideCmd := buildBlendModeCommand("divide")
	blendCmd.AddCommand(divideCmd)

	exclusionCmd := buildBlendModeCommand("exclusion")
	blendCmd.AddCommand(exclusionCmd)

	lightenCmd := buildBlendModeCommand("lighten")
	blendCmd.AddCommand(lightenCmd)

	linearburnCmd := buildBlendModeCommand("linearburn")
	blendCmd.AddCommand(linearburnCmd)

	linearLightCmd := buildBlendModeCommand("linearLight")
	blendCmd.AddCommand(linearLightCmd)

	multiplyCmd := buildBlendModeCommand("multiply")
	blendCmd.AddCommand(multiplyCmd)

	normalCmd := buildBlendModeCommand("normal")
	blendCmd.AddCommand(normalCmd)

	opacityCmd := buildBlendModeCommand("opacity")
	blendCmd.AddCommand(opacityCmd)

	overlayCmd := buildBlendModeCommand("overlay")
	blendCmd.AddCommand(overlayCmd)

	screenCmd := buildBlendModeCommand("screen")
	blendCmd.AddCommand(screenCmd)

	softlightCmd := buildBlendModeCommand("softlight")
	blendCmd.AddCommand(softlightCmd)

	subtractCmd := buildBlendModeCommand("subtract")
	blendCmd.AddCommand(subtractCmd)

	return blendCmd
}
