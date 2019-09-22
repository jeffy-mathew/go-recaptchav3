package go_recaptcha

import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"time"
)
type RecaptchaResult struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

type Recaptcha struct {
	Client 		*http.Client
	SecretKey	string  // secret key from google captcha admin console
}

var recaptcha Recaptcha

func Init(captcha Recaptcha){
	recaptcha = captcha
}

func Verfiy(captchaResponse string, clientIP string) (captchaResult RecaptchaResult, err error){
	url := fmt.Sprintf("https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s&remoteip=%s", recaptcha.SecretKey, captchaResponse, clientIP)
	req, err := http.NewRequest("POST", url, nil)
	resp, err := recaptcha.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &captchaResult)
	return
}
