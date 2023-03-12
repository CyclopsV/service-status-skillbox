package storages

import (
	"encoding/json"
	"fmt"
	"github.com/CyclopsV/service-status-skillbox/internal/mms"
	"io"
	"net/http"
	"sort"
)

type MMSStorage []*mms.MMS

func (ms *MMSStorage) Add(obj *mms.MMS) {
	*ms = append(*ms, obj)
}

func NewMMSStorage() (*MMSStorage, error) {
	url := "http://127.0.0.1:8383/mms"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка получения данных MMS:\n%v", err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var buf []map[string]interface{}
	if err = json.Unmarshal(content, &buf); err != nil {
		return nil, fmt.Errorf("ошибка чтения данных MMS:\n%v", err)
	}
	ms := MMSStorage{}
	for _, el := range buf {
		m := mms.New(el["country"].(string), el["provider"].(string), el["bandwidth"].(string), el["response_time"].(string))
		if m != nil {
			ms = append(ms, m)
		}
	}

	return &ms, nil
}

func (ms MMSStorage) SortCountry() {
	sortF := func(i, j int) bool {
		return ms[i].Country < ms[j].Country
	}
	sort.SliceStable(ms, sortF)
}

func (ms MMSStorage) SortProvider() {
	sortF := func(i, j int) bool {
		return ms[i].Provider < ms[j].Provider
	}
	sort.SliceStable(ms, sortF)
}
