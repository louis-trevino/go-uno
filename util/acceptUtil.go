package util

import (
	"bufio"
	"os"
	"regexp"
)

/*
 * Author: Louis Trevino
 * Copyright(C) Torino Consulting, 2020.
 *
 * Compiled and tested using Go version go1.13.8 windows/amd64
 *
 *
 *   - fmt.Scan() is ignoring characters after 1st space, e.g. in "I have a cat" only "I" will be read.
 *   - fmt.Scanln() throwing warning with "Enter" (newline) or at least if running Windows 10 Home Edition.
 */
func AcceptLn() (usrInput string, err error) {
	reader := bufio.NewReader(os.Stdin)
	if usrInput, err = reader.ReadString('\n'); err != nil {
		return usrInput, err
	} else {
		//usrInput = strings.Replace(usrInput, "\r\n", "", -1)
		re := regexp.MustCompile("\r?\n")
		usrInput = re.ReplaceAllString(usrInput, "")

	}
	//fmt.Printf("usrInput: %s \n", usrInput)
	return usrInput, err
}
