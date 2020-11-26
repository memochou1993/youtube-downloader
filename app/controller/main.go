package controller

import (
	"context"
	"fmt"
	"github.com/memochou1993/youtube-downloader/app"
	"github.com/memochou1993/youtube-downloader/app/model"
	"log"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := &app.Client{}
	id := r.URL.Query().Get("id")

	if id == "" {
		return
	}

	video := client.GetVideo(ctx, id)

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	formats := video.StreamingData.Formats

	if len(formats) == 0 {
		return
	}

	url := findBestFormat(video.StreamingData.Formats).URL

	if url == "" {
		return
	}

	content := client.GetBody(ctx, url)

	download(w, video.VideoDetails.Title, content)
}

func findBestFormat(formats []model.Format) model.Format {
	index := 0
	size := 0

	for i, format := range formats {
		s := format.Height * format.Width

		if s > size {
			index = i
			size = s
		}
	}

	return formats[index]
}

func download(w http.ResponseWriter, filename string, data []byte) {
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.mp4\"", filename))

	if _, err := w.Write(data); err != nil {
		log.Println(err.Error())
	}
}
