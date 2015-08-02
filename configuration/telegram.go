package configuration

type Telegram struct {
	Token        string `json:"token"`
	AllowedUsers []int  `json:"allowed_users"`
}
