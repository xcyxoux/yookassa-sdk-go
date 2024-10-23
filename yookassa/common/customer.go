package yoocommon

type Customer struct {
	// all optional
	FullName string `json:"full_name"`
	INN      string `json:"inn"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
