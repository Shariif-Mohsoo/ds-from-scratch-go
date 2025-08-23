package main

import (
	"log"
	"time"

	"github.com/anthdm/foreverstore/p2p"
)

//	func OnPeer(peer p2p.Peer) error {
//		peer.Close()
//		// fmt.Println("Doing some logic with the peer outside of TCPTransport")
//		return nil
//	}
func main() {
	// tcpOpts := p2p.TCPTransportOpts{
	// 	ListenAddr:    ":3000",
	// 	HandshakeFunc: p2p.NOPHandShakeFunc,
	// 	Decoder:       p2p.DefaultDecoder{},
	// 	OnPeer:        OnPeer,
	// }

	// tr := p2p.NewTCPTransport(tcpOpts)

	// go func() {
	// 	for {
	// 		msg := <-tr.Consume()
	// 		fmt.Printf("%+v\n", msg)
	// 	}
	// }()

	// if err := tr.ListenAndAccept(); err != nil {
	// 	log.Fatal(tr.ListenAndAccept())
	// }
	// select {}

	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		//Todo: onPeer func
	}

	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}
	s := NewFileServer(fileServerOpts)

	go func() {
		time.Sleep(time.Second * 3)
		log.Println("Stopping server now...")
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	// select {}

}
