package manager

import (
	"go-pugs/internal/models"
	"go-pugs/internal/tools/httpBuilder"
	"go-pugs/internal/tools/wrapper"
	"net/http"
)

func (api *API) getFile(w http.ResponseWriter, r *http.Request) {
	ret := &models.File{ID: 115}
	if err := api.fileService.Select(ret); err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	req := httpBuilder.NewRequest()
	req.Method = "GET"
	req.Url = ret.Url

	if data, err := req.Do(""); err != nil {
		wrapper.ErrorResponse(w, err)
		return
	} else {
		wrapper.Response(w, data)
	}
}
