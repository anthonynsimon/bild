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

func rotate() *cobra.Command {
	var angle float64
	var pivot string
	var resizeBounds bool

	var cmd = &cobra.Command{
		Use:     "rotate",
		Short:   "rotate an image by the given angle in degrees",
		Args:    cobra.ExactArgs(2),
		Example: "rotate --angle 90 --resize-bounds input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			var opts *transform.RotationOptions
			if pivot != "" || resizeBounds {
				opts = &transform.RotationOptions{ResizeBounds: resizeBounds}
				if pivot != "" {
					s, err := parseSizeStr(pivot)
					exitIfNotNil(err)
					opts.Pivot = &image.Point{X: s.Width, Y: s.Height}
				}
			}

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return transform.Rotate(img, angle, opts), nil
			})
		}}

	cmd.Flags().Float64VarP(&angle, "angle", "a", 0, "rotation angle in degrees (clockwise)")
	cmd.Flags().StringVarP(&pivot, "pivot", "p", "", "pivot point as XxY (e.g. 100x100)")
	cmd.Flags().BoolVar(&resizeBounds, "resize-bounds", false, "resize image bounds to fit rotated content")

	return cmd
}

func fliph() *cobra.Command {
	return &cobra.Command{
		Use:     "fliph",
		Short:   "flip an image horizontally",
		Args:    cobra.ExactArgs(2),
		Example: "fliph input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			apply(args[0], args[1], func(img image.Image) (image.Image, error) {
				return transform.FlipH(img), nil
			})
		},
	}
}

func flipv() *cobra.Command {
	return &cobra.Command{
		Use:     "flipv",
		Short:   "flip an image vertically",
		Args:    cobra.ExactArgs(2),
		Example: "flipv input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			apply(args[0], args[1], func(img image.Image) (image.Image, error) {
				return transform.FlipV(img), nil
			})
		},
	}
}

func resize() *cobra.Command {
	var width, height int
	var filter string

	var cmd = &cobra.Command{
		Use:     "resize",
		Short:   "resize an image to the given dimensions",
		Args:    cobra.ExactArgs(2),
		Example: "resize --width 800 --height 600 --filter lanczos input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			f, err := parseResampleFilter(filter)
			exitIfNotNil(err)

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return transform.Resize(img, width, height, f), nil
			})
		}}

	cmd.Flags().IntVarP(&width, "width", "w", 0, "target width in pixels")
	cmd.Flags().IntVarP(&height, "height", "h", 0, "target height in pixels")
	cmd.Flags().StringVarP(&filter, "filter", "f", "linear", "resampling filter (nearestneighbor, box, linear, gaussian, mitchellnetravali, catmullrom, lanczos)")

	return cmd
}

func crop() *cobra.Command {
	var rect string

	var cmd = &cobra.Command{
		Use:     "crop",
		Short:   "crop an image to the given rectangle",
		Args:    cobra.ExactArgs(2),
		Example: "crop --rect 0x0+512x256 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			r, err := parseRectStr(rect)
			exitIfNotNil(err)

			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return transform.Crop(img, r), nil
			})
		}}

	cmd.Flags().StringVarP(&rect, "rect", "r", "", "crop rectangle as X0xY0+X1xY1 (e.g. 0x0+512x256)")

	return cmd
}

func translate() *cobra.Command {
	var dx, dy int

	var cmd = &cobra.Command{
		Use:     "translate",
		Short:   "shift an image by dx and dy pixels",
		Args:    cobra.ExactArgs(2),
		Example: "translate --dx 100 --dy 50 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			apply(args[0], args[1], func(img image.Image) (image.Image, error) {
				return transform.Translate(img, dx, dy), nil
			})
		}}

	cmd.Flags().IntVar(&dx, "dx", 0, "horizontal shift in pixels (positive moves right)")
	cmd.Flags().IntVar(&dy, "dy", 0, "vertical shift in pixels (positive moves up)")

	return cmd
}

func shearh() *cobra.Command {
	var angle float64

	var cmd = &cobra.Command{
		Use:     "shearh",
		Short:   "apply horizontal shear transformation",
		Args:    cobra.ExactArgs(2),
		Example: "shearh --angle 30 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			apply(args[0], args[1], func(img image.Image) (image.Image, error) {
				return transform.ShearH(img, angle), nil
			})
		}}

	cmd.Flags().Float64VarP(&angle, "angle", "a", 0, "shear angle in degrees")

	return cmd
}

func shearv() *cobra.Command {
	var angle float64

	var cmd = &cobra.Command{
		Use:     "shearv",
		Short:   "apply vertical shear transformation",
		Args:    cobra.ExactArgs(2),
		Example: "shearv --angle 30 input.jpg output.jpg",
		Run: func(cmd *cobra.Command, args []string) {
			apply(args[0], args[1], func(img image.Image) (image.Image, error) {
				return transform.ShearV(img, angle), nil
			})
		}}

	cmd.Flags().Float64VarP(&angle, "angle", "a", 0, "shear angle in degrees")

	return cmd
}

func createTransform() *cobra.Command {
	var transformCmd = &cobra.Command{
		Use:   "transform",
		Short: "apply geometric transformations to images",
	}

	transformCmd.AddCommand(zoom())
	transformCmd.AddCommand(rotate())
	transformCmd.AddCommand(fliph())
	transformCmd.AddCommand(flipv())
	transformCmd.AddCommand(resize())
	transformCmd.AddCommand(crop())
	transformCmd.AddCommand(translate())
	transformCmd.AddCommand(shearh())
	transformCmd.AddCommand(shearv())

	return transformCmd
}
