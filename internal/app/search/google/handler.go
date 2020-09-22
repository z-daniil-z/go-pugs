package google

import (
	"fmt"
	"go-pugs/internal/middleware"
	"go-pugs/internal/models"
	"go-pugs/internal/pkg/errors"
	"go-pugs/internal/pkg/tools/httpBuilder"
	"go-pugs/internal/pkg/tools/validation"
	"go-pugs/internal/pkg/tools/wrapper"
	val "go-pugs/internal/pkg/validation"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func (api *API) SearchRequest(ctx middleware.PugContext, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context().Value("costum_key"))
	sr := &val.Request{}
	inter, err := validation.Parameters(r, sr)
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}
	sr = inter.(*val.Request)

	req := httpBuilder.NewRequest()
	req.Method = "GET"
	reg := regexp.MustCompile("[\n\t]*")
	req.Headers = map[string]string{
		"User-Agent": reg.ReplaceAllString(`Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0`, ""),
	}
	req.Url = "http://www.google.com/search"

	ret := make([]string, 0)

	total, err := strconv.Atoi(sr.Count)
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	for i := 0; i < total; i += 100 {
		req.Query = map[string]string{
			"num":   "100",
			"start": strconv.FormatInt(int64(i), 10),
			"hl":    "en",
			"meta":  "",
			"q":     fmt.Sprintf("site:%s filetype:%s", sr.WebSite, sr.DocType),
		}
		data, err := req.Do("")
		if err != nil {
			wrapper.ErrorResponse(w, err)
			return
		}
		if err := api.checkBlock(data); err != nil {
			wrapper.ErrorResponse(w, err)
			return
		}
		tmp, err := api.checkAdditionalLink(data)
		if err != nil {
			wrapper.ErrorResponse(w, err)
			return
		}
		if tmp != nil {
			data = tmp
		}
		urls, err := api.findGoodUrls(data, sr.DocType)
		if err != nil {
			wrapper.ErrorResponse(w, err)
			return
		}
		ret = append(ret, urls...)
	}
	wrapper.Response(w, ret)
}

func (api API) checkBlock(data []byte) error {
	reg := regexp.MustCompile("Our systems have detected unusual traffic from your computer network")
	block := reg.Match(data)
	if block {
		return errors.ErrBlock
	}
	return nil
}

func (api *API) checkAdditionalLink(data []byte) ([]byte, error) {
	reg := regexp.MustCompile("If you like, you can.*repeat the search with the omitted results included")
	rawUrl := reg.FindAllString(string(data), -1)
	if len(rawUrl) == 0 {
		return nil, nil
	}
	split := strings.Split(rawUrl[0], `"`)
	if len(split) != 9 {
		return nil, errors.ErrWrongParse
	}
	path := split[3][len(`/url?q=`):]
	req := httpBuilder.NewRequest()
	req.Method = "GET"
	reg = regexp.MustCompile("[\n\t]*")
	req.Headers = map[string]string{
		"User-Agent": reg.ReplaceAllString(`Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0`, ""),
	}
	path, err := url.QueryUnescape(path)
	if err != nil {
		return nil, err
	}
	req.Url = fmt.Sprintf("http://www.google.com%s", path)
	data, err = req.Do("")
	if err != nil {
		return nil, err
	}
	if err := api.checkBlock(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (api *API) findGoodUrls(data []byte, docType string) ([]string, error) {
	reg := regexp.MustCompile("[\n\t]*")
	reg = regexp.MustCompile(reg.ReplaceAllString(`href="/url\?q=https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}
		\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)\.`+docType, ""))
	rawUrls := reg.FindAllString(string(data), -1)
	for i := range rawUrls {
		rawUrls[i] = rawUrls[i][len(`href="/url\?q=`)-1:]
		if err := api.fileService.InsertOrUpdate(&models.File{
			Type: docType,
			Url:  rawUrls[i],
		}); err != nil {
			return nil, err
		}
	}
	return rawUrls, nil
}
