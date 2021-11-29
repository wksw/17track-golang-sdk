package track17

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"time"
)

// userAgent gets user agent
// It has the SDK version information, OS information and GO version
func userAgent() string {
	sys := getSysInfo()
	return fmt.Sprintf("17track-sdk-go/%s (%s/%s/%s;%s)", SDK_VERSION, sys.name,
		sys.release, sys.machine, runtime.Version())
}

type sysInfo struct {
	name    string // OS name such as windows/Linux
	release string // OS version 2.6.32-220.23.2.ali1089.el5.x86_64 etc
	machine string // CPU type amd64/x86_64
}

// getSysInfo gets system info
// gets the OS information and CPU type
func getSysInfo() sysInfo {
	name := runtime.GOOS
	release := "-"
	machine := runtime.GOARCH
	if out, err := exec.Command("uname", "-s").CombinedOutput(); err == nil {
		name = string(bytes.TrimSpace(out))
	}
	if out, err := exec.Command("uname", "-r").CombinedOutput(); err == nil {
		release = string(bytes.TrimSpace(out))
	}
	if out, err := exec.Command("uname", "-m").CombinedOutput(); err == nil {
		machine = string(bytes.TrimSpace(out))
	}
	return sysInfo{name: name, release: release, machine: machine}
}

// HexEncodeSHA256Hash returns hexcode of sha256
func hexEncodeSHA256Hash(body []byte) (string, error) {
	hash := sha256.New()
	if body == nil {
		body = []byte("")
	}
	_, err := hash.Write(body)
	return fmt.Sprintf("%x", hash.Sum(nil)), err
}

func randStr(size int) string {
	rand.Seed(time.Now().UnixNano())
	kinds, result := [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		ikind := rand.Intn(3)
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func hmac256(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
