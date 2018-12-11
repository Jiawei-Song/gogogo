package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("conn.Read, err", err)
		return
	}
	fmt.Println("读到的buf", buf[:4])

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail, err =", err)
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("序列化失败, err = ", err)
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write 失败", err)
		return
	}
	n, err = conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.Write 失败", err)
		return
	}

	fmt.Printf("客户端，发送的消息长度为%d\n", pkglen)
	return
}
