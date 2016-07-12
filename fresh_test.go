package fresh

import (
	"net/http"
	"testing"
	"time"

	"github.com/RainInFall/assert"
)

func TestCheck(t *testing.T) {
	assert.Init(t)

	func() {
		var req http.Header
		var res http.Header

		assert.Ok(!Check(req, res))
	}()

	func() {
		var req http.Header = map[string][]string{
			http.CanonicalHeaderKey("if-none-match"): []string{"tobi"},
		}
		var res http.Header = map[string][]string{
			http.CanonicalHeaderKey("etag"): []string{"tobi"},
		}

		assert.Ok(Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "tobi")
		res.Set("etag", "tobi")

		assert.Ok(Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "tobi")
		res.Set("etag", "luna")

		assert.Ok(!Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "tobi")

		assert.Ok(!Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "W/\"foo\"")
		res.Set("etag", "W/\"foo\"")

		assert.Ok(Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "W/\"foo\"")
		res.Set("etag", "\"foo\"")

		assert.Ok(Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "\"foo\"")
		res.Set("etag", "\"foo\"")

		assert.Ok(Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "\"foo\"")
		res.Set("etag", "W/\"foo\"")

		assert.Ok(!Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "*")
		res.Set("etag", "hey")

		assert.Ok(Check(req, res))
	}()

	func() {
		now := time.Now()
		req := make(http.Header)
		res := make(http.Header)

		duration, err := time.ParseDuration("-4s")
		req.Set("if-modified-since", now.Add(duration).Format(http.TimeFormat))
		assert.Ok(err == nil)
		duration, err = time.ParseDuration("-2s")
		assert.Ok(err == nil)
		res.Set("last-modified", now.Add(duration).Format(http.TimeFormat))

		assert.Ok(!Check(req, res))
	}()

	func() {
		now := time.Now()
		req := make(http.Header)
		res := make(http.Header)

		duration, err := time.ParseDuration("-2s")
		req.Set("if-modified-since", now.Add(duration).Format(http.TimeFormat))
		assert.Ok(err == nil)
		duration, err = time.ParseDuration("-4s")
		assert.Ok(err == nil)
		res.Set("last-modified", now.Add(duration).Format(http.TimeFormat))

		assert.Ok(Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", time.Now().Format(http.TimeFormat))

		assert.Ok(!Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", "foo")

		assert.Ok(!Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("if-none-match", time.Now().Format(http.TimeFormat))
		req.Set("modified-since", "foo")

		assert.Ok(!Check(req, res))
	}()

	func() {
		now := time.Now()
		req := make(http.Header)
		res := make(http.Header)

		duration, err := time.ParseDuration("-2s")
		assert.Ok(err == nil)
		req.Set("if-modified-since", now.Add(duration).Format(http.TimeFormat))
		req.Set("if-none-match", "tobi")

		duration, err = time.ParseDuration("-4s")
		assert.Ok(err == nil)
		res.Set("last-modified", now.Add(duration).Format(http.TimeFormat))
		res.Set("etag", "tobi")

		assert.Ok(Check(req, res))
	}()

	func() {
		now := time.Now()
		req := make(http.Header)
		res := make(http.Header)

		duration, err := time.ParseDuration("-4s")
		assert.Ok(err == nil)
		req.Set("if-modified-since", now.Add(duration).Format(http.TimeFormat))
		req.Set("if-none-match", "tobi")

		duration, err = time.ParseDuration("-2s")
		assert.Ok(err == nil)
		res.Set("last-modified", now.Add(duration).Format(http.TimeFormat))
		res.Set("etag", "tobi")

		assert.Ok(!Check(req, res))
	}()

	func() {
		now := time.Now()
		req := make(http.Header)
		res := make(http.Header)

		duration, err := time.ParseDuration("-2s")
		assert.Ok(err == nil)
		req.Set("if-modified-since", now.Add(duration).Format(http.TimeFormat))
		req.Set("if-none-match", "tobi")

		duration, err = time.ParseDuration("-4s")
		assert.Ok(err == nil)
		res.Set("last-modified", now.Add(duration).Format(http.TimeFormat))
		res.Set("etag", "luna")

		assert.Ok(!Check(req, res))
	}()

	func() {
		now := time.Now()
		req := make(http.Header)
		res := make(http.Header)

		duration, err := time.ParseDuration("-4s")
		assert.Ok(err == nil)
		req.Set("if-modified-since", now.Add(duration).Format(http.TimeFormat))
		req.Set("if-none-match", "tobi")

		duration, err = time.ParseDuration("-2s")
		assert.Ok(err == nil)
		res.Set("last-modified", now.Add(duration).Format(http.TimeFormat))
		res.Set("etag", "tobi")

		assert.Ok(!Check(req, res))
	}()

	func() {
		req := make(http.Header)
		res := make(http.Header)

		req.Set("cache-control", "no-cache")

		assert.Ok(!Check(req, res))
	}()
}
