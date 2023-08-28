// Code generated by goa v3.12.4, DO NOT EDIT.
//
// tools client HTTP transport
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the tools service endpoint HTTP clients.
type Client struct {
	// Available Doer is the HTTP client used to make requests to the available
	// endpoint.
	AvailableDoer goahttp.Doer

	// Installed Doer is the HTTP client used to make requests to the installed
	// endpoint.
	InstalledDoer goahttp.Doer

	// Install Doer is the HTTP client used to make requests to the install
	// endpoint.
	InstallDoer goahttp.Doer

	// Remove Doer is the HTTP client used to make requests to the remove endpoint.
	RemoveDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the tools service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AvailableDoer:       doer,
		InstalledDoer:       doer,
		InstallDoer:         doer,
		RemoveDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Available returns an endpoint that makes HTTP requests to the tools service
// available server.
func (c *Client) Available() goa.Endpoint {
	var (
		decodeResponse = DecodeAvailableResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAvailableRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AvailableDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tools", "available", err)
		}
		return decodeResponse(resp)
	}
}

// Installed returns an endpoint that makes HTTP requests to the tools service
// installed server.
func (c *Client) Installed() goa.Endpoint {
	var (
		decodeResponse = DecodeInstalledResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildInstalledRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.InstalledDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tools", "installed", err)
		}
		return decodeResponse(resp)
	}
}

// Install returns an endpoint that makes HTTP requests to the tools service
// install server.
func (c *Client) Install() goa.Endpoint {
	var (
		encodeRequest  = EncodeInstallRequest(c.encoder)
		decodeResponse = DecodeInstallResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildInstallRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.InstallDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tools", "install", err)
		}
		return decodeResponse(resp)
	}
}

// Remove returns an endpoint that makes HTTP requests to the tools service
// remove server.
func (c *Client) Remove() goa.Endpoint {
	var (
		encodeRequest  = EncodeRemoveRequest(c.encoder)
		decodeResponse = DecodeRemoveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRemoveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tools", "remove", err)
		}
		return decodeResponse(resp)
	}
}
