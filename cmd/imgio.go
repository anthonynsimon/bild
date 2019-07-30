package cmd

import (
	"image"

	"github.com/spf13/cobra"
)

func encode() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "encode",
		Short:   "encodes the input image using the desired encoding set by the destination file extension",
		Args:    cobra.ExactArgs(2),
		Example: "encode input.jpg output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fin := args[0]
			fout := args[1]

			// Apply takes care of resolving the destination encoder
			apply(fin, fout, func(img image.Image) (image.Image, error) {
				return img, nil
			})
		}}
	return cmd
}

func createImgio() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "imgio",
		Short: "i/o operations on images",
	}

	cmd.AddCommand(encode())

	return cmd
}
