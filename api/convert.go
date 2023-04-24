package api

import (
	"log"
	"time"
)

type ConvertInfo struct {
	Download string
}

//go:generate genopts --params --function Convert --extends Base track:string verbose
func (c *Client) Convert(optss ...ConvertOption) (*ConvertInfo, error) {
	opts := MakeConvertOptions(optss...)
	verbose := true || opts.Verbose()

	id, err := c.GetID(GetIDTrack(opts.Track()))
	if err != nil {
		return nil, err
	}
	if verbose {
		log.Printf("id: %+v", id)
	}

	d, err := c.Download(DownloadId(id.ID))
	if err != nil {
		return nil, err
	}
	if verbose {
		log.Printf("download: %+v", d)
	}

	for {
		p, err := c.Progress(ProgressTaskID(d.TaskID))
		if err != nil {
			return nil, err
		}
		if verbose {
			log.Printf("progress: %+v", p)
		}
		if p.Status == "finished" {
			return &ConvertInfo{
				Download: p.Download,
			}, nil
		}

		time.Sleep(1 * time.Second)
	}

	panic("unreachable")
}
