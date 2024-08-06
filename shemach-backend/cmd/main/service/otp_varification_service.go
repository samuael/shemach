package service

import (
	"os"
	"strconv"
	"time"

	tm "github.com/buger/goterm"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/subscriber"
	"github.com/samuael/shemach/shemach-backend/pkg/user"
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
	ticker := time.NewTicker(time.Second * 30)

	agentConfirmationTicker := time.NewTicker(time.Hour * 1)

	pendingConfirmationsDuration, er := strconv.Atoi(os.Getenv("PENDING_CONFIRMATION_DURATION"))
	if er != nil {
		pendingConfirmationsDuration = 30
	}
	pendingEmailConfirmationsDuration, er := strconv.Atoi(os.Getenv("PENDING_EMAIL_CONFIRMATION_DURATION"))
	if er != nil {
		pendingEmailConfirmationsDuration = 20
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
				timestamp := ((time.Now().Unix()) - int64(pendingConfirmationsDuration*60))
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
		case <-agentConfirmationTicker.C:
			{
				// Phone Confirmation Deleting expired confimation messages
				etimestamp := time.Now().Unix() - int64(pendingEmailConfirmationsDuration*60*60)
				counts, _ := otpSer.UserService.RemoveExpiredCXPConfirmations(uint64(etimestamp))
				if counts > 0 {
					tm.Printf(tm.Bold("CXP Expire Deletion %d Records"), counts)
				}
			}
		case mresp := <-otpSer.OTPResponse:
			{
				val, _ := strconv.Atoi(mresp.RemainingSMS)
				otpSer.NumberOfOtpMessagesLeft = uint(val)
			}

		}

	}
}
