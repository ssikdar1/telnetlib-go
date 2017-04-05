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
