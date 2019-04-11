// Copyright 2019 Layer5.io
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

package consul

type supportedOperation struct {
	// a friendly name
	name string
	// the template file name
	templateName string
}

const (
	customOpCommand = "custom"
)

var supportedOps = map[string]supportedOperation{
	customOpCommand: {
		name: "Custom YAML",
	},
}