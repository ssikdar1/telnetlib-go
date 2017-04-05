//TO DO package telnetlib
package main

/**
Partial port of python telnetlib.py

https://svn.python.org/projects/python/tags/r221/Lib/telnetlib.py


>>> from telnetlib import Telnet
>>> tn = Telnet('www.python.org', 79)   # connect to finger port
>>> tn.write('guido\r\n')
>>> print tn.read_all()


GOAL for gobotics:

https://docs.python.org/3/library/telnetlib.html#telnet-example

#!/usr/bin/env python
import getpass
import sys
import telnetlib

HOST = "localhost"
user = raw_input("Enter your remote account: ")
password = getpass.getpass()

tn = telnetlib.Telnet(HOST)

tn.read_until("login: ")
tn.write(user + "\n")
if password:
    tn.read_until("Password: ")
        tn.write(password + "\n")

        tn.write("ls\n")
        tn.write("exit\n")

        print tn.read_all()

**/

import (
	"fmt"
	"io/ioutil"
	"net"
)

type Telnet struct {
	host string
	port int
	eof  int
	conn net.Conn
	/*
	   self.debuglevel = DEBUGLEVEL
	   self.host = host
	   self.port = port
	   self.sock = None
	   self.rawq = ''
	   self.irawq = 0
	   self.cookedq = ''
	   self.eof = 0
	   self.option_callback = None
	*/
}

// Telnet protocol options code (don't change)
const BINARY rune = rune(0)          // 8-bit data path
const ECHO rune = rune(1)            // echo
const RCP rune = rune(2)             // prepare to reconnect
const SGA rune = rune(3)             // suppress go ahead
const NAMS rune = rune(4)            // approximate message size
const STATUS rune = rune(5)          // give status
const TM rune = rune(6)              // timing mark
const RCTE rune = rune(7)            // remote controlled transmission and echo
const NAOL rune = rune(8)            // negotiate about output line width
const NAOP rune = rune(9)            // negotiate about output page size
const NAOCRD rune = rune(10)         // negotiate about CR disposition
const NAOHTS rune = rune(11)         // negotiate about horizontal tabstops
const NAOHTD rune = rune(12)         // negotiate about horizontal tab disposition
const NAOFFD rune = rune(13)         // negotiate about formfeed disposition
const NAOVTS rune = rune(14)         // negotiate about vertical tab stops
const NAOVTD rune = rune(15)         // negotiate about vertical tab disposition
const NAOLFD rune = rune(16)         // negotiate about output LF disposition
const XASCII rune = rune(17)         // extended ascii character set
const LOGOUT rune = rune(18)         // force logout
const BM rune = rune(19)             // byte macro
const DET rune = rune(20)            // data entry terminal
const SUPDUP rune = rune(21)         // supdup protocol
const SUPDUPOUTPUT rune = rune(22)   // supdup output
const SNDLOC rune = rune(23)         // send location
const TTYPE rune = rune(24)          // terminal type
const EOR rune = rune(25)            // end or record
const TUID rune = rune(26)           // TACACS user identification
const OUTMRK rune = rune(27)         // output marking
const TTYLOC rune = rune(28)         // terminal location number
const VT3270REGIME rune = rune(29)   // 3270 regime
const X3PAD rune = rune(30)          // X.3 PAD
const NAWS rune = rune(31)           // window size
const TSPEED rune = rune(32)         // terminal speed
const LFLOW rune = rune(33)          // remote flow control
const LINEMODE rune = rune(34)       // Linemode option
const XDISPLOC rune = rune(35)       // X Display Location
const OLD_ENVIRON rune = rune(36)    // Old - Environment variables
const AUTHENTICATION rune = rune(37) // Authenticate
const ENCRYPT rune = rune(38)        // Encryption option
const NEW_ENVIRON rune = rune(39)    // New - Environment variables
// the following ones come from
// http://www.iana.org/assignments/telnet-options
// Unfortunately, that document does not assign identifiers
// to all of them, so we are making them up
const TN3270E rune = rune(40)             // TN3270E
const XAUTH rune = rune(41)               // XAUTH
const CHARSET rune = rune(42)             // CHARSET
const RSP rune = rune(43)                 // Telnet Remote Serial Port
const COM_PORT_OPTION rune = rune(44)     // Com Port Control Option
const SUPPRESS_LOCAL_ECHO rune = rune(45) // Telnet Suppress Local Echo
const TLS rune = rune(46)                 // Telnet Start TLS
const KERMIT rune = rune(47)              // KERMIT
const SEND_URL rune = rune(48)            // SEND-URL
const FORWARD_X rune = rune(49)           // FORWARD_X
const PRAGMA_LOGON rune = rune(138)       // TELOPT PRAGMA LOGON
const SSPI_LOGON rune = rune(139)         // TELOPT SSPI LOGON
const PRAGMA_HEARTBEAT rune = rune(140)   // TELOPT PRAGMA HEARTBEAT
const EXOPL rune = rune(255)              // Extended-Options-List

func telnet(host string, port int) Telnet {
	ret := Telnet{}
	if host != "" {
		ret = Telnet{host: host, port: port}
		ret.open()
	}
	return ret
}

func (telnet *Telnet) open() {
	telnet.eof = 0

	addr_list, _ := net.LookupHost(telnet.host)
	for _, addr := range addr_list {
		dest := fmt.Sprintf("%s:%d", addr, telnet.port)
		conn, err := net.Dial("tcp", dest)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Println(conn)
			telnet.conn = conn
			break
		}
	}
	if telnet.conn == nil {
		panic("No connection made")
	}
}

//Here python uses read in batches to a buffer but do we need this for read all?
func (telnet Telnet) read_all() {
	p, _ := ioutil.ReadAll(telnet.conn)
	fmt.Println(string(p))
}

func main() {
	tel := telnet("india.colorado.edu", 13)
	tel.read_all()
}
