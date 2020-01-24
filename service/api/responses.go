// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mainflux/mainflux"
)

var (
	_ mainflux.Response = (*removeRes)(nil)
	_ mainflux.Response = (*licenseRes)(nil)
)

type removeRes struct{}

func (res removeRes) Code() int {
	return http.StatusNoContent
}

func (res removeRes) Headers() map[string]string {
	return map[string]string{}
}

func (res removeRes) Empty() bool {
	return true
}

type successRes struct{}

func (res successRes) Code() int {
	return http.StatusOK
}

func (res successRes) Headers() map[string]string {
	return map[string]string{}
}

func (res successRes) Empty() bool {
	return true
}

type licenseRes struct {
	created   bool
	ID        string                 `json:"id,omitempty"`
	Issuer    string                 `json:"issuer,omitempty"`
	DeviceID  string                 `json:"device_id,omitempty"`
	Active    bool                   `json:"active"`
	CreatedAt *time.Time             `json:"created_at,omitempty"`
	ExpiresAt *time.Time             `json:"expires_at,omitempty"`
	UpdatedAt *time.Time             `json:"updated_at,omitempty"`
	UpdatedBy string                 `json:"updated_by,omitempty"`
	Services  []string               `json:"services,omitempty"`
	Plan      map[string]interface{} `json:"plan,omitempty"`
}

type fetchRes struct {
	licenseRes
	Signature []byte `json:"signature,omitempty"`
}

type vewRes struct {
	licenseRes
	Key string `json:"key,omitempty"`
}

func (res licenseRes) Code() int {
	if res.created {
		return http.StatusCreated
	}

	return http.StatusOK
}

func (res licenseRes) Headers() map[string]string {
	ret := make(map[string]string)
	if res.created {
		ret["Location"] = fmt.Sprintf("/licenses/%s", res.ID)
	}

	return ret
}

func (res licenseRes) Empty() bool {
	return res.created
}

type errorRes struct {
	Err string `json:"error"`
}