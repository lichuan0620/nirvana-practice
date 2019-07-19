package descriptors

import (
	"github.com/caicloud/nirvana/definition"

	"github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1/handlers"
)

// CacheDescriptor builds and returns a Descriptor for all Cache APIs.
func CacheDescriptor() definition.Descriptor {
	return definition.Descriptor{
		Path: "/caches",
		Definitions: []definition.Definition{
			{
				Method:   definition.Create,
				Function: handlers.CreateCache,
				Parameters: []definition.Parameter{
					definition.BodyParameterFor("cache object to create"),
				},
				Results: definition.DataErrorResults("the create result (or error)"),
			},
			{
				Method:   definition.List,
				Function: handlers.ListCaches,
				Parameters: []definition.Parameter{
					definition.QueryParameterFor("prefix", "prefix of the cache key"),
				},
				Results: definition.DataErrorResults("the list result (or error)"),
			},
		},
		Children: []definition.Descriptor{
			{
				Path: "/{key}",
				Definitions: []definition.Definition{
					{
						Method:   definition.Get,
						Function: handlers.GetCache,
						Parameters: []definition.Parameter{
							cacheKey(),
						},
						Results: definition.DataErrorResults("the get result (or error)"),
					},
					{
						Method:   definition.Update,
						Function: handlers.UpdateCache,
						Parameters: []definition.Parameter{
							cacheKey(),
							definition.BodyParameterFor(""),
						},
						Results: definition.DataErrorResults("the update result (or error)"),
					},
					{
						Method:   definition.Delete,
						Function: handlers.DeleteCache,
						Parameters: []definition.Parameter{
							cacheKey(),
						},
						Results: []definition.Result{definition.ErrorResult()},
					},
				},
				Description: "single-target Cache APIs",
			},
		},
		Description: "all Cache APIs",
	}
}

func cacheKey() definition.Parameter {
	return definition.PathParameterFor("key", "key of the cache")
}
