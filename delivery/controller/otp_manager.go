package controller

import (
	"github.com/labstack/gommon/log"
	"github.com/tensuqiuwulu/bupda-reset-otp/service"
)

type OtpManagerControllerInterface interface {
	UpdatePhoneLimit() error
}

type OtpManagerControllerImpl struct {
	OtpManagerService service.OtpManagerServiceInterface
}

func NewOtpManagerControllerImpl(otpManagerService service.OtpManagerServiceInterface) OtpManagerControllerInterface {
	return &OtpManagerControllerImpl{OtpManagerService: otpManagerService}
}

func (c *OtpManagerControllerImpl) UpdatePhoneLimit() error {
	log.Info("Run service update phone limit")
	err := c.OtpManagerService.UpdatePhoneLimit()
	if err != nil {
		return err
	}
	return nil
}
