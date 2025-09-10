package handler

import (
	inventory "github.com/budsx/synapsis/inventory-service/proto"
	"github.com/budsx/synapsis/inventory-service/services"
)

type InventoryHandler struct {
	inventory.UnimplementedInventoryServiceServer
	service services.InventoryService
}

func NewInventoryHandler(service services.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		service: service,
	}
}
