package node

type Node struct {
    name  string
    ip_v4 string
    port  int
    status string
}

func New() *Node {
    return &Node{}
}

func (node *Node) SetName(name string) {
    node.name = name
}

func (node *Node) GetName() (string) {
    return node.name
}

func (node *Node) SetIpV4(ip string) {
    node.ip_v4 = ip
}

func (node *Node) GetIpV4() (string) {
    return node.ip_v4
}

func (node *Node) SetPort(port int) {
    node.port = port
}

func (node *Node) GetPort() (int) {
    return node.port
}

func (node *Node) SetUnassigned() {
    node.status = "yellow"
}

func (node *Node) GetStatus() (string) {
    return node.status
}
