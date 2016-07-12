package fresh

import (
	"net/http"
	"strings"
	"time"
)

//go:generate gotemplate github.com/RainInFall/js-like/array ArrayString(string)

func Check(req, res http.Header) bool {
	// defaults
	etagMatches := true
	notModified := true
	// fields
	modifiedSince := req.Get(http.CanonicalHeaderKey("if-modified-since"))
	noneMatch := req[http.CanonicalHeaderKey("if-none-match")]
	lastModified := res.Get(http.CanonicalHeaderKey("last-modified"))
	etag := res.Get(http.CanonicalHeaderKey("etag"))
	cc := req.Get(http.CanonicalHeaderKey("cache-control"))

	// unconditional request
	if 0 == len(modifiedSince) && 0 == len(noneMatch) {
		return false
	}

	// check for non-cache request directives(node) {
	if strings.Contains(cc, "no-cache") {
		return false
	}

	// parse if-none-match
	if len(noneMatch) > 0 {
		etagMatches = ArrayString(noneMatch).Some(
			func(match string, _index int, _array ArrayString) bool {
				return match == "*" || match == etag || match == "W/"+etag
			})
	}

	// if-modified-since
	if len(modifiedSince) > 0 {
		modifiedSince, err1 := time.Parse(http.TimeFormat, modifiedSince)
		lastModified, err2 := time.Parse(http.TimeFormat, lastModified)
		notModified = nil == err1 && nil == err2 && !lastModified.After(modifiedSince)
	}
	return etagMatches && notModified
}
