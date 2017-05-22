//
//
//
package main

//
//
//
import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"net"
	"net/smtp"
	"os"
	"strings"
)

//
//
//
func banner() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("[- MXMAP by ncaio -]")
	fmt.Println(">> caiogore _|_ gmail _|_ com")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("")
}

//
//
//
func openr(h string, u string, he string) {
	fmt.Print("\nCheking for open relay: ")
	c, treta := smtp.Dial(h + ":25")
	if treta != nil {
		fmt.Println(" OMG - Connection refused.")
	}
	c.Hello(he)
	c.Mail(u)
	or := c.Rcpt(u)
	if or != nil {
		color.Green(" [- Access denied -]")
		c.Close()
	} else {
		color.Red(" [- Access allowed -]")
		c.Close()
	}
}

//
//
//
func vrfy(h string, u string, he string) {
	fmt.Print("Testing VRFY " + u + " : ")
	c, treta := smtp.Dial(h + ":25")
	if treta != nil {
		fmt.Println(" OMG - Connection refused.")
	}
	c.Hello(he)
	v := c.Verify(u)
	if v != nil {
		color.Green(" [- VRFY disallowed -]")
		c.Close()
	} else {
		color.Red(" [- VRFY allowed -]")
		c.Close()
	}

}

//
//
//
func rcpt(h string, u string, he string, s string) {
	fmt.Print("Testing RCPT ENUM " + u + ": ")
	c, treta := smtp.Dial(h + ":25")
	if treta != nil {
		fmt.Println(" OMG - Connection refused.")
	}
	c.Hello(h)
	c.Mail(u)
	r := c.Rcpt(u)
	if r != nil {
		color.Green(" [- RCPT enum disallowed -]")
	} else {
		color.Red(" [- RCPT enum allowed -]")
		color.Red("... Maybe a Spoofing attack is possible")
		if strings.Contains(s, "on") {
			fmt.Print("Spoofing: sending mail from " + u + " to " + u)
			spd, err := c.Data()
			if err != nil {
				color.Red(" [- Data error:  (SPD) c.Data() -]")
				c.Close()
			}
			_, err = fmt.Fprintf(spd, "[- MXMAP SPOOFING TEST -]")
			if err != nil {
				color.Red(" [- Data error: (SPD) body -]")
				c.Close()
			}
			color.Green(" [- Email Sended -]")
			c.Close()
		} else {
			color.Red("... When a enumeration is detected you can try a spoofing attack sending a e-mail from " + u + " to " + u + ". to do this, do you need --spoof=on flag. by default is off")
			c.Close()
		}

	}

}

//
//
//
func main() {
	//
	//	HELLO
	//
	banner()
	//
	// Flag domain. --domain=domain.tld or -domain=domain.tld
	//
	wp := flag.String("domain", "localhost", "domain address. Ex: domain.tld")
	un := flag.String("user", "postmaster", "user address. Ex: postmaster")
	sf := flag.String("spoof", "off", "spoofing attack flag")
	he := flag.String("helo", "me", "helo string")
	ud := flag.String("odomain", "evildomain.com", "a domain diferent of default")
	flag.Parse()
	//
	//	DOMAIN <- ARG
	//	USER <-ARG
	//
	domain := *wp
	user := *un
	spoof := *sf
	helo := *he
	opend := *ud
	//
	//	FINDING MX RECORDS
	//
	fmt.Println("Searching for MX records...")
	mx, treta := net.LookupMX(domain)
	if treta != nil {
		color.Red("OMG - MX address not found.")
		os.Exit(1)
	}
	//
	//	ACTION
	//
	fmt.Println("Mx found:", len(mx))
	fmt.Println("")
	fmt.Println("----------------------------------------------------------------------")
	for p := range mx {
		fmt.Print("\nTesting: ", mx[p].Host)
		c, treta := smtp.Dial(mx[p].Host + ":25")
		if treta != nil {
			fmt.Println(" OMG - Connection refused.")
			c.Quit()
		} else {
			color.Green(" [- UP -]")
		}
		//
		//	OPEN RELAY TEST
		//
		openr(mx[p].Host, user+"@"+opend, helo)
		//
		//	VRFY WITHOUT DOMAIN
		//
		vrfy(mx[p].Host, user, helo)
		//
		//	VRFY WITH DOMAIN
		//
		vrfy(mx[p].Host, user+"@"+domain, helo)
		//
		//	RCPT ENUM WITHOUT DOMAIN
		//
		rcpt(mx[p].Host, user, helo, spoof)
		//
		//	RCPT ENUM WITH DOMAIN
		//
		rcpt(mx[p].Host, user+"@"+domain, helo, spoof)
		fmt.Println("")
		fmt.Println("----------------------------------------------------------------------")
	}
}
