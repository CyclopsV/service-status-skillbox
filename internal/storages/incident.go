package storages

import (
	"fmt"
	"github.com/CyclopsV/service-status-skillbox/internal/incident"
	"github.com/CyclopsV/service-status-skillbox/pkg/apiRequest"
	"github.com/CyclopsV/service-status-skillbox/pkg/pars"
)

type IncidentStorage []*incident.Incident

func NewIncidentStorage() (*IncidentStorage, error) {
	resp := apiRequest.Get("http://127.0.0.1:8383/accendent")
	is := IncidentStorage{}
	if resp == nil {
		return &is, fmt.Errorf("ошибка получения данных")
	}
	if err := pars.JSON(&is, resp.Body); err != nil {
		return &is, err
	}
	deleteErrData(is)
	return &is, nil
}
