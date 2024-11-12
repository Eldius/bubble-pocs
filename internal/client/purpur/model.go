package purpur

type GetMinecraftVersionsResponse struct {
	Project  string                       `json:"project"`
	Metadata GetMinecraftVersionsMetadata `json:"metadata"`
	Versions []string                     `json:"versions"`
}

type GetMinecraftVersionsMetadata struct {
	Current string `json:"current"`
}
