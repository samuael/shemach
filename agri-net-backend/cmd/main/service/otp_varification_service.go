package service

import (
	"os"
	"strconv"
	"time"

	tm "github.com/buger/goterm"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
)

type OtpService struct {
	NumberOfOtpMessagesSent        uint
	NumberOfOtpMessagesLeft        uint
	NumberOfOtpActiveConfirmations uint
	SaveOtpCache                   chan *model.TeldaOTP
	SubscriberService              subscriber.ISubscriberService
	OTPResponse                    chan *model.TeldaOTPResponse
	UserService                    user.IUserService
}

func NewOtpService(subscriberService subscriber.ISubscriberService,
	service user.IUserService) *OtpService {
	return &OtpService{
		SubscriberService:              subscriberService,
		SaveOtpCache:                   make(chan *model.TeldaOTP),
		OTPResponse:                    make(chan *model.TeldaOTPResponse),
		NumberOfOtpMessagesSent:        0,
		NumberOfOtpMessagesLeft:        0,
		NumberOfOtpActiveConfirmations: 0,
		UserService:                    service,
	}
}

func (otpSer *OtpService) Run() {
	ticker := time.NewTicker(time.Second * 2)
	pendingConfirmationsDuration, er := strconv.Atoi(os.Getenv("PENDING_CONFIRMATION_DURATION"))
	pendingEmailConfirmationsDuration, er := strconv.Atoi(os.Getenv("PENDING_EMAIL_CONFIRMATION_DURATION"))
	if er != nil {
		pendingConfirmationsDuration = 10
	}
	counter := 0
	for {
		select {
		case <-ticker.C:
			{
				counter++
				tm.Print(tm.Bold(tm.Color(string("|"), tm.GREEN)))
				if counter%10 == 0 {
					tm.Println()
				}
				tm.Clear()
				timestamp := time.Now().Unix() - int64(pendingConfirmationsDuration*60)
				deletedConfirmationMessages, er := otpSer.SubscriberService.RemoveExpiredTempoSubscription(uint64(timestamp))
				otpSer.SubscriberService.DeleteTempoLoginSubscriber(uint64(timestamp))
				if er != nil {
					continue
				}
				otpSer.NumberOfOtpActiveConfirmations -= uint(deletedConfirmationMessages)

				// Email Confirmation Service functionalities
				etimestamp := time.Now().Unix() - int64(pendingEmailConfirmationsDuration*60)
				er = otpSer.UserService.DeletePendingEmailConfirmation(uint64(etimestamp))
			}
		case mresp := <-otpSer.OTPResponse:
			{
				val, _ := strconv.Atoi(mresp.RemainingSMS)
				otpSer.NumberOfOtpMessagesLeft = uint(val)
			}

		}

	}
}
