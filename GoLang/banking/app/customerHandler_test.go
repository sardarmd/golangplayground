package app

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.banking/sardarmd/dto"
// 	"github.banking/sardarmd/mocks/services"
// 	"github.com/golang/mock/gomock"
// 	"github.com/gorilla/mux"
// 	"github.com/sardarmd/banking-lib/errs"
// )

// var ch CustomerHandler
// var mockService *services.MockCustomerService
// var router *mux.Router
// var dummyCustomer = dto.CustomerResponse{Id: "1001", Name: "Sardar", City: "Bangalore", Zip_code: "560067", DateofBirth: "25-08-1987", Status: "1"}

// func setup(t *testing.T) func() {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	mockService = services.NewMockCustomerService(ctrl)
// 	ch = CustomerHandler{mockService}

// 	router = mux.NewRouter()

// 	return func() {
// 		router = nil
// 		defer ctrl.Finish()
// 	}

// }

// func Test_should_return_customer_with_status_code_200(t *testing.T) {

// 	//Arrange
// 	teardown := setup(t)
// 	defer teardown()

// 	mockService.EXPECT().GetCustomer("").Return(&dummyCustomer, nil)
// 	router.HandleFunc("/customers", ch.getCustomer)
// 	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

// 	//Act
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, request)

// 	//Assert
// 	if recorder.Code != http.StatusOK {
// 		t.Error("Failed while testing status ok")
// 	}

// }

// func Test_should_return_Status_500_with_error_message(t *testing.T) {

// 	//Arrange
// 	teardown := setup(t)
// 	defer teardown()

// 	error := errs.NewCustomerException()
// 	mockService.EXPECT().GetCustomer("").Return(nil, error)
// 	router.HandleFunc("/customers", ch.getCustomer)
// 	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

// 	//Act
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, request)

// 	//Assert
// 	if recorder.Code != http.StatusInternalServerError {
// 		t.Error("Failed while testing status Internal server error")
// 	}
// }
