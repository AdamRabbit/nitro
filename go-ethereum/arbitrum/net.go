package arbitrum

import (
	"context"
	"fmt"
)

// PublicNetAPI offers network related RPC methods
type PublicNetAPI struct {
	networkVersion uint64
}

// NewPublicNetAPI creates a new net API instance.
func NewPublicNetAPI(networkVersion uint64) *PublicNetAPI {
	return &PublicNetAPI{networkVersion}
}

// Version returns the current ethereum protocol version.
func (s *PublicNetAPI) Version() string {
	return fmt.Sprintf("%d", s.networkVersion)
}

func (s *PublicNetAPI) HelleWorld(ctx context.Context, name string) string {
	return fmt.Sprintf("%s, hello world!", name)
}
