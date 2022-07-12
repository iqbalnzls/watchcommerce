package order

import (
	"errors"

	"golang.org/x/sync/errgroup"

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

	tx, err := s.orderRepo.BeginDBTrx()
	if err != nil {
		return
	}

	id, err := s.orderRepo.SaveWithDBTrx(tx, &domainOrder.Order{
		Total: total,
	})
	if err != nil {
		_ = s.orderRepo.RollbackDBTrx(tx)
		return
	}

	if err = s.orderDetailsRepo.SaveBulkWithDBTrx(tx, id, orderDetails); err != nil {
		_ = s.orderRepo.RollbackDBTrx(tx)
		return
	}

	for _, product := range products {
		if err = s.productRepo.UpdateByQuantityWithDBTrx(tx, product.ID, product.Quantity); err != nil {
			_ = s.orderRepo.RollbackDBTrx(tx)
			return
		}
	}

	_ = s.orderRepo.CommitDBTrx(tx)

	return
}

func (s *orderService) Get(req *dto.GetOrderRequest) (resp dto.GetOrderResponse, err error) {
	var (
		eg           = new(errgroup.Group)
		order        *domainOrder.Order
		orderDetails []*domainOrderDetails.OrderDetails
	)

	eg.Go(func() (er error) {
		order, er = s.orderRepo.Get(req.OrderID)
		return
	})

	eg.Go(func() (er error) {
		orderDetails, er = s.orderDetailsRepo.GetByOrderID(req.OrderID)
		return
	})

	if err = eg.Wait(); err != nil {
		return
	}

	resp = toGetOrderResponse(order, orderDetails)

	return
}
