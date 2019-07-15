package descriptors

import (
	"github.com/caicloud/nirvana/definition"

	"github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1/handlers"
)

// ProductDescriptor builds and returns a Descriptor for all Product APIs.
func ProductDescriptor() definition.Descriptor {
	return definition.Descriptor{
		Path:        "/products",
		Definitions: []definition.Definition{
			// TODO
		},
		Children: []definition.Descriptor{
			{
				Path: "/{product}",
				Definitions: []definition.Definition{
					{
						Method:   definition.Get,
						Function: handlers.GetProduct,
						Parameters: []definition.Parameter{
							definition.PathParameterFor("product", "name of the product to get"),
						},
						Results: definition.DataErrorResults("the get result (or error)"),
					},
					// TODO
				},
				Description: "single-target Product APIs",
			},
		},
		Description: "all Product APIs",
	}
}
