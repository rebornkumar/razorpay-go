package resources

import (
	"fmt"

	"github.com/rebornkumar/razorpay-go/constants"
	"github.com/rebornkumar/razorpay-go/requests"
)

//Order ...
type Order struct {
	Request *requests.Request
}

// All fetches multiple orders for the given query params.
func (order *Order) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return order.Request.Get(constants.ORDER_URL, queryParams, extraHeaders)
}

// Fetch fetches an order having the given orderID.
func (order *Order) Fetch(orderID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.ORDER_URL, orderID)
	return order.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new order for the given data
func (order *Order) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return order.Request.Post(constants.ORDER_URL, data, extraHeaders)
}

// Edit updates an order having the given orderID.
func (order *Order) Edit(orderID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s", constants.ORDER_URL, orderID)
	return order.Request.Put(url, data, extraHeaders)
}

// Payments fetches the payments for the given orderID.
func (order *Order) Payments(orderID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s/%s/payments", constants.ORDER_URL, orderID)
	return order.Request.Get(url, queryParams, extraHeaders)
}
