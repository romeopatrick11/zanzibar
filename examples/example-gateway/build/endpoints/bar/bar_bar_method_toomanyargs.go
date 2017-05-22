// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bar

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	clientsFooFoo "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/foo/foo"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
)

// TooManyArgsHandler is the handler for "/bar/too-many-args-path"
type TooManyArgsHandler struct {
	Clients *clients.Clients
}

// NewTooManyArgsEndpoint creates a handler
func NewTooManyArgsEndpoint(
	gateway *zanzibar.Gateway,
) *TooManyArgsHandler {
	return &TooManyArgsHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/bar/too-many-args-path".
func (handler *TooManyArgsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	if !req.CheckHeaders([]string{"x-uuid", "x-token"}) {
		return
	}
	var requestBody endpointsBarBar.Bar_TooManyArgs_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := TooManyArgsEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBarBar.BarException:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(sindelar): implement check headers on response

	res.WriteJSON(200, cliRespHeaders, response)
}

// TooManyArgsEndpoint calls thrift client Bar.TooManyArgs
type TooManyArgsEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w TooManyArgsEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_TooManyArgs_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToTooManyArgsClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Token")
	if ok {
		clientHeaders["X-Token"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}

	clientRespBody, cliRespHeaders, err := w.Clients.Bar.TooManyArgs(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBarBar.BarException:
			serverErr := convertTooManyArgsBarException(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	resHeaders.Set("X-Token", cliRespHeaders["X-Token"])
	resHeaders.Set("X-Uuid", cliRespHeaders["X-Uuid"])

	response := convertTooManyArgsClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToTooManyArgsClientRequest(in *endpointsBarBar.Bar_TooManyArgs_Args) *clientsBarBar.Bar_TooManyArgs_Args {
	out := &clientsBarBar.Bar_TooManyArgs_Args{}

	if in.Request != nil {
		out.Request = &clientsBarBar.BarRequest{}
		out.Request.StringField = string(in.Request.StringField)
		out.Request.BoolField = bool(in.Request.BoolField)
	} else {
		out.Request = nil
	}
	if in.Foo != nil {
		out.Foo = &clientsFooFoo.FooStruct{}
		out.Foo.FooString = string(in.Foo.FooString)
		out.Foo.FooI32 = (*int32)(in.Foo.FooI32)
		out.Foo.FooI16 = (*int16)(in.Foo.FooI16)
		out.Foo.FooDouble = (*float64)(in.Foo.FooDouble)
		out.Foo.FooBool = (*bool)(in.Foo.FooBool)
		out.Foo.FooMap = make(map[string]string, len(in.Foo.FooMap))
		for key, value := range in.Foo.FooMap {
			out.Foo.FooMap[key] = string(value)
		}
	} else {
		out.Foo = nil
	}

	return out
}

func convertTooManyArgsBarException(
	clientError *clientsBarBar.BarException,
) *endpointsBarBar.BarException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBarBar.BarException{}
	return serverError
}

func convertTooManyArgsClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[string]int32, len(in.MapIntWithRange))
	for key, value := range in.MapIntWithRange {
		out.MapIntWithRange[key] = int32(value)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key, value := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key] = int32(value)
	}

	return out
}
