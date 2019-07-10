package cmd

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/blend"
	"github.com/spf13/cobra"
)

func buildBlendModeCommand(name string) *cobra.Command {
	var strength float64

	var cmd = &cobra.Command{
		Use:     name,
		Args:    cobra.ExactArgs(3),
		Example: fmt.Sprintf("%s --strength 0.5 input1.jpg input2.jpg output.jpg", name),
		Run: func(cmd *cobra.Command, args []string) {
			fin1 := args[0]
			fin2 := args[1]
			fout := args[2]

			apply2(fin1, fin2, fout, func(img1, img2 image.Image) (image.Image, error) {
				switch name {
				case "add":
					return blend.Opacity(img1, blend.Add(img1, img2), strength), nil
				case "colorburn":
					return blend.Opacity(img1, blend.ColorBurn(img1, img2), strength), nil
				case "colordodge":
					return blend.Opacity(img1, blend.ColorDodge(img1, img2), strength), nil
				case "darken":
					return blend.Opacity(img1, blend.Darken(img1, img2), strength), nil
				case "difference":
					return blend.Opacity(img1, blend.Difference(img1, img2), strength), nil
				case "divide":
					return blend.Opacity(img1, blend.Divide(img1, img2), strength), nil
				case "exclusion":
					return blend.Opacity(img1, blend.Exclusion(img1, img2), strength), nil
				case "lighten":
					return blend.Opacity(img1, blend.Lighten(img1, img2), strength), nil
				case "linearburn":
					return blend.Opacity(img1, blend.LinearBurn(img1, img2), strength), nil
				case "linearLight":
					return blend.Opacity(img1, blend.LinearLight(img1, img2), strength), nil
				case "multiply":
					return blend.Opacity(img1, blend.Multiply(img1, img2), strength), nil
				case "normal":
					return blend.Opacity(img1, blend.Normal(img1, img2), strength), nil
				case "opacity":
					return blend.Opacity(img1, img2, strength), nil
				case "overlay":
					return blend.Opacity(img1, blend.Overlay(img1, img2), strength), nil
				case "screen":
					return blend.Opacity(img1, blend.Screen(img1, img2), strength), nil
				case "softlight":
					return blend.Opacity(img1, blend.SoftLight(img1, img2), strength), nil
				case "subtract":
					return blend.Opacity(img1, blend.Subtract(img1, img2), strength), nil
				}
				return blend.Opacity(img1, blend.Add(img1, img2), strength), nil
			})
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

	blendCmd.AddCommand(buildBlendModeCommand("add"))
	blendCmd.AddCommand(buildBlendModeCommand("colorburn"))
	blendCmd.AddCommand(buildBlendModeCommand("colordodge"))
	blendCmd.AddCommand(buildBlendModeCommand("darken"))
	blendCmd.AddCommand(buildBlendModeCommand("difference"))
	blendCmd.AddCommand(buildBlendModeCommand("divide"))
	blendCmd.AddCommand(buildBlendModeCommand("exclusion"))
	blendCmd.AddCommand(buildBlendModeCommand("lighten"))
	blendCmd.AddCommand(buildBlendModeCommand("linearburn"))
	blendCmd.AddCommand(buildBlendModeCommand("linearLight"))
	blendCmd.AddCommand(buildBlendModeCommand("multiply"))
	blendCmd.AddCommand(buildBlendModeCommand("normal"))
	blendCmd.AddCommand(buildBlendModeCommand("opacity"))
	blendCmd.AddCommand(buildBlendModeCommand("overlay"))
	blendCmd.AddCommand(buildBlendModeCommand("screen"))
	blendCmd.AddCommand(buildBlendModeCommand("softlight"))
	blendCmd.AddCommand(buildBlendModeCommand("subtract"))

	return blendCmd
}
