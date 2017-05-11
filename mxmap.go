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
)

//
//
//
func banner() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("------------------------------------------------------")
	fmt.Println("[- MXMAP by ncaio -]")
	fmt.Println(">> caiogore _|_ gmail _|_ com")
	fmt.Println("------------------------------------------------------")
	fmt.Println("")
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
	wp := flag.String("domain", "localhost", "a string")
	un := flag.String("user", "postmaster", "a string")
	flag.Parse()
	//
	//	DOMAIN <- ARG
	//	USER <-ARG
	//
	domain := *wp
	user := *un
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
	for p := range mx {
		fmt.Print("Testing: ", mx[p].Host)
		c, treta := smtp.Dial(mx[p].Host + ":25")
		if treta != nil {
			fmt.Println(" OMG - Connection refused.")
			c.Quit()
		} else {
			color.Green(" [- UP -]")
			c.Hello("me")
			//
			//	OPEN RELAY TEST
			//
			fmt.Print("Cheking for open relay: ")
			c.Mail(user + "@evildomain.com")
			or := c.Rcpt(user + "@evildomain.com")
			if or != nil {
				color.Green(" [- Access denied -]")
			} else {
				color.Red(" [- Access allowed -]")
			}
			//
			//	VRFY WITH DOMAIN TEST
			//
			fmt.Print("Testing VRFY " + user + "@" + domain + ": ")
			v := c.Verify(user + "@" + domain)
			if v != nil {
				color.Green(" [- VRFY disallowed -]")
			} else {
				color.Red(" [- VRFY allowed -]")
			}
			//
			//	VRFY WITHOUT DOMAIN TEST
			//
			fmt.Print("Testing VRFY postmaster: ")
			vd := c.Verify(user)
			if vd != nil {
				color.Green(" [- VRFY disallowed -]")
			} else {
				color.Red(" [- VRFY allowed -]")
			}
			//
			//	RCPT ENUM WITH DOMAIN
			//
			fmt.Print("Testing RCPT ENUM " + user + "@" + domain + ": ")
			c.Mail(user + "@" + domain)
			r := c.Rcpt(user + "@" + domain)
			if r != nil {
				color.Green(" [- RCPT enum disallowed -]")
			} else {
				color.Red(" [- RCPT enum allowed -]")
				color.Blue("... Spoofing is possible")
				fmt.Print("Spoofing: sending mail from " + user + "@" + domain + " to " + user + "@" + domain)
				spd, err := c.Data()
				if err != nil {
					color.Red(" [- Data error:  (SPD) c.Data() -]")
				}
				_, err = fmt.Fprintf(spd, "[- MXMAP SPOOFING TEST -]")
				if err != nil {
					color.Red(" [- Data error: (SPD) body -]")
				}
				color.Green(" [- Email Sended -]")

			}
			//
			//	RCPT ENUM WITHOUT DOMAIN
			//
			fmt.Print("Testing RCPT ENUM postmaster: ")
			c.Mail(user + "@" + domain)
			rd := c.Rcpt(user)
			if rd != nil {
				color.Green(" [- RCPT enum disallowed -]")
			} else {
				color.Red(" [- RCPT enum allowed -]")
				color.Blue("... Spoofing is possible")
				fmt.Print("Spoofing: sending mail from " + user + "@" + domain + " to " + user + "@" + domain)
				sp, err := c.Data()
				if err != nil {
					color.Red(" [- Data error: (SP) c.Data() -]")
					os.Exit(1)
				}
				_, err = fmt.Fprintf(sp, "[- MXMAP SPOOFING TEST -]")
				if err != nil {
					color.Red(" [- Data error: (SP) body -]")
					os.Exit(1)
				}

				color.Green(" [- Email Sended -]")
			}

		}
		defer c.Close()
		c.Quit()
		fmt.Println("")
	}
}
