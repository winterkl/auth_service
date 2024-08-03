package paginate

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"math"
	"strconv"
)

const defaultLimit = 20

type Paginate struct {
	CurrentPage int `form:"page" json:"current_page"`
	PageSize    int `form:"page_size" json:"page_size"`
	PageCounter int `json:"page_counter"`
}

func New(ctx *gin.Context) (*Paginate, error) {
	pageParam, ok := ctx.GetQuery("page")
	if ok {
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			return nil, err
		}
		if page <= 0 {
			return nil, &InvalidPage{}
		}
		pagination := &Paginate{
			CurrentPage: page,
			PageSize:    defaultLimit,
		}
		// Если есть кастомный параметр размера страницы
		if sizeParam, ok := ctx.GetQuery("page_size"); ok {
			size, err := strconv.Atoi(sizeParam)
			if err != nil {
				return nil, err
			}
			pagination.PageSize = size
		}
		return pagination, nil
	}
	return nil, nil
}

func (p *Paginate) AddPagination(query *bun.SelectQuery) *bun.SelectQuery {
	offset := (p.CurrentPage - 1) * p.PageSize
	return query.Offset(offset).Limit(p.PageSize)
}
func (p *Paginate) CalculatePagesCount(ctx context.Context, query *bun.SelectQuery) error {
	recordCount, err := query.Count(ctx)
	if err != nil {
		return err
	}
	p.PageCounter = int(math.Ceil(float64(recordCount) / float64(p.PageSize)))
	return nil
}
