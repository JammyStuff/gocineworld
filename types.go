package gocineworld

import (
	"strings"
	"time"
)

const showTimeFormat = "2006-01-02T15:04:05"

type Cinema struct {
	RawName string `xml:"name,attr"`
	ID      int    `xml:"id,attr"`
	Films   []Film `xml:"listing>film"`
}

func (c *Cinema) Name() string {
	return strings.TrimPrefix(c.RawName, "Cineworld ")
}

type Film struct {
	Title string `xml:"title,attr"`
	Shows []Show `xml:"shows>show"`
}

type Listings struct {
	Cinemas []Cinema `xml:"cinema"`
}

type Show struct {
	RawTime string `xml:"time,attr"`
	URL     string `xml:"url,attr"`
}

func (s *Show) Time() (time.Time, error) {
	return time.Parse(showTimeFormat, s.RawTime)
}
