package cup

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
	entity Cup,
) (Cup, error) {
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

		&entity.Method,
		&entity.Drink,

		&entity.CoffeeG,
		&entity.BrewMl,
		&entity.WaterMl,
		&entity.MilkMl,
		&entity.SugarG,
		&entity.Vegan,

		&entity.Rating,
		&entity.Timestamp,
	)
	entity.ID = id
	return entity, err
}

func (dao *DAO) List(
	ctx context.Context,
) ([]Cup, error) {
	return dal.FindRows[Cup](ctx, dao.dal.DB(),
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			";",
	)
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

func (dao *DAO) GetAvgRatingForCoffeeID(
	ctx context.Context,
	id int64,
) (float64, error) {
	return dal.GetColumn[float64](ctx, dao.dal.DB(),
		"SELECT AVG(rating)"+
			" FROM "+Table()+
			" WHERE `coffee_id` = ?;",
		id,
	)
}

func (dao *DAO) GetCaffeineForPeriod(
	ctx context.Context,
	from time.Time,
	until time.Time,
) (float64, error) {
	return dal.GetColumn[float64](ctx, dao.dal.DB(),
		"SELECT SUM("+Table()+".`brew_ml` * `methods`.`caffeine_multiplier_per_ml`)"+
			" FROM "+Table()+
			" INNER JOIN `methods` ON `methods`.`id` = "+Table()+".`method`"+
			" WHERE "+Table()+".`timestamp` BETWEEN ? AND ?;",
		from,
		until,
	)
}
