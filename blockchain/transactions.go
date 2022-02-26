package blockchain

import (
	"time"

	"github.com/Nhahan/blockchain/utils"
)

const (
	minerReaward int = 50
)

type Tx struct {
	Id        string
	Timestamp int64
	TxIns     []*TxIn
	TxOuts    []*TxOut
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner  string
	Amount int
}

type TxOut struct {
	Owner  string
	Amount int
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReaward},
	}
	txOuts := []*TxOut{
		{address, minerReaward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: time.Now().Unix(),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}
