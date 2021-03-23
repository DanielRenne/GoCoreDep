// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Counter (CTR) mode with support for continuation.

// CTR converts a block cipher into a stream cipher by
// repeatedly encrypting an incrementing counter and
// xoring the resulting stream of data with the input.

// See NIST SP 800-38A, pp 13-15

package cryptoutils

import (
	"bytes"
	"crypto/cipher"
	"encoding/gob"
)

type betterCTRState struct {
	Ctr     []byte
	Out     []byte
	OutUsed int
}

type BetterCTR struct {
	b       cipher.Block
	ctr     []byte
	out     []byte
	outUsed int
}

const streamBufferSize = 512

// NewBetterCTR returns a Stream which encrypts/decrypts using the given Block in
// counter mode. The length of iv must be the same as the Block's block size.
func NewBetterCTR(block cipher.Block, iv []byte) *BetterCTR {
	if len(iv) != block.BlockSize() {
		panic("cryptoutils.NewBetterCTR: IV length must equal block size")
	}
	bufSize := streamBufferSize
	if bufSize < block.BlockSize() {
		bufSize = block.BlockSize()
	}
	return &BetterCTR{
		b:       block,
		ctr:     dup(iv),
		out:     make([]byte, 0, bufSize),
		outUsed: 0,
	}
}

// NewBetterCTR returns a Stream from state
func NewBetterCTRFromState(block cipher.Block, state []byte) *BetterCTR {
	x := &BetterCTR{
		b: block,
	}
	x.SetState(state)
	return x
}

func (x *BetterCTR) GetState() []byte {
	var state bytes.Buffer

	enc := gob.NewEncoder(&state)

	enc.Encode(betterCTRState{
		Ctr:     x.ctr,
		Out:     x.out,
		OutUsed: x.outUsed,
	})

	return state.Bytes()
}

func (x *BetterCTR) SetState(state []byte) error {
	dec := gob.NewDecoder(bytes.NewBuffer(state))

	var s betterCTRState

	err := dec.Decode(&s)

	if err != nil {
		return err
	}

	x.ctr = s.Ctr
	x.out = s.Out
	x.outUsed = s.OutUsed

	return nil
}

func (x *BetterCTR) refill() {
	remain := len(x.out) - x.outUsed
	if remain > x.outUsed {
		return
	}
	copy(x.out, x.out[x.outUsed:])
	x.out = x.out[:cap(x.out)]
	bs := x.b.BlockSize()
	for remain < len(x.out)-bs {
		x.b.Encrypt(x.out[remain:], x.ctr)
		remain += bs

		// Increment counter
		for i := len(x.ctr) - 1; i >= 0; i-- {
			x.ctr[i]++
			if x.ctr[i] != 0 {
				break
			}
		}
	}
	x.out = x.out[:remain]
	x.outUsed = 0
}

func (x *BetterCTR) XORKeyStream(dst, src []byte) {
	for len(src) > 0 {
		if x.outUsed >= len(x.out)-x.b.BlockSize() {
			x.refill()
		}
		n := xorBytes(dst, src, x.out[x.outUsed:])
		dst = dst[n:]
		src = src[n:]
		x.outUsed += n
	}
}

func dup(p []byte) []byte {
	q := make([]byte, len(p))
	copy(q, p)
	return q
}
