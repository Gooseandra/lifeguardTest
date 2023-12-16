package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"strconv"
	"swagger/restapi/operations"
	"swagger/services"
)

type (
	inventory struct {
		log         *services.Log
		sessions    *services.Sessions
		inventories *services.Inventories
	}
	CreateInventoryItem struct{ inventory }
	ListInventory       struct{ inventory }
	ByID                struct{ crew }
	//UpdateCrew struct{ crew }
)

func NewCreateInventoryItem(l *services.Log, s *services.Sessions, c *services.Inventories) CreateInventoryItem {
	return CreateInventoryItem{inventory: inventory{log: l, sessions: s, inventories: c}}
}

func NewListInventoryItem(l *services.Log, s *services.Sessions, c *services.Inventories) ListInventory {
	return ListInventory{inventory: inventory{log: l, sessions: s, inventories: c}}
}

func NewByIDInventoryItem(l *services.Log, s *services.Sessions, c *services.Inventories) ByID {
	return ListInventory{inventory: inventory{log: l, sessions: s, inventories: c}}
}

func (c CreateInventoryItem) Handle(params operations.CreateInventoryItemParams) middleware.Responder {
	log := c.log.Func("CreateInventoryItem")
	row, err := c.inventories.Create(*params.Body.InventoryType, *params.Body.Name, *params.Body.Description,
		*params.Body.Number)
	if err != nil {
		log.InternalSerer(err.Error())
		return operations.NewCreateInventoryItemInternalServerError()
	}
	log.OK(strconv.FormatUint(row.ID(), 10))
	return operations.NewCreateInventoryItemOK().WithPayload(row.ID())
}

func (c ListInventory) Handle(params operations.ListInventoryItemsParams) middleware.Responder {
	log := c.log.Func("CreateInventoryItem")
	if params.Count == nil {
		log.BadRequest("count is null")
		return operations.NewListUsersBadRequest()
	}
	if params.Skip == nil {
		log.BadRequest("skip is null ")
		return operations.NewListUsersBadRequest()
	}
	row, err := c.inventories.List(*params.Count, *params.Skip)
	if err != nil {
		log.InternalSerer(err.Error())
		return operations.NewCreateInventoryItemInternalServerError()
	}
	payload := make([]*operations.ListInventoryItemsOKBodyItems0, len(row))
	for index, item := range row {
		payload[index] = &operations.ListInventoryItemsOKBodyItems0{ID: item.ID(), Name: item.Name(),
			Number: item.UniqNum(), Description: item.InstanceDesc(), InventoryType: item.TypeName()}
	}
	return operations.NewListInventoryItemsOK().WithPayload(payload)
}
