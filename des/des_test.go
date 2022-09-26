package des

import (
	"encoding/binary"
	"fmt"
	"strings"
	"testing"
)

var (
	correctKeyPerm = "11110000110011001010101011110101010101100110011110001111"
	correctKeys    = [16]string{
		"000110110000001011101111111111000111000001110010",
		"011110011010111011011001110110111100100111100101",
		"010101011111110010001010010000101100111110011001",
		"011100101010110111010110110110110011010100011101",
		"011111001110110000000111111010110101001110101000",
		"011000111010010100111110010100000111101100101111",
		"111011001000010010110111111101100001100010111100",
		"111101111000101000111010110000010011101111111011",
		"111000001101101111101011111011011110011110000001",
		"101100011111001101000111101110100100011001001111",
		"001000010101111111010011110111101101001110000110",
		"011101010111000111110101100101000110011111101001",
		"100101111100010111010001111110101011101001000001",
		"010111110100001110110111111100101110011100111010",
		"101111111001000110001101001111010011111100001010",
		"110010110011110110001011000011100001011111110101",
	}
	correctC0   = "1111000011001100101010101111"
	correctD0   = "0101010101100110011110001111"
	correctC1D1 = "11100001100110010101010111111010101011001100111100011110"
)
func TestSetBit(t *testing.T) {
	var b byte = 0
	b = setBit(b, 6)
	if b != 2 {
		t.Errorf("%d != 2, %08b", b, b)
	}
	b = setBit(b, 4)
	if b != 10 {
		t.Errorf("%d != 10, %08b", b, b)
	}
	b = setBit(b, 5)
	if b != 14 {
		t.Errorf("%d != 14, %08b", b, b)
	}
	b = setBit(b, 7)
	if b != 15 {
		t.Errorf("%d != 15, %08b", b, b)
	}
}

func TestGetBit(t *testing.T) {
	var b byte = 0b01101111
	var b1 byte = 0b01101000
	res := getBit(b, 1)
	//t.Logf("%08b", byte(math.Pow(2, 7-1)))
	//t.Logf("%08b", 0b01101111&byte(math.Pow(2, 7-1)))
	if res != 1 {
		t.Errorf("%d != 1, %08b", res, res)
	}
	res = getBit(b1, 0)
	if res != 0 {
		t.Errorf("%d != 0, %08b", res, res)
	}
}

func TestPermutateBlock(t *testing.T) {
	block := []byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	result := PermutateBlock(block, initialPerm)
	correctString := "1100110000000000110011001111111111110000101010101111000010101010"
	resultString := ""
	for _, v := range result {
		resultString += fmt.Sprintf("%08b", v)
	}
	if resultString != correctString {
		t.Errorf("wrong permutation\ncorrect: %s\nwrong: %s\n", correctString, resultString)
	}
}

func TestExtend(t *testing.T) {
	block := []byte{0b11110000, 0b10101010, 0b11110000, 0b10101010}
	result := Extend(block)
	correctString := "011110100001010101010101011110100001010101010101"
	resultString := ""
	for _, v := range result {
		resultString += fmt.Sprintf("%08b", v)
	}
	if resultString != correctString {
		t.Errorf("wrong permutation\ncorrect: %s\nwrong: %s\n", correctString, resultString)
	}
}

func TestF(t *testing.T) {
	block := []byte{0b11110000, 0b10101010, 0b11110000, 0b10101010}
	key := [6]byte{0b00011011, 0b00000010, 0b11101111, 0b11111100, 0b01110000, 0b01110010}
	f(block, key)
}

func TestEncrypt(t *testing.T) {
	block := []byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	var keyHex uint64 = 0x133457799BBCDFF1
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, keyHex)
	des, err := NewDES(key)
	if err != nil {
		t.Errorf("error making des: %s", err)
	}
	result, err := des.Encrypt(block)
	if err != nil {
		t.Errorf("error encrypting: %s", err)
	}
	t.Logf(string(result))
}

func TestYes(t *testing.T) {
	for i := 0; i < len(eTable); i++ {
		eTable[i]--
	}
	fmt.Print("{")
	for i := 0; i < len(eTable); i++ {
		fmt.Printf("%d, ", eTable[i])
		if i%8 == 0 {
			fmt.Print("\n")
		}
	}
	fmt.Print("}")
}

func TestNo(t *testing.T) {
	input1 := "13  2   8  4   6 15  11  1  10  9   3 14   5  0  12  7"
	input2 := "1 15  13  8  10  3   7  4  12  5   6 11   0 14   9  2"
	input3 := " 7 11   4  1   9 12  14  2   0  6  10 13  15  3   5  8"
	input4 := "2  1  14  7   4 10   8 13  15 12   9  0   3  5   6 11"
	inputSplit1 := strings.Split(input1, " ")
	inputSplit2 := strings.Split(input2, " ")
	inputSplit3 := strings.Split(input3, " ")
	inputSplit4 := strings.Split(input4, " ")
	temp := []string{}
	for _, v := range inputSplit1 {
		v = strings.TrimSpace(v)
		if v != "" {
			temp = append(temp, v)
		}
	}
	inputSplit1 = temp
	temp = []string{}
	for _, v := range inputSplit2 {
		v = strings.TrimSpace(v)
		if v != "" {
			temp = append(temp, v)
		}
	}
	inputSplit2 = temp
	temp = []string{}
	for _, v := range inputSplit3 {
		v = strings.TrimSpace(v)
		if v != "" {
			temp = append(temp, v)
		}
	}
	inputSplit3 = temp
	temp = []string{}
	for _, v := range inputSplit4 {
		v = strings.TrimSpace(v)
		if v != "" {
			temp = append(temp, v)
		}
	}
	inputSplit4 = temp
	// inputSplit = temp
	fmt.Println(len(inputSplit1))
	fmt.Println(len(inputSplit2))
	fmt.Println(len(inputSplit3))
	fmt.Println(len(inputSplit4))
	fmt.Print("{\n")
	fmt.Print("{")
	for _, v := range inputSplit1 {
		fmt.Printf("%s, ", v)
	}
	fmt.Print("},\n")
	fmt.Print("{")
	for _, v := range inputSplit2 {
		fmt.Printf("%s, ", v)
	}
	fmt.Print("},\n")
	fmt.Print("{")
	for _, v := range inputSplit3 {
		fmt.Printf("%s, ", v)
	}
	fmt.Print("},\n")
	fmt.Print("{")
	for _, v := range inputSplit4 {
		fmt.Printf("%s, ", v)
	}
	fmt.Print("},\n")
	fmt.Print("}")
}

func TestRotateLeft(t *testing.T) {
	//b1 := uint32(0b10100000000000000000000001010000)
	b2 := uint32(0b00000101010101100110011110001111)
	//b := byte(0b00010100)
	b := ((b2 << 1) | (b2 >> (28 - 1))) & 0x0FFFFFFF
	//t.Logf("%08b\n", b)
	//b = (b >> 3) | ((b << (8 - 3)) & 255)
	t.Logf("%028b\n", b)
	t.Logf("%s\n", "1010101011001100111100011110")
}
