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

func (dao *DAO) DB() *dal.DAL {
	return dao.dal
}

func (dao *DAO) Validate(entity Bag) error {
	return helpers.Validate(dao.val, entity)
}

func (dao *DAO) ValidateField(entity Bag, field string) error {
	return helpers.ValidateField(dao.val, entity, field)
}

func (dao *DAO) Create(
	ctx context.Context,
	entity Bag,
) (Bag, error) {
	if err := dao.Validate(entity); err != nil {
		return entity, err
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

func (dao *DAO) ListNonEmpty(
	ctx context.Context,
) ([]Bag, error) {
	return dal.FindRows[Bag](ctx, dao.dal.DB(),
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `empty_date` = NULL"+
			" ORDER BY `open_date` ASC;",
	)
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
	id int64,
) ([]Bag, error) {
	q := fmt.Sprintf(
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `coffee_id` = %d"+
			" ORDER BY `id`;",
		id,
	)
	return dal.FindRows[Bag](ctx, dao.dal.DB(),
		q,
	)
}

func (dao *DAO) FindOpenByCoffeeID(
	ctx context.Context,
	id int64,
) ([]Bag, error) {
	q := fmt.Sprintf(
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `coffee_id` = %d"+
			" AND `empty_date` IS NULL"+
			" ORDER BY `id`;",
		id,
	)
	return dal.FindRows[Bag](ctx, dao.dal.DB(),
		q,
	)
}

func (dao *DAO) FindOpenByCoffeeIDWithAtLeast(
	ctx context.Context,
	id int64,
	atLeastG int,
) ([]Bag, error) {
	q := fmt.Sprintf(
		"WITH consumed_coffee AS ("+
			" SELECT `bag_id`, COALESCE(SUM(coffee_g), 0) AS total_consumed"+
			" FROM `cups`"+
			" GROUP BY `bag_id`"+
			") "+
			"SELECT "+Columns(true)+
			" FROM "+Table()+
			" LEFT JOIN consumed_coffee ON `id` = consumed_coffee.`bag_id`"+
			" WHERE `coffee_id` = %d"+
			" AND `empty_date` IS NULL"+
			" AND (`weight_g` - COALESCE(consumed_coffee.total_consumed, 0)) >= %d"+
			" ORDER BY `id`;",
		id,
		atLeastG,
	)
	return dal.FindRows[Bag](ctx, dao.dal.DB(),
		q,
	)
}
