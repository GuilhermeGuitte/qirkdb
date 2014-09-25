package cluster

import(
    "qirkdb/node"
)

type Cluster struct {
    status string
    nodes map[string]*node.Node
}

/**
 * Constructor
 */
func New() *Cluster {
    return &Cluster{
        status: "yellow",
        nodes:  make(map[string]*node.Node),
    }
}

/**
 * Add nodes to cluster
 */
func (cluster *Cluster) AddNode(nd *node.Node) {
    cluster.nodes[nd.GetName()] = nd
}

/**
 * Quantity of nodes
 */
func (cluster *Cluster) CountNodes() (int) {
    return len(cluster.nodes)
}

/**
 * Removing a node
 */
func (cluster *Cluster) RemoveNode(nd *node.Node) {
    delete(cluster.nodes, nd.GetName())
}
