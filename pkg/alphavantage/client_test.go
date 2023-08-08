package av

import (
    "testing"
)

func TestClientGet(t *testing.T) {
    client := NewClient()

    client.Get()

}
