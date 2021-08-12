ETH-DHT
===

A standalone DHT built for Ethereum.

## Table of Contents

[TOC]

## Beginners Guide

If you are a total beginner to this, start here!

1. Learn more about DHT's
2. Understand the difference between Kademlia and DHT's
3. Look at the specs for: RLPx, eth65, libp2p docs (if you can)
4. Start writing note!

### Packages

1. Node
2. Kademlia

**DHT:**
- A Distributed Hash Table is a P2P networking primitive that, in its most basic formulation, permits storage and lookup of key, value pairs.
- Its a hash table spread across many nodes.
- Kademlia is a good example because its pretty simple
- Internal state is also simple
- Kademlia makes tradeoffs (vs. Pastry Chord), its not practical to implement pubsub.
- [Read More - Kademlia detailed spec](http://xlattice.sourceforge.net/components/protocol/kademlia/specs.html)

**Note:** Node ID's are very important in DHT's. Here its a 160 bil opaque identifier, used to identify nodes on the network, as well as content hashes.
