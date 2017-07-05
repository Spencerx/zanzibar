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

// Package bazClient is generated code used to make or handle TChannel calls using Thrift.
package bazClient

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/uber/tchannel-go"
	"github.com/uber/zanzibar/runtime"

	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
)

// Client defines baz client interface.
type Client interface {
	EchoBinary(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoBinary_Args,
	) ([]byte, map[string]string, error)
	EchoBool(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoBool_Args,
	) (bool, map[string]string, error)
	EchoDouble(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoDouble_Args,
	) (float64, map[string]string, error)
	EchoEnum(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoEnum_Args,
	) (clientsBazBaz.Fruit, map[string]string, error)
	EchoI16(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoI16_Args,
	) (int16, map[string]string, error)
	EchoI32(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoI32_Args,
	) (int32, map[string]string, error)
	EchoI64(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoI64_Args,
	) (int64, map[string]string, error)
	EchoI8(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoI8_Args,
	) (int8, map[string]string, error)
	EchoList(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoList_Args,
	) ([]string, map[string]string, error)
	EchoMap(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoMap_Args,
	) (map[clientsBazBase.UUID]*clientsBazBase.BazResponse, map[string]string, error)
	EchoSet(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoSet_Args,
	) (map[string]struct{}, map[string]string, error)
	EchoString(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoString_Args,
	) (string, map[string]string, error)
	EchoUUID(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoUUID_Args,
	) (clientsBazBase.UUID, map[string]string, error)
	EchoUUIDList(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SecondService_EchoUUIDList_Args,
	) ([]clientsBazBase.UUID, map[string]string, error)
	Call(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Call_Args,
	) (map[string]string, error)
	Compare(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Compare_Args,
	) (*clientsBazBase.BazResponse, map[string]string, error)
	Ping(
		ctx context.Context,
		reqHeaders map[string]string,
	) (*clientsBazBase.BazResponse, map[string]string, error)
	DeliberateDiffNoop(
		ctx context.Context,
		reqHeaders map[string]string,
	) (map[string]string, error)
}

// NewClient returns a new TChannel client for service baz.
func NewClient(gateway *zanzibar.Gateway) Client {
	serviceName := gateway.Config.MustGetString("clients.baz.serviceName")
	sc := gateway.Channel.GetSubChannel(serviceName, tchannel.Isolated)

	ip := gateway.Config.MustGetString("clients.baz.ip")
	port := gateway.Config.MustGetInt("clients.baz.port")
	sc.Peers().Add(ip + ":" + strconv.Itoa(int(port)))

	timeout := time.Millisecond * time.Duration(
		gateway.Config.MustGetInt("clients.baz.timeout"),
	)
	timeoutPerAttempt := time.Millisecond * time.Duration(
		gateway.Config.MustGetInt("clients.baz.timeoutPerAttempt"),
	)

	client := zanzibar.NewTChannelClient(gateway.Channel,
		&zanzibar.TChannelClientOption{
			ServiceName:       serviceName,
			Timeout:           timeout,
			TimeoutPerAttempt: timeoutPerAttempt,
		},
	)

	return &bazClient{
		client: client,
	}
}

// bazClient is the TChannel client for downstream service.
type bazClient struct {
	client zanzibar.TChannelClient
}

