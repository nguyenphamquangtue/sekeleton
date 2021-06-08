package storages

// User reflects users data from DB
type User struct {
	ID       string `pg:"id,type:text,pk" json:"id"`
	Password string `pg:"password,type:text" json:"password"`
}

func (User) TableName() string {
	return "users"
}

type EntityUseCase interface {
	GetAuthToken(args map[string]string) (map[string]interface{}, error)
	RegisterUser(user User) (map[string]interface{}, error)
}

type EntityRepository interface {
	Login(userID string, pwd string) error
	InsertUser(user *User) error
	GetUserByID(username string) (User, error)
}
