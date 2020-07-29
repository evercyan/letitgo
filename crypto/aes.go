package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func pKCS7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

func pKCS7UnPadding(text []byte) []byte {
	length := len(text)
	unpadding := int(text[length-1])
	return text[:(length - unpadding)]
}

func AesEncrypt(text string, key string) (string, error) {
	textByte, keyByte := []byte(text), []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	textByte = pKCS7Padding(textByte, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyByte[:blockSize])
	encrypt := make([]byte, len(textByte))
	blockMode.CryptBlocks(encrypt, textByte)
	return Base64Encode(string(encrypt)), nil
}

func AesDecrypt(text string, key string) (string, error) {
	textByte, keyByte := []byte(Base64Decode(text)), []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyByte[:blockSize])
	decrypt := make([]byte, len(textByte))
	blockMode.CryptBlocks(decrypt, textByte)
	return string(pKCS7UnPadding(decrypt)), nil
}
