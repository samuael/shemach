package service

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/samuael/shemach/backend/pkg/session"
	"github.com/samuael/shemach/backend/pkg/user"
)

type OtpService struct {
	NumberOfOtpMessagesSent        uint
	NumberOfOtpMessagesLeft        uint
	NumberOfOtpActiveConfirmations uint
	NumberOfPendingRegistration    uint
	UserService                    user.IUserService
	SessionService                 session.ISessionService
	DeleteChannel                  chan bool

	PendingConfirmationCodeDuration int // phone number short code confirmation duration
	RegistrationWindowDuration      int // confirmed phone number user registration duration in minutes
}

func NewOtpService(service user.IUserService, sessionService session.ISessionService) *OtpService {
	return &OtpService{
		NumberOfOtpMessagesSent:        0,
		NumberOfOtpMessagesLeft:        0,
		NumberOfOtpActiveConfirmations: 0,
		UserService:                    service,
		SessionService:                 sessionService,
		DeleteChannel:                  make(chan bool),
	}
}

func (otpSer *OtpService) Run() {
	ticker := time.NewTicker(time.Second * 15)
	confirmedPhoneWaitTicker := time.NewTicker(time.Second * 30)
	sessionExpirationTicker := time.NewTicker(time.Second * 1)

	var err error
	otpSer.PendingConfirmationCodeDuration, err = strconv.Atoi(os.Getenv("PENDING_CONFIRMATION_DURATION"))
	if err != nil {
		otpSer.PendingConfirmationCodeDuration = 30
	}
	otpSer.RegistrationWindowDuration, err = strconv.Atoi(os.Getenv("CONFIRMED_PHONES_WAIT_DURATION_IN_MINUTES"))
	if err != nil {
		otpSer.RegistrationWindowDuration = 1
	}

	forgotPasswordShortcodeDurationInMinutes, err := strconv.ParseInt(os.Getenv("PENDING_FORGOT_PASSWORD_SHORTCODE_DURATION_IN_MINUTES"), 10, 64)
	if err != nil {
		forgotPasswordShortcodeDurationInMinutes = 1
	}
	forgotPasswordTableCleanerTicker := time.NewTicker(time.Duration(int64(int64(time.Second) * forgotPasswordShortcodeDurationInMinutes)))
	defer func() {
		ticker.Stop()
		sessionExpirationTicker.Stop()
		confirmedPhoneWaitTicker.Stop()
		forgotPasswordTableCleanerTicker.Stop()
	}()
	// sessionExpirationDurationInHours, err := strconv.Atoi(os.Getenv("SESSION_EXPIRATION_DURATION_IN_DAYS"))
	// if err != nil {
	// 	sessionExpirationDurationInHours = 2
	// }

	counter := 0
	for {
		select {
		case <-ticker.C:
			{
				counter++
				// tm.Print(tm.Bold(tm.Color(string("|"), tm.GREEN)))
				// if counter%10 == 0 {
				// 	tm.Println()
				// }
				// tm.Clear()
				timestamp := ((time.Now().Unix()) - int64(otpSer.PendingConfirmationCodeDuration))
				deletedCount, err := otpSer.UserService.DeleteExpiredTempoUsers(context.Background(), timestamp)
				if err != nil {
					continue
				}
				otpSer.NumberOfOtpActiveConfirmations -= uint(deletedCount)
			}
		case <-confirmedPhoneWaitTicker.C:
			timestamp := ((time.Now().Unix()) - int64(otpSer.RegistrationWindowDuration)*60)
			deletedCount, err := otpSer.UserService.DeleteExpiredVerifiedPhones(context.Background(), timestamp)
			if err != nil {
				continue
			}
			otpSer.NumberOfPendingRegistration -= uint(deletedCount)

		case <-sessionExpirationTicker.C:
			timestamp := ((time.Now().Unix()) - int64(otpSer.PendingConfirmationCodeDuration))
			deletedCount, err := otpSer.SessionService.DeleteSessionBeforeTimestamp(context.Background(), timestamp)
			if err != nil {
				continue
			}
			println("Deleted Sessions.", deletedCount)
		case <-otpSer.DeleteChannel:
			return
		case <-forgotPasswordTableCleanerTicker.C:
			timestamp := time.Now().Add(-time.Second * time.Duration(forgotPasswordShortcodeDurationInMinutes))
			count, err := otpSer.UserService.DeleteExpiredForgotPasswordSession(context.Background(), uint64(timestamp.Unix()))
			if err != nil {
				panic(err)
				log.Fatal(err)
			}
			println("Deleted Forgot password sessions: ", count)
		}
	}
}
