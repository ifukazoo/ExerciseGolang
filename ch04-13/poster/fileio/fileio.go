package fileio

import "os"

// WriteFile :ファイル書き込み
func WriteFile(filename string, data []byte) error {
	var (
		f   *os.File
		err error
	)
	if f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644); err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(data); err != nil {
		return err
	}
	return nil
}
