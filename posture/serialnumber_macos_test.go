// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo && darwin && !ios

package posture

import (
	"fmt"
	"testing"

	"tailscale.com/types/logger"
	"tailscale.com/util/cibuild"
)

func TestGetSerialNumberMac(t *testing.T) {
	// Do not run this test on CI, it can only be ran on macOS
	// and we currenty only use Linux runners.
	if cibuild.On() {
		t.Skip()
	}

	sns, err := GetSerialNumbers(logger.Discard)
	if err != nil {
		t.Fatalf("failed to get serial number: %s", err)
	}

	if len(sns) != 1 {
		t.Errorf("expected list of one serial number, got %v", sns)
	}

	if len(sns[0]) <= 0 {
		t.Errorf("expected a serial number with more than zero characters, got %s", sns[0])
	}

	fmt.Printf("serials: %v\n", sns)
}
