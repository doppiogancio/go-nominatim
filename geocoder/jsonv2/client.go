package jsonv2

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

const format = "jsonv2"

func New(baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
	}
}

func (c *Client) Search(req shared.SearchRequest) ([]*Place, error) {
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

	var list []*Place

	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, errors.New("address not found")
	}

	return list, nil
}
