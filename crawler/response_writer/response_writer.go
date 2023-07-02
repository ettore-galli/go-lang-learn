package response_writer

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

type ResponseWriter struct {
	fileNameBuilder func(string) string
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func BuildFileNameFromContent(content string) string {
	return fmt.Sprintf("Response_%v", GetMD5Hash(content))
}

func PerformFileWrite(fileFqn string, content string) error {
	wfile, err := os.Create(fileFqn)
	if err != nil {
		return err
	}
	defer wfile.Close()

	_, werr := wfile.WriteString(content)
	if werr != nil {
		wfile.Close()
		return err
	}

	return nil

}

func (rp *ResponseWriter) WriteResponse(outputDir string, response string) error {
	filename := rp.fileNameBuilder(response)
	fileFqn := filepath.Join(outputDir, filename)
	return PerformFileWrite(fileFqn, response)
}

func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{fileNameBuilder: BuildFileNameFromContent}
}
