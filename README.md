### IPFS, IPFS-Cluster, and Underlying Modules

### IPFS
The InterPlanetary File System (IPFS) is a peer-to-peer (p2p) network and protocol designed to make the web faster, safer, and more open. IPFS is an open-source project by Protocol Labs with contributions from thousands of individuals around the world. It aims to supplement, or possibly even replace, the Hypertext Transfer Protocol (HTTP) that currently powers the web, by creating a more reliable and robust system.

Instead of relying on centralized servers to deliver web content, IPFS connects devices with the same system of files. In some ways, this makes IPFS similar to the World Wide Web. But IPFS could be seen as a single BitTorrent swarm, exchanging objects within one Git repository.

IPFS uses content-addressing to uniquely identify each file in the global namespace. This system of files leads to completely distributed applications. It brings file systems and the web closer together with their technological advances.

### IPFS-Cluster
IPFS-Cluster is a standalone application and a CLI client which allocates, replicates, and tracks pins across a cluster of IPFS daemons. It has been designed with high reliability, automatic allocation, operational simplicity, and scalability in mind, offering a friendly REST API and leaving out advanced IPFS features like the Mutable File System (MFS).

IPFS-Cluster serves as an orchestration layer for managing a large number of nodes and coordinating the storage of data. It allows users to pin data across multiple IPFS nodes, ensuring that data is reliably stored and redundantly backed up across multiple machines.

### Underlying Modules
The core components of IPFS and IPFS-Cluster consist of several different protocols and technologies. Some of the primary ones include:

* Content-addressing: IPFS uses a method called content-addressing to store and retrieve data. This means that each piece of data is identified based on its content, not its location.

* Distributed Hash Tables (DHTs): IPFS uses DHTs to achieve a decentralized system of nodes. DHTs allow the network to find peers and content quickly and efficiently.

* BitSwap: It's the exchange protocol within IPFS, which lets nodes exchange blocks of data with each other. It works a bit like BitTorrent, where you give and get data from multiple sources at the same time.

* MerkleDAG: IPFS combines all the data into a MerkleDAG, a tree data structure where each node contains a hash of its content, which allows for secure and fast access to large distributed datasets.

* Libp2p: This is a modular network stack that allows you to build your own peer-to-peer applications. IPFS uses libp2p for its networking capabilities, but libp2p is designed to be used in many other systems and projects.

* IPLD: InterPlanetary Linked Data (IPLD) is a data model and suite of protocols for linking and traversing distributed data across different systems. It's used in IPFS to allow for interoperation between different content-addressed systems.

These modules and components come together to form IPFS and IPFS-Cluster, creating a decentralized, efficient, and reliable system for storing and retrieving data across the Internet.
