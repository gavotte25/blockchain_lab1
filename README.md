# blockchain_lab1
### Some references
https://copyprogramming.com/howto/convert-slice-of-int64-to-byte-array-and-vice-versa-in-golang
https://www.youtube.com/watch?v=VjoPRuVbhCg
https://www.quora.com/What-is-the-difference-between-light-nodes-and-full-nodes-in-blockchain-technology-What-are-the-benefits-and-drawbacks-of-each-type-of-node-Which-type-of-node-do-you-think-will-be-more-popular-among-users-customers
https://electrum.readthedocs.io/en/latest/spv.html

client validates transaction instead of letting fullnode do everything, avoiding malicious acts

It's easier for the servers to return the Merkle path than to check very transaction to reduce the work-load. Hence, the whole process works out.
https://stackoverflow.com/questions/49012000/how-is-a-merkle-tree-path-generated

https://bitcoin.stackexchange.com/questions/37899/how-do-i-find-out-what-block-a-transaction-is-in
https://bitcoin.stackexchange.com/questions/88763/finding-a-transaction-in-the-blockchain

All miners are fullnodes, but not all fullnodes are miners.
Wallet apps are lightnodes.

# Command line interface documentation
Before using CLI, you should install Cobra - a library we use to create CLI - by running this command `go get -u github.com/spf13/cobra@latest`, then run this command: `go run && go build` to create commands.

The root command we use is `blockchain_lab1`. And you can always use the flag `-h` to show help.

As a server administrator:
- `blockchain_lab1 createchain --name (-n) myblockchain ` allows you to create a new blockchain with your custom name, i.e `myblockchain`, and initial first block with an empty-content transaction.
- `blockchain_lab1 deletechain --name (-n) myblockchain` allows you to delete permanently a blockchain with name `myblockchain`. Be careful when using this command.

As a client:
- `blockchain_lab1 addtransaction --name (-n) myblockchain --data (-d) "A send to B an amount of 2 BTC" "B send to C an amount of 3 BTC"` will add a single or multiple of transactions which content defined in `--data` flag to the lastest block in the blockchain `myblockchain`.
- `blockchain_lab1 addblock --name (-n) myblockchain --data (-d) "A send to B an amount of 3 BTC"` allows you to add a new block to the blockchain `myblockchain` with a first transaction which content defined in `--data` flag.
- `blockchain_lab1 chaininfo --name (-n) myblockchain` simply prints the information of all blocks. The information contains "Block address", "Block size", "Time stamp", "Number of transaction".
- `blockchain_lab1 getNumberOfBlock --name (-n) myblockchain` retrieves the total number of blocks in a specified blockchain.
- `blockchain_lab1 numtransaction --name (-n) myblockchain` the "numtransaction" command retrieves the number of transactions on a specified blockchain at all blocks.