package resources

import (
	"fmt"

	"github.com/rebornkumar/razorpay-go/constants"
	"github.com/rebornkumar/razorpay-go/requests"
)

//PaymentLink ...
type PaymentLink struct {
	Request *requests.Request
}

// All fetches multiple PaymentLink for the given queryParams.
func (pl *PaymentLink) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return pl.Request.Get(constants.PaymentLink_URL, queryParams, extraHeaders)
}

// Fetch fetches an PaymentLink having the given paymentLinkID.
func (pl *PaymentLink) Fetch(paymentLinkID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.PaymentLink_URL, paymentLinkID)

	return pl.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new PaymentLink  for the given data.
func (pl *PaymentLink) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return pl.Request.Post(constants.PaymentLink_URL, data, extraHeaders)
}
