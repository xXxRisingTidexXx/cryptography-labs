package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"io/ioutil"
)

func main() {
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
	text, err := ioutil.ReadFile(*textPath)
	if err != nil {
		log.Fatalf("main: gost2814789 failed to read the text, %v", err)
	}
	key, err := ioutil.ReadFile(*keyPath)
	if err != nil {
		log.Fatalf("main: gost2814789 failed to read the key, %v", err)
	}
	textRunes := []rune(string(text))
	if len(textRunes)%2 == 1 {
		last := textRunes[len(textRunes)-1]
		log.Debugf("main: gost2814789 got odd-length message, appended %q to the end", last)
		textRunes = append(textRunes, last)
	}
	cipher, output := cryptolabs.NewGOST2814789(cryptolabs.NewSBox([8][16]byte{})), ""
	if *isDecryption {
		output = cipher.Decrypt(string(textRunes), string(key))
	} else {
		output = cipher.Encrypt(string(textRunes), string(key))
	}
	if err := ioutil.WriteFile(*outputPath, []byte(output), 0600); err != nil {
		log.Fatalf("main: gost2814789 failed to write the output, %v", err)
	}
}
