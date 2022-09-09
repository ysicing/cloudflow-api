// Copyright (c) 2022 ysicing(i@ysicing.me) All rights reserved.
// Use of this source code is covered by the following licenses:
// (1) Affero General Public License 3.0 (AGPL 3.0)
// license that can be found in the LICENSE file.

package apis

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
