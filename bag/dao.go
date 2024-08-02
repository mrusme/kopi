package bag

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/helpers"
)

type DAO struct {
	dal *dal.DAL
	val *validator.Validate
}

func NewDAO(dal *dal.DAL) *DAO {
	dao := new(DAO)

	dao.dal = dal
	dao.val = validator.New(validator.WithRequiredStructEnabled())

	return dao
}

func (dao *DAO) Create(
	ctx context.Context,
	entity Bag,
) (Bag, error) {
	if err := dao.val.Struct(entity); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return entity, validationErrors
	}

	entity.Timestamp = time.Now()
	id, err := dal.Create(ctx, dao.dal.DB(),
		"INSERT INTO "+Table()+
			" ("+Columns(false)+")"+
			" VALUES ("+helpers.QueryArgRepeat(ColumnsNumber(false))+");",
		&entity.CoffeeID,

		&entity.WeightG,
		&entity.Grind,

		&entity.RoastDate,
		&entity.OpenDate,
		&entity.EmptyDate,
		&entity.PurchaseDate,

		&entity.PriceUSDct,
		&entity.PriceSats,

		&entity.Timestamp,
	)
	entity.ID = id
	return entity, err
}

func (dao *DAO) GetByID(
	ctx context.Context,
	id int64,
) (Bag, error) {
	return dal.GetRow[Bag](ctx, dao.dal.DB(),
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
) ([]Bag, error) {
	placeholders, args := dal.InArgs(ids)
	q := fmt.Sprintf(
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `id` IN (%s)"+
			" ORDER BY `id`;",
		placeholders,
	)
	return dal.FindRows[Bag](ctx, dao.dal.DB(),
		q,
		args...,
	)
}

func (dao *DAO) FindByCoffeeID(
	ctx context.Context,
	ids []int64,
) ([]Bag, error) {
	q := fmt.Sprintf(
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `coffee_id` = ?"+
			" ORDER BY `id`;",
		ids,
	)
	return dal.FindRows[Bag](ctx, dao.dal.DB(),
		q,
	)
}
