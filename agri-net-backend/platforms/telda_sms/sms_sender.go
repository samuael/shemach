package telda_sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("accept", "text/plain")
	req.Header.Add("Content-Type", "application/json")
	APIKey := os.Getenv("API_KEY")
	req.Header.Add("Authorization", "Basic "+APIKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	response := &model.TeldaOTPResponse{}
	decoder := json.NewDecoder(bytes.NewReader(body))
	err = decoder.Decode(response)
	if err != nil {
		println(err.Error())
	}
	return response, err
}
