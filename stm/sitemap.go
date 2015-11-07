package stm

import (
	"log"
	"runtime"
)

func NewSitemap() *Sitemap {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

	sm := &Sitemap{
		opts: NewOptions(),
	}
	return sm
}

type Sitemap struct {
	opts  *Options
	bldr  Builder
	bldrs Builder
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

func (sm *Sitemap) Create() *Sitemap {
	sm.bldrs = NewBuilderIndexfile(sm.opts.Location())
	// go sm.bldr.run()
	// go sm.bldrs.run()
	return sm
}

func (sm *Sitemap) Add(url interface{}) *Sitemap {
	if sm.bldr == nil {
		sm.bldr = NewBuilderFile(sm.opts.Location())
	}

	err := sm.bldr.Add(url)
	if err != nil {
		if err.FinalizedError() {
			sm.bldr = NewBuilderFile(sm.opts.Location())
			return sm.Add(url)
		} else if err.FullError() {
			sm.finalizeFile()
			return sm.Add(url)
		}
	}

	return sm
}

func (sm *Sitemap) finalize() {
	sm.finalizeFile()
	sm.finalizeIndexfile()
}

func (sm *Sitemap) finalizeFile() {
	sm.bldr.Finalize()
	sm.bldrs.Add(sm.bldr)
}

func (sm *Sitemap) finalizeIndexfile() {
	sm.bldrs.Finalize()
	sm.bldrs.Write()
}
