// Copyright (c) 2022 ysicing(i@ysicing.me) All rights reserved.
// Use of this source code is covered by the following licenses:
// (1) Affero General Public License 3.0 (AGPL 3.0)
// license that can be found in the LICENSE file.

package apis

import "github.com/ysicing/cloudflow-api/apps/v1beta1"

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1beta1.SchemeBuilder.AddToScheme)
}
