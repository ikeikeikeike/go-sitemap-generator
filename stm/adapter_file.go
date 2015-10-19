package stm

import (
	"compress/zlib"
	"log"
	"os"
	"regexp"
)

var gzipPtn = regexp.MustCompile(".gz$")

func NewFileAdapter() *FileAdapter {
	adapter := &FileAdapter{}
	return adapter
}

type FileAdapter struct{}

func (a *FileAdapter) Write(loc *Location, data []byte) {
	dir := loc.Directory()
	fi, err := os.Stat(dir)
	if err != nil {
		_ = os.MkdirAll(dir, 0755)
	} else if !fi.IsDir() {
		log.Fatal("%s should be a directory", dir)
	}

	file, _ := os.Open(loc.Path())
	fi, err = file.Stat()
	if err != nil {
		log.Fatal("%s file not exists", loc.Path())
	} else if !fi.Mode().IsRegular() {
		log.Fatal("%s should be a filename", loc.Path())
	}

	if gzipPtn.MatchString(loc.Path()) {
		a.gzip(file, data)
	} else {
		a.plain(file, data)
	}
}

func (a *FileAdapter) gzip(file *os.File, data []byte) {
	gz := zlib.NewWriter(file)
	defer gz.Close()
	gz.Write(data)
}

func (a *FileAdapter) plain(file *os.File, data []byte) {
	file.Write(data)
	defer file.Close()
}
