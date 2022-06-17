package pmapi

import (
	"context"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
)

type Filter struct {
	ID       string
	Name     string
	Status   int
	Priority int
	Sieve    string
	Tree     []map[string]interface{}
	Version  int
}

func (c *client) ListFilters(ctx context.Context) (filters []Filter, err error) {
	var res struct {
		Filters []Filter
	}

	if _, err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get("/filters")
	}); err != nil {
		return nil, err
	}

	return res.Filters, nil
}

type FilterReq struct {
	Filter
}

func (c *client) CreateFilter(ctx context.Context, filter Filter) (Filter, error) {
	if filter.Name == "" {
		return Filter{}, errors.New("name is required")
	}
	if filter.Sieve == "" {
		return Filter{}, errors.New("sieve is required")
	}

	var res struct {
		Filter Filter
	}

	if _, err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetBody(&FilterReq{
			Filter: filter,
		}).SetResult(&res).Post("/filters")
	}); err != nil {
		return Filter{}, err
	}

	return res.Filter, nil
}

func (c *client) UpdateFilter(ctx context.Context, filter Filter) (Filter, error) {
	if filter.Name == "" {
		return Filter{}, errors.New("name is required")
	}
	if filter.Sieve == "" {
		return Filter{}, errors.New("sieve is required")
	}

	var res struct {
		Filter Filter
	}

	if _, err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetBody(&FilterReq{
			Filter: filter,
		}).SetResult(&res).Put("/filters/" + filter.ID)
	}); err != nil {
		return Filter{}, err
	}

	return res.Filter, nil
}

func (c *client) DeleteFilter(ctx context.Context, filterID string) error {
	if _, err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.Delete("/filters/" + filterID)
	}); err != nil {
		return err
	}
	return nil
}
