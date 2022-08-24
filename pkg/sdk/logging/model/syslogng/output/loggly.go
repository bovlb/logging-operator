// Copyright © 2022 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output

import "github.com/banzaicloud/operator-tools/pkg/secret"

// +kubebuilder:object:generate=true
// Documentation: https://github.com/syslog-ng/syslog-ng/blob/master/scl/loggly/loggly.conf

type Loggly struct {
	SyslogOutput `json:",inline"`
	Tag          string         `json:"tag,omitempty"`
	Token        *secret.Secret `json:"token"`
}