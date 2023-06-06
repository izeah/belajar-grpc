package services

import (
	"context"

	paginationPb "belajar-grpc/pb/pagination"
	productPb "belajar-grpc/pb/product"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
}

func (s *ProductService) GetProducts(context.Context, *productPb.Empty) (*productPb.Products, error) {
	products := &productPb.Products{
		Pagination: &paginationPb.Pagination{
			Total:       10,
			PerPage:     5,
			CurrentPage: 1,
			LastPage:    2,
		},
		Data: []*productPb.Product{
			{
				Id:    1,
				Name:  "Product A",
				Price: 10000,
				Stock: 99,
				Category: &productPb.Category{
					Id:   1,
					Name: "T-Shirt",
				},
			},
			{
				Id:    2,
				Name:  "Product B",
				Price: 10000,
				Stock: 99,
				Category: &productPb.Category{
					Id:   1,
					Name: "T-Shirt",
				},
			},
			{
				Id:    3,
				Name:  "Product C",
				Price: 10000,
				Stock: 99,
				Category: &productPb.Category{
					Id:   1,
					Name: "T-Shirt",
				},
			},
		},
	}

	return products, status.Error(codes.OK, "Get data successfull")
}
