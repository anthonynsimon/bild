package cmd

import (
	"image"

	"github.com/anthonynsimon/bild/effect"
	"github.com/spf13/cobra"
)

func grayscale() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "grayscale",
		Aliases: []string{"gray"},
		Short:   "applies the grayscale effect",
		Args:    cobra.ExactArgs(2),
		Example: "grayscale input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Grayscale(img), nil
			})
		}}
	return cmd
}

func sepia() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "sepia",
		Short:   "applies the sepia effect",
		Args:    cobra.ExactArgs(2),
		Example: "sepia input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Sepia(img), nil
			})
		}}
	return cmd
}

func sharpen() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "sharpen",
		Aliases: []string{"sharp"},
		Short:   "applies the sharpen effect",
		Args:    cobra.ExactArgs(2),
		Example: "sharpen input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Sharpen(img), nil
			})
		}}
	return cmd
}

func sobel() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "sobel",
		Short:   "applies the sobel effect",
		Args:    cobra.ExactArgs(2),
		Example: "sobel input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Sobel(img), nil
			})
		}}
	return cmd
}

func invert() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "invert",
		Short:   "applies the invert effect",
		Args:    cobra.ExactArgs(2),
		Example: "invert input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Invert(img), nil
			})
		}}
	return cmd
}

func median() *cobra.Command {
	var radius float64

	var cmd = &cobra.Command{
		Use:     "median",
		Short:   "applies the median effect",
		Args:    cobra.ExactArgs(2),
		Example: "median --radius 2.5 input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Median(img, radius), nil
			})
		}}

	cmd.Flags().Float64VarP(&radius, "radius", "r", 3, "the effect's radius")

	return cmd
}

func erode() *cobra.Command {
	var radius float64

	var cmd = &cobra.Command{
		Use:     "erode",
		Short:   "applies the erode effect",
		Args:    cobra.ExactArgs(2),
		Example: "erode --radius 0.5 input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Erode(img, radius), nil
			})
		}}

	cmd.Flags().Float64VarP(&radius, "radius", "r", 0.5, "the effect's radius")

	return cmd
}

func dilate() *cobra.Command {
	var radius float64

	var cmd = &cobra.Command{
		Use:     "dilate",
		Short:   "applies the dilate effect",
		Args:    cobra.ExactArgs(2),
		Example: "dilate --radius 0.5 input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Dilate(img, radius), nil
			})
		}}

	cmd.Flags().Float64VarP(&radius, "radius", "r", 0.5, "the effect's radius")

	return cmd
}

func edgedetection() *cobra.Command {
	var radius float64

	var cmd = &cobra.Command{
		Use:     "edgedetection",
		Short:   "applies the edgedetection effect",
		Args:    cobra.ExactArgs(2),
		Example: "edgedetection --radius 0.5 input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.EdgeDetection(img, radius), nil
			})
		}}

	cmd.Flags().Float64VarP(&radius, "radius", "r", 0.5, "the effect's radius")

	return cmd
}

func emboss() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "emboss",
		Short:   "applies the emboss effect",
		Args:    cobra.ExactArgs(2),
		Example: "emboss input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.Emboss(img), nil
			})
		}}

	return cmd
}

func unsharpmask() *cobra.Command {
	var radius float64
	var amount float64

	var cmd = &cobra.Command{
		Use:     "unsharpmask",
		Short:   "applies the unsharpmask effect",
		Args:    cobra.ExactArgs(2),
		Example: "unsharpmask input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return effect.UnsharpMask(img, radius, amount), nil
			})
		}}

	cmd.Flags().Float64VarP(&radius, "radius", "r", 25, "the effect's radius")
	cmd.Flags().Float64VarP(&radius, "amount", "a", 1, "the effect's amount")

	return cmd
}

func createEffect() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "effect",
		Short: "apply effects on images",
	}

	cmd.AddCommand(grayscale())
	cmd.AddCommand(sepia())
	cmd.AddCommand(sharpen())
	cmd.AddCommand(sobel())
	cmd.AddCommand(invert())
	cmd.AddCommand(median())
	cmd.AddCommand(erode())
	cmd.AddCommand(dilate())
	cmd.AddCommand(edgedetection())
	cmd.AddCommand(emboss())
	cmd.AddCommand(unsharpmask())

	return cmd
}
