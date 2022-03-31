package entity

import (
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
	PageToken *string `json:"page_token" valid:"-"`
	PageSize  *int    `json:"page_size" valid:"-"`
}

func NewPagination(pageToken *string, pageSize *int) (*Pagination, error) {
	if pageSize == nil {
		pageSize = utils.PInt(10)
	}

	e := Pagination{
		PageToken: pageToken,
		PageSize:  pageSize,
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
