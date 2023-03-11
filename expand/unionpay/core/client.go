package core

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"upi/expand/unionpay/core/consts"
)

// Client 微信支付API v3 基础 Client
type Client struct {
	httpClient *http.Client
	//credential auth.Credential
	//validator  auth.Validator
	//signer     auth.Signer
	//cipher     cipher.Cipher
}

// APIResult 请求结果
type APIResult struct {
	// 本次请求所使用的 HTTPRequest
	Request *http.Request
	// 本次请求所获得的 HTTPResponse
	Response *http.Response
}

var (
	regJSONTypeCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	regXMLTypeCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// NewClient 初始化一个微信支付API v3 HTTPClient
//
// 初始化的时候你可以传递多个配置信息
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: consts.DefaultTimeout,
		},
	}
}

// Request 向银联发送请求
//
// 相比于 Get / Post / Put / Patch / Delete 方法，本方法可以设置更多内容
// 特别地，如果需要为当前请求设置 Header，可以使用本方法
func (client *Client) Request(
	ctx context.Context,
	method, requestPath string,
	headerParams http.Header,
	queryParams url.Values,
	postBody interface{},
	contentType string,
) (result *APIResult, err error) {

	// Setup path and query parameters
	varURL, err := url.Parse(requestPath)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := varURL.Query()
	for k, values := range queryParams {
		for _, v := range values {
			query.Add(k, v)
		}
	}

	// Encode the parameters.
	varURL.RawQuery = query.Encode()

	if postBody == nil {
		return client.doRequest(ctx, method, varURL.String(), headerParams, contentType, nil, "")
	}

	// Detect postBody type and set body content
	if contentType == "" {
		contentType = consts.ApplicationJSON
	}
	var body *bytes.Buffer
	body, err = setBody(postBody, contentType)
	if err != nil {
		return nil, err
	}
	return client.doRequest(ctx, method, varURL.String(), headerParams, contentType, body, body.String())
}

// Set Request body from an interface
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if regJSONTypeCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if regXMLTypeCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

func (client *Client) doRequest(
	ctx context.Context,
	method string,
	requestURL string,
	header http.Header,
	contentType string,
	reqBody io.Reader,
	signBody string,
) (*APIResult, error) {

	var (
		err           error
		request       *http.Request
	)

	// Construct Request
	if request, err = http.NewRequestWithContext(ctx, method, requestURL, reqBody); err != nil {
		return nil, err
	}

	// Header Setting Priority:
	// Fixed Headers > Per-Request Header Parameters

	// Add Request Header Parameters
	for key, values := range header {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	// Set Fixed Headers
	request.Header.Set(consts.Accept, "*/*")
	request.Header.Set(consts.ContentType, contentType)

	ua := fmt.Sprintf(consts.UserAgentFormat, consts.Version, runtime.GOOS, runtime.Version())
	request.Header.Set(consts.UserAgent, ua)

	// Send HTTP Request
	result, err := client.doHTTP(request)
	if err != nil {
		return result, err
	}
	// Check if Success
	if err = CheckResponse(result.Response); err != nil {
		return result, err
	}
	// Validate WechatPay Signature
	return result, nil
}

// CheckResponse 校验请求是否成功
//
// 当http回包的状态码的范围不是200-299之间的时候，会返回相应的错误信息，主要包括http状态码、回包错误码、回包错误信息提示
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid response, read body error: %w", err)
	}
	_ = resp.Body.Close()

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(slurp))
	apiError := &APIError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(slurp),
	}
	// 忽略 JSON 解析错误，均返回 apiError
	_ = json.Unmarshal(slurp, apiError)
	return apiError
}

func (client *Client) doHTTP(req *http.Request) (result *APIResult, err error) {
	result = &APIResult{
		Request: req,
	}

	result.Response, err = client.httpClient.Do(req)
	return result, err
}