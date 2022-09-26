package des

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

func NewDES(key []byte) (*DES, error) {
	if len(key) != 8 {
		return nil, fmt.Errorf("key size is not 8 but %d, key:`%s`", len(key), string(key))
	}
	des := DES{}
	k := [8]byte{}
	for i := 0; i < 8; i++ {
		k[i] = key[i]
	}
	des.gensubkeys(k)
	return &des, nil
}

type DES struct {
	subkeys [16][6]byte
}

func (des *DES) gensubkeys(key [8]byte) {

	c0d0 := [7]byte{}
	result := ""
	for i := 0; i < len(c0d0Perm); i++ {
		cur := getBit(key[c0d0Perm[i]/8], c0d0Perm[i]%8)
		result += strconv.Itoa(int(cur))
		if cur == 1 {
			c0d0[i/8] = setBit(c0d0[i/8], byte(i%8))
		} else {
			c0d0[i/8] = clearBit(c0d0[i/8], byte(i%8))
		}
	}

	c0 := binary.BigEndian.Uint32(c0d0[0:4]) & 0xFFFFFFF0
	d0 := (binary.BigEndian.Uint32(c0d0[3:])) & 0x0FFFFFFF

	ci := c0
	di := d0
	for i := 0; i < 16; i++ {
		if i == 0 || i == 1 || i == 8 || i == 15 {
			ci = (ci << 1) | (ci>>(28-1))&0xFFFFFFF0
			di = ((di << 1) | (di >> (28 - 1))) & 0x0FFFFFFF
		} else {
			ci = (ci << 2) | (ci>>(28-2))&0xFFFFFFF0
			di = ((di << 2) | (di >> (28 - 2))) & 0x0FFFFFFF
		}
		cibytes := make([]byte, 4)
		dibytes := make([]byte, 4)
		binary.BigEndian.PutUint32(cibytes, ci)
		binary.BigEndian.PutUint32(dibytes, di)
		cidi := [7]byte{}
		for j := 0; j < 3; j++ {
			cidi[j] = cibytes[j]
		}
		cidi[3] = (cibytes[3] & 0xF0) | (dibytes[0] & 0x0F)
		for j := 1; j < 4; j++ {
			cidi[j+3] = dibytes[j]
		}
		for j := 0; j < 48; j++ {
			cur := getBit(cidi[kiPerm[j]/8], kiPerm[j]%8)
			if cur == 1 {
				des.subkeys[i][j/8] = setBit(des.subkeys[i][j/8], byte(j%8))
			} else {
				des.subkeys[i][j/8] = clearBit(des.subkeys[i][j/8], byte(j%8))
			}
		}
	}

}

func (d *DES) Encrypt(block []byte) ([]byte, error) {
	if len(block) != 8 {
		return nil, fmt.Errorf("block size != 8 but %d", len(block))
	}
	ip := PermutateBlock(block, initialPerm)
	if debug {
		fmt.Printf("After initial permutation: %s\n", BlockToHex(ip))
	}
	lprev := ip[:4]
	rprev := ip[4:]
	if debug {
		fmt.Printf("After splitting: L0=%s R0=%s\n", BlockToHex(lprev), BlockToHex(rprev))
	}
	for i := 0; i < 16; i++ {
		li := rprev
		ri := make([]byte, 4)
		binary.BigEndian.PutUint32(ri, binary.BigEndian.Uint32(lprev)^binary.BigEndian.Uint32(f(rprev, d.subkeys[i])))
		lprev = li
		rprev = ri
		if debug {
			fmt.Printf("Round %d %s %s %s\n", i, BlockToHex(lprev), BlockToHex(rprev), BlockToHex(d.subkeys[i][:]))
		}
	}
	r16l16 := append(rprev, lprev...)
	result := make([]byte, 8)
	for i := byte(0); i < 64; i++ {
		idx := ipinverse[i] - 1
		cur := getBit(r16l16[idx/8], idx%8)
		if cur == 1 {
			result[i/8] = setBit(result[i/8], i%8)
		} else {
			result[i/8] = clearBit(result[i/8], i%8)
		}
	}
	return result, nil
}

func (d *DES) Decrypt(block []byte) ([]byte, error) {
	if len(block) != 8 {
		return nil, fmt.Errorf("block size != 8 but %d", len(block))
	}
	ip := PermutateBlock(block, initialPerm)
	if debug {
		fmt.Printf("After initial permutation: %s\n", BlockToHex(ip))
	}
	lprev := ip[:4]
	rprev := ip[4:]
	if debug {
		fmt.Printf("After splitting: L0=%s R0=%s\n", BlockToHex(lprev), BlockToHex(rprev))
	}
	for i := 15; i >= 0; i-- {
		li := rprev
		ri := make([]byte, 4)
		binary.BigEndian.PutUint32(ri, binary.BigEndian.Uint32(lprev)^binary.BigEndian.Uint32(f(rprev, d.subkeys[i])))
		lprev = li
		rprev = ri
		if debug {
			fmt.Printf("Round %d %s %s %s\n", i, BlockToHex(lprev), BlockToHex(rprev), BlockToHex(d.subkeys[i][:]))
		}
	}
	r16l16 := append(rprev, lprev...)
	result := make([]byte, 8)
	for i := byte(0); i < 64; i++ {
		idx := ipinverse[i] - 1
		cur := getBit(r16l16[idx/8], idx%8)
		if cur == 1 {
			result[i/8] = setBit(result[i/8], i%8)
		} else {
			result[i/8] = clearBit(result[i/8], i%8)
		}
	}
	return result, nil
}
