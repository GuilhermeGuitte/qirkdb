package node

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestGettersAndSetter(t *testing.T) {
    // Range
    nd := New()

    // Act
    nd.SetName("node-1")
    nd.SetIpV4("192.168.0.1")
    nd.SetPort(9000)
    nd.SetUnassigned()

    assert.Equal(t, "node-1", nd.GetName())
    assert.Equal(t, "192.168.0.1", nd.GetIpV4())
    assert.Equal(t, 9000, nd.GetPort())
    assert.Equal(t, "yellow", nd.GetStatus())
}
