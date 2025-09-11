package handler

import (
	"context"

	"github.com/budsx/synapsis/inventory-service/entity"
	inventory "github.com/budsx/synapsis/inventory-service/proto"
)

func (h *InventoryHandler) CheckStock(ctx context.Context, req *inventory.CheckStockRequest) (*inventory.CheckStockResponse, error) {
	stock, err := h.service.CheckStock(ctx, &entity.CheckStockRequest{
		ProductID: req.ProductId,
	})
	if err != nil {
		return nil, err
	}

	return &inventory.CheckStockResponse{
		Stock: stock.Stock,
	}, nil
}

func (h *InventoryHandler) ReserveStock(ctx context.Context, req *inventory.ReserveStockRequest) (*inventory.ReserveStockResponse, error) {
	// success, err := h.service.ReserveStock(ctx, req.ProductId, req.Quantity)
	// if err != nil {
	// 	return nil, err
	// }

	return &inventory.ReserveStockResponse{
		// Success: success,
	}, nil
}

func (h *InventoryHandler) ReleaseStock(ctx context.Context, req *inventory.ReleaseStockRequest) (*inventory.ReleaseStockResponse, error) {
	// success, err := h.service.ReleaseStock(ctx, req.ProductId, req.Quantity)
	// if err != nil {
	// 	return nil, err
	// }

	return &inventory.ReleaseStockResponse{
		// Success: success,
	}, nil
}

func (h *InventoryHandler) GetProductByID(ctx context.Context, req *inventory.GetProductByIDRequest) (*inventory.GetProductByIDResponse, error) {
	product, err := h.service.GetProductByID(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return &inventory.GetProductByIDResponse{
			Product: nil,
		}, nil
	}

	return &inventory.GetProductByIDResponse{
		Product: &inventory.Product{
			Id:        product.ID,
			Name:      product.Name,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		},
	}, nil
}
