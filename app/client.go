package app

import (
	"context"
	"github.com/memochou1993/youtube-downloader/app/model"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	host = "https://youtube.com"
)

type Client struct {
	HTTPClient *http.Client
}

func (c *Client) New() *http.Client {
	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

	return c.HTTPClient
}

func (c *Client) Get(ctx context.Context, url string) (*http.Response, error) {
	client := c.New()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func (c *Client) GetBody(ctx context.Context, url string) []byte {
	resp, err := c.Get(ctx, url)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return body
}

func (c *Client) GetVideo(ctx context.Context, id string) *model.Video {
	body := c.GetBody(ctx, host+"/get_video_info?video_id="+id)

	video := &model.Video{}
	video.ParseVideoInfo(string(body))

	return video
}
