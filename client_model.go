package huawei

type Ret struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type GetTokenParams struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
