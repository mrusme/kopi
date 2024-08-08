package equipment

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

func (dao *DAO) DB() *dal.DAL {
	return dao.dal
}

func (dao *DAO) Validate(entity Equipment) error {
	return helpers.Validate(dao.val, entity)
}

func (dao *DAO) ValidateField(entity Equipment, field string) error {
	return helpers.ValidateField(dao.val, entity, field)
}

func (dao *DAO) Create(
	ctx context.Context,
	entity Equipment,
) (Equipment, error) {
	if err := dao.Validate(entity); err != nil {
		return entity, err
	}

	entity.Timestamp = time.Now()
	id, err := dal.Create(ctx, dao.dal.DB(),
		"INSERT INTO "+Table()+
			" ("+Columns(false)+")"+
			" VALUES ("+helpers.QueryArgRepeat(ColumnsNumber(false))+");",
		&entity.Name,
		&entity.Description,

		&entity.Type,

		&entity.PurchaseDate,
		&entity.DecommissionDate,

		&entity.PriceUSDct,
		&entity.PriceSats,

		&entity.Timestamp,
	)
	entity.ID = id
	return entity, err
}

func (dao *DAO) List(
	ctx context.Context,
) ([]Equipment, error) {
	return dal.FindRows[Equipment](ctx, dao.dal.DB(),
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			";",
	)
}

func (dao *DAO) GetByID(
	ctx context.Context,
	id int64,
) (Equipment, error) {
	return dal.GetRow[Equipment](ctx, dao.dal.DB(),
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
) ([]Equipment, error) {
	placeholders, args := dal.InArgs(ids)
	q := fmt.Sprintf(
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `id` IN (%s)"+
			" ORDER BY `id`;",
		placeholders,
	)
	return dal.FindRows[Equipment](ctx, dao.dal.DB(),
		q,
		args...,
	)
}
