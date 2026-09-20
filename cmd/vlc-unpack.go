package cmd

import (
	"Archiver/pkg/vlc"
	"Archiver/utils"
	"io"
	"path/filepath"
	"strings"
	"os"

	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use: "vlc",
	Short: "Unpack file using variable-length code",
	Run: unpack,
}

func unpack(_ *cobra.Command, args []string) {

	if len(args) == 0 || args[0] == "" {
		HandleError(utils.ErrEmptyPath)
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

	packed := vlc.Decode(data)

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
	unpackCmd.AddCommand(vlcUnpackCmd)
}