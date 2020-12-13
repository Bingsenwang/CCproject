# CCproject
project for CC Bingsen - Zhile
We implemented IKNP protocol and KK13 protocol.
Using sha256 as random oracle and OT based on ECC.
Require github.com/ethereum/go-ethereum/ and golang.org/x/crypto

### To run the protocol

Using command "go build" in both /client_2.0 and /server_2.0
Run server_2.0 in /server_2.0 first, when see "start listen", run client_2.0 in /client_2.0

### Change the parameters

To change the parameters(k,n,l,server_address...), modify the variables in main.go in both directories.


