package delivery

import (
	infraPsql "github.com/iqbalnzls/watchcommerce/src/infrastructure/repository/psql"
	"github.com/iqbalnzls/watchcommerce/src/pkg/config"
	"github.com/iqbalnzls/watchcommerce/src/pkg/database"
	"github.com/iqbalnzls/watchcommerce/src/pkg/validator"
	usecaseBrand "github.com/iqbalnzls/watchcommerce/src/usecase/brand"
	usecaseOrder "github.com/iqbalnzls/watchcommerce/src/usecase/order"
	usecaseProduct "github.com/iqbalnzls/watchcommerce/src/usecase/product"
)

type Container struct {
	BrandService   usecaseBrand.BrandServiceIFace
	ProductService usecaseProduct.ProductServiceIFace
	OrderService   usecaseOrder.OrderServiceIFace
	Config         *config.Config
	Validator      *validator.DataValidator
}

func SetupContainer() *Container {
	//init config
	cfg := config.NewConfig("./resources/config.json")

	//init validator
	v := validator.NewValidator()

	//init database
	db := database.NewDatabase(&cfg.Database)

	//init repository
	brandRepo := infraPsql.NewRepositoryBrand(db)
	productRepo := infraPsql.NewProductRepository(db)
	orderRepo := infraPsql.NewOrderRepository(db)
	orderDetailsRepo := infraPsql.NewOrderDetailsRepository(db)

	//init service
	brandService := usecaseBrand.NewBrandService(brandRepo)
	productService := usecaseProduct.NewProductService(productRepo)
	orderService := usecaseOrder.NewOrderService(productRepo, orderRepo, orderDetailsRepo)

	return &Container{
		BrandService:   brandService,
		ProductService: productService,
		OrderService:   orderService,
		Config:         cfg,
		Validator:      v,
	}
}
