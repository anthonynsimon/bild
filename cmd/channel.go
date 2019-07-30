package cmd

import (
	"fmt"
	"image"

	"github.com/anthonynsimon/bild/channel"
	"github.com/spf13/cobra"
)

func extractChannel() *cobra.Command {
	var channels string

	var cmd = &cobra.Command{
		Use:     "extract",
		Short:   "extracts RGBA channels from an input image",
		Args:    cobra.ExactArgs(2),
		Example: "extract --channels rba input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				onlyChannels := []channel.Channel{}
				for _, c := range channels {
					switch c {
					case 'r':
						onlyChannels = append(onlyChannels, channel.Red)
					case 'g':
						onlyChannels = append(onlyChannels, channel.Green)
					case 'b':
						onlyChannels = append(onlyChannels, channel.Blue)
					case 'a':
						onlyChannels = append(onlyChannels, channel.Alpha)
					default:
						return nil, fmt.Errorf("unknown channel alias '%c'", c)
					}
				}

				result := channel.ExtractMultiple(img, onlyChannels...)
				return result, nil
			})
		}}

	cmd.Flags().StringVarP(&channels, "channels", "c", "rgba", "the channels to include in the histogram")

	return cmd
}

func createChannel() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "channel",
		Short: "channel operations on images",
	}

	cmd.AddCommand(extractChannel())

	return cmd
}
