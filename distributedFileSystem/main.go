package main

import (
	"bytes"
	"log"
	"time"

	"github.com/anthdm/foreverstore/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		//Todo: onPeer func
	}

	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootStrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)
	tcpTransport.OnPeer = s.OnPeer

	return s

}

//	func OnPeer(peer p2p.Peer) error {
//		peer.Close()
//		// fmt.Println("Doing some logic with the peer outside of TCPTransport")
//		return nil
//	}
func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")
	go func() {
		log.Fatal(s1.Start())
	}()
	time.Sleep(1 * time.Second)

	go s2.Start()
	time.Sleep(1 * time.Second)

	data := bytes.NewReader([]byte("My big data file here!"))
	s2.StoreData("myPrivateData", data)

	select {}
}
