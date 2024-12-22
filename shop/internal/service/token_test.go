package service

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TokenTestSuite struct {
	suite.Suite
	service TokenService
}

func (s *TokenTestSuite) SetupSuite() {
	s.service = NewTokenService("key", time.Duration(time.Hour))
}

func (s *TokenTestSuite) TestGenerateToken() {
	token, err := s.service.GenerateAccessToken("1")
	s.Require().NoError(err)
	s.NotEmpty(token)
}

func TestTokenSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
