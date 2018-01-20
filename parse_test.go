package gocineworld

import (
	"io/ioutil"
	"testing"
)

func TestParseListings(t *testing.T) {
	listings := parseTestListingsXml(t)
	t.Run("Number of cinemas", func(t *testing.T) { testParseListingsNumCinemas(t, &listings) })
	t.Run("Cinema name", func(t *testing.T) { testParseListingsCinemaRawName(t, &listings) })
	t.Run("Cinema ID", func(t *testing.T) { testParseListingsCinemaID(t, &listings) })
	t.Run("Number of films", func(t *testing.T) { testParseListingsNumFilms(t, &listings) })
	t.Run("Film title", func(t *testing.T) { testParseListingsFilmTitle(t, &listings) })
	t.Run("Show time", func(t *testing.T) { testParseListingsShowRawTime(t, &listings) })
	t.Run("Show URL", func(t *testing.T) { testParseListingsShowURL(t, &listings) })
}

func testParseListingsCinemaID(t *testing.T, listings *Listings) {
	cinema := listings.Cinemas[0]
	if cinema.ID != 1 {
		t.Fatal("Cinema ID was not correct")
	}
}

func testParseListingsCinemaRawName(t *testing.T, listings *Listings) {
	cinema := listings.Cinemas[0]
	if cinema.RawName != "Cineworld Aberdeen - Queens Links" {
		t.Fatal("Cinema name was not correct")
	}
}

func testParseListingsFilmTitle(t *testing.T, listings *Listings) {
	cinema := listings.Cinemas[0]
	film := cinema.Films[0]
	if film.Title != "All The Money In The World" {
		t.Fatal("Film title was not correct")
	}
}

func testParseListingsNumCinemas(t *testing.T, listings *Listings) {
	numCinemas := len(listings.Cinemas)
	if numCinemas != 97 {
		t.Fatal("Expected 97 cinemas, got %d", numCinemas)
	}
}

func testParseListingsNumFilms(t *testing.T, listings *Listings) {
	cinema := listings.Cinemas[0]
	numFilms := len(cinema.Films)
	if numFilms != 25 {
		t.Fatalf("Expected 25 films, got %d", numFilms)
	}
}

func testParseListingsShowRawTime(t *testing.T, listings *Listings) {
	cinema := listings.Cinemas[0]
	film := cinema.Films[0]
	show := film.Shows[0]
	if show.RawTime != "2018-01-18T11:15:00" {
		t.Fatal("Show time was not correct")
	}
}

func testParseListingsShowURL(t *testing.T, listings *Listings) {
	cinema := listings.Cinemas[0]
	film := cinema.Films[0]
	show := film.Shows[0]
	if show.URL != "https://booking.cineworld.co.uk/booking/8014/95482" {
		t.Fatal("Show URL was not correct")
	}
}

func parseTestListingsXml(t *testing.T) Listings {
	listingsXml, err := readTestListingsXml()
	if err != nil {
		t.Fatalf("Error reading test listings XML: %v", err)
	}

	listings, err := ParseListings(listingsXml)
	if err != nil {
		t.Fatalf("Error parsing listings: %v", err)
	}

	return listings
}

func readTestListingsXml() ([]byte, error) {
	return ioutil.ReadFile("resources/listings.xml")
}
