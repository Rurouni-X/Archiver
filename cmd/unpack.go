package cmd

import (
	"Archiver/pkg/compression"
	"Archiver/pkg/compression/vlc"
	"Archiver/utils"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Pack file",
	Run:   unpack,
}

func unpack(cmd *cobra.Command, args []string) {

	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		HandleError(utils.ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("unknow method decompression")
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		HandleError(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		HandleError(err)
	}

	packed := decoder.Decode(data)

	err = os.WriteFile(unpackFileName(filePath), []byte(packed), 0644)
	if err != nil {
		HandleError(err)
	}
}

func unpackFileName(path string) string {

	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	return baseName + "." + utils.UnpackedExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("method", "m", "", "decompression method")

	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
