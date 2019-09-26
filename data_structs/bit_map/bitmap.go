package bit_map

import (
	"fmt"
	"math"
	"strings"
)

const BitSize = 8

type BitMap struct {
	bitArray []byte
	Size     uint64
}

func (bm *BitMap) Init() {
	if bm.Size > 1<<31 {
		panic("bit map Size over flow")
	}
	var r uint32
	if bm.Size <= 1 {
		r = 1
	} else {
		r = uint32(math.Ceil(float64(bm.Size) / BitSize))
	}
	bm.bitArray = make([]byte, r)
}

func (bm *BitMap) calIdx(i uint64) (idx uint64, pos uint64) {
	idx = i >> 3
	pos = i & (BitSize - 1)
	return
}

func (bm *BitMap) Set(i uint64) error {
	if i > bm.Size {
		return fmt.Errorf("bit map Size over flow, max Size=[%d], error position=[%d]", bm.Size, i)
	}
	idx, pos := bm.calIdx(i)
	bm.bitArray[idx] |= 1 << (pos - 1)
	return nil
}

func (bm *BitMap) Exists(i uint64) bool {
	if i > bm.Size {
		return false
	}
	idx, pos := bm.calIdx(i)
	return bm.bitArray[idx]&(1<<(pos-1)) == 1<<(pos-1)
}

func (bm *BitMap) String() string {
	strs := make([]string, 0, len(bm.bitArray))
	for _, b := range bm.bitArray {
		strs = append(strs, byte2String(b))
	}
	return strings.Join(strs, "-")
}

func byte2String(b byte) string {
	var magicBit uint32 = 1
	str := ""
	for i := 0; i < BitSize; i++ {
		u := magicBit << uint32(i)
		if uint32(b)&u == u {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}
