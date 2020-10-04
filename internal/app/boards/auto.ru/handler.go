package auto_ru

import (
	"encoding/json"
	"fmt"
	"go-pugs/internal/app/boards/auto.ru/response"
	"go-pugs/internal/middleware"
	"go-pugs/internal/pkg/requests/boards"
	"go-pugs/internal/pkg/tools/generator"
	"go-pugs/internal/pkg/tools/httpBuilder"
	"go-pugs/internal/pkg/tools/validation"
	"go-pugs/internal/pkg/tools/wrapper"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func (api *API) GetAnnouncementInfo(ctx middleware.PugContext, w http.ResponseWriter, r *http.Request) {

}

func (api *API) SearchRequest(ctx middleware.PugContext, w http.ResponseWriter, r *http.Request) {
	sr := &boards.AutoRequest{}
	inter, err := validation.Parameters(r, sr)
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}
	sr = inter.(*boards.AutoRequest)

	reg := regexp.MustCompile("[\n\t]*")
	req := httpBuilder.NewRequest()
	req.Method = "GET"
	left, right := generator.RandomSHA256(), generator.RandomSHA256()
	req.Headers = map[string]string{
		"X-Features": reg.ReplaceAllString(`search_exclude_catalog_filter_trucks,sms_retriever,
			search_catalog_filter_trucks,search_with_delivery,chat_bot,search_exclude_catalog_filter_cars,
			search_catalog_filter_moto,search_grouping_id,phone_redirects_in_commercial_form,
			search_dynamic_catalog_equipment,search_no_premium_ads_in_new_listing,search_exclude_catalog_filter_moto,
			search_trucks_by_new_type,autoru_expert_landing,vin_resolution_with_history,search_special_trucks,
			chat_bot_checkup_button,search_mmng_multichoice,search_only_nds,search_catalog_filter_cars`, ""),
		"X-Device-Uid": fmt.Sprintf("%s.%s", left, right),
		"X-Session-Id": fmt.Sprintf("a:%s.%s|%s.%s.EFLH_2RAY_%s.%s-%s_%s", left, right,
			strconv.FormatInt(time.Now().Unix(), 10), strconv.FormatInt(int64(rand.Uint64()), 10),
			generator.RandomSHA256()[0:12], generator.RandomSHA256()[0:7], generator.RandomSHA256(),
			generator.RandomSHA256()[0:5]),
		"X-Vertis-Platform": "android/xxhdpi",
		"X-Authorization":   fmt.Sprintf("Vertis android-%s", generator.RandomSHA256()),
		"X-Client-Date":     strconv.FormatInt(time.Now().Unix(), 10),
		"X-Timezone-Name":   "Europe/Moscow",
		"Host":              "apiauto.ru",
		"Cookie":            "X-Vertis-DC=myt",
		"User-Agent":        "okhttp/3.12.1",
	}
	req.Query = map[string]string{
		"category":  "CARS",
		"page":      sr.Page,
		"sort":      "relevance-exp1-desc",
		"mark":      sr.Mark,
		"super_gen": strconv.FormatInt(int64(rand.Uint64()), 10),
		"page_size": "10",
	}

	if sr.Model != "" {
		req.Query["model"] = sr.Model
	}

	req.Url = "https://apiauto.ru/1.0/reviews/auto/listing"

	data, err := req.Do("")
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}
	ret := &response.Search{}
	if err := json.Unmarshal(data, ret); err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

}
