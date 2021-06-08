package libs

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserAuthKey int8

func UserIDFromCtx(ctx context.Context) (string, bool) {

	v := ctx.Value(UserAuthKey(0))
	id, ok := v.(string)

	return id, ok
}

func Value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func HashPassword(password string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	result := string(hashedPassword)
	return &result, nil
}

func ComparePassword(hashpassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	if err != nil {
		return errors.New("Password mismatch!")
	}
	return nil
}
