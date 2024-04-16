package core

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bububa/biz-weibo/core/internal/debug"
	"github.com/bububa/biz-weibo/model"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient  object
type SDKClient struct {
	appID      string
	debug      bool
	httpClient *http.Client
}

// NewSDKClient init sdk client
func NewSDKClient(appID string) *SDKClient {
	return &SDKClient{
		appID:      appID,
		httpClient: defaultHttpClient(),
	}
}

// AppID get appID
func (c SDKClient) AppID() string {
	return c.appID
}

// SetDebug set debug mode
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

func (c SDKClient) SetHTTPClient(clt *http.Client) {
	c.httpClient = clt
}

func (c SDKClient) RequestURL(apiPath string, payload *model.Payload) string {
	if strings.HasPrefix(apiPath, "/") {
		apiPath = apiPath[1:]
	}
	ret := fmt.Sprintf("%s/%s", Gateway, apiPath)
	if payload == nil || payload.IsPost() || len(payload.Values()) == 0 {
		return ret
	}
	return fmt.Sprintf("%s?%s", ret, payload.Values().Encode())
}

// Do execute api request
func (c *SDKClient) Do(accessToken string, req model.Request, resp model.Response) error {
	var (
		payload = req.Payload()
		reqURL  = c.RequestURL(req.URL(), payload)
		err     error
	)
	if payload != nil && payload.IsPost() {
		err = c.Post(accessToken, reqURL, payload.Data(), &resp)
	} else {
		err = c.Get(accessToken, reqURL, &resp)
	}
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

// Post data through api
func (c *SDKClient) Post(accessToken string, reqUrl string, reqBytes []byte, resp interface{}) error {
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}
	httpReq.Header.Add("Accept", "application/json,application/text+gw2.0")
	httpReq.Header.Add("Content-Type", "application/json")
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = new(model.BaseResponse)
	}
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return nil
}

// Get data through api
func (c *SDKClient) Get(accessToken string, reqUrl string, resp interface{}) error {
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}
	httpReq.Header.Add("Accept", "application/json,application/text+gw2.0")
	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if resp == nil {
		resp = new(model.BaseResponse)
	}
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	return nil
}
