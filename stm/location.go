package stm

import (
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
)

func NewLocation(opts *Options) *Location {
	loc := &Location{
		opts: opts,
	}
	return loc
}

type Location struct {
	adp  Adapter
	opts *Options
}

func (loc *Location) Directory() string {
	return filepath.Join(
		loc.opts.publicPath,
		loc.opts.sitemapsPath,
	)
}

func (loc *Location) Path() string {
	return filepath.Join(
		loc.opts.publicPath,
		loc.opts.sitemapsPath,
		loc.Filename(),
	)
}

func (loc *Location) PathInPublic() string {
	return filepath.Join(
		loc.opts.sitemapsPath,
		loc.Filename(),
	)
}

func (loc *Location) URL() string {
	base, _ := url.Parse(loc.opts.sitemapsHost)

	var u *url.URL
	for _, ref := range []string{
			loc.opts.sitemapsPath, loc.Filename()} {
		u, _ = url.Parse(ref)
		base.ResolveReference(u)
	}

	return base.String()
}

func (loc *Location) Filesize() int64 {
	f, _ := os.Open(loc.Path())
	defer f.Close()
	fi, _ := f.Stat()
	return fi.Size()
}

func (loc *Location) Filename() string {
	if loc.opts.filename == "" && loc.opts.Namer() == nil {
		log.Fatal("No filename or namer set")
	}

	if loc.opts.filename == "" {
		loc.opts.SetFilename(loc.opts.Namer().String())
	}
	return loc.opts.filename
}

func (loc *Location) ReserveName() string {
	nmr := loc.opts.Namer()
	if nmr != nil {
		loc.Filename()
		nmr.Next()
	}

	return loc.opts.filename
}

func (loc *Location) IsReservedName() bool {
	if loc.opts.filename == "" {
		return false
	}
	return true
}

func (loc *Location) Namer() *Namer {
	return loc.opts.Namer()
}

func (loc *Location) IsVerbose() bool {
	return loc.opts.verbose
}

func (loc *Location) Write(data []byte, linkCount int) {
	loc.adp.Write(loc, data)
	if loc.IsVerbose() {
		pp.Println(loc.Summary(linkCount))
	}
}

func (loc *Location) Summary(linkCount int) string {
	// filesize = number_to_human_size(loc.Filesize())
	// width = self.class::PATH_OUTPUT_WIDTH
	// path = SitemapGenerator::Utilities.ellipsis(self.path_in_public, width)
	// fmt.Sprintf("+ #{('%-'+width.to_s+'s') % path} #{'%10s' % link_count} links / #{'%10s' % filesize}")
	return ""
}
