package cluster

import(
    "testing"
    "github.com/stretchr/testify/assert"
    "qirkdb/node"
)

func TestShouldAddNode(t *testing.T) {
    // Range
    cluster := New()
    nd    := node.New()

    nd.SetName("node1")

    // Act
    cluster.AddNode(nd)
    count := cluster.CountNodes()

    // Assert
    assert.Equal(t, 1, count)
}


func TestShouldRemoveNode(t *testing.T) {
    // Range
    cluster := New()
    nd    := node.New()

    nd.SetName("node1")

    // Act
    cluster.AddNode(nd)
    count := cluster.CountNodes()
    assert.Equal(t, 1, count)

    cluster.RemoveNode(nd)

    // Assert
    count = cluster.CountNodes()
    assert.Equal(t, 0, count)
}