// EchoBinary is a client RPC call for method "SecondService::EchoBinary"
func (c *bazClient) EchoBinary(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoBinary_Args,
) ([]byte, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoBinary_Result
	var resp []byte

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoBinary", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoBinary")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoBinary_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoBool is a client RPC call for method "SecondService::EchoBool"
func (c *bazClient) EchoBool(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoBool_Args,
) (bool, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoBool_Result
	var resp bool

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoBool", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoBool")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoBool_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoDouble is a client RPC call for method "SecondService::EchoDouble"
func (c *bazClient) EchoDouble(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoDouble_Args,
) (float64, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoDouble_Result
	var resp float64

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoDouble", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoDouble")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoDouble_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoEnum is a client RPC call for method "SecondService::EchoEnum"
func (c *bazClient) EchoEnum(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoEnum_Args,
) (clientsBazBaz.Fruit, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoEnum_Result
	var resp clientsBazBaz.Fruit

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoEnum", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoEnum")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoEnum_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoI16 is a client RPC call for method "SecondService::EchoI16"
func (c *bazClient) EchoI16(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI16_Args,
) (int16, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoI16_Result
	var resp int16

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoI16", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoI16")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoI16_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoI32 is a client RPC call for method "SecondService::EchoI32"
func (c *bazClient) EchoI32(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI32_Args,
) (int32, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoI32_Result
	var resp int32

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoI32", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoI32")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoI32_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoI64 is a client RPC call for method "SecondService::EchoI64"
func (c *bazClient) EchoI64(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI64_Args,
) (int64, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoI64_Result
	var resp int64

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoI64", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoI64")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoI64_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoI8 is a client RPC call for method "SecondService::EchoI8"
func (c *bazClient) EchoI8(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoI8_Args,
) (int8, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoI8_Result
	var resp int8

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoI8", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoI8")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoI8_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoList is a client RPC call for method "SecondService::EchoList"
func (c *bazClient) EchoList(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoList_Args,
) ([]string, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoList_Result
	var resp []string

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoList", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoList")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoList_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoMap is a client RPC call for method "SecondService::EchoMap"
func (c *bazClient) EchoMap(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoMap_Args,
) (map[clientsBazBase.UUID]*clientsBazBase.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoMap_Result
	var resp map[clientsBazBase.UUID]*clientsBazBase.BazResponse

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoMap", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoMap")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoMap_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoSet is a client RPC call for method "SecondService::EchoSet"
func (c *bazClient) EchoSet(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoSet_Args,
) (map[string]struct{}, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoSet_Result
	var resp map[string]struct{}

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoSet", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoSet")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoSet_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoString is a client RPC call for method "SecondService::EchoString"
func (c *bazClient) EchoString(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoString_Args,
) (string, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoString_Result
	var resp string

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoString", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoString")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoString_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoUUID is a client RPC call for method "SecondService::EchoUUID"
func (c *bazClient) EchoUUID(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoUUID_Args,
) (clientsBazBase.UUID, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoUUID_Result
	var resp clientsBazBase.UUID

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoUUID", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoUUID")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoUUID_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// EchoUUIDList is a client RPC call for method "SecondService::EchoUUIDList"
func (c *bazClient) EchoUUIDList(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_EchoUUIDList_Args,
) ([]clientsBazBase.UUID, map[string]string, error) {
	var result clientsBazBaz.SecondService_EchoUUIDList_Result
	var resp []clientsBazBase.UUID

	success, respHeaders, err := c.client.Call(
		ctx, "SecondService", "EchoUUIDList", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for EchoUUIDList")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SecondService_EchoUUIDList_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// Call is a client RPC call for method "SimpleService::Call"
func (c *bazClient) Call(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Call_Args,
) (map[string]string, error) {
	var result clientsBazBaz.SimpleService_Call_Result

	success, respHeaders, err := c.client.Call(
		ctx, "SimpleService", "Call", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for Call")
		}
	}
	if err != nil {
		return nil, err
	}

	return respHeaders, err
}

// Compare is a client RPC call for method "SimpleService::Compare"
func (c *bazClient) Compare(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Compare_Args,
) (*clientsBazBase.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SimpleService_Compare_Result
	var resp *clientsBazBase.BazResponse

	success, respHeaders, err := c.client.Call(
		ctx, "SimpleService", "Compare", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		case result.OtherAuthErr != nil:
			err = result.OtherAuthErr
		default:
			err = errors.New("bazClient received no result or unknown exception for Compare")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_Compare_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// Ping is a client RPC call for method "SimpleService::Ping"
func (c *bazClient) Ping(
	ctx context.Context,
	reqHeaders map[string]string,
) (*clientsBazBase.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SimpleService_Ping_Result
	var resp *clientsBazBase.BazResponse

	args := &clientsBazBaz.SimpleService_Ping_Args{}
	success, respHeaders, err := c.client.Call(
		ctx, "SimpleService", "Ping", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("bazClient received no result or unknown exception for Ping")
		}
	}
	if err != nil {
		return resp, nil, err
	}

	resp, err = clientsBazBaz.SimpleService_Ping_Helper.UnwrapResponse(&result)
	return resp, respHeaders, err
}

// DeliberateDiffNoop is a client RPC call for method "SimpleService::SillyNoop"
func (c *bazClient) DeliberateDiffNoop(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error) {
	var result clientsBazBaz.SimpleService_SillyNoop_Result

	args := &clientsBazBaz.SimpleService_SillyNoop_Args{}
	success, respHeaders, err := c.client.Call(
		ctx, "SimpleService", "SillyNoop", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.AuthErr != nil:
			err = result.AuthErr
		case result.ServerErr != nil:
			err = result.ServerErr
		default:
			err = errors.New("bazClient received no result or unknown exception for SillyNoop")
		}
	}
	if err != nil {
		return nil, err
	}

	return respHeaders, err
}
