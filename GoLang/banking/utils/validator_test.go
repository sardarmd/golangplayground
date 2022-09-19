package validator

import (
	"net/http"
	"testing"

	"github.banking/sardarmd/dto"
)

func Test_should_return_error_when_amount_less_than_5000(t *testing.T) {
	request := dto.NewAccountRequest{CustomerId: "1230", Amount: 200, AccountType: "saving"}

	appError := ValidateNewAccount(request)

	if appError.Message != "To open the account minimum amount should be 5000" {
		t.Error("Invalid message when amount is less than 5000")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid error code when amount is less than 5000")
	}
}
