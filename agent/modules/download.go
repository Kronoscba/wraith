package modules

import (
	"io"
	"net"
	"os"
)

// DownloadFile receives a file from the server and saves it locally.
func DownloadFile(conn net.Conn, destPath string) error {
	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, conn)
	return err
}
