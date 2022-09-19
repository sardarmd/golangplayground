package service

// import (
// 	"testing"

// 	"github.banking/sardarmd/dto"
// 	// mockaccounts "github.banking/sardarmd/mocks/domain/accounts"
// 	"github.com/golang/mock/gomock"
// )

// var mockRepo *mockaccounts.MockAccountRepository

// // var accountService AccountService

// func setup(t *testing.T) func() {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	// mockRepo = mockaccounts.NewMockAccountRepository(ctrl)
// 	// NewAccountService(mockRepo)

// 	return func() {
// 		// accountService = nil
// 		defer ctrl.Finish()
// 	}

// }

// func Test_Should_return_validation_error_when_request_not_validate(t *testing.T) {
// 	//Arrange
// 	setup(t)

// }

// // --------------------------------------------------------------------------------------------------------------------------
// //                                      Helper functions
// // ---------------------------------------------------------------------------------------------------------------------------

// func NewAccountRequest() *dto.NewAccountRequest {
// 	request := dto.NewAccountRequest{
// 		CustomerId: "", AccountType: "saving",
// 		Amount: 0,
// 	}
// 	return &request
// }
