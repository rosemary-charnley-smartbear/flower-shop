/*
 * Flower Shop -- Billing
 *
 * This is a sample Flower Shop server.   This includes the Billing REST APIs. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). 
 *
 * API version: 2.0.0
 * Contact: rosemary.charnley@smartbear.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Bill struct {
	Id string `json:"id"`
	Person *Person `json:"person,omitempty"`
	BillingAddress *Address `json:"billingAddress,omitempty"`
	DiscountCode string `json:"discountCode,omitempty"`
	PaymentType string `json:"paymentType,omitempty"`
	Discount float64 `json:"discount,omitempty"`
	Subtotal float64 `json:"subtotal,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	Total float64 `json:"total,omitempty"`
}
