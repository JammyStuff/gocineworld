package gocineworld

type Cinema struct {
	Name  string `xml:"name,attr"`
	ID    int    `xml:"id,attr"`
	Films []Film `xml:"listing>film"`
}

type Film struct {
	Title string `xml:"title,attr"`
	Shows []Show `xml:"shows>show"`
}

type Listings struct {
	Cinemas []Cinema `xml:"cinema"`
}

type Show struct {
	Time string `xml:"time,attr"`
	URL  string `xml:"url,attr"`
}
