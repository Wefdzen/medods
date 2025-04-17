package db

type User struct {
	ID               uint
	Guid             string `json:"guid"`
	RefreshTokenHash string `json:"refreshTokenHash"`
	IpClient         string `json:"ipClient"`
	LiveToken        string `json:"liveToken"`
	UnicCode         string `json:"unicCode"`
}
