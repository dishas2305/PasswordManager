package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"passmanager/config"

	logger "github.com/sirupsen/logrus"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

type Header struct {
	Key   string
	Value string
}

func GetHeader() []Header {
	headers := []Header{}
	headers = append(headers, Header{Key: "Content-type", Value: "application/json"})

	return headers
}

func POSTMethod(URL string, body interface{}, headers []Header) (*http.Response, error) {
	logger.Info("func_POSTMethod: body data: ", body, URL)
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		logger.Error("func_POSTMethod: Error in marshal: ", err)
		return nil, err
	}

	// Making post call
	request, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		logger.Error("func_POSTMethod: Error in http new request: ", err)
		return nil, err
	}
	for _, header := range headers {
		request.Header.Add(header.Key, header.Value)
	}
	logger.Info("func_POSTMethod: Request data: ", request)
	response, err := Client.Do(request)
	if err != nil {
		logger.Error("func_POSTMethod: URL: ", URL, " Error: ", err)
		return nil, err
	}
	logger.Info("func_POSTMethod: response data: ", response)
	// End - making post call
	// defer response.Body.Close()

	return response, nil
}

func GETMethod(URL string, headers []Header) (*http.Response, error) {
	// Making get call
	request, err := http.NewRequest(http.MethodGet, URL, bytes.NewBuffer(nil))
	if err != nil {
		logger.Error("func_GETMethod: Error in http.NewRequest", err)
		return nil, err
	}
	for _, header := range headers {
		request.Header.Add(header.Key, header.Value)
	}
	logger.Info("func_GETMethod: Request data: ", request)
	response, err := Client.Do(request)
	if err != nil {
		logger.Error("func_GETMethod: URL: ", URL, " Error: ", err)
		return nil, err
	}
	logger.Info("func_GETMethod: response data: ", response)
	// End - making get call

	return response, nil
}

func DELETEMethod(URL string, headers []Header) (*http.Response, error) {
	// Making delete call
	request, err := http.NewRequest(http.MethodDelete, URL, bytes.NewBuffer(nil))
	if err != nil {
		logger.Error("func_DELETEMethod: Error in http.NewRequest", err)
		return nil, err
	}
	for _, header := range headers {
		request.Header.Add(header.Key, header.Value)
	}
	logger.Info("func_DELETEMethod: Request data: ", request)
	response, err := Client.Do(request)
	if err != nil {
		logger.Error("func_DELETEMethod: URL: ", URL, " Error: ", err)
		return nil, err
	}
	logger.Info("func_DELETEMethod: response data: ", response)
	// End - making delete call

	return response, nil
}

func PUTMethod(URL string, body interface{}, headers []Header) (*http.Response, error) {
	logger.Info("func_PUTMethod: body data: ", body, URL)
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		logger.Error("func_PUTMethod: Error in marshal: ", err)
		return nil, err
	}

	// Making put call
	request, err := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		logger.Error("func_PUTMethod: Error in http new request: ", err)
		return nil, err
	}
	for _, header := range headers {
		request.Header.Add(header.Key, header.Value)
	}
	logger.Info("func_PUTMethod: Request data: ", request)
	response, err := Client.Do(request)
	if err != nil {
		logger.Error("func_PUTMethod: URL: ", URL, " Error: ", err)
		return nil, err
	}
	logger.Info("func_PUTMethod: response data: ", response)
	// End - making put call
	// defer response.Body.Close()

	return response, nil
}

func MSHttpClientCall(httpMethod string, URL string, body interface{}, headers []Header) (*http.Response, error) {
	var httpClientRes *http.Response

	if httpMethod == "POST" {
		response, err := POSTMethod(URL, body, headers)
		if err != nil {
			logger.Error("func_MSHttpClientCall: Error from post method: ", err)
			return nil, err
		}
		httpClientRes = response
	} else if httpMethod == "GET" {
		response, err := GETMethod(URL, headers)
		if err != nil {
			logger.Error("func_MSHttpClientCall : Error from get method")
			return nil, err
		}
		httpClientRes = response
	} else {
		return nil, config.ErrInvalidHttpMethod
	}

	// Check status code
	if err := CheckStatusCode(httpClientRes); err != nil {
		resBody, readAllErr := ioutil.ReadAll(httpClientRes.Body)
		if readAllErr != nil {
			logger.Error("func_MSHttpClientCall: Error in reading response: ", err)
		}
		logger.Error("func_MSHttpClientCall: URL: ", URL, " Error in check status code: ", err, " And response: ", string(resBody))
		return nil, err
	}
	// End - Check status code

	return httpClientRes, nil
}

func CheckStatusCode(response *http.Response) error {
	switch response.StatusCode {
	case 200:
		return nil
	case 400:
		return config.ErrHttpCallBadRequest
	case 401:
		return config.ErrHttpCallUnauthorized
	case 404:
		return config.ErrHttpCallNotFound
	case 500:
		return config.ErrHttpCallInternalServerError
	default:
		logger.Error("new status code found: ", response.StatusCode)
		return config.ErrWentWrong
	}
}

func PATCHMethod(URL string, body interface{}, headers []Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		logger.Error("func_PATCHMethod: Error in marshal: ", err)
		return nil, err
	}

	// Making delete call
	request, err := http.NewRequest(http.MethodPatch, URL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		logger.Error("func_PATCHMethod: Error in http.NewRequest", err)
		return nil, err
	}
	for _, header := range headers {
		request.Header.Add(header.Key, header.Value)
	}
	fmt.Println("func_PATCHMethod: Request data: ", request)
	response, err := Client.Do(request)
	if err != nil {
		logger.Error("func_PATCHMethod: URL: ", URL, " Error: ", err)
		return nil, err
	}
	fmt.Println("func_PATCHMethod: response data: ", response)
	// End - making delete call

	return response, nil
}
