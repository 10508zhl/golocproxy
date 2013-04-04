package main

import (
	"../util"
	"flag"
	"log"
	"net"
)

var (
	local  = flag.String("l", "127.0.0.1:80", "Access of the local app service")
	remote = flag.String("r", "127.0.0.1:8010", "Access of the golocproxy server")
)

func main() {
	flag.Usage = util.Usage
	flag.Parse()

	println("golocproxy client starting: ", local, "->", remote)
	proxy, err := net.Dial("tcp", *remote)
	if err != nil {
		log.Fatal("CAN'T CONNECT:", *remote, " err:", err)
	}
	defer proxy.Close()
	proxy.Write([]byte(util.C2P_CONNECT))

	var buf [util.TOKEN_LEN]byte
	for {
		n, err := proxy.Read(buf[0:])
		if err != nil {
			log.Fatal("CAN'T READ,", " err:", err)
		}
		token := string(buf[0:n])
		if token == util.P2C_NEW_SESSION {
			go session()
		}
	}
}

func session() {
	session, err := net.Dial("tcp", *remote)
	if err != nil {
		println("Can't' connect:", remote, " err:", err)
		return
	}
	defer session.Close()

}