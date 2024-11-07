package libs

type ErrorMessage struct {
	Message string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type UserDB struct {
	UserCode string `gorm:"column:UserCode"`
	Password string `gorm:"column:Password"`
}

type Product struct {
	ID         int16
	Remark     string
	FirstPrice int64
}
