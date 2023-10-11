package database

import "gorm.io/gorm"

type paginate struct {
	limit int
	page  int
}

func NewPaginate(limit int, page int) *paginate {
	return &paginate{limit: limit, page: page}
}

func (p *paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit

	return db.
		Order("create_date desc, (joy + sorrow + anger + surprise + under_exposed + blurred + headwear) desc").
		Offset(offset).
		Limit(p.limit)
}
