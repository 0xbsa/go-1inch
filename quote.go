package go1inch

import (
	"context"
	"errors"
	"net/http"
)

// Quote gets quote for an aggregated swap which can be used with a web3 provider to send the transaction
func (c *Client) Quote(ctx context.Context, network Network, src, dst, amount string, opts *QuoteOpts) (*QuoteRes, int, http.Header, error) {
	endpoint := "/quote"

	if src == "" || dst == "" || amount == "" {
		return nil, 0, nil, errors.New("required parameter is missing")
	}

	queries := make(map[string]interface{})

	queries["src"] = src
	queries["dst"] = dst
	queries["amount"] = amount

	if opts != nil {

		if opts.Fee != "" {
			queries["fee"] = opts.Fee
		}
		if opts.Protocols != "" {
			queries["protocols"] = opts.Protocols
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
		if opts.MainRouteParts != "" {
			queries["mainRouteParts"] = opts.MainRouteParts
		}
		if opts.Parts != "" {
			queries["parts"] = opts.Parts
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

	var dataRes QuoteRes
	statusCode, headers, err := c.doRequest(ctx, network, endpoint, "GET", &dataRes, nil, queries)
	if err != nil {
		return nil, statusCode, headers, err
	}
	return &dataRes, statusCode, headers, nil
}
