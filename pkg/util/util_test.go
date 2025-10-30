package util_test

import (
	"os"
	"package-name/pkg/util"
	"path"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	go_test_ "github.com/pefish/go-test"
)

func init() {
	projectRoot, _ := go_test_.ProjectRoot()
	envMap, err := godotenv.Read(path.Join(projectRoot, ".env"))
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return
		}
		panic(err)
	}
	for k, v := range envMap {
		os.Setenv(k, v)
	}
	// init something...
}

func TestTest(t *testing.T) {
	r := util.Test(1, 2)
	go_test_.Equal(t, r, uint64(3))
}
