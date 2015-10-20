package stm

import "runtime"

func NewSitemap() *Sitemap {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sm := &Sitemap{
		opts: NewOptions(),
		adp:  NewFileAdapter(),
	}
	return sm
}

type Sitemap struct {
	opts  *Options
	loc   *Location
	namer *Namer
	bldr  Builder
	adp   Adapter
}

func (sm *Sitemap) SetDefaultHost(host string) {
	sm.opts.SetDefaultHost(host)
}

func (sm *Sitemap) SetSitemapsPath(path string) {
	sm.opts.SetSitemapsPath(path)
}

func (sm *Sitemap) SetAdapter(adapter Adapter) {
	sm.adp = adapter
}

func (sm *Sitemap) Create() Builder {
	sm.bldr = NewBuilderFile()
	go sm.bldr.run()
	return sm.bldr
}

func (sm *Sitemap) Location() *Location {
	loc := NewLocation(
		sm.opts.SitemapsHost(),
		sm.Namer(),
		sm.opts.publicPath,
		sm.opts.sitemapsPath,
		sm.adp,
		sm.opts.verbose,
		sm.opts.compress,
	)
	return loc
}

func (sm *Sitemap) Namer() *Namer {
	if sm.namer == nil {
		if sm.bldr == nil {
			sm.namer = sm.bldr.loc.namer
		} else {
			sm.namer = NewNamer(sm.opts.filename)
		}
	}
	return sm.namer
}
