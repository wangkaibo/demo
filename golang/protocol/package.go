package protocol

import (
	"encoding/binary"
	"io"
)

type Package struct {
	Version [2]byte
	IpInt uint32
	Port int16
	Length uint16
	Content []byte
}

func (p *Package) Packet(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.Version)
	err = binary.Write(writer, binary.BigEndian, &p.Length)
	err = binary.Write(writer, binary.BigEndian, &p.IpInt)
	err = binary.Write(writer, binary.BigEndian, &p.Port)
	err = binary.Write(writer, binary.BigEndian, &p.Content)

	return err;
}

func (p *Package) UnPacket(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.Version)
	err = binary.Read(reader, binary.BigEndian, &p.Length)
	err = binary.Read(reader, binary.BigEndian, &p.IpInt)
	err = binary.Read(reader, binary.BigEndian, &p.Port)
	p.Content = make([]byte, p.Length - 6)
	err = binary.Read(reader, binary.BigEndian, &p.Content)

	return err;
}