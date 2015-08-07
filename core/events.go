// Copyright 2014 The go-expanse Authors
// This file is part of the go-expanse library.
//
// The go-expanse library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-expanse library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-expanse library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"math/big"

	"github.com/expanse-project/go-expanse/common"
	"github.com/expanse-project/go-expanse/core/state"
	"github.com/expanse-project/go-expanse/core/types"
)

// TxPreEvent is posted when a transaction enters the transaction pool.
type TxPreEvent struct{ Tx *types.Transaction }

// TxPostEvent is posted when a transaction has been processed.
type TxPostEvent struct{ Tx *types.Transaction }

// NewBlockEvent is posted when a block has been imported.
type NewBlockEvent struct{ Block *types.Block }

// NewMinedBlockEvent is posted when a block has been imported.
type NewMinedBlockEvent struct{ Block *types.Block }

// ChainSplit is posted when a new head is detected
type ChainSplitEvent struct {
	Block *types.Block
	Logs  state.Logs
}

type ChainEvent struct {
	Block *types.Block
	Hash  common.Hash
	Logs  state.Logs
}

type ChainSideEvent struct {
	Block *types.Block
	Logs  state.Logs
}

type PendingBlockEvent struct {
	Block *types.Block
	Logs  state.Logs
}

type ChainUncleEvent struct {
	Block *types.Block
}

type ChainHeadEvent struct{ Block *types.Block }

type GasPriceChanged struct{ Price *big.Int }

// Mining operation events
type StartMining struct{}
type TopMining struct{}
