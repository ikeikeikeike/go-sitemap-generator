package sitemap

import "runtime"

func NewSitemap() *Sitemap {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sm := &Sitemap{opts: NewOptions()}
	return sm
}

type Sitemap struct {
	opts *Options
}

func (sm *Sitemap) SetDefaultHost(host string) {
	sm.opts.SetDefaultHost(host)
}

func (sm *Sitemap) SetSitemapsPath(path string) {
	sm.opts.SetSitemapsPath(path)
}

func (sm *Sitemap) Create() Builder {
	bldr := NewBuilderFile()
	go bldr.run()
	return bldr
}
