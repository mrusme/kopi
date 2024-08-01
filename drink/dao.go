package drink

import (
	"context"

	"github.com/mrusme/kopi/dal"
)

type DAO struct {
	dal *dal.DAL
}

func NewDAO(dal *dal.DAL) *DAO {
	dao := new(DAO)

	dao.dal = dal

	return dao
}

func (dao *DAO) List(
	ctx context.Context,
) (Drink, error) {
	return dal.GetRow[Drink](ctx, dao.dal.DB(),
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			";",
	)
}

func (dao *DAO) GetByID(
	ctx context.Context,
	id int64,
) (Drink, error) {
	return dal.GetRow[Drink](ctx, dao.dal.DB(),
		"SELECT "+Columns(true)+
			" FROM "+Table()+
			" WHERE `id` = ?"+
			" LIMIT 1;",
		id,
	)
}
