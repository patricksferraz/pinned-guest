package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/pinned-guest/utils"
	"github.com/paemuri/brdoc"
)

func init() {
	govalidator.TagMap["doc"] = govalidator.Validator(func(str string) bool {
		return brdoc.IsCNPJ(str) || brdoc.IsCPF(str)
	})

	govalidator.SetFieldsRequiredByDefault(true)
}

type Pagination struct {
	Last     *time.Time `json:"last" valid:"-"`
	PageSize *int       `json:"page_size" valid:"-"`
}

func NewPagination(last, pageSize *int) (*Pagination, error) {
	var _last time.Time

	if pageSize == nil {
		pageSize = utils.PInt(10)
	}

	if *last != 0 {
		_last = time.UnixMicro(int64(*last))
	}
	e := Pagination{
		Last:     &_last,
		PageSize: pageSize,
	}

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Pagination) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
