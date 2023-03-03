package repository

import (
	"github.com/tensuqiuwulu/bupda-reset-otp/config"
	mysqlmodel "github.com/tensuqiuwulu/bupda-reset-otp/model/mysql_model"
	"gorm.io/gorm"
)

type OtpManagerRepository interface {
	GetAllPhoneLimit(db *gorm.DB) ([]mysqlmodel.OtpManager, error)
	UpdatePhoneLimit(db *gorm.DB, phone string, limit int) error
}

type OtpManagerRepositoryImpl struct {
	DB *config.Database
}

func NewOtpManagerRepositoryImpl(db *config.Database) OtpManagerRepository {
	return &OtpManagerRepositoryImpl{DB: db}
}

func (r *OtpManagerRepositoryImpl) GetAllPhoneLimit(db *gorm.DB) ([]mysqlmodel.OtpManager, error) {
	var otpManager []mysqlmodel.OtpManager
	err := db.Where("phone_limit < 5").Find(&otpManager).Error
	if err != nil {
		return nil, err
	}
	return otpManager, nil
}

func (r *OtpManagerRepositoryImpl) UpdatePhoneLimit(db *gorm.DB, phone string, limit int) error {
	err := db.Model(&mysqlmodel.OtpManager{}).Where("phone = ?", phone).Update("phone_limit", limit).Error
	if err != nil {
		return err
	}
	return nil
}
