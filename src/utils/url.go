package utils

import (
	"fyne.io/fyne/v2"
	"net/url"
)

func ParseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	return link
}
