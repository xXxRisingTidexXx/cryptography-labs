package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now()
	sboxPath := flag.String("s", "examples/gost2814789.sbx", "S-Box file path")
	textPath := flag.String("t", "examples/gost2814789.txt", "The input message file path")
	keyPath := flag.String("k", "examples/gost2814789.key", "The key file path")
	outputPath := flag.String("o", "examples/gost2814789.out", "The cipher output file path")
	isDecryption := flag.Bool("d", false, "Decrypt the input message")
	isVerbose := flag.Bool("v", false, "Run the program in verbose mode")
	flag.Parse()
	if *isVerbose {
		log.SetLevel(log.DebugLevel)
	}
	sbox, err := ioutil.ReadFile(*sboxPath)
	if err != nil {
		log.Fatalf("main: gost2814789 failed to read the sbox, %v", err)
	}
	if len(sbox) != 256 {
		log.Fatalf("main: gost2814789 got sbox with invalid length, %d != 255", len(sbox))
	}
	text, err := ioutil.ReadFile(*textPath)
	if err != nil {
		log.Fatalf("main: gost2814789 failed to read the text, %v", err)
	}
	key, err := ioutil.ReadFile(*keyPath)
	if err != nil {
		log.Fatalf("main: gost2814789 failed to read the key, %v", err)
	}
	if modulo := len(text) % 8; modulo != 0 {
		log.Debug("main: gost2814789 got incomplete message, padded with the last rune")
		for i := modulo; i < 8; i++ {
			text = append(text, text[len(text)-1])
		}
	}
	table := [8][16]byte{}
	for i := range sbox {
		table[i/16][i%16] = sbox[i]
	}
	cipher, output := cryptolabs.NewGOST2814789(cryptolabs.NewSBox(table)), ""
	if *isDecryption {
		output = cipher.Decrypt(string(text), string(key))
	} else {
		output = cipher.Encrypt(string(text), string(key))
	}
	if err := ioutil.WriteFile(*outputPath, []byte(output), 0600); err != nil {
		log.Fatalf("main: gost2814789 failed to write the output, %v", err)
	}
	log.Debugf("main: gost2814789 finished in %s", time.Since(start))
}
