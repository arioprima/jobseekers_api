package schemas

import "time"

type SchemaDataUser struct {
	ID           string    `json:"id" validate:"uuid"`
	BiodataId    string    `json:"biodata_id" validate:"uuid"`
	Firstname    string    `json:"firstname" validate:"required"`
	Lastname     string    `json:"lastname"`
	Email        string    `json:"email" validate:"required,email"`
	Phone        string    `json:"phone" validate:"required"`
	BirthDate    time.Time `json:"birth_date"`
	BirthPlace   string    `json:"birth_place"`
	ProvinceId   string    `json:"province_id"`
	CityId       string    `json:"city_id"`
	DistrictId   string    `json:"district_id"`
	Address      string    `json:"address"`
	EducationId  string    `json:"education_id"`
	Password     string    `json:"password" validate:"required,min=3,max=100"`
	ProfileImage string    `json:"profile_image"`
	RoleId       string    `json:"role_id"`
	Summary      string    `json:"summary"`
	OtpCode      string    `json:"otp_code"`
	Token        string    `json:"token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ExpiredAt    time.Time `json:"expired_at"`
}

type LoginUserResponse struct {
	ID           string    `json:"id"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname,omitempty"`
	Email        string    `json:"email"`
	RoleId       string    `json:"role_id"`
	RoleName     string    `json:"role_name"`
	ProfileImage *string   `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OtpEmailResponse struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	OtpCode string `json:"otp_code"`
}

type RegisterResponse struct {
	ID        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	RoleId    string    `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}
