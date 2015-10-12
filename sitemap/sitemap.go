package sitemap

func NewSitemap() *Sitemap {
	sm := &Sitemap{opts: NewOptions()}
	return sm
}

type Sitemap struct {
	opts *Options
	bld Builder
}

func (sm *Sitemap) SetDefaultHost(host string) {
	sm.opts.SetDefaultHost(host)
}

func (sm *Sitemap) SetSitemapsPath(path string) {
	sm.opts.SetSitemapsPath(path)
}

func (sm *Sitemap) Create() Builder {
	sm.bld = &BuilderURL{}
	return sm.bld
}
