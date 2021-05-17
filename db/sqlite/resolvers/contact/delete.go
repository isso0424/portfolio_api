package contact

import (
	"github.com/isso0424/portfolio_api/db/sqlite/model"
	"github.com/isso0424/portfolio_api/domain"
	"github.com/isso0424/portfolio_api/types"
	"github.com/isso0424/portfolio_api/types/internal_error"
)

func(db *ContactDB) Delete(id string) (*domain.Contact, types.APIError) {
	newData := model.Contact{}
	newData.ID = id
	err := db.db.Delete(&newData).Error
	if err != nil {
		return nil, internal_error.New(err.Error())
	}

	return &newData.Contact, nil
}
