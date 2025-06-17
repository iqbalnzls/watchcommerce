package order

import (
	"errors"

	"golang.org/x/sync/errgroup"

	domainOrder "github.com/iqbalnzls/watchcommerce/src/domain"
	"github.com/iqbalnzls/watchcommerce/src/dto"
	orderRepo "github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/order"
	orderDetailsRepo "github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/order_details"
	productRepo "github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/product"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
)

type orderService struct {
	productRepo      productRepo.RepositoryIFace
	orderRepo        orderRepo.RepositoryIFace
	orderDetailsRepo orderDetailsRepo.RepositoryIFace
}

func NewOrderService(productRepo productRepo.RepositoryIFace, orderRepo orderRepo.RepositoryIFace, orderDetailsRepo orderDetailsRepo.RepositoryIFace) ServiceIFace {
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

func (s *orderService) Save(appCtx *appContext.AppContext, req *dto.CreateOrderRequest) (err error) {
	var (
		orderDetails = make([]domainOrder.OrderDetails, 0)
		products     = make([]*domainOrder.Product, 0)
		total        int64
	)

	for _, order := range req.OrderDetails {
		product, er := s.productRepo.GetByID(appCtx, order.ProductID)
		if er != nil {
			return er
		}
		if order.Quantity > product.Quantity {
			err = errors.New(constant.ErrorStockNotAvailable)
			return
		}
		orderDetail := domainOrder.OrderDetails{
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

	id, err := s.orderRepo.SaveWithDBTrx(appCtx, tx, &domainOrder.Order{
		Total: total,
	})
	if err != nil {
		_ = s.orderRepo.RollbackDBTrx(appCtx, tx)
		return
	}

	if err = s.orderDetailsRepo.SaveBulkWithDBTrx(appCtx, tx, id, orderDetails); err != nil {
		_ = s.orderRepo.RollbackDBTrx(appCtx, tx)
		return
	}

	for _, product := range products {
		if err = s.productRepo.UpdateByQuantityWithDBTrx(appCtx, tx, product.ID, product.Quantity); err != nil {
			_ = s.orderRepo.RollbackDBTrx(appCtx, tx)
			return
		}
	}

	_ = s.orderRepo.CommitDBTrx(appCtx, tx)

	return
}

func (s *orderService) Get(appCtx *appContext.AppContext, req *dto.GetOrderRequest) (resp dto.GetOrderResponse, err error) {
	var (
		eg           = new(errgroup.Group)
		order        *domainOrder.Order
		orderDetails []*domainOrder.OrderDetails
	)

	eg.Go(func() (er error) {
		order, er = s.orderRepo.Get(appCtx, req.OrderID)
		return
	})

	eg.Go(func() (er error) {
		orderDetails, er = s.orderDetailsRepo.GetByOrderID(appCtx, req.OrderID)
		return
	})

	if err = eg.Wait(); err != nil {
		return
	}

	resp = toGetOrderResponse(order, orderDetails)

	return
}
