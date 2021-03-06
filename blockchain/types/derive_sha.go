// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"encoding/binary"
	"github.com/cosmos/iavl"
	"github.com/idena-network/idena-go/common"
	db "github.com/tendermint/tm-db"
)

type DerivableList interface {
	Len() int
	GetBytes(i int) []byte
}

func DeriveSha(list DerivableList) common.Hash {
	tree, _ := iavl.NewMutableTree(db.NewMemDB(), 1024)
	for i := 0; i < list.Len(); i++ {
		key := make([]byte, 4)
		binary.LittleEndian.PutUint32(key, uint32(i))
		tree.Set(key, list.GetBytes(i))
	}
	var result common.Hash
	copy(result[:], tree.WorkingHash())
	return result
}
