package model

type Subscriber struct {
	ID            uint64 `json:"id"`
	Fullname      string `json:"fullname"`
	Lang          string `json:"lang"`
	Subscriptions []int  `json:"subscriptions"`
	Role          int
	Phone         string `json:"phone"`
}

type TempoSubscriber struct {
	ID               int    `json:"id"`
	Fullname         string `json:"fullname"`
	Lang             string `json:"lang"`
	Role             uint8  `json:"role"`
	Phone            string `json:"phone"`
	ConfirmationCode string `json:"confirmation_number"`
	Unix             int64  `json:"time_stamp_of_creation"`
	Trials           uint8  `json:"trials"`
}

func (ts *TempoSubscriber) GetSubscriber() *Subscriber {
	return &Subscriber{
		Fullname:      ts.Fullname,
		Lang:          ts.Lang,
		Subscriptions: []int{},
		Role:          int(ts.Role),
		Phone:         ts.Phone,
	}
}

type TempoLoginSubscriber struct {
	ID           int    `json:"id"`
	Phone        string `json:"phone"`
	Unix         uint64 `json:"unix"`
	Confirmation string `json:"confirmation_code"`
}

type TeldaOTP struct {
	Phone      string `json:"phone"`
	OTP        string `json:"otp"`
	SenderName string `json:"senderName"`
	Remark     string `json:"remark"`
}

/*
{
	"phone": "+251992078204",
	"otp": "21222",
	"senderName": "Telda",
	"remark": "This is your confirmation code."
  }
*/

type TeldaOTPResponse struct {
	MsGReceiverPhone string `json:"msG_RecieverPhone"`
	MsgOperator      string `json:"msG_Operator"`
	MsgText          string `json:"msG_Text"`
	// MsgISO3166       string `json:"msG_ISO3166"`
	MsgSendDate     string `json:"msG_SendDate"`
	MsgReference    uint   `json:"msG_Reference"`
	MsgErrorCode    string `json:"msG_ErrorCode"`
	MsgShortMessage string `json:"msG_ShortMessage"`
	MsgLongMessage  string `json:"msG_LongMessage"`
	RemainingSMS    string `json:"remaining_SMS"`
}

/*
	{
    "msG_RecieverPhone": "+251992078204",
    "msG_Operator": "+251910737676",
    "msG_Text": "<#> Your Telda otp:\n21222\n\nThis is your confirmation code.",
    "msG_ISO3166": null,
    "msG_SendDate": "2022-04-08T09:34:28.0601036+03:00",
    "msG_Reference": 9601,
    "msG_ErrorCode": "00001",
    "msG_ShortMessage": "Success",
    "msG_LongMessage": "Message Composed Successfully",
    "remaining_SMS": "499"
}

*/
