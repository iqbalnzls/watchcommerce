package delivery

import (
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/brand"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/order"
	"github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/order_details"
	infraPsql "github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql/product"
	"github.com/iqbalnzls/watchcommerce/src/pkg/config"
	"github.com/iqbalnzls/watchcommerce/src/pkg/database"
	"github.com/iqbalnzls/watchcommerce/src/pkg/validator"
	usecaseBrand "github.com/iqbalnzls/watchcommerce/src/usecase/brand"
	usecaseOrder "github.com/iqbalnzls/watchcommerce/src/usecase/order"
	usecaseProduct "github.com/iqbalnzls/watchcommerce/src/usecase/product"
)

type Container struct {
	ProductService usecaseProduct.ProductServiceIFace
	OrderService   usecaseOrder.OrderServiceIFace
	BrandService   usecaseBrand.BrandServiceIFace
	Validator      *validator.DataValidator
	Config         *config.Config
}

func SetupContainer() *Container {
	//init config
	cfg := config.NewConfig("./resources/config.json")

	//init validator
	v := validator.NewValidator()

	//init database
	db := database.NewDatabase(&cfg.Database)

	//init repository
	brandRepo := brand.NewRepositoryBrand(db)
	productRepo := infraPsql.NewProductRepository(db)
	orderRepo := order.NewOrderRepository(db)
	orderDetailsRepo := order_details.NewOrderDetailsRepository(db)

	//init service
	brandService := usecaseBrand.NewBrandService(brandRepo)
	productService := usecaseProduct.NewProductService(productRepo)
	orderService := usecaseOrder.NewOrderService(productRepo, orderRepo, orderDetailsRepo)

	return &Container{
		ProductService: productService,
		OrderService:   orderService,
		BrandService:   brandService,
		Config:         cfg,
		Validator:      v,
	}
}
