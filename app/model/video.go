package model

import (
	"encoding/json"
	"log"
	"net/url"
)

type Video struct {
	VideoDetails struct {
		VideoID          string `json:"videoId"`
		Title            string `json:"title"`
		ShortDescription string `json:"shortDescription"`
		Author           string `json:"author"`
	} `json:"videoDetails"`
	StreamingData struct {
		Formats []Format `json:"formats"`
	} `json:"streamingData"`
}

type Format struct {
	URL              string `json:"url"`
	MimeType         string `json:"mimeType"`
	Bitrate          int    `json:"bitrate"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	LastModified     string `json:"lastModified"`
	ContentLength    string `json:"contentLength"`
	Quality          string `json:"quality"`
	Fps              int    `json:"fps"`
	QualityLabel     string `json:"qualityLabel"`
	ProjectionType   string `json:"projectionType"`
	AverageBitrate   int    `json:"averageBitrate"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
}

func (v *Video) ParseVideoInfo(info string) {
	data, err := url.ParseQuery(info)

	if err != nil {
		log.Println(err.Error())
		return
	}

	playerResponse := data.Get("player_response")

	if err := json.Unmarshal([]byte(playerResponse), v); err != nil {
		log.Println(err.Error())
	}
}
