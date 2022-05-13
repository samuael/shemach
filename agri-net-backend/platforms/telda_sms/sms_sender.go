package telda_sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

func SendOtp(otp *model.TeldaOTP) (*model.TeldaOTPResponse, error) {
	url := "https://api.telda.com.et/api/write/SendOTP"
	method := "POST"

	payload := strings.NewReader(string(func() []byte {
		if bytes, er := json.Marshal(otp); er != nil {
			println(er.Error())
			return []byte{}
		} else {
			return bytes
		}
	}()))

	OTPEnabled, _ := strconv.ParseBool(os.Getenv("OTP_ENABLE"))

	var res *http.Response
	var err error
	if OTPEnabled {
		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			return nil, err
		}
		req.Header.Add("accept", "text/plain")
		req.Header.Add("Content-Type", "application/json")
		APIKey := os.Getenv("API_KEY")
		req.Header.Add("Authorization", "Basic "+APIKey)
		res, err = client.Do(req)
	} else {
		res = &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.0",
		}
	}
	if err != nil {
		return nil, err
	}
	if OTPEnabled {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	var body []byte
	if OTPEnabled {
		body, err = ioutil.ReadAll(res.Body)
	} else {
		body = []byte(`{"msG_RecieverPhone":"` + otp.Phone + `","msG_Operator":"+251910737676","msG_Text":"\u003c#\u003e Your Agri-Net otp:\n` + otp.OTP + `\n\nThis is your confirmation and temporary password ","msG_SendDate":"2022-05-09T12:07:23.0686171+03:00","msG_Reference":11788,"msG_ErrorCode":"00001","msG_ShortMessage":"Success","msG_LongMessage":"Message Composed Successfully","remaining_SMS":"496"}`)
	}
	if err != nil {
		return nil, err
	}
	response := &model.TeldaOTPResponse{}
	decoder := json.NewDecoder(bytes.NewReader(body))
	err = decoder.Decode(response)
	return response, err
}
