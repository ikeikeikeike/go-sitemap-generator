package stm

import "runtime"

func NewSitemap() *Sitemap {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sm := &Sitemap{
		opts: NewOptions(),
	}
	return sm
}

type Sitemap struct {
	opts *Options
	bldr Builder
}

func (sm *Sitemap) SetDefaultHost(host string) {
	sm.opts.SetDefaultHost(host)
}

func (sm *Sitemap) SetSitemapsPath(path string) {
	sm.opts.SetSitemapsPath(path)
}

func (sm *Sitemap) SetAdapter(adp Adapter) {
	sm.opts.SetAdapter(adp)
}

func (sm *Sitemap) Create() Builder {
	sm.bldr = NewBuilderFile(sm.opts.Location())
	go sm.bldr.run()
	return sm.bldr
}
