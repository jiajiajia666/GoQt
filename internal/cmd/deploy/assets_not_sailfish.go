// +build !sailfish

package deploy

import "errors"

func sailfish_ssh(port, login string, cmd ...string) error {
	return errors.New("please run \"go install -v -tags=sailfish github.com/jiajiajia666/GoQt/cmd/...\" to enable sailfish deployments")
}
