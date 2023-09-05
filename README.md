## LibP2P playground

The goal of this repo is to create as complicated a network possible and test various scenerios:

1. Peer Discovery - so every node is connected to every other node
2. Peer Failover - each node which node is not available
3. Peer Rediscovery - if a failed node comes back online, then all other nodes need to reconnect with it just like nothing happened
4. Store some shared state, which can keep getting updated
