// Copyright (c) i-am-smolli, CorentinPtrl.
// SPDX-License-Identifier: MPL-2.0

package provider

import "testing"

func TestParseLabPathAndID(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantLabPath string
		wantID      int64
		wantErr     bool
	}{
		{
			name:        "valid import id",
			input:       "/test-lab.unl|42",
			wantLabPath: "/test-lab.unl",
			wantID:      42,
		},
		{
			name:    "missing separator",
			input:   "/test-lab.unl",
			wantErr: true,
		},
		{
			name:    "invalid id",
			input:   "/test-lab.unl|not-a-number",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			labPath, id, err := parseLabPathAndID(tc.input, "eveng_network")
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if labPath != tc.wantLabPath {
				t.Fatalf("unexpected lab path: got %q want %q", labPath, tc.wantLabPath)
			}
			if id != tc.wantID {
				t.Fatalf("unexpected id: got %d want %d", id, tc.wantID)
			}
		})
	}
}
