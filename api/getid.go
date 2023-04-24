package api

import (
	"fmt"
)

type GetIDInfo struct {
	ID      string `json:"id"`
	Method  string `json:"method"`
	Success bool   `json:"success"`
}

//go:generate genopts --params --function GetID --extends Base track:string
func (c *Client) GetID(optss ...GetIDOption) (*GetIDInfo, error) {
	opts := MakeGetIDOptions(optss...)

	path := fmt.Sprintf("getId/%s", opts.Track())

	var payload GetIDInfo
	if err := c.get(path, &payload); err != nil {
		return nil, err
	}
	return &payload, nil

}
