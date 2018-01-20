package gocineworld

import "encoding/xml"

func ParseListings(listingsXml []byte) (Listings, error) {
	var listings Listings
	err := xml.Unmarshal(listingsXml, &listings)
	return listings, err
}
