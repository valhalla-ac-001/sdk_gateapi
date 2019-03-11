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

type MarginCurrencyPair struct {
	// Currency pair
	Id string `json:"id,omitempty"`
	// Base currency
	Base string `json:"base,omitempty"`
	// Quote currency
	Quote string `json:"quote,omitempty"`
	// Leverage
	Leverage int32 `json:"leverage,omitempty"`
	// Minimum base currency to loan, `null` means no limit
	MinBaseAmount string `json:"min_base_amount,omitempty"`
	// Minimum quote currency to loan, `null` means no limit
	MinQuoteAmount string `json:"min_quote_amount,omitempty"`
	// Maximum borrowable amount for quote currency. Base currency limit is calculated by quote maximum and market price. `null` means no limit
	MaxQuoteAmount string `json:"max_quote_amount,omitempty"`
}
