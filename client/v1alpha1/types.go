package v1alpha1

import ()

type Cache struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
