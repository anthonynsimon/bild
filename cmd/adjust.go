package cmd

import (
	"image"

	"github.com/anthonynsimon/bild/adjust"

	"github.com/spf13/cobra"
)

func brightness() *cobra.Command {
	var change float64

	var cmd = &cobra.Command{
		Use:     "brightness",
		Short:   "adjust the relative brightness of an image",
		Args:    cobra.ExactArgs(2),
		Example: "brightness --change 0.5 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return adjust.Brightness(img, change), nil
			})
		}}

	cmd.Flags().Float64VarP(&change, "change", "c", 0, "adjust change")

	return cmd
}

func contrast() *cobra.Command {
	var change float64

	var cmd = &cobra.Command{
		Use:     "contrast",
		Short:   "adjust the relative contrast of an image",
		Args:    cobra.ExactArgs(2),
		Example: "contrast --change 0.5 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return adjust.Contrast(img, change), nil
			})
		}}

	cmd.Flags().Float64VarP(&change, "change", "c", 0, "adjust change")

	return cmd
}

func gamma() *cobra.Command {
	var change float64

	var cmd = &cobra.Command{
		Use:     "gamma",
		Short:   "adjust the gamma of an image",
		Args:    cobra.ExactArgs(2),
		Example: "gamma --gamma 1.1 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return adjust.Gamma(img, change), nil
			})
		}}

	cmd.Flags().Float64VarP(&change, "gamma", "g", 0, "set the gamma of the image")

	return cmd
}

func hue() *cobra.Command {
	var change int

	var cmd = &cobra.Command{
		Use:     "hue",
		Short:   "adjust the hue of an image",
		Args:    cobra.ExactArgs(2),
		Example: "hue --change 170 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return adjust.Hue(img, change), nil
			})
		}}

	cmd.Flags().IntVarP(&change, "change", "c", 0, "adjust change")

	return cmd
}

func saturation() *cobra.Command {
	var change float64

	var cmd = &cobra.Command{
		Use:     "saturation",
		Short:   "adjust the saturation of an image",
		Args:    cobra.ExactArgs(2),
		Example: "saturation --change 170 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return adjust.Saturation(img, change), nil
			})
		}}

	cmd.Flags().Float64VarP(&change, "change", "c", 0, "adjust change")

	return cmd
}

func createAdjust() *cobra.Command {
	adjustCmd := &cobra.Command{
		Use:   "adjust",
		Short: "adjust basic image features like brightness or contrast",
	}

	adjustCmd.AddCommand(brightness())
	adjustCmd.AddCommand(contrast())
	adjustCmd.AddCommand(gamma())
	adjustCmd.AddCommand(hue())
	adjustCmd.AddCommand(saturation())

	return adjustCmd
}
