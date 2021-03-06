package core

import (
	"golang.org/x/net/context"
	"github.com/OpenBazaar/openbazaar-go/pb"
	peer "gx/ipfs/QmbyvM8zRFDkbFdYyt1MnevUMJ62SiSGbfDFZ3Z8nkrzr4/go-libp2p-peer"
)

func (n *OpenBazaarNode) GetPeerStatus(peerId string) string {
	p, err := peer.IDB58Decode(peerId)
	if err != nil {
		return "error parsing peerId"
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	m := pb.Message{MessageType: pb.Message_PING,}
	_, err = n.Service.SendRequest(ctx, p, &m)
	if err != nil {
		return "offline"
	}
	return "online"
}
