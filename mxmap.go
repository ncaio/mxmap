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
	//
	//
	banner()
	//
	//
	//
	wp := flag.String("domain", "localhost", "a string")
	flag.Parse()
	//
	//
	//
	domain := *wp
	//
	//
	//
	fmt.Println("Searching for MX records...")
	mx, treta := net.LookupMX(domain)
	if treta != nil {
		fmt.Println("OMG - MX address not found.")
	}
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
			fmt.Print("Cheking for open relay: ")
			c.Mail("postmaster@evildomain.com")
			or := c.Rcpt("postmaster@evildomain.com")
			if or != nil {
				color.Green(" [- Access denied -]")
			} else {
				color.Red(" [- Access allowed -]")
			}
			//
			//
			//
			fmt.Print("Testing VRFY postmaster@" + domain + ": ")
			v := c.Verify("postmaster@" + domain)
			if v != nil {
				color.Green(" [- VRFY disallowed -]")
			} else {
				color.Red(" [- VRFY allowed -]")
			}
			fmt.Print("Testing VRFY postmaster: ")
			vd := c.Verify("postmaster")
			if vd != nil {
				color.Green(" [- VRFY disallowed -]")
			} else {
				color.Red(" [- VRFY allowed -]")
			}
			//
			//
			//
			fmt.Print("Testing RCPT ENUM postmaster@" + domain + ": ")
			c.Mail("postmaster@" + domain)
			r := c.Rcpt("postmaster@" + domain)
			if r != nil {
				color.Green(" [- RCPT enum disallowed -]")
			} else {
				color.Red(" [- RCPT enum allowed -]")
			}
			//
			//
			//
			fmt.Print("Testing RCPT ENUM postmaster: ")
			c.Mail("postmaster@" + domain)
			rd := c.Rcpt("postmaster")
			if rd != nil {
				color.Green(" [- RCPT enum disallowed -]")
			} else {
				color.Red(" [- RCPT enum allowed -]")
			}

		}
		defer c.Close()
		c.Quit()
		fmt.Println("")
	}
}
