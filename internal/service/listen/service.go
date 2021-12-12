package listen

import (
	"context"
	"encoding/json"
	"net/url"
	"sync"
	"time"

	"github.com/lisabestteam/password-svc/internal/database"
	"github.com/spf13/cast"
)

func (l listen) Run(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()

	log := l.log.WithField("service", "listener")

	log.Info("run listener")

	ticker := time.NewTicker(5 * time.Second)

	maxId, err := l.passwords.MaxId()
	if err != nil {
		l.log.WithError(err).Error("failed to get max id")
		return
	}

	links := value{
		Type:   strprt("1"),
		Cursor: maxId,
	}.Encode()

	for ; ; <-ticker.C {
		list, err := l.client.GetDataList(links)
		if err != nil {
			log.WithError(err).Error("failed to get data")
			continue
		}

		if len(list.Data) == 0 {
			log.Debug("data is empty")
			continue
		}

		for _, data := range list.Data {
			var password database.Password
			if err = json.Unmarshal(data.Attributes.Value, &password); err != nil {
				log.WithError(err).Error("failed to unmarshal data value")
				continue
			}
			password.Id = cast.ToUint64(data.Id)

			if err = l.db.CreatePassword(password); err != nil {
				log.WithError(err).Error("failed to create password in database")
			}
		}

		links = list.Links.Next
	}
}

func strprt(string2 string) *string {
	return &string2
}

type value struct {
	Order  *string
	Limit  *int
	Number *int
	Type   *string
	Cursor *uint64
}

func (v value) Encode() string {
	values := url.Values{}

	if v.Limit != nil {
		values.Add("page[limit]", cast.ToString(v.Limit))
	}

	if v.Order != nil {
		values.Add("page[order]", cast.ToString(v.Order))
	}

	if v.Number != nil {
		values.Add("page[number]", cast.ToString(v.Number))
	}

	if v.Type != nil {
		values.Add("filter[type]", cast.ToString(v.Type))
	}

	if v.Cursor != nil {
		values.Add("page[cursor]", cast.ToString(v.Cursor))
	}

	return "/v3/data?" + values.Encode()
}
