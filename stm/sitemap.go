package stm

import "runtime"

func NewSitemap() *Sitemap {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sm := &Sitemap{opts: NewOptions()}
	return sm
}

type Sitemap struct {
	opts  *Options
	loc   *Location
	bldr  Builder
	namer *Namer
}

func (sm *Sitemap) SetDefaultHost(host string) {
	sm.opts.SetDefaultHost(host)
}

func (sm *Sitemap) SetSitemapsPath(path string) {
	sm.opts.SetSitemapsPath(path)
}

func (sm *Sitemap) Create() Builder {
	sm.bldr = NewBuilderFile()
	go sm.bldr.run()
	return sm.bldr
}

// func (sm *Sitemap) Location() *Location {
// loc := NewLocation(
// host:  sm.opts.SitemapsHost(),
// namer:  sm.Namer(),
// public_path:  sm.opts.publicPath,
// sitemaps_path:  sm.opts.sitemapsPath,
// adapter:  sm.opts.adapter,
// verbose:  verbose,
// compress:  @compress
// )
// return loc
// }

// func (sm *Sitemap) Namer() *Namer {
// if sm.namer == nil {
// if sm.bldr == nil {
// sm.namer = sm.bldr.loc.namer
// } else {
// sm.namer = NewNamer(sm.opts.filename)
// }
// }
// return sm.namer
// }
