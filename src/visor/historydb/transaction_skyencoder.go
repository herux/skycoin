// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package historydb

import (
	"errors"
	"math"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/coin"
)

// EncodeSizeTransaction computes the size of an encoded object of type Transaction
func EncodeSizeTransaction(obj *Transaction) uint64 {
	i0 := uint64(0)

	// obj.Txn.Length
	i0 += 4

	// obj.Txn.Type
	i0++

	// obj.Txn.InnerHash
	i0 += 32

	// obj.Txn.Sigs
	i0 += 4
	{
		i1 := uint64(0)

		// x
		i1 += 65

		i0 += uint64(len(obj.Txn.Sigs)) * i1
	}

	// obj.Txn.In
	i0 += 4
	{
		i1 := uint64(0)

		// x
		i1 += 32

		i0 += uint64(len(obj.Txn.In)) * i1
	}

	// obj.Txn.Out
	i0 += 4
	{
		i1 := uint64(0)

		// x.Address.Version
		i1++

		// x.Address.Key
		i1 += 20

		// x.Coins
		i1 += 8

		// x.Hours
		i1 += 8

		i0 += uint64(len(obj.Txn.Out)) * i1
	}

	// obj.BlockSeq
	i0 += 8

	return i0
}

// EncodeTransaction encodes an object of type Transaction to the buffer in encoder.Encoder.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func EncodeTransaction(buf []byte, obj *Transaction) error {
	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Txn.Length
	e.Uint32(obj.Txn.Length)

	// obj.Txn.Type
	e.Uint8(obj.Txn.Type)

	// obj.Txn.InnerHash
	e.CopyBytes(obj.Txn.InnerHash[:])

	// obj.Txn.Sigs maxlen check
	if len(obj.Txn.Sigs) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Txn.Sigs length check
	if uint64(len(obj.Txn.Sigs)) > math.MaxUint32 {
		return errors.New("obj.Txn.Sigs length exceeds math.MaxUint32")
	}

	// obj.Txn.Sigs length
	e.Uint32(uint32(len(obj.Txn.Sigs)))

	// obj.Txn.Sigs
	for _, x := range obj.Txn.Sigs {

		// x
		e.CopyBytes(x[:])

	}

	// obj.Txn.In maxlen check
	if len(obj.Txn.In) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Txn.In length check
	if uint64(len(obj.Txn.In)) > math.MaxUint32 {
		return errors.New("obj.Txn.In length exceeds math.MaxUint32")
	}

	// obj.Txn.In length
	e.Uint32(uint32(len(obj.Txn.In)))

	// obj.Txn.In
	for _, x := range obj.Txn.In {

		// x
		e.CopyBytes(x[:])

	}

	// obj.Txn.Out maxlen check
	if len(obj.Txn.Out) > 65535 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Txn.Out length check
	if uint64(len(obj.Txn.Out)) > math.MaxUint32 {
		return errors.New("obj.Txn.Out length exceeds math.MaxUint32")
	}

	// obj.Txn.Out length
	e.Uint32(uint32(len(obj.Txn.Out)))

	// obj.Txn.Out
	for _, x := range obj.Txn.Out {

		// x.Address.Version
		e.Uint8(x.Address.Version)

		// x.Address.Key
		e.CopyBytes(x.Address.Key[:])

		// x.Coins
		e.Uint64(x.Coins)

		// x.Hours
		e.Uint64(x.Hours)

	}

	// obj.BlockSeq
	e.Uint64(obj.BlockSeq)

	return nil
}

// DecodeTransaction decodes an object of type Transaction from the buffer in encoder.Decoder.
// Returns the number of bytes used from the buffer to decode the object.
func DecodeTransaction(buf []byte, obj *Transaction) (int, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Txn.Length
		i, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.Txn.Length = i
	}

	{
		// obj.Txn.Type
		i, err := d.Uint8()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.Txn.Type = i
	}

	{
		// obj.Txn.InnerHash
		if len(d.Buffer) < len(obj.Txn.InnerHash) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}
		copy(obj.Txn.InnerHash[:], d.Buffer[:len(obj.Txn.InnerHash)])
		d.Buffer = d.Buffer[len(obj.Txn.InnerHash):]
	}

	{
		// obj.Txn.Sigs

		ul, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return len(buf) - len(d.Buffer), encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Txn.Sigs = make([]cipher.Sig, length)

			for z2 := range obj.Txn.Sigs {
				{
					// obj.Txn.Sigs[z2]
					if len(d.Buffer) < len(obj.Txn.Sigs[z2]) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Txn.Sigs[z2][:], d.Buffer[:len(obj.Txn.Sigs[z2])])
					d.Buffer = d.Buffer[len(obj.Txn.Sigs[z2]):]
				}

			}
		}
	}

	{
		// obj.Txn.In

		ul, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return len(buf) - len(d.Buffer), encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Txn.In = make([]cipher.SHA256, length)

			for z2 := range obj.Txn.In {
				{
					// obj.Txn.In[z2]
					if len(d.Buffer) < len(obj.Txn.In[z2]) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Txn.In[z2][:], d.Buffer[:len(obj.Txn.In[z2])])
					d.Buffer = d.Buffer[len(obj.Txn.In[z2]):]
				}

			}
		}
	}

	{
		// obj.Txn.Out

		ul, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}

		if length > 65535 {
			return len(buf) - len(d.Buffer), encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Txn.Out = make([]coin.TransactionOutput, length)

			for z2 := range obj.Txn.Out {
				{
					// obj.Txn.Out[z2].Address.Version
					i, err := d.Uint8()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Txn.Out[z2].Address.Version = i
				}

				{
					// obj.Txn.Out[z2].Address.Key
					if len(d.Buffer) < len(obj.Txn.Out[z2].Address.Key) {
						return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
					}
					copy(obj.Txn.Out[z2].Address.Key[:], d.Buffer[:len(obj.Txn.Out[z2].Address.Key)])
					d.Buffer = d.Buffer[len(obj.Txn.Out[z2].Address.Key):]
				}

				{
					// obj.Txn.Out[z2].Coins
					i, err := d.Uint64()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Txn.Out[z2].Coins = i
				}

				{
					// obj.Txn.Out[z2].Hours
					i, err := d.Uint64()
					if err != nil {
						return len(buf) - len(d.Buffer), err
					}
					obj.Txn.Out[z2].Hours = i
				}

			}
		}
	}

	{
		// obj.BlockSeq
		i, err := d.Uint64()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.BlockSeq = i
	}

	return len(buf) - len(d.Buffer), nil
}
