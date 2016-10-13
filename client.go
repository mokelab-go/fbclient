package fbclient

type Client interface {
	GetAccessToken(code string) (AccessToken, error)
	GetMyInfo(token, fields string) (map[string]interface{}, error)
}

type AccessToken struct {
	Token     string
	TokenType string
	ExpiresIn float64
}
