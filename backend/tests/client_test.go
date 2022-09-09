package tests

import (
	"testing"

	clientservices "casino.website/pkg/services/client.services"
)

func NewClientTest(t *testing.T) {
	client := clientservices.NewClient("test", 1)

	if client.Username != "test" {
		t.Errorf("Expected Username(%s) is not same as"+
			" actual username (%s)", "test", client.Username)
	}

	if client.Wallet != 0 {
		t.Errorf("Expected Wallet(%d) is not the same as"+
			" actual wallet (%d)", 0, client.Wallet)
	}

	if client.AccessType != 1 {
		t.Errorf("Expected AccessType(%d) is not the same as"+
			" actual accessType (%d)", 1, client.AccessType)
	}
}
