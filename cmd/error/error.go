package error

import (
	"fmt"

	cmdutil "github.com/i2tsuki/mkr-graph/cmd/cmdutil"

	errors "github.com/pkg/errors"
)

func NewError(f cmdutil.Factory, err error) error {
	err = errors.Wrap(err, fmt.Sprintf(
		f.App.Msg,
		f.App.Name,
		f.App.Version,
	))

	return err
}
