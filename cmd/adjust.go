package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func buildAdjustCommand(name string) *cobra.Command {
	var change float64

	var cmd = &cobra.Command{
		Use:     name,
		Short:   fmt.Sprintf("adjust the relative %s of an image", name),
		Args:    cobra.ExactArgs(1),
		Example: fmt.Sprintf("%s --change 0.5 image.jpg", name),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.Flags().Float64VarP(&change, "change", "c", 0, "adjust change")

	return cmd
}

func createAdjust() *cobra.Command {
	adjustCmd := &cobra.Command{
		Use:   "adjust",
		Short: "adjust basic image features like brightness or contrast",
	}

	brightnessCmd := buildAdjustCommand("brightness")
	adjustCmd.AddCommand(brightnessCmd)

	contrastCmd := buildAdjustCommand("contrast")
	adjustCmd.AddCommand(contrastCmd)

	gammaCmd := buildAdjustCommand("gamma")
	adjustCmd.AddCommand(gammaCmd)

	hueCmd := buildAdjustCommand("hue")
	adjustCmd.AddCommand(hueCmd)

	saturationCmd := buildAdjustCommand("saturation")
	adjustCmd.AddCommand(saturationCmd)

	return adjustCmd
}
