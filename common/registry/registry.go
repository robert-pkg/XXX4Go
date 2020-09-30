package registry

import (
	"github.com/pkg/errors"
	"github.com/robert-pkg/micro-go/registry"
	consul_registry "github.com/robert-pkg/micro-go/registry/consul"
)

// Config .
type Config struct {
	RegistryName string // consul, etcd
	Addrs        []string
}

// InitRegistry .
func InitRegistry(c *Config) registry.Registry {

	// 根据系统配置，使用etcd， consul
	registry, err := consul_registry.NewRegistry()
	if err != nil {
		panic(errors.Wrap(err, "grpc server start fail"))
	}

	return registry
}

// InitRegistryAsDefault .
func InitRegistryAsDefault(c *Config) registry.Registry {
	registry.DefaultRegistry = InitRegistry(c)
	return registry.DefaultRegistry
}
