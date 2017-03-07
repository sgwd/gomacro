// this file was generated by gomacro command: import "crypto/cipher"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/cipher"
)

func init() {
	Binds["crypto/cipher"] = map[string]Value{
		"NewCBCDecrypter":	ValueOf(cipher.NewCBCDecrypter),
		"NewCBCEncrypter":	ValueOf(cipher.NewCBCEncrypter),
		"NewCFBDecrypter":	ValueOf(cipher.NewCFBDecrypter),
		"NewCFBEncrypter":	ValueOf(cipher.NewCFBEncrypter),
		"NewCTR":	ValueOf(cipher.NewCTR),
		"NewGCM":	ValueOf(cipher.NewGCM),
		"NewGCMWithNonceSize":	ValueOf(cipher.NewGCMWithNonceSize),
		"NewOFB":	ValueOf(cipher.NewOFB),
	}
	Types["crypto/cipher"] = map[string]Type{
		"AEAD":	TypeOf((*cipher.AEAD)(nil)).Elem(),
		"Block":	TypeOf((*cipher.Block)(nil)).Elem(),
		"BlockMode":	TypeOf((*cipher.BlockMode)(nil)).Elem(),
		"Stream":	TypeOf((*cipher.Stream)(nil)).Elem(),
		"StreamReader":	TypeOf((*cipher.StreamReader)(nil)).Elem(),
		"StreamWriter":	TypeOf((*cipher.StreamWriter)(nil)).Elem(),
	}
	Proxies["crypto/cipher"] = map[string]Type{
		"AEAD":	TypeOf((*AEAD_crypto_cipher)(nil)).Elem(),
		"Block":	TypeOf((*Block_crypto_cipher)(nil)).Elem(),
		"BlockMode":	TypeOf((*BlockMode_crypto_cipher)(nil)).Elem(),
		"Stream":	TypeOf((*Stream_crypto_cipher)(nil)).Elem(),
	}
}

// --------------- proxy for crypto/cipher.AEAD ---------------
type AEAD_crypto_cipher struct {
	NonceSize_	func() int
	Open_	func(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error)
	Overhead_	func() int
	Seal_	func(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte
}
func (Obj AEAD_crypto_cipher) NonceSize() int {
	return Obj.NonceSize_()
}
func (Obj AEAD_crypto_cipher) Open(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error) {
	return Obj.Open_(dst, nonce, ciphertext, additionalData)
}
func (Obj AEAD_crypto_cipher) Overhead() int {
	return Obj.Overhead_()
}
func (Obj AEAD_crypto_cipher) Seal(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte {
	return Obj.Seal_(dst, nonce, plaintext, additionalData)
}

// --------------- proxy for crypto/cipher.Block ---------------
type Block_crypto_cipher struct {
	BlockSize_	func() int
	Decrypt_	func(dst []byte, src []byte) 
	Encrypt_	func(dst []byte, src []byte) 
}
func (Obj Block_crypto_cipher) BlockSize() int {
	return Obj.BlockSize_()
}
func (Obj Block_crypto_cipher) Decrypt(dst []byte, src []byte)  {
	Obj.Decrypt_(dst, src)
}
func (Obj Block_crypto_cipher) Encrypt(dst []byte, src []byte)  {
	Obj.Encrypt_(dst, src)
}

// --------------- proxy for crypto/cipher.BlockMode ---------------
type BlockMode_crypto_cipher struct {
	BlockSize_	func() int
	CryptBlocks_	func(dst []byte, src []byte) 
}
func (Obj BlockMode_crypto_cipher) BlockSize() int {
	return Obj.BlockSize_()
}
func (Obj BlockMode_crypto_cipher) CryptBlocks(dst []byte, src []byte)  {
	Obj.CryptBlocks_(dst, src)
}

// --------------- proxy for crypto/cipher.Stream ---------------
type Stream_crypto_cipher struct {
	XORKeyStream_	func(dst []byte, src []byte) 
}
func (Obj Stream_crypto_cipher) XORKeyStream(dst []byte, src []byte)  {
	Obj.XORKeyStream_(dst, src)
}