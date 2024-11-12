package purpur

import (
	"iter"
	"sort"
)

type GetMinecraftVersionsResponse struct {
	Project  string                       `json:"project"`
	Metadata GetMinecraftVersionsMetadata `json:"metadata"`
	Versions []string                     `json:"versions"`
}

type GetMinecraftVersionsMetadata struct {
	Current string `json:"current"`
}

func (r *GetMinecraftVersionsResponse) SortVersions() {
	sort.Slice(r.Versions[:], func(i, j int) bool {
		return r.Versions[i] > r.Versions[j]
	})
}

func (r *GetMinecraftVersionsResponse) AllVersions() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, version := range r.Versions {
			if !yield(version) {
				return
			}
		}
	}
}
