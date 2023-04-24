// Copyright 2023 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package client

import (
	_context "context"
	"fmt"
	"io"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// Linger please
var (
	_ _context.Context
)

// CallAPIService CallAPI service
type CallAPIService service

// Call Call invokes an arbitrary, network-specific procedure call with network-specific parameters.
// The guidance for what this endpoint should or could do is purposely left vague. In Ethereum, this
// could be used to invoke eth_call to implement an entire Rosetta API interface for some smart
// contract that is not parsed by the implementation creator (like a DEX). This endpoint could also
// be used to provide access to data that does not map to any Rosetta models instead of requiring an
// integrator to use some network-specific SDK and call some network-specific endpoint (like
// surfacing staking parameters). Call is NOT a replacement for implementing Rosetta API endpoints
// or mapping network-specific data to Rosetta models. Rather, it enables developers to build
// additional Rosetta API interfaces for things they care about without introducing complexity into
// a base-level Rosetta implementation. Simply put, imagine that the average integrator will use
// layered Rosetta API implementations that each surfaces unique data.
func (a *CallAPIService) Call(
	ctx _context.Context,
	callRequest *types.CallRequest,
) (*types.CallResponse, *types.Error, error) {
	var (
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/call"
	localVarHeaderParams := make(map[string]string)

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = callRequest

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to prepare request: %w", err)
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, nil, fmt.Errorf("failed to call API: %w", err)
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			_ = fmt.Sprintf("unable to close response body: %w", err)
		}
	}(localVarHTTPResponse.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response: %w", err)
	}

	switch localVarHTTPResponse.StatusCode {
	case _nethttp.StatusOK:
		var v types.CallResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf(
				"failed to decode when hit status code 200, response body %s: %w",
				string(localVarBody),
				err,
			)
		}

		return &v, nil, nil
	case _nethttp.StatusInternalServerError:
		var v types.Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, fmt.Errorf(
				"failed to decode when hit status code 500, response body %s: %w",
				string(localVarBody),
				err,
			)
		}

		return nil, &v, fmt.Errorf("error %+v", v)
	case _nethttp.StatusBadGateway,
		_nethttp.StatusServiceUnavailable,
		_nethttp.StatusGatewayTimeout,
		_nethttp.StatusRequestTimeout:
		return nil, nil, fmt.Errorf(
			"status code %d, response body %s: %w",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
			ErrRetriable,
		)
	default:
		return nil, nil, fmt.Errorf(
			"invalid status code %d, response body %s",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
		)
	}
}
