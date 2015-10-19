package stm

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
)

func NewLocation() *Location {
	loc := &Location{
		adapter:    NewFileAdapter(),
		publicPath: "public/",
	}
	return loc
}

type Location struct {
	adapter Adapter

	verbose      bool
	host         string
	publicPath   string
	sitemapsPath string
}

func (loc *Location) SetPublicPath(path string) {
	loc.publicPath = path
}

func (loc *Location) SetSitemapsPath(path string) {
	loc.sitemapsPath = path
}

// func (loc *Location) with(opts={})
// self.merge(opts)
// }

func (loc *Location) Directory() string {
	return filepath.Join(loc.publicPath, loc.sitemapsPath)
}

func (loc *Location) Path() string {
	return filepath.Join(loc.publicPath, loc.sitemapsPath, loc.Filename())
}

func (loc *Location) PathInPublic() string {
	return filepath.Join(loc.sitemapsPath, loc.Filename())
}

func (loc *Location) URL() string {
	base, _ := url.Parse(loc.host)

	var u *url.URL
	for _, ref := range []string{loc.sitemapsPath, loc.Filename()} {
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
	return ""

	// raise SitemapGenerator::SitemapError, "No filename or namer set" unless self[:filename] || self[:namer]
	// unless self[:filename]
	// self.send(:[]=, :filename, self[:namer].to_s, :super => true)

	// if self[:compress] == false || (self[:namer] && self[:namer].start? && self[:compress] == :all_but_first) {
	// self[:filename].gsub!(/\.gz$/, '')
	// }
	// self[:filename]
}

// func (loc *Location) ReserveName() {
// if self[:namer]
// filename
// self[:namer].next
// end
// self[:filename]
// }

// func (loc *Location) IsReservedName() bool {
// !!self[:filename]
// }

// func (loc *Location) namer() {
// self[:namer]
// }

func (loc *Location) IsVerbose() bool {
	return loc.verbose
}

// func (loc *Location) []=(key, value, opts={})
// if !opts[:super]
// case key
// when :namer
// super(:filename, nil)
// when :filename
// super(:namer, nil)
// end
// end
// super(key, value)
// }

func (loc *Location) Write(data []byte, linkCount int) {
	loc.adapter.Write(loc, data)
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
