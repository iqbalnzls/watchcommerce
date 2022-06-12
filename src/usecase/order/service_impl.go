package order

import (
	"errors"

	domainOrder "github.com/iqbalnzls/watchcommerce/src/domain/order"
	domainOrderDetails "github.com/iqbalnzls/watchcommerce/src/domain/order_details"
	domainProduct "github.com/iqbalnzls/watchcommerce/src/domain/product"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	"github.com/iqbalnzls/watchcommerce/src/pkg/utils"
)

type orderService struct {
	productRepo      domainProduct.ProductRepositoryIFace
	orderRepo        domainOrder.OrderRepositoryIFace
	orderDetailsRepo domainOrderDetails.OrderDetailsRepositoryIFace
}

func NewOrderService(productRepo domainProduct.ProductRepositoryIFace, orderRepo domainOrder.OrderRepositoryIFace, orderDetailsRepo domainOrderDetails.OrderDetailsRepositoryIFace) OrderServiceIFace {
	if productRepo == nil {
		panic("product repository is nil")
	}
	if orderRepo == nil {
		panic("order repository is nil")
	}
	if orderDetailsRepo == nil {
		panic("order details repository is nil")
	}

	return &orderService{
		productRepo:      productRepo,
		orderRepo:        orderRepo,
		orderDetailsRepo: orderDetailsRepo,
	}
}

func (s *orderService) Save(req *dto.CreateOrderRequest) (err error) {
	var (
		orderDetails = make([]domainOrderDetails.OrderDetails, 0)
		products     = make([]*domainProduct.Product, 0)
		total        int64
	)

	for _, order := range req.OrderDetails {
		product, er := s.productRepo.GetByID(order.ProductID)
		if er != nil {
			return er
		}
		if order.Quantity > product.Quantity {
			err = errors.New(constant.ErrorStockNotAvailable)
			utils.Error(err)
			return
		}
		orderDetail := domainOrderDetails.OrderDetails{
			Quantity:  order.Quantity,
			Price:     product.Price * order.Quantity,
			ProductID: product.ID,
		}
		orderDetails = append(orderDetails, orderDetail)
		total += orderDetail.Price
		product.Quantity -= order.Quantity
		products = append(products, product)
	}

	id, err := s.orderRepo.Save(&domainOrder.Order{
		Total: total,
	})
	if err != nil {
		return
	}

	if err = s.orderDetailsRepo.SaveBulk(id, orderDetails); err != nil {
		return
	}

	for _, product := range products {
		if err = s.productRepo.UpdateByQuantity(product.ID, product.Quantity); err != nil {
			return
		}
	}

	return
}

func (s *orderService) Get(req *dto.GetOrderRequest) (resp dto.GetOrderResponse, err error) {
	order, err := s.orderRepo.Get(req.OrderID)
	if err != nil {
		return
	}

	orderDetails, err := s.orderDetailsRepo.GetByOrderID(order.ID)
	if err != nil {
		return
	}

	resp = toGetOrderResponse(order, orderDetails)

	return
}
