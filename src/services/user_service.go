package services

import (
	"github.com/salmantaghooni/golang-car-web-api/src/api/dto"
	"github.com/salmantaghooni/golang-car-web-api/src/common"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     *logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	return &UserService{logger: &logger, cfg: config.GetConfig(), otpService: NewOtpService(cfg)}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	digits := s.cfg.OTP.Digits
	otp := common.GenereateOTP(digits)
	err := s.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
