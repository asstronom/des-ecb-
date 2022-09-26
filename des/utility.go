package des

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Sets the bit at pos in the integer n.
func setBit(n byte, pos byte) byte {
	n |= (1 << (7 - pos))
	return n
}

func getBit(n byte, pos byte) byte {
	return (n & byte(math.Pow(2, float64(7-pos)))) >> (7 - pos)
}

// Clears the bit at pos in n.
func clearBit(n byte, pos byte) byte {
	mask := ^(byte(1) << (7 - pos))
	n &= mask
	return n
}

func PermutateBlock(block []byte, perm [64]byte) []byte {
	result := make([]byte, 8)
	for i := byte(0); i < 64; i++ {
		cur := getBit(block[perm[i]/8], perm[i]%8)
		if cur == 1 {
			result[i/8] = setBit(result[i/8], i%8)
		} else {
			result[i/8] = clearBit(result[i/8], i%8)
		}
	}
	return result
}

func Extend(block []byte) []byte {
	result := make([]byte, 6)
	for i := byte(0); i < 48; i++ {
		cur := getBit(block[eTable[i]/8], eTable[i]%8)
		if cur == 1 {
			result[i/8] = setBit(result[i/8], i%8)
		} else {
			result[i/8] = clearBit(result[i/8], i%8)
		}
	}
	return result
}

func f(block []byte, subkey [6]byte) []byte {
	extBlock := Extend(block)
	blockxor := binary.BigEndian.Uint64(append(extBlock, 0, 0))
	kxor := binary.BigEndian.Uint64(append(subkey[:], 0, 0))
	xorres := make([]byte, 8)
	binary.BigEndian.PutUint64(xorres, blockxor^kxor)
	xorres = xorres[:6]

	sb := make([]byte, 8)

	for i := byte(0); i < 8; i++ {
		idx := i * 6
		row := getBit(xorres[idx/8], idx%8)*2 + getBit(xorres[(idx+5)/8], (idx+5)%8)
		column := getBit(xorres[(idx+1)/8], (idx+1)%8)*8 + getBit(xorres[(idx+2)/8], (idx+2)%8)*4 +
			getBit(xorres[(idx+3)/8], (idx+3)%8)*2 + getBit(xorres[(idx+4)/8], (idx+4)%8)
		sb[i] = s[i][row][column]
	}

	for i := 0; i < 8; i += 2 {
		sb[i/2] = (sb[i] << 4) + sb[i+1]
	}
	sb = sb[:4]
	fRes := make([]byte, 4)
	for i := byte(0); i < 32; i++ {
		idx := pPerm[i] - 1
		cur := getBit(sb[idx/8], idx%8)
		if cur == 1 {
			fRes[i/8] = setBit(fRes[i/8], i%8)
		} else {
			fRes[i/8] = clearBit(fRes[i/8], i%8)
		}
	}
	return fRes
}

func BlockToHex(block []byte) string {
	result := ""
	for i := range block {
		result += fmt.Sprintf("%02x", block[i])
	}
	return result
}
