package go1inch

import (
	"context"
	"errors"
	"net/http"
)

// Swap gets swap for an aggregated swap which can be used with a web3 provider to send the transaction
func (c *Client) Swap(ctx context.Context, network Network, src, dst, amount, fromAddress string, slippage int64, opts *SwapOpts) (*SwapRes, int, http.Header, error) {
	endpoint := "/swap"

	if src == "" || dst == "" || amount == "" || fromAddress == "" {
		return nil, 0, nil, errors.New("required parameter is missing")
	}

	queries := make(map[string]interface{})

	queries["src"] = src
	queries["dst"] = dst
	queries["amount"] = amount
	queries["from"] = fromAddress
	queries["slippage"] = slippage

	if opts != nil {
		queries["burnChi"] = opts.BurnChi
		queries["allowPartialFill"] = opts.AllowPartialFill
		queries["disableEstimate"] = opts.DisableEstimate

		if opts.Protocols != "" {
			queries["protocols"] = opts.Protocols
		}
		if opts.Receiver != "" {
			queries["receiver"] = opts.Receiver
		}
		if opts.ReferrerAddress != "" {
			queries["referrer"] = opts.ReferrerAddress
		}
		if opts.Fee != "" {
			queries["fee"] = opts.Fee
		}
		if opts.GasPrice != "" {
			queries["gasPrice"] = opts.GasPrice
		}
		if opts.ComplexityLevel != "" {
			queries["complexityLevel"] = opts.ComplexityLevel
		}
		if opts.ConnectorTokens != "" {
			queries["connectorTokens"] = opts.ConnectorTokens
		}
		if opts.GasLimit != "" {
			queries["gasLimit"] = opts.GasLimit
		}
		if opts.Parts != "" {
			queries["parts"] = opts.Parts
		}
		if opts.VirtualParts != "" {
			queries["virtualParts"] = opts.VirtualParts
		}
		if opts.MainRouteParts != "" {
			queries["mainRouteParts"] = opts.MainRouteParts
		}

		if opts.IncludeGas {
			queries["includeGas"] = true
		}
		if opts.IncludeProtocols {
			queries["includeProtocols"] = true
		}
		if opts.IncludeTokensInfo {
			queries["includeTokensInfo"] = true
		}

	}

	var dataRes SwapRes
	statusCode, headers, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, headers, err
	}
	return &dataRes, statusCode, headers, nil
}
