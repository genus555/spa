package clientloop

import (
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	src_file, err := os.Open(src)
	if err != nil {return err}
	defer src_file.Close()

	dst_file, err := os.Create(dst)
	if err != nil {return err}
	defer dst_file.Close()

	_, err = io.Copy(dst_file, src_file)
	if err != nil {return err}
	return nil
}