package validator

import (
	"github.banking/sardarmd/dto"
	"github.com/sardarmd/banking-lib/errs"
)

func ValidateNewAccount(r dto.NewAccountRequest) *errs.AppErrors {

	if r.Amount < 5000 {
		return errs.NewValidationError("To open the account minimum amount should be 5000")
	} else {
		return nil
	}

}
