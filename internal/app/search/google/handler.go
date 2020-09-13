package google

import (
	"go-pugs/internal/models"
	"go-pugs/internal/tools/httpBuilder"
	"go-pugs/internal/tools/wrapper"
	"net/http"
	"regexp"
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
		"q":     "site:www.spbstu.ru filetype:pdf",
	}

	data, err := req.Do("")
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}
	reg := regexp.MustCompile(`href="/url\?q=https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)\.pdf`)
	rawUrls := reg.FindAllString(string(data), -1)
	for i := range rawUrls {
		rawUrls[i] = rawUrls[i][len(`href="/url\?q=`)-1:]
		if err := api.fileService.Insert(&models.File{
			Type: "pdf",
			Url:  rawUrls[i],
		}); err != nil {
			wrapper.ErrorResponse(w, err)
			return
		}
	}

	wrapper.Response(w, rawUrls)
}
