package api

import (
	"time"
)

type ConvertInfo struct {
	Download string
}

//go:generate genopts --params --function Convert --extends Base track:string
func (c *Client) Convert(optss ...ConvertOption) (*ConvertInfo, error) {
	opts := MakeConvertOptions(optss...)
	verbose := opts.Verbose()
	track := opts.Track()

	if verbose {
		c.logger.Printf("requesting id for track: %s", opts.Track())
	}

	id, err := c.GetID(GetIDTrack(track))
	if err != nil {
		return nil, err
	}
	if verbose {
		c.logger.Printf("have id: %+v", id)
	}

	d, err := c.Download(DownloadId(id.ID), DownloadVerbose(verbose))
	if err != nil {
		return nil, err
	}
	if verbose {
		c.logger.Printf("starting download for %s", d.TaskID)
	}

	for {
		p, err := c.Progress(ProgressTaskID(d.TaskID))
		if err != nil {
			return nil, err
		}
		if verbose {
			c.logger.Printf("progress: %+v", p)
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
