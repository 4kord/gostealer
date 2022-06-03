package browsers

type PasswordEntry struct {
	OriginUrl string `json:"origin_url"`
	Username  string `json:"username_value"`
	Password  string `json:"password"`
}

type CookieEntry struct {
	Host   string `json:"host"`
	Path   string `json:"path"`
	Expiry int    `json:"expiry"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type AutofillEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
