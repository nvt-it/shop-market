package block

import (
	"time"
)

type OrderBlock struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *OrderBlock) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *OrderBlock {
	block := &OrderBlock{
		time.Now.Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
	}

	block.SetHash()

	return block
}
