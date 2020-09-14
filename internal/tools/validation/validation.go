package validation

import (
	"encoding/json"
	"go-pugs/internal/tools/validation/const"
	"net/http"
	"reflect"
)

type ValidatingStruct interface {
	Struct() interface{}
	Validate() error
}

func Parameters(r *http.Request, v ValidatingStruct) (interface{}, error) {
	inter := v.Struct()
	tmp := map[string]interface{}{}
	ref := reflect.TypeOf(inter)
	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		val, in := field.Tag.Lookup(_const.From)
		if !in {
			val = _const.Query
		}
		if val == _const.Query {
			name, in := field.Tag.Lookup(_const.Json)
			if !in {
				name = field.Name
			}
			tmp[name] = r.URL.Query().Get(name)
		} //Add here new parameters
	}

	data, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, inter); err != nil {
		return nil, err
	}
	if err := v.Validate(); err != nil {
		return nil, err
	}
	return inter, nil
}
