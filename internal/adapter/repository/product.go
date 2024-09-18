package repository

import (
	"sistemaLogistica/internal/controllers"
	"sistemaLogistica/internal/entity"
)

type repoProduct struct {
	repoSql controllers.RepoSql
}

func (r repoProduct) Insert(info entity.Product) (entity.Product, error) {
	var infoReturn entity.Product

	err := r.repoSql.QueryRow(`INSERT INTO products (client_id ,item, quantity, kind, date) VALUES ($1,$2,$3,$4,$5) RETURNING client_id, item, quantity, kind, date`, []any{info.ClientId, info.Item, info.Quantity, info.Kind, info.Date}, &infoReturn.ClientId, &infoReturn.Item, &infoReturn.Quantity, &infoReturn.Kind, &infoReturn.Date)

	return infoReturn, err
}

func (r repoProduct) Get(info int) (entity.Product, error) {
	var infoReturn entity.Product

	err := r.repoSql.QueryRow(`SELECT * FROM products WHERE client_id=$1`, []any{info}, &infoReturn.ClientId, &infoReturn.Item, &infoReturn.Quantity, &infoReturn.Kind, &infoReturn.Date)

	return infoReturn, err
}

func (r repoProduct) GetByName(name string) (entity.Product, error) {
	var infoReturn entity.Product

	err := r.repoSql.QueryRow(`SELECT * FROM products WHERE kind=$1`, []any{name}, &infoReturn.ClientId, &infoReturn.Item, &infoReturn.Quantity, &infoReturn.Kind, &infoReturn.Date)

	return infoReturn, err
}

func (r repoProduct) Update(info entity.Product) (entity.Product, error) {
	var infoReturn entity.Product

	err := r.repoSql.QueryRow(`UPDATE products SET client_id=$1, quantity=$3, kind=$4, date=$5 WHERE item=$2 RETURNING id,item,quantity,kind,date`, []any{info.ClientId, info.Item, info.Quantity, info.Kind, info.Date}, &infoReturn.ClientId, &infoReturn.Item, &infoReturn.Quantity, &infoReturn.Kind, &infoReturn.Date)

	return infoReturn, err

}

func (r repoProduct) Delete(info entity.Product) error {

	err := r.repoSql.QueryRow(`DELETE FROM products WHERE id=$1`, []any{info.ClientId})

	return err
}

func NewProductRepo(r controllers.RepoSql) controllers.ProductRepository {
	return repoProduct{repoSql: r}
}
