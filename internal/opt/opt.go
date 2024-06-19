package opt

import (
	"fmt"
	"mistlib/internal/dht"

	"mistlib/internal/webrtc"
)

// Chunk represents a 3D coordinate
type Chunk struct {
	X int
	Y int
	Z int
}

type ChunkOffer struct {
	Chunk Chunk
	Offer string
}

type PeerDataStore struct {
	// chunkを入れるとWebRTCのpeerを返す
	peer map[string]*webrtc.PeerData
	k    *dht.Kademlia
}

func (p PeerDataStore) connect(chunk Chunk) {
	chunkString := fmt.Sprintf("%d, %d, %d", chunk.X, chunk.Y, chunk.Z)
	fmt.Println("Connecting to", chunkString)
	id := dht.NewKademliaID(chunkString)
	fmt.Println("ID:", id.String())

	findNode := p.k.FindNode(*id)
	fmt.Println("Found node:", findNode)
	// 近いノードを探す

	for _, contact := range findNode {
		// TODO: ここでのfindNodeは近いノードであるという前提、違うのであれば修正したい
		// おそらくDHTを理解できていないかも
		fmt.Println("Contact:", contact)
		// 接続を試みる
		peer := p.peer[chunkString]
		peer.Connect(contact.Address)
	}

	// findNodeがない場合はBootstrapNodeに接続する
	// idが見つかった場合は、webrtcのpeerのdatachannelを利用して、接続のリクエストを送る
}
