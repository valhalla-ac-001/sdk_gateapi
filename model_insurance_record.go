/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.7.2
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type InsuranceRecord struct {
	// Unix timestamp in seconds
	T int64 `json:"t,omitempty"`
	// Insurance balance
	B string `json:"b,omitempty"`
}
