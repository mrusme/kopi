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

func (dao *DAO) DB() *dal.DAL {
	return dao.dal
}

func (dao *DAO) List(
	ctx context.Context,
	onlyNonEmpty bool,
) ([]Label, error) {
	cond := ""
	if onlyNonEmpty {
		cond = " WHERE `bags`.`empty_date` IS NULL"
	}
	return dal.FindRows[Label](ctx, dao.dal.DB(),
		"SELECT"+
			" `bags`.`id` AS `bag_id`, `bags`.`coffee_id` AS `coffee_id`,"+
			" `coffees`.`roaster` AS `roaster`,"+
			" `coffees`.`name` AS `name`,"+
			" `coffees`.`origin` AS `origin`,"+
			" `coffees`.`altitude_lower_m` AS `altitude_lower_m`,"+
			" `coffees`.`altitude_upper_m` AS `altitude_upper_m`,"+
			" `coffees`.`level` AS `level`,"+
			" `coffees`.`flavors` AS `flavors`,"+
			" `coffees`.`info` AS `info`,"+
			" `coffees`.`decaf` AS `decaf`,"+

			" `bags`.`weight_g` AS `weight_g`,"+
			" `bags`.`grind` AS `grind`,"+

			" `bags`.`roast_date` AS `roast_date`,"+
			" `bags`.`open_date` AS `open_date`,"+
			" `bags`.`empty_date` AS `empty_date`,"+
			" `bags`.`purchase_date` AS `purchase_date`,"+

			" `bags`.`price_usd_ct` AS `price_usd_ct`,"+
			" `bags`.`price_sats` AS `price_sats`"+

			" FROM `bags`"+
			" INNER JOIN `coffees` ON `coffees`.`id` = `bags`.`coffee_id`"+
			cond+
			" ORDER BY `bags`.`open_date` ASC;",
	)
}
