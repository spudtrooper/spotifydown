package api

import (
	"fmt"
)

type MetadataInfo struct {
	Album   string `json:"album"`
	Artists string `json:"artists"`
	Cache   bool   `json:"cache"`
	Cover   string `json:"cover"`
	ID      string `json:"id"`
	Isrc    string `json:"isrc"`
	Success bool   `json:"success"`
	Title   string `json:"title"`
}

//go:generate genopts --params --function Metadata --extends Base track:string
func (c *Client) Metadata(optss ...MetadataOption) (*MetadataInfo, error) {
	opts := MakeMetadataOptions(optss...)

	path := fmt.Sprintf("metadata/track/%s", opts.Track())

	var payload MetadataInfo
	if err := c.get(path, &payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
