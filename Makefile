build:
	@go build -o p2p

test:	build
	./p2p -port 6666

node1:
	./p2p -port 5434 -priv CAMSeTB3AgEBBCA+P7yPsaWQKwUefEvcAjNd0eyZLtyLoNPcZqaig+/pB6AKBggqhkjOPQMBB6FEA0IABJFpkamZPf7Q4/EJSeFXmozOap987FW6uNw9m2lb5D9irgdLskYizJroFAjtQ7bPYMLt3+oI9tQCeBg7Ss0wjPw=

node2: 
	./p2p -port 5555 -peer /ip4/127.0.0.1/tcp/5434/p2p/Qma7HB5QJ4fzHHfSUPJi4ghMx6iTLdKNoYHhmvz9GAXYfK -priv CAMSeTB3AgEBBCBn23a67eSJa+pCyY+IiEVCHVuD/OOr9BZK2hjCBwrvkqAKBggqhkjOPQMBB6FEA0IABOZCchDryMK5b2ima23p55r8+plJDWtRbbhzSR6uoswck4nPpVqDWUI0Xn81eWKJxK2qLwLrdoE4nTQWxIoCklI=

node3: 
	./p2p -port 5556 -peer /ip4/127.0.0.1/tcp/5555/p2p/QmSfRbwNTxutoB7cHQFQsVDjpKHWhaYkmF2urEauZ6QVHA  -priv  CAMSeTB3AgEBBCCdwYCWE6drPatossd76aEumvCdLoxN2nq9wbww+AN6QqAKBggqhkjOPQMBB6FEA0IABKB3uI2k1AkGa8f+sxxkdGTI0YJZI4ofOmCw5IT957PQiEbuvVGESGhMS1AubrAOKW30FNleiMl1E2zQxrPW4pM=

node4: 
	./p2p -port 5557 -peer /ip4/127.0.0.1/tcp/5555/p2p/QmSfRbwNTxutoB7cHQFQsVDjpKHWhaYkmF2urEauZ6QVHA  -priv  CAMSeTB3AgEBBCBWnkflgxAKP9xM6GFl7sfqW6Tc8I3Z/K4NTsaNThum4qAKBggqhkjOPQMBB6FEA0IABK/2/Zye5ViOr2f4ggPmCaFS6aaXrVrPDMM/ii9G6R50UWRzENhImGkr87vGbT+PD7aSxZhymoBvSxcpFBxQUQc=







