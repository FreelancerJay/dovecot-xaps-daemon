//
// The MIT License (MIT)
//
// Copyright (c) 2015 Stefan Arentz <stefan@arentz.ca>
// Copyright (c) 2017 Frederik Schwan <frederik dot schwan at linux dot com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
//

package main

import (
	"flag"
	"github.com/freswa/dovecot-xaps-daemon/aps"
	"github.com/freswa/dovecot-xaps-daemon/database"
	"github.com/freswa/dovecot-xaps-daemon/logger"
	"github.com/freswa/dovecot-xaps-daemon/socket"
	log "github.com/sirupsen/logrus"
)

const Version = "1.1"

var logLevel = flag.String("loglevel", "warn", "Loglevel: debug, error, fatal, info, panic")
var socketpath = flag.String("socket", "/var/run/xapsd/xapsd.sock", "path to the socketpath for Dovecot")
var checkDelayedInterval = flag.Int("delayCheckInterval", 20, "interval to check for delayed push notifications to send")
var delayMessageTime = flag.Int("delayTime", 30, "seconds to wait until a notification for a non NewMessage event gets sent")
var databasefile = flag.String("database", "/var/lib/xapsd/databasefile.json", "path to the databasefile file")
var key = flag.String("key", "/etc/xapsd/key.pem", "path to the pem file containing the private key")
var certificate = flag.String("certificate", "/etc/xapsd/certificate.pem", "path to the pem file containing the certificate")



func main() {
	flag.Parse()
	logger.ParseLoglevel(*logLevel)

	log.Debugln("Opening databasefile at", *databasefile)
	db, err := database.NewDatabase(*databasefile)
	if err != nil {
		log.Fatal("Cannot open databasefile: ", *databasefile)
	}
	topic := aps.NewApns(*certificate, *key, *checkDelayedInterval, *delayMessageTime, db)

	log.Printf("Starting xapsd %s on %s", Version, *socketpath)
	socket.NewSocket(*socketpath, db, topic)
}
