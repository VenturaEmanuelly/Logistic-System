package repository

import (
	"sistemaLogistica/internal/controllers"
	"sistemaLogistica/internal/entity"
)



type repoTotalProduct struct {
	repoSql controllers.RepoSql
}

func (r repoTotalProduct) Upsert(info entity.TotalProduct) (entity.TotalProduct, error) {
	var infoReturn entity.TotalProduct

	err := r.repoSql.QueryRow(`INSERT INTO total_products (name, total) VALUES ($1, 1) ON CONFLICT (name) DO UPDATE SET name = $1, total = total + 1 RETURNING name, total`, []any{info.Name}, &infoReturn.Name, &infoReturn.Total)

	return infoReturn, err
}

func (r repoTotalProduct) Get(info string) (entity.TotalProduct, error) {
	var infoReturn entity.TotalProduct

	err := r.repoSql.QueryRow(`SELECT * FROM products WHERE name=$1`, []any{info}, &infoReturn.Name, &infoReturn.Total)

	return infoReturn, err
}

func NewTotalProductRepo(r controllers.RepoSql) controllers.TotalProductRepository {
	return repoTotalProduct{r}
}
