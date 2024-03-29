package v1alpha1

import (
	"time"

	meta "github.com/lichuan0620/nirvana-practice/pkg/apis/meta/v1"
)

type Product struct {
	meta.Metadata `json:",inline"`
	Spec          *ProductSpec   `json:"spec,omitempty"`
	Status        *ProductStatus `json:"status,omitempty"`
}

type ProductSpec struct {
	Category string   `json:"category,omitempty"`
	Price    *float64 `json:"price,omitempty"`
}

type ProductStatus struct {
	Sold          *bool      `json:"sold,omitempty"`
	SoldTimestamp *time.Time `json:"soldTimestamp,omitempty"`
}
