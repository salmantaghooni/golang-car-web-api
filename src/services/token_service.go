package services

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/salmantaghooni/golang-car-web-api/src/api/dto"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/service_errors"
)

type TokenService struct {
	logger *logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	UserName     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	return &TokenService{
		cfg:    cfg,
		logger: &logger,
	}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {
	td := &dto.TokenDetail{}
	td.AccessTokenExpireTime = time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	td.RefreshTokenExpireTime = time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	act := jwt.MapClaims{}

	act["user_id"] = token.UserId
	act["first_name"] = token.FirstName
	act["last_name"] = token.LastName
	act["username"] = token.UserName
	act["mobile_number"] = token.MobileNumber
	act["email"] = token.Email
	act["roles"] = token.Roles
	act["exp"] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, act)
	var err error
	td.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}
	rtc["user_id"] = token.UserId
	rtc["exp"] = td.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)
	td.RefreshToken, err = rt.SignedString([]byte(s.cfg.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(s.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}
