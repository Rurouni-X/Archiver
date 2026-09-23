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

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

func pack(cmd *cobra.Command, args []string) {

	var encoder compression.Encoder
	if len(args) == 0 || args[0] == "" {
		HandleError(utils.ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknow method compression")
	}

	filePath := args[0] // путь до файла

	r, err := os.Open(filePath)
	if err != nil {
		HandleError(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		HandleError(err)
	}

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packFileName(filePath), packed, 0644)
	if err != nil {
		HandleError(err)
	}
}

func packFileName(path string) string {

	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	return baseName + "." + utils.PackedExtension
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
