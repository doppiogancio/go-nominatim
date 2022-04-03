package geocodejson

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/doppiogancio/go-nominatim/shared"
	"io/ioutil"
	"net/http"
)

type (
	Client struct {
		baseUrl string
	}
)

const format = "geocodejson"

func New(baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
	}
}

func (c *Client) Search(req shared.SearchRequest) (*FeatureCollection, error) {
	url := fmt.Sprintf("%s/search.php?q=%s&format=%s&accept-language=%s",
		c.baseUrl,
		req.Q,
		format,
		req.AcceptLanguage,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var collection *FeatureCollection

	err = json.Unmarshal(body, &collection)
	if err != nil {
		return nil, err
	}

	if len(collection.Features) == 0 {
		return nil, errors.New("address not found")
	}

	return collection, nil
}

func (c *Client) Reverse(req shared.ReverseGeocodeRequest) (*FeatureCollection, error) {
	url := fmt.Sprintf(
		"%s/reverse?format=%s&lat=%f&lon=%f&accept-language=%s",
		c.baseUrl,
		format,
		req.Latitude,
		req.Longitude,
		req.AcceptLanguage,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var errResponse struct {
		Error string
	}

	err = json.Unmarshal(body, &errResponse)
	if err != nil {
		return nil, err
	}

	if errResponse.Error != "" {
		return nil, errors.New(errResponse.Error)
	}

	var collection *FeatureCollection

	err = json.Unmarshal(body, &collection)
	if err != nil {
		return nil, err
	}

	return collection, nil
}
