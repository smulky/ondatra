// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ondatra

import (
	"github.com/openconfig/ondatra/binding"
)

// ATEDevice is an automated test equipment.
type ATEDevice struct {
	*Device
	otg *OTG
}

// OTG returns a handle to the OTG API.
func (a *ATEDevice) OTG() *OTG {
	if a.otg == nil {
		a.otg = &OTG{devID: a.ID(), ate: a.res.(binding.ATE)}
	}
	return a.otg
}

// Topology returns a handle to the topology API.
func (a *ATEDevice) Topology() *Topology {
	return &Topology{a.res.(binding.ATE)}
}

// Traffic returns a handle to the traffic API.
func (a *ATEDevice) Traffic() *Traffic {
	return &Traffic{a.res.(binding.ATE)}
}
