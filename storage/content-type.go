package storage

import (
	"regexp"
)

type MediaTypeStruct struct {
	HTML  string
	CSS   string
	IMAGE string
	JS    *regexp.Regexp
	JSON  *regexp.Regexp
}

var MediaType MediaTypeStruct = MediaTypeStruct{
	HTML:  "text/html",
	CSS:   "text/css",
	IMAGE: "image/svg+xml",
	JS:    regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"),
	JSON:  regexp.MustCompile("[/+]json$"),
}
