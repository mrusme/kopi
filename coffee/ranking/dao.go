package ranking

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

func (dao *DAO) GetRanking(
	ctx context.Context,
) ([]Ranking, error) {
	return dal.FindRows[Ranking](ctx, dao.dal.DB(),
		"WITH AvgRatings AS ("+
			" SELECT `bags`.`coffee_id`, AVG(`rating`) AS `avg_rating`"+
			" FROM `cups`"+
			" INNER JOIN `bags` ON `bags`.`id` = `cups`.`bag_id`"+
			" GROUP BY `coffee_id`"+
			"),"+
			"RankedCoffees AS ("+
			" SELECT `coffee_id`, `avg_rating`,"+
			" RANK() OVER (ORDER BY `avg_rating` DESC) AS `ranking`"+
			" FROM AvgRatings"+
			")"+
			"SELECT `coffee_id`, `avg_rating`, `ranking`"+
			" FROM RankedCoffees"+
			" ORDER BY `ranking`;",
	)
}
