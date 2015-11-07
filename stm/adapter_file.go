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

func (adp *FileAdapter) Write(loc *Location, data []byte) {
	dir := loc.Directory()
	fi, err := os.Stat(dir)
	if err != nil {
		_ = os.MkdirAll(dir, 0755)
	} else if !fi.IsDir() {
		log.Fatalf("%s should be a directory", dir)
	}

	file, _ := os.OpenFile(loc.Path(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	fi, err = file.Stat()
	if err != nil {
		log.Fatalf("%s file not exists", loc.Path())
	} else if !fi.Mode().IsRegular() {
		log.Fatalf("%s should be a filename", loc.Path())
	}
	defer file.Close()

	if gzipPtn.MatchString(loc.Path()) {
		adp.gzip(file, data)
	} else {
		adp.plain(file, data)
	}
}

func (adp *FileAdapter) gzip(file *os.File, data []byte) {
	gz := zlib.NewWriter(file)
	defer gz.Close()
	gz.Write(data)
}

func (adp *FileAdapter) plain(file *os.File, data []byte) {
	file.Write(data)
	defer file.Close()
}
