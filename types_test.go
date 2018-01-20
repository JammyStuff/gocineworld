package gocineworld

import "testing"

func TestCinemaName(t *testing.T) {
	cinema := Cinema{
		RawName: "Cineworld Aberdeen - Queens Links",
		ID:      1,
	}
	name := cinema.Name()
	if name != "Aberdeen - Queens Links" {
		t.Fatalf("Expected cinema name to be Aberdeen - Queens Links, got %s", name)
	}
}

func TestShowTime(t *testing.T) {
	rawTime := "2018-01-18T11:15:00"
	show := Show{
		RawTime: rawTime,
		URL:     "https://booking.cineworld.co.uk/booking/8014/95482",
	}
	time, err := show.Time()
	if err != nil {
		t.Fatalf("Error parsing show time: %v", err)
	}
	formattedTime := time.Format(showTimeFormat)
	if formattedTime != rawTime {
		t.Fatalf("Expected show time to be %s, got %s", rawTime, formattedTime)
	}
}
