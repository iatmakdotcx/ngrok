// +build !release,!autoupdate

package assets

import (
	"io/ioutil"
)

func Asset(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}
