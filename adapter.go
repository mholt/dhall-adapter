// Copyright 2015 Matthew Holt and The Caddy Authors
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

package dhalladapter

import (
	"encoding/json"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/philandstuff/dhall-golang/v6"
)

func init() {
	caddyconfig.RegisterAdapter("dhall", Adapter{})
}

// Adapter adapts Dhall to Caddy JSON.
type Adapter struct{}

// Adapt converts the Dhall config in body to Caddy JSON.
func (a Adapter) Adapt(body []byte, options map[string]interface{}) (result []byte, warnings []caddyconfig.Warning, err error) {
	var decoded interface{}
	err = dhall.Unmarshal(body, &decoded)
	if err != nil {
		return
	}
	result, err = json.Marshal(decoded)
	return
}

// Interface guard
var _ caddyconfig.Adapter = (*Adapter)(nil)
