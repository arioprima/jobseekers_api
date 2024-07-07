package models

import "time"

type Biodata struct {
	ID          string    `json:"id" validate:"uuid"`
	Firstname   string    `json:"firstname" validate:"required"`
	Lastname    string    `json:"lastname,omitempty"`
	Email       string    `json:"email" validate:"required,email"`
	BirthDate   time.Time `json:"birth_date,omitempty"`
	BirthPlace  string    `json:"birth_place,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	ProvinceId  string    `json:"province_id,omitempty" gorm:"column:province_id"`
	CityId      string    `json:"city_id,omitempty" gorm:"column:city_id"`
	DistrictId  string    `json:"district_id,omitempty" gorm:"column:district_id"`
	Address     string    `json:"address,omitempty"`
	EducationId string    `json:"education_id,omitempty" gorm:"column:education_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Biodata) tableName() string {
	return "biodata"
}
