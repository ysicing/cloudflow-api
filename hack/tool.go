// Copyright (c) 2022 ysicing(i@ysicing.me) All rights reserved.
// Use of this source code is covered by the following licenses:
// (1) Affero General Public License 3.0 (AGPL 3.0)
// license that can be found in the LICENSE file.

//go:build tools
// +build tools

// This package imports things required by build scripts, to force `go mod` to see them as dependencies

package tools

import _ "k8s.io/code-generator"
