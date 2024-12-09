package verifier

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/textproto"
	"crypto/tls"
	"strings"
)

const smtpServer = "smtp.gmail.com:465"

const sender_email = "amanrajmathematics@gmail.com"


// Verifies if the email truly exists
func VerifyDomain(target_email string) (bool, error) {
	
	// Dial the connection
	conn, err := net.Dial("tcp", smtpServer)

	if err != nil {
		return false, errors.New("ERR_CONNECTION_REFUSED")
	}

	// Defer till the end 
	defer conn.Close()

	conn = tls.Client(conn, &tls.Config{
		InsecureSkipVerify : true, 
	})

	// Buffer reader initialization
	bufferReader := bufio.NewReader(conn) 
	tpReader := textproto.NewReader(bufferReader)

	// Welcome message
	// Status : 220 
	if _, err := tpReader.ReadLine(); err != nil {
		return false, errors.New("ERR_READLINE_WELCOME")
	}

	// Send EHLO to the server to initiate communication
	if _, err := fmt.Fprintf(conn, "EHLO smtp.gmail.com\r\n"); err != nil {
		return false, errors.New("ERR_SENDING_EHLO")
	}

	_, err = tpReader.ReadLine()
	if err != nil {
		return false, errors.New("ERR_READING_EHLO_RESPONSE")
	}

	// MAIL 
	if _, err := fmt.Fprintf(conn, "MAIL FROM:<%s>\r\n", sender_email); err != nil {
		return false, errors.New("ERR_SENDING_MAIL")
	}

	// Server response
	if _, err := tpReader.ReadLine(); err != nil {
		return false, errors.New("ERR_READING_MAIL_RESPONSE")
	}

	// RCPT
	if _, err := fmt.Fprintf(conn, "RCPT TO:<%s>\r\n", target_email); err != nil {
		return false, errors.New("ERR_SENDING_RCPT_REQUEST")
	}

	// Server response
	response, err := tpReader.ReadLine()

	if err != nil {
		return false, errors.New("ERR_READING_RCPT_RESPONSE")
	} 

	if strings.HasPrefix(response, "250") {
		return true, nil 
	} else {
		return false, nil 
	}

}
