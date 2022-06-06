package omdb

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type OMDBImpl struct {
	APIKey string
	Domain string
}

func NewOMDBService(domain, apiKey string) *OMDBImpl {
	return &OMDBImpl{
		APIKey: apiKey,
		Domain: domain,
	}
}

func (s *OMDBImpl) FindAll(p Params) (*FindAllResponse, error) {

	v, _ := query.Values(p)
	url := s.Domain + "/?apikey=" + s.APIKey + "&" + v.Encode()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var resp FindAllResponse
	e := json.NewDecoder(r.Body).Decode(&resp)
	if e != nil {
		return nil, e
	}

	return &resp, nil
}

func (s *OMDBImpl) Detail(ID string) (*DetailResponse, error) {

	url := s.Domain + "/?apikey=" + s.APIKey + "&i=" + ID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var resp DetailResponse
	e := json.NewDecoder(r.Body).Decode(&resp)
	if e != nil {
		return nil, e
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}
	return &resp, nil
}
