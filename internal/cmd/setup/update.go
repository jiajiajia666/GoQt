package setup

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/jiajiajia666/GoQt/internal/utils"
)

func Update() {
	utils.Log.Info("running: 'qtsetup update'")

	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/jiajiajia666/GoQt/cmd/..."), "run \"go clean cmd\"")
	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/jiajiajia666/GoQt/internal/..."), "run \"go clean internal\"")

	fetch := exec.Command("git", "fetch", "-f", "--all")
	fetch.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "jiajiajia666", "qt")
	utils.RunCmd(fetch, "run \"git fetch\"")

	checkoutCmd := exec.Command("git", "checkout", "-f", "--", utils.GoQtPkgPath("cmd"))
	checkoutCmd.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "jiajiajia666", "qt")
	utils.RunCmd(checkoutCmd, "run \"git checkout cmd\"")

	checkoutInternal := exec.Command("git", "checkout", "-f", "--", utils.GoQtPkgPath("internal"))
	checkoutInternal.Dir = filepath.Join(utils.MustGoPath(), "src", "github.com", "jiajiajia666", "qt")
	utils.RunCmd(checkoutInternal, "run \"git checkout internal\"")

	hash := "please install git"
	if _, err := exec.LookPath("git"); err == nil {
		cmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
		cmd.Dir = utils.GoQtPkgPath()
		hash = strings.TrimSpace(utils.RunCmdOptional(cmd, "get git hash"))
	}

	utils.RunCmd(exec.Command("go", "install", "-v", fmt.Sprintf("-ldflags=\"-X=github.com/jiajiajia666/GoQt/internal/cmd.buildVersion=%v\"", hash), fmt.Sprintf("github.com/jiajiajia666/GoQt/cmd/...")), "run \"go install\"")

	Prep(runtime.GOOS)
}

func Upgrade() {
	utils.Log.Info("running: 'qtsetup upgrade'")

	utils.RunCmd(exec.Command("go", "clean", "-i", "github.com/jiajiajia666/GoQt/..."), "run \"go clean\"")
	utils.RemoveAll(utils.GoQtPkgPath())

	utils.RunCmd(exec.Command("go", "get", "-v", "-d", "github.com/jiajiajia666/GoQt/cmd/..."), "run \"go get\"")

	hash := "please install git"
	if _, err := exec.LookPath("git"); err == nil {
		cmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
		cmd.Dir = utils.GoQtPkgPath()
		hash = strings.TrimSpace(utils.RunCmdOptional(cmd, "get git hash"))
	}

	utils.RunCmd(exec.Command("go", "install", "-v", fmt.Sprintf("-ldflags=\"-X=github.com/jiajiajia666/GoQt/internal/cmd.buildVersion=%v\"", hash), fmt.Sprintf("github.com/jiajiajia666/GoQt/cmd/...")), "run \"go install\"")
}
