package uereader

import (
	"encoding/binary"
	"io"
	"sync/atomic"
)

type Reader struct {
	err   error
	name  string
	index atomic.Int32

	r     io.ReadSeeker
	order binary.ByteOrder
}

func NewReader(name string, r io.ReadSeeker) *Reader {
	return &Reader{
		name: name,

		r:     r,
		order: binary.LittleEndian,
	}
}

func (r *Reader) Err() error {
	if r.err == nil {
		return nil
	}

	return &Error{
		err:     r.err,
		Name:    r.name,
		IndexID: r.index.Load(),
	}
}

func SubReader[T any](ar *Reader, name string, fn func(r *Reader) (T, error)) (data T) {
	if ar.err != nil {
		return data
	}

	ar.index.Add(1)
	newAr := NewReader(name, ar.r)

	data, err := fn(newAr)
	if err != nil {
		ar.err = err
	}

	if newAr.Err() != nil {
		ar.err = newAr.Err()
	}

	return data
}
