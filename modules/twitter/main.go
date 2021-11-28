package twitter

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	helper "swapper/src"

	"github.com/go-resty/resty/v2"
)

type GuestTokenStruct struct {
	GuestToken string `json:"guest_token"`
}

func Login(username, password string) {
	client := resty.New()
	token := helper.RandString(23)
	client.SetCookie(&http.Cookie{
		Name:  "_mb_tk",
		Value: token,
	})

	client.SetHeaders(map[string]string{
		"Host":         "twitter.com",
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36",
		"Accept":       "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	})

	resp, _ := client.R().
		SetFormData(map[string]string{
			"authenticity_token":         token,
			"session[username_or_email]": username,
			"session[password]":          password,
		}).
		Post("https://twitter.com/sessions")
	fmt.Println(resp.RawResponse.Cookies())
}

func GetGuestToken() string {
	res := GuestTokenStruct{}
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Host":          "api.twitter.com",
		"Content-Type":  "application/x-www-form-urlencoded",
		"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36",
		"Accept":        "*/*",
		"Authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
	})
	resp, err := client.R().
		SetResult(&res).
		Post("https://api.twitter.com/1.1/guest/activate.json")

	if err != nil || resp.StatusCode() != 200 {
		fmt.Println("something went wrong getting the guest token")
	}
	return res.GuestToken

}

func CheckUser(username, guestToken string) (available bool, ratelimit int64) {
	res := GuestTokenStruct{}
	client := resty.New()
	client.SetHeaders(map[string]string{
		"Host":          "twitter.com",
		"Content-Type":  "application/json",
		"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36",
		"Accept":        "*/*",
		"Authorization": "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		"x-guest-token": guestToken,
	})

	resp, _ := client.R().
		SetResult(&res).
		Post("https://twitter.com/i/api/graphql/7mjxD3-C6BxitPMVQ6w0-Q/UserByScreenName?variables=" + url.QueryEscape(`{"screen_name":"`+username+`","withSafetyModeUserFields":false,"withSuperFollowsUserFields":false}`))
	if len(string(resp.Body())) > 30 {
		available = false
	} else {
		available = true
	}
	ratelimit, _ = strconv.ParseInt(resp.Header()["X-Rate-Limit-Remaining"][0], 10, 0)
	return
}
