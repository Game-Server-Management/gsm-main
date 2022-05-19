package utilities

import (
	"io"
	"io/fs"
	"net/http"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func DownloadFile(path string, url string, mode fs.FileMode) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	err = os.Chmod(path, mode)
	return err
}
