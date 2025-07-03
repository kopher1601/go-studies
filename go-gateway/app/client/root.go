package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/kopher1601/go-studies/go-gateway/common"
	"github.com/kopher1601/go-studies/go-gateway/config"
	"github.com/kopher1601/go-studies/go-gateway/kafka"
)

const (
	_defaultBatchTime = 2
)

type HttpClient struct {
	client *resty.Client
	cfg    config.App

	producer kafka.Producer
}

func NewHttpClient(
	cfg config.App,
	producer map[string]kafka.Producer,
) HttpClient {
	batchTime := cfg.Producer.BatchTime

	if batchTime == 0 {
		batchTime = _defaultBatchTime
	}

	client := resty.New().
		SetJSONMarshaler(common.JsonHandler.Marshal).
		SetJSONUnmarshaler(common.JsonHandler.Unmarshal)
}
