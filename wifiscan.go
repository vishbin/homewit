package main
/*
	UTIL to get the name and some more detail of the wifi this server is connected .

 */
import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"fmt"
)

const osxCmd = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"
const osxArgs = "-I"
const linuxCmd = "iwgetid"
const linuxArgs = "--raw"

func WifiName() string {
	platform := runtime.GOOS
	if platform == "darwin" {
		return scanOSX()
	}else{
		return scanLinux()

	}


	return ""
}

func scanLinux() string {
	cmd := exec.Command(linuxCmd, linuxArgs)
	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	var str string

	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}

	name := strings.Replace(str, "\n", "", -1)
	return name
}

func scanOSX() string {

	cmd := exec.Command(osxCmd, osxArgs)

	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	var str string

	if b, err := ioutil.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}
	fmt.Println(str)

	r := regexp.MustCompile(`s*SSID: (.+)s*`)

	name := r.FindAllStringSubmatch(str, -1)
	fmt.Println(name)

	if len(name) <= 1 {
		return "Not Found"
	} else {
		return name[1][1]
	}
}

func panicIf(err error) {
	if err != nil {
fmt.Println(err)		
panic(err)
	}
}

func main(){

	fmt.Println(WifiName())

}
