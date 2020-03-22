package main

import (
	"bufio"
	"bytes"
	"demo/golang/leetcode"
	"demo/golang/protocol"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//mode := os.Args[1]
	//if mode == "1" {
	//	listen()
	//} else {
	//	connect()
	//}
	l1 := leetcode.InitListNode([]int{1,2,3,4,5,5})
	l2 := leetcode.InitListNode([]int{4,5,6,7,3,5})
	res := leetcode.MergeTwoLists(l1, l2)

	log.Println(res)

	return
}


func listen() {
	ln,err := net.Listen("tcp",":9527")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listen...")
	for {
		log.Println("waiting...")
		conn,err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		packages,_ := protocol.Read(conn)
		//protocol.Proxy(scPackage)
		buf := new(bytes.Buffer)
		for _,pkg := range packages {
			_ = pkg.Packet(buf)
		}
		_, _ = conn.Write(buf.Bytes())
		conn.Close()
	}
}


func connect() {
	conn,err := net.Dial("tcp",":9527")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	_ = conn.(*net.TCPConn).SetNoDelay(false)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	//p,_ := reader.ReadBytes('\n')
	p,_ := reader.ReadBytes('\n')
	pack := &protocol.Package{
		Version: [2]byte{'V', '1'},
		Content: p,
	}
	s := "255.0.255.1"
	ip := net.ParseIP(s).To4()
	pack.IpInt = binary.BigEndian.Uint32(ip)
	pack.Port = 8888
	pack.Length = uint16(len(pack.Content) + 6)
	log.Println("包内容：")
	log.Println(pack)
	buf := new(bytes.Buffer)
	_ = pack.Packet(buf)
	//_ = pack.Packet(buf)
	//_ = pack.Packet(buf)
	//_ = pack.Packet(buf)
	log.Println(buf.Bytes())
	_,_ = conn.Write(buf.Bytes())
	_,_ = conn.Write(buf.Bytes())
	_,_ = conn.Write(buf.Bytes())
	_,_ = conn.Write(buf.Bytes())
	buffer := new(bytes.Buffer)
	_, _ = io.Copy(buffer, conn)

	resPack := new(protocol.Package)
	_ = resPack.UnPacket(buffer)
	fmt.Print("Message from server: ")
	fmt.Print(resPack)
}