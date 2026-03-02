package cmd

import (
	"image"

	"github.com/anthonynsimon/bild/transform"
	"github.com/spf13/cobra"
)

func zoom() *cobra.Command {
	var factor float64
	var pivot string

	var cmd = &cobra.Command{
		Use:     "zoom",
		Short:   "apply zoom transformation to an input image",
		Args:    cobra.ExactArgs(2),
		Example: "zoom --factor 2.0 --pivot 280x0 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			var opts *transform.ZoomOptions
			if pivot != "" {
				s, err := parseSizeStr(pivot)
				exitIfNotNil(err)
				opts = &transform.ZoomOptions{Pivot: &image.Point{X: s.Width, Y: s.Height}}
			}

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return transform.Zoom(img, factor, opts), nil
			})
		}}

	cmd.Flags().Float64VarP(&factor, "factor", "f", 1.0, "the zoom factor (>1.0 zooms in, <1.0 zooms out)")
	cmd.Flags().StringVarP(&pivot, "pivot", "p", "", "pivot point as XxY (e.g. 280x0 for top-right)")

	return cmd
}

func createTransform() *cobra.Command {
	var transformCmd = &cobra.Command{
		Use:   "transform",
		Short: "apply geometric transformations to images",
	}

	transformCmd.AddCommand(zoom())

	return transformCmd
}
