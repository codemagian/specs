/*
 * Qase.io API
 *
 * You can use our API to access Qase.io API endpoints, which allows to retrieve information about entities stored in database and perform actions with them. The API is organized around REST.
 *
 * API version: 1.0.0
 * Contact: support@qase.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package qase

type InlineResponse200Result struct {
	Total    int32      `json:"total,omitempty"`
	Filtered int32      `json:"filtered,omitempty"`
	Count    int32      `json:"count,omitempty"`
	Entities *[]Project `json:"entities,omitempty"`
}
