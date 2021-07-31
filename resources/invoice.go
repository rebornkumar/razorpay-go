package resources

import (
	"fmt"

	"github.com/rebornkumar/razorpay-go/constants"
	"github.com/rebornkumar/razorpay-go/requests"
)

//Invoice ...
type Invoice struct {
	Request *requests.Request
}

// All fetches multiple invoices for the given queryParams.
func (inv *Invoice) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return inv.Request.Get(constants.INVOICE_URL, queryParams, extraHeaders)
}

// Fetch fetches an invoice having the given invoiceID.
func (inv *Invoice) Fetch(invoiceID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.INVOICE_URL, invoiceID)

	return inv.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new invoice for the given data.
func (inv *Invoice) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return inv.Request.Post(constants.INVOICE_URL, data, extraHeaders)
}
