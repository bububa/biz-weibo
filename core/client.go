package core

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/bububa/biz-weibo/core/internal/debug"
	"github.com/bububa/biz-weibo/model"
)

// SDKClient  object
type SDKClient struct {
	appID string
	debug bool
}

// NewSDKClient init sdk client
func NewSDKClient(appID string) *SDKClient {
	return &SDKClient{
		appID: appID,
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
	debug.PrintPostJSONRequest(reqUrl, reqBytes, c.debug)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	if accessToken != "" {
		httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}
	httpReq.Header.Add("Accept", "application/json,application/text+gw2.0")
	httpReq.Header.Add("Content-Type", "application/json")
	httpResp, err := http.DefaultClient.Do(httpReq)
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
	httpResp, err := http.DefaultClient.Do(httpReq)
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
