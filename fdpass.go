//+build !linux

package clamd

import (
	"errors"
	"os"
)

/*
Scan file using file descriptor
*/
func (c *Clamd) ScanFileFdpass(file *os.File) (chan *ScanResult, error) {
	return nil, errors.New("not supported")
}
