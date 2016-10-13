package fbclient

import (
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	appID := "<App ID>"
	rediredtURI := "<Redirect URI>"
	clientSecret := "<Client Secret>"
	code := "<Code parameter>"

	var client Client = NewClient(appID, rediredtURI, clientSecret)
	token, err := client.GetAccessToken(code)
	if err != nil {
		t.Errorf("%s Error=%s", token, err)
	}
	t.Errorf("%s", token)
}

func TestGetMyInfo(t *testing.T) {
	appID := "<App ID>"
	rediredtURI := "<Redirect URI>"
	clientSecret := "<Client Secret>"
	token := "<Access token>"

	var client Client = NewClient(appID, rediredtURI, clientSecret)
	obj, err := client.GetMyInfo(token, "email,gender")
	if err != nil {
		t.Errorf("Error %s", err)
		return
	}
	t.Errorf("%s", obj)
}
