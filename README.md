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

Why light nodes need blockchain header? -> May be they also store full block info but in of not all blocks

https://viblo.asia/p/phan-loai-va-tam-quan-trong-cua-cac-node-trong-mang-blockchain-6J3Zg0JAlmB