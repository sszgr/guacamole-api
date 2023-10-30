package guacamole

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Guacamole struct {
	BaseUrl  string
	Username string
	Password string

	SkipVerify bool

	dataSource string
	token      string
	timeout    time.Duration
	client     *http.Client
}

func (g *Guacamole) getHttpClient() *http.Client {
	if g.client != nil {
		return g.client
	}
	g.client = &http.Client{
		Timeout: g.timeout,
	}

	if g.SkipVerify {
		g.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	return g.client
}

func (g *Guacamole) Connect() error {
	g.token = ""
	return g.RefreshToken()
}

func (g *Guacamole) RefreshToken() error {

	values := url.Values{"token": {g.token}}
	if g.token == "" {
		values = url.Values{
			"username": {g.Username},
			"password": {g.Password},
		}
	}
	client := g.getHttpClient()
	resp, err := client.PostForm(g.BaseUrl+"/api/tokens", values)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var tokenResp ConnectResponse

	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return err
	}
	g.dataSource = tokenResp.DataSource
	g.token = tokenResp.AuthToken
	return nil
}

func (g *Guacamole) Call(method, uri string, xq map[string]string, body interface{}) ([]byte, error) {
	err := g.RefreshToken()
	if err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest(method, g.BaseUrl+uri, bytes.NewReader(reqBody))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// URL query params
	q := req.URL.Query()
	q.Add("token", g.token)
	if len(xq) > 0 {
		for k, v := range xq {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	client := g.getHttpClient()
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	if resp.Body == nil {
		return nil, errors.New(ErrBodyIsNil)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
