package cmd

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/channel"
	"github.com/anthonynsimon/bild/histogram"
	"github.com/spf13/cobra"
)

func newHisto() *cobra.Command {
	var channels string

	var cmd = &cobra.Command{
		Use:     "new",
		Short:   "creates a RBG histogram from an input image",
		Args:    cobra.ExactArgs(2),
		Example: "new --channels rgb input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				onlyChannels := []channel.Channel{channel.Alpha}
				for _, c := range channels {
					switch c {
					case 'r':
						onlyChannels = append(onlyChannels, channel.Red)
					case 'g':
						onlyChannels = append(onlyChannels, channel.Green)
					case 'b':
						onlyChannels = append(onlyChannels, channel.Blue)
					default:
						return nil, fmt.Errorf("unknown channel alias '%c'", c)
					}
				}

				hist := histogram.NewRGBAHistogram(img)
				result := channel.ExtractMultiple(hist.Image(), onlyChannels...)
				return result, nil
			})
		}}

	cmd.Flags().StringVarP(&channels, "channels", "c", "rgb", "the channels to include in the histogram")

	return cmd
}

func createHistogram() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "histogram",
		Short: "histogram operations on images",
	}

	cmd.AddCommand(newHisto())

	return cmd
}
