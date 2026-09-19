package cmd

import (
	"Archiver/pkg/vlc"
	"Archiver/utils"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcPackCmd = &cobra.Command{
	Use: "vlc",
	Short: "Pack file using variable-length code",
	Run: pack,
}

func pack(_ *cobra.Command, args []string) {

	if len(args) == 0 || args[0] == "" {
		HandleError(utils.ErrEmptyPath)
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

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packFileName(filePath), []byte(packed), 0644)
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
	packCmd.AddCommand(vlcPackCmd)
}