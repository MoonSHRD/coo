package main

import (
	"context"
	"flag"
	"fmt"
	mrand "math/rand"

	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	"github.com/libp2p/go-libp2p-core/crypto"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	listenPort := flag.Int("port", 53100, "Port for waiting for incoming connections")
	flag.Parse()

	ctx := context.Background()

	sourceMultiAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", *listenPort))

	r := mrand.New(mrand.NewSource(int64(10)))
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		panic(err)
	}
	host, err := libp2p.New(
		ctx,
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
		libp2p.EnableRelay(circuit.OptHop),
	)
	if err != nil {
		panic(err)
	}

	//fmt.Println("This node: ", host.ID().Pretty(), " ", host.Addrs())
	for _, ips := range host.Addrs() {
		fmt.Printf("%s/p2p/%s\n", ips, host.ID())
	}

	_, err = dht.New(ctx, host)
	if err != nil {
		panic(err)
	}

	select {}
}
