package client

import (
	v1alpha1 "github.com/lichuan0620/nirvana-practice/client/v1alpha1"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes a versioned client.
type Interface interface {
	// V1Alpha1 returns v1alpha1 client.
	V1Alpha1() v1alpha1.Interface
}

// Client contains versioned clients.
type Client struct {
	v1alpha1 *v1alpha1.Client
}

// NewClient creates a new client.
func NewClient(cfg *rest.Config) (Interface, error) {
	c := &Client{}
	var err error

	c.v1alpha1, err = v1alpha1.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// MustNewClient creates a new client or panic if an error occurs.
func MustNewClient(cfg *rest.Config) Interface {
	return &Client{
		v1alpha1: v1alpha1.MustNewClient(cfg),
	}
}

// V1 returns a versioned client.
func (c *Client) V1Alpha1() v1alpha1.Interface {
	return c.v1alpha1
}
