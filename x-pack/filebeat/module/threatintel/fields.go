// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package threatintel

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "threatintel", asset.ModuleFieldsPri, AssetThreatintel); err != nil {
		panic(err)
	}
}

// AssetThreatintel returns asset data.
// This is the base64 encoded zlib format compressed contents of module/threatintel.
func AssetThreatintel() string {
	return "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"
}
