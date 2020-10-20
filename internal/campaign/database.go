package campaign

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"net/http"
	"time"
)

func NewDBConnection(conf *Settings) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("%v://%v:%v", conf.DatabaseProto, conf.DatabaseHost, conf.DatabasePort)),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(time.Duration(conf.Timeout)*time.Second),
		elastic.SetHttpClient(&http.Client{Timeout: time.Second * time.Duration(conf.DatabaseTimeout)}))
	if err != nil {
		return nil, err
	}

	return client, nil
}
