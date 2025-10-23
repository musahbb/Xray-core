package anytls

import (
	"context"

	"github.com/xtls/xray-core/common"
)

const protocolName = "anytls"

func init() {
	common.Must(common.RegisterConfig((*ServerConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewServer(ctx, config.(*ServerConfig))
	}))
}
