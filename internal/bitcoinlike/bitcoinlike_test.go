package bitcoinlike

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v := &Validator{
		SupportedTypes: []string{"P2PKH", "P2SH"},
	}

	msg := fmt.Sprintf("supported types are: %+q", v.SupportedTypes)
	unsupportedErrMsg := errors.WithMessage(ErrUnsupportedTypes, msg).Error()

	err := v.CheckTypes([]string{"P2PKH", "P2SH"})
	assert.Nil(t, err)

	err = v.CheckTypes([]string{"P", "P2PKH"})
	assert.EqualError(t, err, unsupportedErrMsg)

	err = v.CheckTypes(nil)
	assert.EqualError(t, err, ErrEmptyEnabledTypes.Error())
}
