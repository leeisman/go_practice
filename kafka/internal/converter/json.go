package converter

import (
	"encoding/json"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.silkrode.com.tw/golang/errors"
)

// JSON ...
func JSON(in interface{}, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return errors.Wrap(errors.ErrInternalError, err.Error())
	}
	err = json.Unmarshal(b, out)
	if err != nil {
		return errors.Wrap(errors.ErrInternalError, err.Error())
	}
	return nil
}

// JSONB ...
func JSONB(in interface{}) (postgres.Jsonb, error) {
	jsonb := postgres.Jsonb{}
	b, err := json.Marshal(in)
	if err != nil {
		return jsonb, errors.Wrapf(errors.ConvertPostgresError(err), "%v", err)
	}
	jsonb.RawMessage = b
	return jsonb, nil
}
