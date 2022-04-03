package httputils

import (
	"net/http"
	"strings"
)

func GetBearer(cookie []*http.Cookie) string {
	for _, c := range cookie {
		if c.Name == "reddit_session" {
			return c.Value
		}
	}
	return ""
}

func RefreshBearer(user, passwd, clientid, secret string) string {
	req, _ := http.NewRequest("POST", "https://ssl.reddit.com/api/v1/access_token",
		strings.NewReader("grant_type=password&username="+user+"&password="+passwd+"&client_id="))
	req.SetBasicAuth(clientid, secret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	return GetBearer(resp.Cookies())
}
