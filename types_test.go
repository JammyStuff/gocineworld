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
