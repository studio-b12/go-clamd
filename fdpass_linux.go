package clamd

import (
	"errors"
	"fmt"
	"net"
	"os"
	"syscall"
)

/*
Scan file using file descriptor
*/
func (c *Clamd) ScanFileFdpass(file *os.File) (chan *ScanResult, error) {

	conn, err := c.newConnection()
	if err != nil {
		return nil, err
	}

	fds := []int{int(file.Fd())}
	rights := syscall.UnixRights(fds...)

	command := fmt.Sprintf("FILDES")
	conn.sendCommand(command)

	unixConn, ok := conn.Conn.(*net.UnixConn)
	if !ok {
		return nil, errors.New("fpass only supported by unix sockets")
	}
	socketFile, err := unixConn.File()
	if err != nil {
		return nil, err
	}

	socket := int(socketFile.Fd())
	defer socketFile.Close()

	err = syscall.Sendmsg(socket, nil, rights, nil, 0)
	if err != nil {
		return nil, err
	}

	ch, wg, err := conn.readResponse()

	go func() {
		wg.Wait()
		conn.Close()
	}()

	return ch, err
}
