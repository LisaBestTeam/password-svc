package horizon

import (
	"encoding/json"
)

func (h horizon) GetDataList(values string) (DataList, error) {
	var result DataList
	err := h.get(&result, values)
	return result, err
}

type DataList struct {
	Data []struct {
		Id            string `json:"id"`
		Type          string `json:"type"`
		Relationships struct {
			Owner struct {
				Data struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"owner"`
		} `json:"relationships"`
		Attributes struct {
			Type  int             `json:"type"`
			Value json.RawMessage `json:"value"`
		} `json:"attributes"`
	} `json:"data"`
	Included []struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"included"`
	Links struct {
		Next string `json:"next"`
		Self string `json:"self"`
	} `json:"links"`
}
