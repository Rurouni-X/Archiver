package cmd

import "github.com/spf13/cobra"

var vlcUnpackCmd = &cobra.Command{
	Use: "vlc",
	Short: "Unpack file using variable-length code",
	Run: unpack,
}

func unpack(_ *cobra.Command, args []string) {

}