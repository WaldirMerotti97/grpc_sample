package service

import (
	"context"
	"grpc_sample/internal/database"
	"grpc_sample/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryServiceClient(category database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: category,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: "",
	}

	return categoryResponse, nil

}
