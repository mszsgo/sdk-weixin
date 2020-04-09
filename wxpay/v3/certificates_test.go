package v3

import "testing"

func TestProtocol_V3Certificates(t *testing.T) {
	err := NewProtocol(NewCfg()).V3Certificates()
	t.Error(err)
}
