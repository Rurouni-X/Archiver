package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const PackedExtension = "vlc"
var ErrEmptyPath = errors.New("path to file is not specified")

var vlcCmd = &cobra.Command{
	Use: "vlc",
	Short: "Pack file using variable-length code",
	Run: pack,
}

func pack(_ *cobra.Command, args []string) {

	if len(args) == 0 || args[0] == "" {
		HandleError(ErrEmptyPath)
	}

	filePath := args[0] // путь до файла

	r, err := os.Open(filePath)
	if err != nil {
		HandleError(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		HandleError(err)
	}

	packed := ""

	fmt.Println(string(data))

	err = os.WriteFile(packFileName(filePath), []byte(packed), 0644)
	if err != nil {
		HandleError(err)
	}
}

func packFileName(path string) string {

	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, ext)

	return baseName + "." + PackedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}