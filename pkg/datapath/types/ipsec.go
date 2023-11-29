// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package types

import "context"

type NodeUpdater interface {
	UpdateLocalNode()
}

type IPsecKeyCustodian interface {
	AuthKeySize() int
	SPI() uint8
	StartBackgroundJobs(context.Context, NodeUpdater, NodeHandler) error
}