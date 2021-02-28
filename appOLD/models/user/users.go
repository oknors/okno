package handlers

//User struct declaration
type User struct {
	ID       string
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Gender   string `json:"Gender"`
	Password string `json:"Password"`
}

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}
