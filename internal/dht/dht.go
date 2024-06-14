package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
)

const (
	IDLength = 20
)

type KademliaID [IDLength]byte

type Contact struct {
	ID      KademliaID
	Address string
}

type RoutingTable struct {
	Self    Contact
	Buckets [IDLength * 8][]Contact
	mutex   sync.Mutex
}

type Kademlia struct {
	routingTable *RoutingTable
	dataStore    map[KademliaID][]byte
	mutex        sync.Mutex
}

func NewKademliaID(data string) *KademliaID {
	hash := sha1.Sum([]byte(data))
	id := KademliaID(hash)
	return &id
}

func (id KademliaID) String() string {
	return hex.EncodeToString(id[:])
}

func (rt *RoutingTable) AddContact(contact Contact) {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()
	bucketIndex := rt.bucketIndex(contact.ID)
	rt.Buckets[bucketIndex] = append(rt.Buckets[bucketIndex], contact)
}

func (rt *RoutingTable) bucketIndex(id KademliaID) int {
	distance := rt.Self.ID.Xor(id)
	return distance.BitLen() - 1
}

func (id KademliaID) Xor(other KademliaID) *big.Int {
	distance := new(big.Int)
	distance.SetBytes(id[:])
	otherInt := new(big.Int)
	otherInt.SetBytes(other[:])
	distance.Xor(distance, otherInt)
	return distance
}

func (k *Kademlia) Store(key, value string) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	id := NewKademliaID(key)
	k.dataStore[*id] = []byte(value)
}

func (k *Kademlia) FindValue(key string) ([]byte, bool) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	id := NewKademliaID(key)
	value, found := k.dataStore[*id]
	return value, found
}

func (k *Kademlia) FindNode(id KademliaID) []Contact {
	bucketIndex := k.routingTable.bucketIndex(id)
	fmt.Println("Bucket index:", bucketIndex)
	return k.routingTable.Buckets[bucketIndex]
}
