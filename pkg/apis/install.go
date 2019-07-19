package apis

import (
	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/plugins/reqlog"

	"github.com/lichuan0620/nirvana-practice/client"
	"github.com/lichuan0620/nirvana-practice/pkg/apis/middlewares"
	"github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1/descriptors"
)

// Install configures the given Nirvana Config object with the API Descriptor.
func Install(config *nirvana.Config, client client.Interface) {
	config.Configure(
		reqlog.Default(),
		nirvana.Descriptor(
			ProductAPIDescriptor(client),
			//CacheAPIDescriptor(),
		),
	)
}

func CacheInstall(config *nirvana.Config) {
	config.Configure(
		reqlog.Default(),
		nirvana.Descriptor(
			CacheAPIDescriptor(),
		),
	)
}

// APIDescriptor builds and returns a descriptor of all APIs.
func ProductAPIDescriptor(client client.Interface) definition.Descriptor {
	return definition.Descriptor{
		Path:     "/api",
		Consumes: []string{definition.MIMEJSON},
		Produces: []string{definition.MIMEJSON},
		Middlewares: []definition.Middleware{
			middlewares.WithClientInterface(client),
		},
		Children: []definition.Descriptor{
			{
				Path: "/v1alpha1",
				Children: []definition.Descriptor{
					descriptors.ProductDescriptor(),
				},
				Description: "all v1alpha1 APIs",
			},
		},
		Description: "all APIs",
	}
}

func CacheAPIDescriptor() definition.Descriptor {
	return definition.Descriptor{
		Path:     "/api",
		Consumes: []string{definition.MIMEJSON},
		Produces: []string{definition.MIMEJSON},
		Children: []definition.Descriptor{
			{
				Path: "/v1alpha1",
				Children: []definition.Descriptor{
					//descriptors.ProductDescriptor(),
					descriptors.CacheDescriptor(),
				},
				Description: "all v1alpha1 APIs",
			},
		},
		Description: "all APIs",
	}
}
