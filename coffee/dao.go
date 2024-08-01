package coffee

import (
	"context"
	"fmt"

	"github.com/mrusme/kopi/dal"
	"github.com/mrusme/kopi/helpers"
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
	entity Coffee,
) (Coffee, error) {
	id, err := dal.Create(ctx, dao.dal.DB(),
		"INSERT INTO "+Table()+
			" ("+Columns(false)+")"+
			" VALUES ("+helpers.QueryArgRepeat(ColumnsNumber(false))+");",
		&entity.Roaster,
		&entity.Name,
		&entity.Origin,
		&entity.AltitudeLower,
		&entity.AltitudeUpper,
		&entity.Level,
		&entity.Flavors,
		&entity.Info,
		&entity.RoastingDate,
		&entity.Timestamp,
	)
	entity.ID = id
	return entity, err
}

func (dao *DAO) GetByID(
	ctx context.Context,
	id int64,
) (Coffee, error) {
	return dal.GetRow[Coffee](ctx, dao.dal.DB(),
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
) ([]Coffee, error) {
	placeholders, args := dal.InArgs(ids)
	q := fmt.Sprintf(
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `id` IN (%s)"+
			" ORDER BY `id`;",
		placeholders,
	)
	return dal.FindRows[Coffee](ctx, dao.dal.DB(),
		q,
		args...,
	)
}

func (dao *DAO) FindByOrigin(
	ctx context.Context,
	origin string,
) ([]Coffee, error) {
	return dal.FindRows[Coffee](ctx, dao.dal.DB(),
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `origin` LIKE ?;",
		"%"+origin+"%",
	)
}
