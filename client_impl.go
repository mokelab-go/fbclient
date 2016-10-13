package fbclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type impl struct {
	appID        string
	redirectURI  string
	clientSecret string
}

func NewClient(appID, redirectURI, clientSecret string) Client {
	return &impl{
		appID:        appID,
		redirectURI:  redirectURI,
		clientSecret: clientSecret,
	}
}

func (c *impl) GetAccessToken(code string) (AccessToken, error) {
	q := make(url.Values)
	q.Set("client_id", c.appID)
	q.Set("redirect_uri", c.redirectURI)
	q.Set("client_secret", c.clientSecret)
	q.Set("code", code)
	statusCode, obj, err := sendRequest("GET", "/v2.8/oauth/access_token", q)
	if err != nil {
		return AccessToken{}, err
	}
	if statusCode != 200 {
		return AccessToken{}, findErrorMessage(obj)
	}
	token := getString(obj, "access_token")
	tokenType := getString(obj, "token_type")
	expiresIn := getFloat64(obj, "expires_in")
	return AccessToken{
		Token:     token,
		TokenType: tokenType,
		ExpiresIn: expiresIn,
	}, nil
}

func (c *impl) GetMyInfo(token, fields string) (map[string]interface{}, error) {
	q := make(url.Values)
	q.Set("access_token", token)
	q.Set("fields", fields)
	statusCode, obj, err := sendRequest("GET", "/v2.8/me", q)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if statusCode != 200 {
		return map[string]interface{}{}, findErrorMessage(obj)
	}
	return obj, nil
}

func sendRequest(method, path string, query url.Values) (int, map[string]interface{}, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "graph.facebook.com",
		Path:   path,
	}
	u.RawQuery = query.Encode()
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return -1, nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, nil, err
	}
	doc := json.NewDecoder(resp.Body)
	var obj map[string]interface{}
	err = doc.Decode(&obj)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, obj, nil
}

func findErrorMessage(obj map[string]interface{}) error {
	e := getObject(obj, "error")
	if e == nil {
		return errors.New("unknown error")
	}
	return errors.New(getString(e, "message"))
}

func getObject(obj map[string]interface{}, key string) map[string]interface{} {
	if val, ok := obj[key].(map[string]interface{}); ok {
		return val
	} else {
		return nil
	}
}

func getString(obj map[string]interface{}, key string) string {
	if val, ok := obj[key].(string); ok {
		return val
	} else {
		return ""
	}
}

func getFloat64(obj map[string]interface{}, key string) float64 {
	if val, ok := obj[key].(float64); ok {
		return val
	} else {
		return 0
	}
}
