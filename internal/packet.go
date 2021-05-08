package internal

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type Packet interface {
	Bool() (bool, error)
	Float() (float32, error)
	UInt64() (uint64, error)
	UInt32() (uint32, error)
	UInt16() (uint16, error)
	UInt8() (uint8, error)
	Int64() (int64, error)
	Int32() (int32, error)
	Int16() (int16, error)
	Int8() (int8, error)
}

func NewPacket(data []byte) Packet {
	return &packet{
		reader: bufio.NewReader(bytes.NewReader(data)),
	}
}

type packet struct {
	reader *bufio.Reader
}

func (p *packet) read() (byte, error) {
	return p.reader.ReadByte()
}

func (p *packet) readN(n int) ([]byte, error) {
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		b, err :=  p.read()
		if err != nil {
			return nil, fmt.Errorf("unable to read %v bytes: %v", n, err)
		}
		out[i] = b
	}

	return out, nil
}

func (p *packet) Bool() (bool, error) {
	b, err := p.read()
	if err != nil {
		return false, fmt.Errorf("failed to read byte: %v", err)
	}
	if b < 0 || b > 1 {
		return false, fmt.Errorf("unexpected byte value %v is <0 or >1", b)
	}
	return b == 1, nil
}

func (p *packet) Float() (float32, error) {
	b, err := p.readN(4)
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return math.Float32frombits(binary.LittleEndian.Uint32(b)), nil
}

func (p *packet) UInt64() (uint64, error) {
	b, err := p.readN(8)
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return binary.LittleEndian.Uint64(b), nil
}

func (p *packet) UInt32() (uint32, error) {
	b, err := p.readN(4)
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return binary.LittleEndian.Uint32(b), nil
}

func (p *packet) UInt16() (uint16, error) {
	b, err := p.readN(2)
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return binary.LittleEndian.Uint16(b), nil
}

func (p *packet) UInt8() (uint8, error) {
	b, err := p.read()
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return b, nil
}

func (p *packet) Int64() (int64, error) {
	b, err := p.readN(8)
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return int64(binary.LittleEndian.Uint64(b)), nil
}

func (p *packet) Int32() (int32, error) {
	b, err := p.readN(4)
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return int32(binary.LittleEndian.Uint32(b)), nil
}

func (p *packet) Int16() (int16, error) {
	b, err := p.readN(2)
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return int16(binary.LittleEndian.Uint16(b)), nil
}

func (p *packet) Int8() (int8, error) {
	b, err := p.read()
	if err != nil {
		return 0, fmt.Errorf("failed to read bytes: %v", err)
	}

	return int8(b), nil
}
