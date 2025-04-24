package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bytes"
	"encoding/json"

	"github.com/skip2/go-qrcode"
)

func main() {
	// Ambil argumen dari CLI
	inputPath := flag.String("in", "", "Path ke file input (wajib)")
	outputPath := flag.String("out", "output.png", "Nama file output (default: output.png)")
	size := flag.Int("size", 512, "Ukuran QR code dalam pixel (default: 512)")
	flag.Parse()

	if *inputPath == "" {
		fmt.Println("Gunakan -in untuk menentukan file input")
		flag.Usage()
		os.Exit(1)
	}

	// Baca isi file
	data, err := ioutil.ReadFile(*inputPath)
	if err != nil {
		log.Fatalf("Gagal membaca file: %v\n", err)
	}
	
	// Compact JSON
	var buffer bytes.Buffer

	err = json.Compact(&buffer, data)
	if err != nil {
		log.Fatalf("Gagal melakukan compact JSON: %v\n", err)
	}
	data = buffer.Bytes()

	// Generate QR code
	err = qrcode.WriteFile(string(data), qrcode.Medium, *size, *outputPath)
	if err != nil {
		log.Fatalf("Gagal membuat QR code: %v\n", err)
	}

	fmt.Printf("QR code berhasil dibuat: %s\n", *outputPath)
}
