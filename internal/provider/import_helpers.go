// Copyright (c) i-am-smolli, CorentinPtrl.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLabPathAndID(importID string, resourceType string) (string, int64, error) {
	parts := strings.Split(importID, "|")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid import ID for %s: expected \"<lab_path>|<id>\", got %q", resourceType, importID)
	}

	labPath := strings.TrimSpace(parts[0])
	idPart := strings.TrimSpace(parts[1])
	if labPath == "" || idPart == "" {
		return "", 0, fmt.Errorf("invalid import ID for %s: lab_path and id must not be empty", resourceType)
	}

	id, err := strconv.ParseInt(idPart, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("invalid import ID for %s: id must be an integer, got %q", resourceType, idPart)
	}

	return labPath, id, nil
}
