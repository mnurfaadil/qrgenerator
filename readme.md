# QR Generator

## Build on your own

Tambahkan pustaka yang digunakan

```shell
go get github.com/skip2/go-qrcode
```

Jalankan perintah dibawah untuk membangun sendiri program

```shell
make build
```

## How To Use

Jalankan perintah ini untuk membuat qr

JSON

```shell
./bin/qrgenerator -in ./data/input.json -out ./output/output.png
```

TXT

```shell
./bin/qrgenerator -in ./data/raw.txt -out ./output/raw.png
```
