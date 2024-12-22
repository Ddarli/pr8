package handler

import (
	"context"
	mocks "github.com/Ddarli/app/shop/mocks/service"
	"github.com/Ddarli/utils/models"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HandlerSuite struct {
	suite.Suite
	handler      Handler
	tokenService *mocks.TokenService
	service      *mocks.Service
}

func (s *HandlerSuite) SetupSuite() {
	s.service = new(mocks.Service)
	s.tokenService = new(mocks.TokenService)

	s.handler = NewHttpHandler(s.service, s.tokenService)
}

func (s *HandlerSuite) TestGetProducts() {

	s.service.EXPECT().GetAll(context.Background()).Return([]*models.Product{}, nil)

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/api/v1/products", nil)
	s.NoError(err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request
	s.handler.(*httpHandler).router.ServeHTTP(rr, req)

	// Check the status code
	s.Equal(http.StatusNotFound, rr.Code)
}

func TestService(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}
