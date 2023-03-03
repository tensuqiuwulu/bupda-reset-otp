package service

import (
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/bupda-reset-otp/repository"
	"gorm.io/gorm"
)

type OtpManagerServiceInterface interface {
	UpdatePhoneLimit() error
}

type OtpManagerServiceImpl struct {
	DB                   *gorm.DB
	Logger               *logrus.Logger
	OtpManagerRepository repository.OtpManagerRepository
}

func NewOtpManagerServiceImpl(db *gorm.DB, logger *logrus.Logger, otpManagerRepository repository.OtpManagerRepository) OtpManagerServiceInterface {
	return &OtpManagerServiceImpl{DB: db, Logger: logger, OtpManagerRepository: otpManagerRepository}
}

func (s *OtpManagerServiceImpl) UpdatePhoneLimit() error {
	otpManager, err := s.OtpManagerRepository.GetAllPhoneLimit(s.DB)
	if err != nil {
		return err
	}

	if len(otpManager) == 0 {
		log.Info("No phone limit")
	}

	for _, v := range otpManager {
		err := s.OtpManagerRepository.UpdatePhoneLimit(s.DB, v.Phone, 5)
		if err != nil {
			return err
		}
	}
	return nil
}
