package cmd

import (
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/noise"
	"github.com/spf13/cobra"
)

func generateNoise() *cobra.Command {
	size := ""
	mono := false

	var cmd = &cobra.Command{
		Use:     "new",
		Short:   "generates an image filled with noise",
		Args:    cobra.ExactArgs(1),
		Example: "new -s 100x100 output.png",
		Run: func(cmd *cobra.Command, args []string) {
			fout := args[0]

			size, err := parseSizeStr(size)
			exitIfNotNil(err)

			result := noise.Generate(size.Width, size.Height, &noise.Options{
				Monochrome: mono,
			})

			encoder := resolveEncoder(fout, imgio.PNGEncoder())
			err = imgio.Save(fout, result, encoder)
			exitIfNotNil(err)
		}}

	cmd.Flags().StringVarP(&size, "size", "s", "512x512", "the width and height of the output image")
	cmd.Flags().BoolVarP(&mono, "monochrome", "m", false, "output monochrome noise")

	return cmd
}

func createNoise() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "noise",
		Short: "noise generators",
	}

	cmd.AddCommand(generateNoise())

	return cmd
}
