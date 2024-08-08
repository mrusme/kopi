package label

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

func (dao *DAO) GetLabelsOfNonEmptyBags(
	ctx context.Context,
) ([]Label, error) {
	return dal.FindRows[Label](ctx, dao.dal.DB(),
		"SELECT `bags`.`id` AS `bag_id`, `bags`.`coffee_id` AS `coffee_id`,"+
			" `coffees`.`roaster` AS `roaster`, `coffees`.`name` AS `name`"+
			" FROM `bags`"+
			" INNER JOIN `coffees` ON `coffees`.`id` = `bags`.`coffee_id`"+
			" WHERE `bags`.`empty_date` IS NULL"+
			" ORDER BY `bags`.`open_date` ASC;",
	)
}
