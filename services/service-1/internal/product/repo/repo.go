package repo

import "github.com/hyuti/clean-slice-template/services/service-1/ent"

func New(c *ent.Client) IProductRepo {
	return &ProductRepo{client: c}
}
