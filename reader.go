package uereader

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/google/uuid"
)

var ErrFStringNotNullTerminated = errors.New("string is not null terminated")

func ReadBytes(r io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)

	_, err := r.Read(data)

	return data, err
}

func ReadByte(r io.Reader) (byte, error) {
	data, err := ReadBytes(r, 1)
	return data[0], err
}

func ReadBool(r io.Reader) (bool, error) {
	data, err := ReadByte(r)

	if data != 0 && data != 1 {
		return false, fmt.Errorf("invalid boolean value %d", data)
	}

	return data != 0, err
}

func ReadFBool(r io.Reader) (bool, error) {
	data, err := ReadInt32(r, binary.LittleEndian)

	if data != 0 && data != 1 {
		return false, fmt.Errorf("invalid boolean value %d", data)
	}

	return data != 0, err
}

func ReadUInt8(r io.Reader) (uint8, error) {
	return ReadByte(r)
}

func ReadUInt16(r io.Reader, order binary.ByteOrder) (uint16, error) {
	data, err := ReadBytes(r, 2)
	if err != nil {
		return 0, err
	}

	return order.Uint16(data), nil
}

func ReadUInt32(r io.Reader, order binary.ByteOrder) (uint32, error) {
	data, err := ReadBytes(r, 4)
	if err != nil {
		return 0, err
	}

	return order.Uint32(data), nil
}

func ReadUInt64(r io.Reader, order binary.ByteOrder) (uint64, error) {
	data, err := ReadBytes(r, 8)
	if err != nil {
		return 0, err
	}

	return order.Uint64(data), nil
}

func ReadInt8(r io.Reader) (int8, error) {
	data, err := ReadUInt8(r)
	return int8(data), err
}

func ReadInt16(r io.Reader, order binary.ByteOrder) (int16, error) {
	data, err := ReadUInt16(r, order)
	return int16(data), err
}

func ReadInt32(r io.Reader, order binary.ByteOrder) (int32, error) {
	data, err := ReadUInt32(r, order)
	return int32(data), err
}

func ReadInt64(r io.Reader, order binary.ByteOrder) (int64, error) {
	data, err := ReadUInt64(r, order)
	return int64(data), err
}

func ReadUUID(r io.Reader) (guid uuid.UUID, err error) {
	_, err = r.Read(guid[:])
	return guid, err
}

func ReadBigEndianUUID(r io.Reader) (guid uuid.UUID, err error) {
	for i := 0; i < 4; i++ {

	}
	data := make([]uint32, 4)
	if err = binary.Read(r, binary.BigEndian, &data); err != nil {
		return guid, err
	}

	for i, v := range data {
		binary.LittleEndian.PutUint32(guid[i*4:(i+1)*4], v)
	}

	return guid, nil
}

func ReadShaHash(r io.Reader) (hash [20]byte, err error) {
	_, err = r.Read(hash[:])
	return hash, err
}

func ReadString(r io.Reader) (string, error) {
	length, err := ReadInt32(r, binary.LittleEndian)
	if err != nil {
		return "", err
	} else if length == 0 {
		return "", nil
	}

	if length < 0 {
		length = -length
	}

	data, err := ReadBytes(r, int(length))
	if err != nil {
		return "", err
	} else if data[len(data)-1] != 0x0 { // ensure it's null-terminated
		return "", ErrFStringNotNullTerminated
	}

	// avoid the null character while returning
	return string(data[:len(data)-1]), nil
}
