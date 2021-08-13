package huawei

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/google/go-querystring/query"
)

type HuaweiClient struct {
	clientID       string
	clinetSecret   string
	client         *http.Client
	token          string
	tokenExpiresIn int64
	baseURL        *url.URL
}

// 代码中使用这个构造方法
func NewHuaweiClient(clinetID, clientSecret string) *HuaweiClient {
	return &HuaweiClient{
		clientID:     clinetID,
		clinetSecret: clientSecret,
		client:       &http.Client{},
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "connect-api.cloud.huawei.com",
		},
	}
}

// 单元测试中使用这个构造方法
func NewHuaweiClientWithEnv() (*HuaweiClient, error) {
	var clientID string
	var clientSecret string

	if clientID = os.Getenv("HUAWEI_CLIENT_ID"); clientID == "" {
		return nil, fmt.Errorf("HUAWEI_CLIENT_ID is empty")
	}

	if clientSecret = os.Getenv("HUAWEI_CLIENT_SECRET"); clientSecret == "" {
		return nil, fmt.Errorf("HUAWEI_CLIENT_SECRET is empty")
	}

	return &HuaweiClient{
		clientID:     clientID,
		clinetSecret: clientSecret,
		client:       &http.Client{},
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "connect-api.cloud.huawei.com",
		},
	}, nil
}

// 创建一个Request
func (c *HuaweiClient) NewRequest(method, path string, queryParams interface{}, bodyParams interface{}) (*http.Request, error) {

	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	c.baseURL.Path = unescaped

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Content-Type", "application/json")

	var body []byte
	if bodyParams != nil {
		body, err = json.Marshal(bodyParams)
		if err != nil {
			return nil, err
		}
	}

	if queryParams != nil {
		q, err := query.Values(queryParams)
		if err != nil {
			return nil, err
		}
		c.baseURL.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, c.baseURL.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

// 发送请求
func (c *HuaweiClient) Do(req *http.Request, v interface{}) (*http.Response, error) {

	if c.token == "" || c.tokenExpiresIn == 0 || time.Now().Unix() >= c.tokenExpiresIn {
		err := refreshToken(c)
		if err != nil {
			panic(err)
		}
	}

	req.Header.Add("client_id", c.clientID)
	req.Header.Add("Authorization", "Bearer "+c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("the huawei client received an unhealthy status code: %d, message: %s", resp.StatusCode, resp.Status)
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}

// 刷新Token
func refreshToken(c *HuaweiClient) error {
	params := &GetTokenParams{
		GrantType:    "client_credentials",
		ClientId:     c.clientID,
		ClientSecret: c.clinetSecret,
	}

	path := "/api/oauth2/v1/token"
	c.baseURL.Path = path

	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	resp, err := http.Post(c.baseURL.String(), "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var token GetTokenRes
	err = json.Unmarshal(content, &token)
	if err != nil {
		return err
	}

	if token.AccessToken == "" {
		return fmt.Errorf("get token failed: %s", string(content))
	}

	c.token = token.AccessToken
	c.tokenExpiresIn = token.ExpiresIn + time.Now().Unix()

	return nil
}
