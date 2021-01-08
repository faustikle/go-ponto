package pontotel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/faustikle/go-ponto/internal/fileconfig"
	"github.com/faustikle/go-ponto/internal/ponto"
)

const host = "https://back.pontotel.com.br/web"

type Client struct {
	CompMan      string
	FingerPrint  string
	Password     string
	sessionToken string
	httpClient   http.Client
}

type ClientConfig struct {
	CompMan     string
	FingerPrint string
	Password    string
}

func NewClient(cfg ClientConfig) *Client {
	_ = fileconfig.Config{}
	return &Client{
		CompMan:     cfg.CompMan,
		FingerPrint: cfg.FingerPrint,
		Password:    cfg.Password,
		httpClient: http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) Login() (LoginResponse, error) {
	loginReq := LoginRequest{
		CompMan:     c.CompMan,
		Fingerprint: c.FingerPrint,
		Pwd:         c.Password,
	}

	body, err := json.Marshal(loginReq)
	if err != nil {
		return LoginResponse{}, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/checkpwd", host), bytes.NewReader(body))
	if err != nil {
		return LoginResponse{}, err
	}

	request.Header = setHeaders(request.Header)

	res, err := c.httpClient.Do(request)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("o sistema de ponto está sobrecarregado, tente novamente mais tarde")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return LoginResponse{}, err
	}

	err = c.validateResponse(res, resBody)
	if err != nil {
		return LoginResponse{}, err
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(resBody, &loginResponse)
	if err != nil {
		return LoginResponse{}, err
	}

	c.sessionToken = loginResponse.Success.SessionToken

	return loginResponse, nil
}

func (c *Client) Register(p ponto.Entity, kind, lat, lon string) (RegisterResponse, error) {
	fmt.Println("Batendo ponto...")
	if kind == "" {
		kind = p.NextKind()
	}

	if kind == "" {
		return RegisterResponse{}, fmt.Errorf("todos os pontos já foram batidos")
	}

	formData := url.Values{
		"employee":     {p.Employee},
		"sessionToken": {c.sessionToken},
		"fingerprint":  {c.FingerPrint},
		"kind":         {kind},
		"lat":          {lat},
		"lon":          {lon},
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/savetimelog", host), strings.NewReader(formData.Encode()))
	if err != nil {
		return RegisterResponse{}, err
	}

	request.Header = setHeaders(request.Header)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(request)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("o sistema de ponto está sobrecarregado, tente novamente mais tarde")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return RegisterResponse{}, err
	}

	err = c.validateResponse(res, resBody)
	if err != nil {
		return RegisterResponse{}, err
	}

	var r RegisterResponse
	err = json.Unmarshal(resBody, &r)
	if err != nil {
		return RegisterResponse{}, err
	}

	return r, nil
}

func (c *Client) Logged() bool {
	if c.sessionToken == "" {
		return false
	}

	return true
}

func (c *Client) validateResponse(res *http.Response, resBody []byte) error {
	if res.StatusCode >= 400 {
		return fmt.Errorf("response error status %d with body %s", res.StatusCode, string(resBody))
	}

	return nil
}

func setHeaders(h http.Header) http.Header {
	newHeaders := http.Header{
		"User-Agent":   []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36"},
		"Content-Type": []string{"application/json"},
		"Accept":       []string{"*/*"},
	}

	for k, v := range newHeaders {
		for _, vv := range v {
			h.Add(k, vv)
		}
	}

	return h
}
