package uereader

import (
	"io"

	"github.com/google/uuid"
)

func (r *Reader) Skip(size int64) (skipped int64) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	skipped, r.err = r.r.Seek(size, io.SeekCurrent)
	return
}

func (r *Reader) Bytes(size int) (data []byte) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadBytes(r.r, size)
	return data
}

func (r *Reader) Byte() (data byte) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadByte(r.r)
	return data
}

func (r *Reader) Bool() (data bool) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadBool(r.r)
	return data
}

func (r *Reader) FBool() (data bool) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadFBool(r.r)
	return data
}

func (r *Reader) UInt8() (data uint8) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadByte(r.r)
	return data
}

func (r *Reader) UInt16() (data uint16) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadUInt16(r.r, r.order)
	return data
}

func (r *Reader) UInt32() (data uint32) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadUInt32(r.r, r.order)
	return data
}

func (r *Reader) UInt64() (data uint64) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadUInt64(r.r, r.order)
	return data
}

func (r *Reader) Int8() (data int8) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadInt8(r.r)
	return data
}

func (r *Reader) Int16() (data int16) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadInt16(r.r, r.order)
	return data
}

func (r *Reader) Int32() (data int32) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadInt32(r.r, r.order)
	return data
}

func (r *Reader) Int64() (data int64) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadInt64(r.r, r.order)
	return data
}

func (r *Reader) UUID() (data uuid.UUID) {
	if r.err != nil {
		return data
	}
	r.index.Add(1)

	data, r.err = ReadUUID(r.r)
	return data
}

func (r *Reader) BigEndianUUID() (data uuid.UUID) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadBigEndianUUID(r.r)
	return data
}

func (r *Reader) ShaHash() (data [20]byte) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadShaHash(r.r)
	return data
}

func (r *Reader) String() (data string) {
	if r.err != nil {
		return
	}
	r.index.Add(1)

	data, r.err = ReadString(r.r)
	return data
}

type ArrayReaderFunc[T any] func(ar *Reader) (T, error)

func ReadArray[T any](ar *Reader, size int32, fn ArrayReaderFunc[T]) []T {
	if ar.err != nil || size <= 0 {
		return nil
	}

	data := make([]T, size)
	for i := int32(0); i < size; i++ {
		ar.index.Add(1)

		value, err := fn(ar)
		if err != nil {
			ar.err = err
			return nil
		}

		data[i] = value
	}

	return data
}

func ReadSlice[T any](ar *Reader, fn ArrayReaderFunc[T]) []T {
	return ReadArray(ar, ar.Int32(), fn)
}
