package application

type WhiteList struct {
	AllowedOrigins []string
}

func NewWhiteList() WhiteList {
	return WhiteList{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://app.gomemon.bary822.me",
		},
	}
}

func (wl *WhiteList) IsAllowedOrigin(origin string) bool {
	for _, allowed_origin := range wl.AllowedOrigins {
		if allowed_origin == origin {
			return true
		}
	}

	return false
}
