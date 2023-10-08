![image](https://github.com/solace-labs/keynet/assets/103751566/f71e2064-d78b-42df-affe-65f7f1081294)

## Solace Phase02 Keynet

The goal of this repo is to create as complicated a network possible and test various scenerios:

1. Peer Discovery - so every node is connected to every other node
2. Peer Failover - each node which node is not available
3. Peer Rediscovery - if a failed node comes back online, then all other nodes need to reconnect with it just like nothing happened
4. Store some shared state, which can keep getting updated

## Squad

- [ ] Periodically keep checking if the squad of ID has the right peers - If updated, trigger resharing
- [ ] Store multiple wallet addresses under one squad

## TODO

- [x] Create a smart-contract querying interface - responsible for just that
  - Querying Squad State
  - Querying Public Key Management state as well
- [x] Make all nodes a part of the same squad (For Testing)
- [x] Implement external request handlers, which can trigger events
- [x] Port this over to the TSS network or port that to this network (DKG + Signing)
- [x] Handle storing & retreving of signed messages (Partial)
- [ ] Gracefully handle panics

## Interface

- [x] Mock smart-contract which has the state of peers + wallets under management
- [x] Plan for Node / Network restarts (SaveData caching)
- [x] DKG should be done as and when the network is inititalized. DKG should be incrementally done as and when new wallets come under management
- [x] Signing Verification should take place, with the SC data as the source of truth (till Phase03)
- [x] Specify the ethereum transaction format for the Transaction Wrapper
- [x] Unwrap, Verify and Sign
- [x] Signatures and metadata should be stored by individual nodes on their respective DBs (Squad level DBs)

## Wrappers

- [x] Define how the Rules look like
- [x] Interface for Declaring Rules
- [x] Define how a tx looks like
- [x] Interface for checking if a Tx fits the rules
- [x] Transaction Parser module - Interface
- [x] Rules Parser Format
- [x] Rules / Transaction Enforcer

## Rules

- [ ] Multichain Support

- [x] Spend Cap
- [x] Sender => Max <Token, Value>
- [x] Store spend data
- [x] Goup wise Spend limits

- [x] Sender Groups

- [x] Recepient Based Rules
- [x] Groups are Allowed to send tokens to certain recepients

- [x] Static way to generate transaction hashes
- [x] Handle multiple signatures with the same TX signature - Use Nonce!
- [x] Sender based Nonce management

- [ ] Add URD for rules
- [ ] Add CRUD for Spending Caps

- [ ] Read local save data instead of running DKG again

## Metrics

- [x] Peer ID List
- [x] Squad IDs / Wallet Addresses
- [x] Stored Signatures per Squad

- [x] Fetching sigantures based on Tx ID

## Smart Contract Wallet Integration

- [x] Use external RPCs to check for SCW Owners
- [ ] Accepts New Rules for wallets, only from Owners
- [x] Request to get the public key after DKG
