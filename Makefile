.Phony: test run stop mock

test:
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go fmt ./... && go test -coverprofile coverage.cov -cover ./... # use -v for verbose
	@echo "\n"
	@echo "=================================================================================="
	@echo "All Package Coverage"
	@echo "=================================================================================="
	go tool cover -func coverage.cov

run:
	docker compose up --build -d

stop:
	docker compose down

mock:
	@echo "Generate Mock Interface.."
	mockery --name=BrandRepositoryIFace --structname=BrandRepositoryIFaceMock --filename=brand_repository_mock.go --output=./src/pkg/mock/infrastructure/repository/psql --dir=./src/infrastructure/repository/psql/brand/
	mockery --name=OrderRepositoryIFace --structname=OrderRepositoryIFaceMock --filename=order_repository_mock.go --output=./src/pkg/mock/infrastructure/repository/psql --dir=./src/infrastructure/repository/psql/order/
	mockery --name=OrderDetailsRepositoryIFace --structname=OrderDetailsRepositoryIFaceMock --filename=order_details_repository_mock.go --output=./src/pkg/mock/infrastructure/repository/psql --dir=./src/infrastructure/repository/psql/order_details/
	mockery --name=ProductRepositoryIFace --structname=ProductRepositoryIFaceMock --filename=product_repository_mock.go --output=./src/pkg/mock/infrastructure/repository/psql --dir=./src/infrastructure/repository/psql/product/
	mockery --name=BrandServiceIFace --structname=BrandServiceIFaceMock --filename=service_mock.go --output=./src/pkg/mock/usecase/brand --dir=./src/usecase/brand
	mockery --name=ProductServiceIFace --structname=ProductServiceIFaceMock --filename=service_mock.go --output=./src/pkg/mock/usecase/product --dir=./src/usecase/product
	mockery --name=OrderServiceIFace --structname=OrderServiceIFaceMock --filename=service_mock.go --output=./src/pkg/mock/usecase/order --dir=./src/usecase/order