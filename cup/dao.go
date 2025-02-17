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
	dao.val.RegisterValidation("is_idslist", helpers.IDsListValidation)

	return dao
}

func (dao *DAO) DB() *dal.DAL {
	return dao.dal
}

func (dao *DAO) Validate(entity Cup) error {
	return helpers.Validate(dao.val, entity)
}

func (dao *DAO) ValidateField(entity Cup, field string) error {
	return helpers.ValidateField(dao.val, entity, field)
}

func (dao *DAO) Create(
	ctx context.Context,
	entity Cup,
) (Cup, error) {
	if err := dao.Validate(entity); err != nil {
		return entity, err
	}

	entity.Timestamp = time.Now()
	id, err := dal.Create(ctx, dao.dal.DB(),
		"INSERT INTO "+Table()+
			" ("+Columns(false)+")"+
			" VALUES ("+helpers.QueryArgRepeat(ColumnsNumber(false))+");",
		&entity.BagID,

		&entity.Method,
		&entity.Drink,

		&entity.EquipmentIDs,

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

func (dao *DAO) GetAvgRatingForBagID(
	ctx context.Context,
	id int64,
) (float64, error) {
	return dal.GetColumn[float64](ctx, dao.dal.DB(),
		"SELECT AVG(rating)"+
			" FROM "+Table()+
			" WHERE `bag_id` = ?;",
		id,
	)
}

func (dao *DAO) GetAvgRatingForCoffeeID(
	ctx context.Context,
	id int64,
) (float64, error) {
	return dal.GetColumn[float64](ctx, dao.dal.DB(),
		"SELECT AVG(rating)"+
			" FROM "+Table()+
			" INNER JOIN `bags` ON `bags`.`id` = "+Table()+".`bag_id`"+
			" WHERE `bags`.`coffee_id` = ?;",
		id,
	)
}

func (dao *DAO) GetCupsForPeriod(
	ctx context.Context,
	from time.Time,
	until time.Time,
) (int64, error) {
	return dal.GetColumn[int64](ctx, dao.dal.DB(),
		"SELECT IFNULL(COUNT(), 0)"+
			" FROM "+Table()+
			" WHERE "+Table()+".`timestamp` BETWEEN ? AND ?;",
		from,
		until,
	)
}

func (dao *DAO) GetCupsForPeriodByBagID(
	ctx context.Context,
	from time.Time,
	until time.Time,
	id int64,
) (int64, error) {
	return dal.GetColumn[int64](ctx, dao.dal.DB(),
		"SELECT IFNULL(COUNT(), 0)"+
			" FROM "+Table()+
			" WHERE "+Table()+".`bag_id` = ? AND "+Table()+".`timestamp` BETWEEN ? AND ?;",
		id,
		from,
		until,
	)
}

func (dao *DAO) GetCaffeineForPeriod(
	ctx context.Context,
	from time.Time,
	until time.Time,
) (float64, error) {
	return dal.GetColumn[float64](ctx, dao.dal.DB(),
		"SELECT"+
			" IFNULL("+
			" SUM("+Table()+".`coffee_g` * `methods`.`caffeine_mg_extraction_yield_per_g` * "+
			" (CASE `coffees`.`level`"+
			"  WHEN 'light' THEN 0.95"+
			"  WHEN 'medium' THEN 1.00"+
			"  WHEN 'dark' THEN 1.10"+
			" END)"+
			" * (1 - `methods`.`caffeine_loss_factor`))"+
			", 0)"+
			" FROM "+Table()+
			" INNER JOIN `methods` ON `methods`.`id` = "+Table()+".`method`"+
			" INNER JOIN `bags` ON `bags`.`id` = "+Table()+".`bag_id`"+
			" INNER JOIN `coffees` ON `coffees`.`id` = `bags`.`coffee_id`"+
			" WHERE `coffees`.`decaf` = FALSE"+
			"   AND "+Table()+".`timestamp` BETWEEN ? AND ?;",
		from,
		until,
	)
}

func (dao *DAO) GetWaterForPeriod(
	ctx context.Context,
	from time.Time,
	until time.Time,
) (int64, error) {
	return dal.GetColumn[int64](ctx, dao.dal.DB(),
		"SELECT IFNULL(SUM(`brew_ml` + `water_ml`), 0)"+
			" FROM "+Table()+
			" WHERE "+Table()+".`timestamp` BETWEEN ? AND ?;",
		from,
		until,
	)
}

func (dao *DAO) GetMilkForPeriod(
	ctx context.Context,
	from time.Time,
	until time.Time,
) (int64, error) {
	return dal.GetColumn[int64](ctx, dao.dal.DB(),
		"SELECT IFNULL(SUM(`milk_ml`), 0)"+
			" FROM "+Table()+
			" WHERE "+Table()+".`timestamp` BETWEEN ? AND ?;",
		from,
		until,
	)
}

func (dao *DAO) GetRealMilkForPeriod(
	ctx context.Context,
	from time.Time,
	until time.Time,
) (int64, error) {
	return dao.GetTypeMilkForPeriod(ctx, false, from, until)
}

func (dao *DAO) GetPlantMilkForPeriod(
	ctx context.Context,
	from time.Time,
	until time.Time,
) (int64, error) {
	return dao.GetTypeMilkForPeriod(ctx, true, from, until)
}

func (dao *DAO) GetTypeMilkForPeriod(
	ctx context.Context,
	vegan bool,
	from time.Time,
	until time.Time,
) (int64, error) {
	return dal.GetColumn[int64](ctx, dao.dal.DB(),
		"SELECT IFNULL(SUM(`milk_ml`), 0)"+
			" FROM "+Table()+
			" WHERE `vegan` = ? AND "+Table()+".`timestamp` BETWEEN ? AND ?;",
		vegan,
		from,
		until,
	)
}

func (dao *DAO) GetCoffeeLeftByBagID(
	ctx context.Context,
	id int64,
) (int64, error) {
	return dal.GetColumn[int64](ctx, dao.dal.DB(),
		"SELECT `bags`.`weight_g` - IFNULL(SUM(`coffee_g`), 0)"+
			" FROM "+Table()+
			" INNER JOIN `bags` ON `bags`.`id` = "+Table()+".`bag_id`"+
			" WHERE `bag_id` = ?;",
		id,
	)
}
