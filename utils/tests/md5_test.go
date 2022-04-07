package utils_tests

import (
	"crypto/sha512"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/breadinator/swkshp/utils"
	"github.com/stretchr/testify/assert"
)

var paths [3]string = [...]string{
	"./md5_test.go",
	"./goquery_test.go",
	`D:\Screen Captures\Desktop\Desktop Screenshot 2022.03.27 - 20.47.10.75.png`,
}

func Test_GetFileMD5_DIFFERENT(t *testing.T) {
	for i := 1; i < len(paths); i++ {
		a, err := utils.GetFileMD5(paths[i], 512)
		assert.Nil(t, err)
		b, err := utils.GetFileMD5(paths[i-1], 512)
		assert.Nil(t, err)
		assert.False(t, utils.SlicesEqual(a, b))
		fmt.Printf("%x\n", a)
	}
}

func Test_GetFileMD5_BUFSIZES(t *testing.T) {
	for _, path := range paths {
		powers := [...]uint16{0, 1, 7, 9, 13}
		sums := *new([len(powers)][]byte)
		var err error
		for i, power := range powers {
			sums[i], err = utils.GetFileMD5(path, 2^power)
			assert.Nil(t, err)
			if i != 0 {
				assert.Equal(t, sums[i], sums[i-1])
			}
		}
	}
}

func Test_GetFileMD5_404(t *testing.T) {
	_, err := utils.GetFileMD5("C:/this/doesn't/exist", 512)
	assert.NotNil(t, err)
}

func Test_GetFileMD5_CertUtil(t *testing.T) {
	re := regexp.MustCompile(`:\s*([a-f\d]+)\s*CertUtil:`)
	for _, path := range paths {
		a, _ := utils.GetFileMD5(path, 512)

		abs, _ := filepath.Abs(path)
		output, _ := exec.Command("CertUtil", "-hashfile", abs, "md5").Output()
		b := re.FindStringSubmatch(string(output))[1]

		assert.Equal(t, b, fmt.Sprintf("%x", a))
	}
}

func Test_GetFileHash_SHA512(t *testing.T) {
	for _, path := range paths {
		_, err := utils.GetFileHash(path, 512, sha512.New())
		assert.Nil(t, err)
	}
}
