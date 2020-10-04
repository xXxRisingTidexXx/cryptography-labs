package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"io/ioutil"
)

func main() {
	textPath := flag.String("t", "examples/feistel.txt", "The input message file path")
	keyPath := flag.String("k", "examples/feistel.key", "The key file path")
	outputPath := flag.String("o", "examples/feistel.out", "The cipher output file path")
	fName := flag.String("f", "magma", "F-function transforming runes")
	isDecryption := flag.Bool("d", false, "Decrypt the input message")
	isVerbose := flag.Bool("v", false, "Run the program in verbose mode")
	flag.Parse()
	if *isVerbose {
		log.SetLevel(log.DebugLevel)
	}
	text, err := ioutil.ReadFile(*textPath)
	if err != nil {
		log.Fatalf("main: feistel failed to read the text, %v", err)
	}
	key, err := ioutil.ReadFile(*keyPath)
	if err != nil {
		log.Fatalf("main: feistel failed to read the key, %v", err)
	}
	textRunes := []rune(string(text))
	if len(textRunes)%2 == 1 {
		last := textRunes[len(textRunes)-1]
		log.Debugf("main: feistel got odd-length message, appended %q to the end", last)
		textRunes = append(textRunes, last)
	}
	f := cryptolabs.Magma
	switch *fName {
	case "des":
		f = cryptolabs.DES
	case "and":
		f = cryptolabs.AND
	default:
	}
	cipher, output := cryptolabs.NewFeistel(f), ""
	if *isDecryption {
		output = cipher.Decrypt(string(textRunes), string(key))
	} else {
		output = cipher.Encrypt(string(textRunes), string(key))
	}
	if err := ioutil.WriteFile(*outputPath, []byte(output), 0600); err != nil {
		log.Fatalf("main: feistel failed to write the output, %v", err)
	}
}
