package tarball

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"path/filepath"
)

var (
	// ErrGzipFileNotFound the file not found in the gzip
	ErrGzipFileNotFound = errors.New("file not found in the gzip")
	// ErrNotGzipType the file is not a gzip
	ErrNotGzipType = errors.New("file is not a gzip type")
	// ErrInvalidFileName the file name is invalid
	ErrInvalidFileName = errors.New("invalid file name")
)

// ExtractFile founds and reads a specific file into a gzip file and folders recursively
func ExtractFile(reader io.Reader, out io.Writer, fileName string) (string, error) {
	if fileName == "" {
		return "", ErrInvalidFileName
	}
	archive, err := gzip.NewReader(reader)
	if err == io.EOF || err == gzip.ErrHeader {
		return "", ErrNotGzipType
	} else if err != nil {
		return "", err
	}
	defer archive.Close()

	tarReader := tar.NewReader(archive)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			return "", ErrGzipFileNotFound
		} else if err != nil {
			return header.Name, err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			name := filepath.Base(header.Name)
			if fileName == name {
				_, err := io.Copy(out, tarReader)
				return header.Name, err
			}
		default:
			continue
		}
	}
}
