package bytebase

import (
	"github.com/go-resty/resty/v2"
)

// https://www.bytebase.com/docs/api/sql-advise/ api documentation

type byteBase struct {
	Url        string
	debug      bool
	restClient *resty.Client
}

type IByteBaseClient interface {
	IsDebug() bool
	HealthCheck() error
}

func (o *byteBase) HealthCheck() error {
	_, err := o.get(o.Url, nil)
	if err != nil {
		return err
	}
	return nil
}

func (o *byteBase) IsDebug() bool {
	return o.debug
}

// Resty Methods

func (o *byteBase) postNoBody(url string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *byteBase) post(url string, body interface{}) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *byteBase) get(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}