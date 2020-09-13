package google

import (
	"go-pugs/internal/tools/httpBuilder"
	"go-pugs/internal/tools/wrapper"
	"net/http"
)

func (api *API) searchRequest(w http.ResponseWriter, r *http.Request) {
	req := httpBuilder.NewRequest()
	req.Method = "GET"
	req.Headers = map[string]string{
		"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36`,
	}
	req.Url = "http://www.google.com/search"
	req.Query = map[string]string{
		"num":   "100",
		"start": "0",
		"hl":    "en",
		"meta":  "",
		"q":     "@\\spbstu\\",
	}
	if data, err := req.Do("1"); err != nil {
		wrapper.ErrorResponse(w, err)
		return
	} else {
		wrapper.Response(w, data)
	}
}
