package bitcoinlike

import (
	"fmt"
	"testing"

	"github.com/LanfordCai/ava/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v := &Validator{
		SupportedTypes: []string{"P2PKH", "P2SH"},
	}

	msg := fmt.Sprintf("Supported types are: %+q", v.SupportedTypes)
	unsupportedErrMsg := errors.WithMessage(common.ErrUnsupportedTypes, msg).Error()

	err := v.CheckTypes([]string{"P2PKH", "P2SH"})
	assert.Nil(t, err)

	err = v.CheckTypes([]string{"P", "P2PKH"})
	assert.EqualError(t, err, unsupportedErrMsg)

	err = v.CheckTypes([]string{})
	assert.EqualError(t, err, common.ErrEmptyEnabledTypes.Error())
}
