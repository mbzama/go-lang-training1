/*
 * Generate rsa keys.
 */

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey

	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	filepath := "output/" + formatted
	os.MkdirAll(filepath, os.ModePerm)

	saveInPlainFormat(filepath+"/private.txt", key)
	saveGobKey(filepath+"/private.key", key)
	savePEMKey(filepath+"/private.pem", key)

	saveGobKey(filepath+"/public.key", publicKey)
	savePublicPEMKey(filepath+"/public.pem", publicKey)
}

func saveInPlainFormat(fileName string, key *rsa.PrivateKey) {
	b := x509.MarshalPKCS1PrivateKey(key)
	keybase64 := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("\n" + keybase64 + "\n\n")

	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	outFile.WriteString(keybase64)
	checkError(err)
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	privateKey := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err = pem.Encode(outFile, privateKey)

	checkError(err)
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
