package log

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

func (dao *DAO) Validate(entity Log) error {
	return helpers.Validate(dao.val, entity)
}

func (dao *DAO) ValidateField(entity Log, field string) error {
	return helpers.ValidateField(dao.val, entity, field)
}

func (dao *DAO) Create(
	ctx context.Context,
	entity Log,
) (Log, error) {
	if err := dao.Validate(entity); err != nil {
		return entity, err
	}

	entity.Timestamp = time.Now()
	id, err := dal.Create(ctx, dao.dal.DB(),
		"INSERT INTO "+Table()+
			" ("+Columns(false)+")"+
			" VALUES ("+helpers.QueryArgRepeat(ColumnsNumber(false))+");",
		&entity.EquipmentID,

		&entity.Key,
		&entity.Value,

		&entity.Timestamp,
	)
	entity.ID = id
	return entity, err
}

func (dao *DAO) GetByID(
	ctx context.Context,
	id int64,
) (Log, error) {
	return dal.GetRow[Log](ctx, dao.dal.DB(),
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
) ([]Log, error) {
	placeholders, args := dal.InArgs(ids)
	q := fmt.Sprintf(
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `id` IN (%s)"+
			" ORDER BY `id`;",
		placeholders,
	)
	return dal.FindRows[Log](ctx, dao.dal.DB(),
		q,
		args...,
	)
}

func (dao *DAO) FindByEquipmentID(
	ctx context.Context,
	id int64,
) ([]Log, error) {
	return dal.FindRows[Log](ctx, dao.dal.DB(),
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `equipment_id` = ?"+
			" ORDER BY `id`;",
		id,
	)
}
