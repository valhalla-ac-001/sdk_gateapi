/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.5.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Margin account detail. `base` refers to base currency, while `quotes to quote currency
type MarginAccount struct {
	// Currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
	Base MarginAccountCurrency `json:"base,omitempty"`
	Quote MarginAccountCurrency `json:"quote,omitempty"`
}
