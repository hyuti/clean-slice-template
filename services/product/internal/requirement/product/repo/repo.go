package repo

import "github.com/hyuti/clean-slice-template/services/product/ent"

func New(c *ent.Client) IProductRepo {
	return &ProductRepo{client: c}
}
