package adapter

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"gorgs/member/model"
	"reflect"
)

type HttpRequest struct {
}

func (h HttpRequest) _request(method, url string, headers map[string]interface{}, params interface{}) (model.HttpRequestResponse, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)

	for key, val := range headers {
		req.Header.Add(key, fmt.Sprintf("%v", val))
	}

	req.Header.SetMethod(method)
	if method == "GET" {
		q := req.URI().QueryArgs()
		if params != nil && reflect.TypeOf(params).Kind() == reflect.Map {

			mapParams, _ := params.(map[string]interface{})
			if mapParams != nil {
				for key, val := range mapParams {
					q.Add(key, fmt.Sprintf("%v", val))
				}
			}
		}
	} else {
		if params != nil && reflect.TypeOf(params).Kind() == reflect.String {
			req.SetBodyString(params.(string))
		} else {
			req.Header.Set("Content-Type", "application/json")
			reqByte, _ := json.Marshal(params)
			reqStr := string(reqByte)
			req.SetBodyString(reqStr)

		}
	}

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	response := model.HttpRequestResponse{}
	if err != nil {
		return response, err
	} else {
		response.StatusCode = resp.Header.StatusCode()
		response.ResponseString = resp.String()
		json.Unmarshal(resp.Body(), &response.Body)
		return response, nil
	}

}

func (h HttpRequest) DecodeBody(input interface{}, output interface{}) error {

	jsonString, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonString, &output)
	if err != nil {
		return err
	}
	return nil

}

func (h HttpRequest) DoRequest(method, url string, headers map[string]interface{}, params interface{}) (model.HttpRequestResponse, error) {
	resp, err := h._request(method, url, headers, params)
	return resp, err
}

func (h HttpRequest) Post(url string, headers map[string]interface{}, params interface{}) (model.HttpRequestResponse, error) {
	return h.DoRequest("POST", url, headers, params)
}

func (h HttpRequest) Get(url string, headers map[string]interface{}, params interface{}) (model.HttpRequestResponse, error) {
	return h.DoRequest("GET", url, headers, params)
}

func (h HttpRequest) Put(url string, headers map[string]interface{}, params interface{}) (model.HttpRequestResponse, error) {
	return h.DoRequest("PUT", url, headers, params)
}

func (h HttpRequest) Delete(url string, headers map[string]interface{}, params map[string]interface{}) (model.HttpRequestResponse, error) {
	return h.DoRequest("DELETE", url, headers, params)
}
