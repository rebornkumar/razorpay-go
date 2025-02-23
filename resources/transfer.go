package resources

import (
	"fmt"

	"github.com/rebornkumar/razorpay-go/constants"
	"github.com/rebornkumar/razorpay-go/requests"
)

//Transfer ...
type Transfer struct {
	Request *requests.Request
}

// All fetches collection of transfer for the given queryParams.
func (t *Transfer) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return t.Request.Get(constants.TRANSFER_URL, queryParams, extraHeaders)
}

// Fetch fetches a transfer having the given transferID.
func (t *Transfer) Fetch(transferID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.TRANSFER_URL, transferID)
	return t.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new transfer for the given data.
func (t *Transfer) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return t.Request.Post(constants.TRANSFER_URL, data, extraHeaders)
}

// Edit edits the transfer having the given transferID.
func (t *Transfer) Edit(transferID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.TRANSFER_URL, transferID)
	return t.Request.Patch(url, data, extraHeaders)
}

// Reverse reverses the transfer having the given transferID.
func (t *Transfer) Reverse(transferID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/reversals", constants.TRANSFER_URL, transferID)
	return t.Request.Post(url, data, extraHeaders)
}

// Reversals fetches a collection of transfer associated with the given transferID.
func (t *Transfer) Reversals(transferID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/reversals", constants.TRANSFER_URL, transferID)
	return t.Request.Get(url, queryParams, extraHeaders)
}
