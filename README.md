## LibP2P playground

The goal of this repo is to create as complicated a network possible and test various scenerios:

1. Peer Discovery - so every node is connected to every other node
2. Peer Failover - each node which node is not available
3. Peer Rediscovery - if a failed node comes back online, then all other nodes need to reconnect with it just like nothing happened
4. Store some shared state, which can keep getting updated

## Squad

1. Periodically keep checking if the squad of ID has the right peers - If updated, trigger resharing

## TODO

- [x] Create a smart-contract querying interface - responsible for just that
  - Querying Squad State
  - Querying Public Key Management state as well
- [x] Make all nodes a part of the same squad (For Testing)
- [x] Implement external request handlers, which can trigger events
- [x] Port this over to the TSS network or port that to this network (DKG + Signing)
- [x] Handle storing & retreving of signed messages (Partial)

## Interface

- [x] Mock smart-contract which has the state of peers + wallets under management
- [x] Plan for Node / Network restarts (SaveData caching)
- [ ] DKG should be done as and when the network is inititalized. DKG should be incrementally done as and when new wallets come under management
- [ ] Signing Verification should take place, with the SC data as the source of truth (till Phase03)
- [ ] Specify the ethereum transaction format for the Transaction Wrapper
- [ ] Unwrap, Verify and Sign
- [x] Signatures and metadata should be stored by individual nodes on their respective DBs (Squad level DBs)

## Wrappers

- [ ] Define how the Rules look like
- [ ] Interface for Declaring Rules

- [x] Define how a tx looks like
- [ ] Interface for checking if a Tx fits the rules

- [x] Transaction Parser module - Interface
- [ ] Rules Parser Format
- [ ] Rules / Transaction Enforcer

## Rules

- [ ] Sender Address based Rule
  - [ ] Sender => Max <Token, Value>
  - [ ]
