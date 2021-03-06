package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/common/message"
	"io"
	"net"
)

// Transfer 用于网络传输的对象结构体
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

// ReadPkg 从链接中读取消息的函数
func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("conn.Read, err", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[:4])
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail, err =", err)
	}

	err = json.Unmarshal(transfer.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("序列化失败, err = ", err)
	}
	return
}

// WritePkg 向链接中写入消息的函数
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	var pkglen uint32
	pkglen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(transfer.Buf[0:4], pkglen)
	n, err := transfer.Conn.Write(transfer.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("transfer.Conn.Write 失败", err)
		return
	}
	n, err = transfer.Conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("transfer.Conn.Write 失败", err)
		return
	}
	
	return
}
