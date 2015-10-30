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
	bldrIdx *BuilderIndexFile
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

func (sm *Sitemap) Add() Builder {
	sm.bldr = NewBuilderFile(sm.opts.Location())
	go sm.bldr.run()
	return sm.bldr
}

func (sm *Sitemap) Create() Builder {
	sm.bldrIdx = NewBuilderIndexFile()
	go sm.bldrIdx.run()
	return sm.bldrIdx
}
