package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/asstronom/des-ecb-/des"
)



type Cipher interface {
	Encrypt(block []byte) ([]byte, error)
	Decrypt(block []byte) ([]byte, error)
}

var isDecrypt bool
var isTriple bool
var inputFile string
var outputFile string
var key string

func main() {
	flag.BoolVar(&isTriple, "triple", false, "sets encryption method to TripleDES")
	flag.BoolVar(&isDecrypt, "decrypt", false, "sets mode to decrypt")
	flag.StringVar(&inputFile, "i", "", "sets input file")
	flag.StringVar(&outputFile, "o", "", "sets output file")
	flag.StringVar(&key, "key", "", "sets key (used only for `des` mode)")
	flag.Parse()

	var text bytes.Buffer
	if inputFile == "" {
		words := flag.Args()
		for i := range words {
			text.WriteString(words[i])
		}
	} else {
		bytes, err := os.ReadFile(inputFile)
		//fmt.Println(string(bytes))
		if err != nil {
			log.Fatalln("error opening file", err)
		}
		text.Write(bytes)
	}

	var cipher Cipher

	if isTriple {
		fmt.Print("input keys separated by whitespaces:\n")
		var k1, k2, k3 string
		_, err := fmt.Scanf("%s %s %s\n", &k1, &k2, &k3)
		if err != nil {
			log.Fatalf("error scanning key: %s\n", err)
		}
		cipher, err = des.NewTripleDES([]byte(k1), []byte(k2), []byte(k3))
		if err != nil {
			log.Fatalf("error making tripleDes: %s\n", err)
		}
	} else {
		var err error
		cipher, err = des.NewDES([]byte(key))
		if err != nil {
			log.Fatalf("error making des: %s", err)
		}
	}
	result := make([]byte, 0, text.Len())
	text.Write(make([]byte, text.Len()%8))
	fmt.Println(text.Len(), text.Len()/8)
	for text.Len() != 0 {
		curBlock := make([]byte, 8)
		_, err := text.Read(curBlock)
		if err != nil {
			log.Fatalf("error reading bytes: %s", err)
		}
		if isDecrypt {
			curBlock, err = cipher.Decrypt(curBlock)
			if err != nil {
				log.Fatalf("error decrypting: %s", err)
			}
		} else {
			curBlock, err = cipher.Encrypt(curBlock)
			if err != nil {
				log.Fatalf("error encrypting: %s", err)
			}
		}
		result = append(result, curBlock...)
	}

	if isDecrypt {
		i := len(result) - 1
		for ; result[i] == 0; i-- {
		}
		result = result[:i+1]
	}

	if outputFile == "" {
		fmt.Println(result)
	} else {
		file, err := os.Create(outputFile)
		defer func() {
			err := file.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()
		if err != nil {
			log.Fatalln(err)
		}
		file.Write(result)
	}
}
