package dht

import (
	"bytes"
	"crypto/sha1"
	"math/big"
	"testing"
)

func TestNewKademliaID(t *testing.T) {
	data := "testdata"
	id := NewKademliaID(data)
	expected := sha1.Sum([]byte(data))
	if !bytes.Equal(id[:], expected[:]) {
		t.Errorf("NewKademliaID() = %v, want %v", id, expected)
	}
}

func TestKademliaID_Xor(t *testing.T) {
	id1 := NewKademliaID("node1")
	id2 := NewKademliaID("node2")
	distance := id1.Xor(*id2)

	expectedDistance := new(big.Int).Xor(new(big.Int).SetBytes(id1[:]), new(big.Int).SetBytes(id2[:]))
	if distance.Cmp(expectedDistance) != 0 {
		t.Errorf("KademliaID.Xor() = %v, want %v", distance, expectedDistance)
	}
}

func TestRoutingTable_AddContact(t *testing.T) {
	rt := &RoutingTable{
		Self: Contact{ID: *NewKademliaID("self")},
	}

	contact := Contact{ID: *NewKademliaID("contact1"), Address: "127.0.0.1:8000"}
	rt.AddContact(contact)

	bucketIndex := rt.bucketIndex(contact.ID)
	contacts := rt.Buckets[bucketIndex]

	found := false
	for _, c := range contacts {
		if c.ID == contact.ID {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Contact not found in bucket %d", bucketIndex)
	}
}

func TestKademlia_StoreAndFindValue(t *testing.T) {
	k := &Kademlia{
		routingTable: &RoutingTable{
			Self: Contact{ID: *NewKademliaID("self")},
		},
		dataStore: make(map[KademliaID][]byte),
	}

	key := "key1"
	value := "value1"

	k.Store(key, value)
	storedValue, found := k.FindValue(key)

	if !found {
		t.Errorf("Value not found for key %s", key)
	}

	if string(storedValue) != value {
		t.Errorf("FindValue() = %s, want %s", storedValue, value)
	}
}

func TestKademlia_FindNode(t *testing.T) {
	k := &Kademlia{
		routingTable: &RoutingTable{
			Self: Contact{ID: *NewKademliaID("self")},
		},
		dataStore: make(map[KademliaID][]byte),
	}

	nodeID := *NewKademliaID("node1")
	contact := Contact{ID: nodeID, Address: "127.0.0.1:8000"}
	k.routingTable.AddContact(contact)

	foundContacts := k.FindNode(nodeID)

	found := false
	for _, c := range foundContacts {
		if c.ID == contact.ID {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Node not found in bucket")
	}
}
