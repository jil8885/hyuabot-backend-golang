package responses

type TokenDetails struct {
	AccessToken     string
	RefreshToken    string
	AccessUUID      string
	RefreshUUID     string
	AccessTokenExp  int64
	RefreshTokenExp int64
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AccessDetails struct {
	AccessUUID string
	UserName   string
}
