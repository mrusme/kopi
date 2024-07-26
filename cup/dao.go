package cup

import (
	"context"
	"fmt"

	"github.com/mrusme/balzac/dal"
	"github.com/mrusme/balzac/helpers"
)

type DAO struct {
	dal *dal.DAL
}

func NewDAO(dal *dal.DAL) *DAO {
	dao := new(DAO)

	dao.dal = dal

	return dao
}

func (dao *DAO) Create(
	ctx context.Context,
	entity Cup,
) (Cup, error) {
	id, err := dal.Create(ctx, dao.dal.DB(),
		"INSERT INTO "+Table()+
			" ("+Columns(false)+")"+
			" VALUES ("+helpers.QueryArgRepeat(ColumnsNumber(false))+");",
		&entity.CoffeeID,
		&entity.Drink,
		&entity.Timestamp,
	)
	entity.ID = id
	return entity, err
}

func (dao *DAO) GetByID(
	ctx context.Context,
	id int64,
) (Cup, error) {
	return dal.GetRow[Cup](ctx, dao.dal.DB(),
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `id` = ?"+
			" LIMIT 1;",
		id,
	)
}

func (dao *DAO) Count(
	ctx context.Context,
) (int64, error) {
	return dal.GetColumn[int64](ctx, dao.dal.DB(),
		"SELECT COUNT(*)"+
			" FROM "+Table()+";",
	)
}

func (dao *DAO) FindByIDs(
	ctx context.Context,
	ids []int64,
) ([]Cup, error) {
	placeholders, args := dal.InArgs(ids)
	q := fmt.Sprintf(
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `id` IN (%s)"+
			" ORDER BY `id`;",
		placeholders,
	)
	return dal.FindRows[Cup](ctx, dao.dal.DB(),
		q,
		args...,
	)
}

func (dao *DAO) FindIDsWithDrink(
	ctx context.Context,
	drink string,
) ([]int64, error) {
	return dal.FindColumns[int64](ctx, dao.dal.DB(),
		"SELECT `id`"+
			" FROM "+Table()+
			" WHERE `drink` = ?;",
		drink,
	)
}
