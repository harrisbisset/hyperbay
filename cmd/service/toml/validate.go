package toml

import (
	"errors"
)

func ValidateRelayConfig(rc *RelayConfig) error {
	if rc == nil {
		return errors.New("relay config is nil")
	}

	if rc.Sites == nil {
		return errors.New("relay sites is nil")
	}

	return nil
}
