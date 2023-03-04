package storages

import (
	"github.com/CyclopsV/service-status-skillbox/internal/email"
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
)

type EmailStorage []*email.Email

func (es *EmailStorage) Add(obj *email.Email) {
	*es = append(*es, obj)
}

func NewEmailStorage(filename string) *EmailStorage {
	emailStr := pars.FileToStr(filename)
	es := EmailStorage{}
	for _, s := range emailStr {
		res := email.FromSTR(s)
		if res == nil {
			continue
		}
		es.Add(res)
	}
	return &es
}