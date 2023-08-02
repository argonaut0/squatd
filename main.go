package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/insomniacslk/dhcp/dhcpv4/nclient4"
)

const TARGET_SERVER_IP = "10.246.110.254:67"
const IFACE = "enp6s0"

func main() {
	target, err := net.ResolveUDPAddr("udp4", TARGET_SERVER_IP)
	if err != nil {
		log.Fatal(err)
	}
	client, err := nclient4.New(
		IFACE,
		nclient4.WithDebugLogger(),
		nclient4.WithServerAddr(target))
	if err != nil {
		log.Fatal(err)
	}

	for {
		// only returns first offer, fixme
		offer, err := client.DiscoverOffer(context.Background())
		if err != nil {
			continue
		}
		if !offer.YourIPAddr.Equal(net.IPv4(10, 246, 110, 253)) {
			continue
		}
		client.RequestFromOffer(context.Background(), offer)
		time.Sleep(10 * time.Second)
	}

}
