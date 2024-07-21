package models

type OtpCode struct {
	ID     string `json:"id" gorm:"primaryKey;column:id"`
	UserId string `json:"user_id" gorm:"column:user_id"`
	Code   string `json:"code" gorm:"column:code"`
}

func (r *OtpCode) tableName() string {
	return "otp_codes"
}
