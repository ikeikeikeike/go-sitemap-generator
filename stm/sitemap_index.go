package stm

import (
	"log"
	"runtime"
)

// NewSitemap returns the created the Sitemap's pointer
func NewSitemapIndex() *SitemapIndex {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	sm := &SitemapIndex{
		opts: NewOptions(),
	}
	return sm
}

// Sitemap provides interface for create sitemap xml file and that has convenient interface.
// And also needs to use first this struct if it wants to use this package.
type SitemapIndex struct {
	opts  *Options
	indxbldr *BuilderIndexfile
}

// SetDefaultHost is your website's host name
func (sm *SitemapIndex) SetDefaultHost(host string) {
	sm.opts.SetDefaultHost(host)
}

// SetSitemapsHost is the remote host where your sitemaps will be hosted
func (sm *SitemapIndex) SetSitemapsHost(host string) {
	sm.opts.SetSitemapsHost(host)
}

// SetSitemapsPath sets this to a directory/path if you don't
// want to upload to the root of your `SitemapsHost`
func (sm *SitemapIndex) SetSitemapsPath(path string) {
	sm.opts.SetSitemapsPath(path)
}

// SetPublicPath is the directory to write sitemaps to locally
func (sm *SitemapIndex) SetPublicPath(path string) {
	sm.opts.SetPublicPath(path)
}

// SetAdapter can switch output file storage.
// We have S3Adapter and FileAdapter (default: FileAdapter)
func (sm *SitemapIndex) SetAdapter(adp Adapter) {
	sm.opts.SetAdapter(adp)
}

// SetVerbose can switch verbose output to console.
func (sm *SitemapIndex) SetVerbose(verbose bool) {
	sm.opts.SetVerbose(verbose)
}

// SetCompress can switch compress for the output file.
func (sm *SitemapIndex) SetCompress(compress bool) {
	sm.opts.SetCompress(compress)
}

// SetFilename can apply any name in this method if you wants to change output file name
func (sm *SitemapIndex) SetFilename(filename string) {
	sm.opts.SetFilename(filename)
}

func (sm *SitemapIndex) GetLocation() *Location {
	return sm.indxbldr.loc
}

// Create method must be that calls first this method in that before call to Add method on this struct.
func (sm *SitemapIndex) Create() *SitemapIndex {
	sm.indxbldr = NewBuilderIndexfile(sm.opts.IndexLocation())
	return sm
}

// Add Should call this after call to Create method on this struct.
func (sm *SitemapIndex) Add(link interface{}) *SitemapIndex {
	err := sm.indxbldr.Add(link)
	if err != nil {
		log.Printf("%v", "Could not add index sitemap entry")
	}

	return sm
}

func (sm *SitemapIndex) AddLocation(location *Location) *SitemapIndex {
	err := sm.indxbldr.AddLocation(location)
	if err != nil {
		log.Printf("%v", "Could not add index sitemap entry")
	}

	return sm
}

func (sm *SitemapIndex) Finalize() *SitemapIndex {
	sm.indxbldr.Write()
	return sm
}

// XMLContent returns the XML content of the sitemap
func (sm *SitemapIndex) XMLContent() []byte {
	return sm.indxbldr.XMLContent()
}

// PingSearchEngines requests some ping server.
// It also has that includes PingSearchEngines function.
func (sm *SitemapIndex) PingSearchEngines(urls ...string) {
	PingSearchEngines(sm.opts, urls...)
}
