package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

func Read(buf io.Reader) (pack []Package, err error) {
	scanner := bufio.NewScanner(buf)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == byte('V') && len(data) > 4 {
			var length uint16
			log.Println(data)
			_ = binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &length)
			if 4 + int(length) <= len(data) {
				return int(length) + 4, data[:int(length)+4], nil
			}
		}
		return
	})

	packages := make([]Package, 0)
	for scanner.Scan() {
		scPackage := new(Package)
		err := scPackage.UnPacket(bytes.NewReader(scanner.Bytes()))
		log.Println(scPackage)
		packages = append(packages, *scPackage)
		if err != nil {
			return packages, err
		}
	}
	return packages, nil
}