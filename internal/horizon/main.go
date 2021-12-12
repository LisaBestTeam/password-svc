package horizon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HorizonClient interface {
	GetDataList(values string) (DataList, error)
}

func NewHorizonClient(endpoint string) HorizonClient {
	return &horizon{
		endpoint: endpoint,
	}
}

type horizon struct {
	endpoint string
}

func (h *horizon) get(dest interface{}, values string) error {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", h.endpoint, values), nil)
	if err != nil {
		return err
	}

	do, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if err = json.NewDecoder(do.Body).Decode(dest); err != nil {
		return err
	}

	return nil
}
