package v1alpha1

type Cache struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
