package blockdb

import (
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/coin"
)

// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.

// encodeSizeUxOut computes the size of an encoded object of type UxOut
func encodeSizeUxOut(obj *coin.UxOut) uint64 {
	i0 := uint64(0)

	// obj.Head.Time
	i0 += 8

	// obj.Head.BkSeq
	i0 += 8

	// obj.Body.SrcTransaction
	i0 += 32

	// obj.Body.Address.Version
	i0++

	// obj.Body.Address.Key
	i0 += 20

	// obj.Body.Coins
	i0 += 8

	// obj.Body.Hours
	i0 += 8

	return i0
}

// encodeUxOut encodes an object of type UxOut to a buffer allocated to the exact size
// required to encode the object.
func encodeUxOut(obj *coin.UxOut) ([]byte, error) {
	n := encodeSizeUxOut(obj)
	buf := make([]byte, n)

	if err := encodeUxOutToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeUxOutToBuffer encodes an object of type UxOut to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeUxOutToBuffer(buf []byte, obj *coin.UxOut) error {
	if uint64(len(buf)) < encodeSizeUxOut(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Head.Time
	e.Uint64(obj.Head.Time)

	// obj.Head.BkSeq
	e.Uint64(obj.Head.BkSeq)

	// obj.Body.SrcTransaction
	e.CopyBytes(obj.Body.SrcTransaction[:])

	// obj.Body.Address.Version
	e.Uint8(obj.Body.Address.Version)

	// obj.Body.Address.Key
	e.CopyBytes(obj.Body.Address.Key[:])

	// obj.Body.Coins
	e.Uint64(obj.Body.Coins)

	// obj.Body.Hours
	e.Uint64(obj.Body.Hours)

	return nil
}

// decodeUxOut decodes an object of type UxOut from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeUxOut(buf []byte, obj *coin.UxOut) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Head.Time
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Head.Time = i
	}

	{
		// obj.Head.BkSeq
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Head.BkSeq = i
	}

	{
		// obj.Body.SrcTransaction
		if len(d.Buffer) < len(obj.Body.SrcTransaction) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.Body.SrcTransaction[:], d.Buffer[:len(obj.Body.SrcTransaction)])
		d.Buffer = d.Buffer[len(obj.Body.SrcTransaction):]
	}

	{
		// obj.Body.Address.Version
		i, err := d.Uint8()
		if err != nil {
			return 0, err
		}
		obj.Body.Address.Version = i
	}

	{
		// obj.Body.Address.Key
		if len(d.Buffer) < len(obj.Body.Address.Key) {
			return 0, encoder.ErrBufferUnderflow
		}
		copy(obj.Body.Address.Key[:], d.Buffer[:len(obj.Body.Address.Key)])
		d.Buffer = d.Buffer[len(obj.Body.Address.Key):]
	}

	{
		// obj.Body.Coins
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Body.Coins = i
	}

	{
		// obj.Body.Hours
		i, err := d.Uint64()
		if err != nil {
			return 0, err
		}
		obj.Body.Hours = i
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeUxOutExact decodes an object of type UxOut from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeUxOutExact(buf []byte, obj *coin.UxOut) error {
	if n, err := decodeUxOut(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}
