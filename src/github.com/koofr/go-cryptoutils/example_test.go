package cryptoutils_test

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"github.com/koofr/go-cryptoutils"
	"io"
)

func ExampleNewBetterCTR() {
	key := []byte("example key 1234")
	plaintext := []byte("some plaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cryptoutils.NewBetterCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:aes.BlockSize+4], plaintext[:4])

	state := stream.GetState()

	anotherStream := cryptoutils.NewBetterCTRFromState(block, state)
	anotherStream.XORKeyStream(ciphertext[aes.BlockSize+4:], plaintext[4:])

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	// CTR mode is the same for both encryption and decryption, so we can
	// also decrypt that ciphertext with NewBetterCTR.

	plaintext2 := make([]byte, len(plaintext))
	stream = cryptoutils.NewBetterCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("%s\n", plaintext2)
	// Output: some plaintext
}
