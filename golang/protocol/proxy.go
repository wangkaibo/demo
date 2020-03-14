package protocol

import (
	"encoding/binary"
	"log"
	"net"
)

func Proxy(p *Package) []byte {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, p.IpInt)
	address := net.IP(ipByte).String() + ":" + string(p.Port)
	conn,err := net.Dial("tcp",address)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = conn.Write(p.Content)

	return []byte{'1'}
}