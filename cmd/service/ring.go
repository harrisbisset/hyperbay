package service

import (
	"fmt"

	"github.com/harrisbisset/webrelay/service/toml"
)

type sitePath struct {
	Path string
	toml.Site
}

func GetRing(sites []toml.Site) []sitePath {

	NewSitePath := func(Path string, Site toml.Site, Count int) (sitePath, int) {
		return sitePath{
			Path: Path,
			Site: Site,
		}, Count + 1
	}

	site_len := len(sites)
	sitePaths := make([]sitePath, site_len*2)

	var prev_site toml.Site
	var next_site toml.Site
	j := 0
	for i, s := range sites {
		switch i {
		case 0:
			prev_site = sites[site_len-1]
			next_site = sites[i+1]
		case site_len - 1:
			prev_site = sites[i-1]
			next_site = sites[0]
		default:
			next_site = sites[i+1]
			prev_site = sites[i-1]
		}

		sitePaths[j], j = NewSitePath(fmt.Sprintf("/route/%s/next", s.Slug), next_site, j)
		sitePaths[j], j = NewSitePath(fmt.Sprintf("/route/%s/prev", s.Slug), prev_site, j)
	}

	return sitePaths
}
