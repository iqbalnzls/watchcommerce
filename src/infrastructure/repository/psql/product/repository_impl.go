package product

import (
	"database/sql"
	"errors"

	domainProduct "github.com/iqbalnzls/watchcommerce/src/domain"
	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) RepositoryIFace {
	if db == nil {
		panic("db connection repository is nil")
	}

	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Save(appCtx *appContext.AppContext, domain *domainProduct.Product) (err error) {
	var query = `INSERT INTO commerce.product(brand_id, name, price, quantity) VALUES($1, $2, $3, $4)`

	_, err = r.db.Exec(query, domain.BrandID, domain.Name, domain.Price, domain.Quantity)
	if err != nil {
		appCtx.Logger.Error(err.Error())
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	return
}

func (r *productRepo) GetByID(appCtx *appContext.AppContext, id int64) (domain *domainProduct.Product, err error) {
	var (
		query = `SELECT p.id,
       p.brand_id,
       p.name,
       p.price,
       p.quantity,
       p.created_at,
       p.updated_at,
       b.id,
       b.name,
       b.created_at,
       b.updated_at
FROM commerce.product p
         JOIN commerce.brand b on b.id = p.brand_id
WHERE p.id = $1;`
		product domainProduct.Product
	)

	if err = r.db.QueryRow(query, id).Scan(&product.ID, &product.BrandID, &product.Name, &product.Price,
		&product.Quantity, &product.CreatedAt, &product.UpdatedAt, &product.Brand.ID, &product.Brand.Name,
		&product.Brand.CreatedAt, &product.Brand.UpdatedAt); err != nil {
		appCtx.Logger.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New(constant.ErrorDataNotFound)
			return
		}
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	domain = &product

	return
}

func (r *productRepo) GetByBrandID(appCtx *appContext.AppContext, brandID int64) (products []*domainProduct.Product, err error) {
	var query = `SELECT p.id,
       p.brand_id,
       p.name,
       p.price,
       p.quantity,
       p.created_at,
       p.updated_at,
       b.id,
       b.name,
       b.created_at,
       b.updated_at
FROM commerce.product p
         JOIN commerce.brand b on p.brand_id = b.id
WHERE p.brand_id = $1`

	rows, err := r.db.Query(query, brandID)
	if err != nil {
		appCtx.Logger.Error(err.Error())
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	for rows.Next() {
		var product domainProduct.Product
		if err = rows.Scan(&product.ID, &product.BrandID, &product.Name, &product.Price,
			&product.Quantity, &product.CreatedAt, &product.UpdatedAt, &product.Brand.ID, &product.Brand.Name,
			&product.Brand.CreatedAt, &product.Brand.UpdatedAt); err != nil {
			appCtx.Logger.Error(err.Error())
			err = errors.New(constant.ErrorDatabaseProblem)
			return
		}

		products = append(products, &product)
	}

	if len(products) == 0 {
		err = errors.New(constant.ErrorDataNotFound)
		appCtx.Logger.Error(err.Error())
		return
	}

	return
}

func (r *productRepo) UpdateByQuantityWithDBTrx(appCtx *appContext.AppContext, tx *sql.Tx, id, quantity int64) (err error) {
	var query = `UPDATE commerce.product SET quantity = $2, updated_at = now() WHERE id = $1`

	_, err = tx.Exec(query, id, quantity)
	if err != nil {
		appCtx.Logger.Error(err.Error())
		err = errors.New(constant.ErrorDatabaseProblem)
		return
	}

	return
}
