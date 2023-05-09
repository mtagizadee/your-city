package countries

import (
	"net/http"
	"your-city/packages/common"
	"your-city/packages/db"
	"your-city/packages/utils"

	"gorm.io/gorm"
)

type Country struct {
	Id int `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Code string `gorm:"not null;unique;size:3" json:"code"`
}

type coutriesService struct {}

func (service *coutriesService) getAll(code string) ([]Country, *common.ErrorType) {	
	var countries []Country

	db := db.GetDB()

	var res *gorm.DB
	if code != "" {
		res = db.Where("code LIKE ?", "%" + code + "%").Find(&countries)
	} else { res = db.Find(&countries) }

	if err := res.Error; err != nil {
		return nil, utils.DefaultError(err)
	}

	return countries, nil
}

func (service *coutriesService) create(dto *createCountryDto) (*Country, *common.ErrorType) {
	db := db.GetDB()
	
	var country Country = Country{
		Name: dto.Name,
		Code: dto.Code,
	}

	if err := db.Create(&country).Error; err != nil {
		if utils.IsUniqueKeyError(err) {
			return nil, &common.ErrorType{Status: http.StatusConflict, Message: "country with this code already exists"}
		}

		return nil, utils.DefaultError(err)
	}

	return &country, nil
}