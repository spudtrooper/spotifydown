package api

import (
	"encoding/json"
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type ProgressInfo struct {
	Status           string `json:"status"`
	DownloadProgress int    `json:"download_progress"`
	ConvertProgress  int    `json:"convert_progress"`
	Download         string `json:"download"`
}

//go:generate genopts --params --function Progress --extends Base taskID:string
func (c *Client) Progress(optss ...ProgressOption) (*ProgressInfo, error) {
	opts := MakeProgressOptions(optss...)

	const uri = "https://ytmp3api.net/iframe/progress.php"

	type body struct {
		TaskID string `json:"taskId"`
	}
	b, err := request.JSONMarshal(body{
		TaskID: opts.TaskID(),
	})
	if err != nil {
		return nil, err
	}

	resp, err := request.Post(uri, nil, strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}

	var res ProgressInfo
	if err := json.Unmarshal(resp.Data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
