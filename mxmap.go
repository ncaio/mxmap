//
//
//
package main

//
//
//
import (
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
	banner()
	domain := os.Args[1]
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
			//
			//
			//
			//conn, err := net.Dial("tcp", mx[p].Host+":25")
			//if err != nil {
			//	fmt.Println(err)
			//	os.Exit(1)
			//}
			//defer conn.Close()
			//var readbuf [512]byte
			//n, _ := conn.Read(readbuf[0:])
			//fmt.Print("Banner: ")
			//os.Stdout.Write(readbuf[0:n])
			//
			//
			//
			fmt.Print("Cheking for open relay: ")
			c.Mail("postmaster@evildomain.com")
			or := c.Rcpt("postmaster@evildomain.com")
			if or != nil {
				//fmt.Println(" [- Access denied -]")
				color.Green(" [- Access denied -]")
			} else {
				//fmt.Println(" [- Access allowed -]")
				color.Red(" [- Access allowed -]")
			}
			//
			//
			//
			fmt.Print("Testing VRFY postmaster@" + domain + ": ")
			v := c.Verify("postmaster@" + domain)
			if v != nil {
				//fmt.Println(" [- VRFY disallowed -]")
				color.Green(" [- VRFY disallowed -]")
			} else {
				//fmt.Println(" [- VRFY allowed -]")
				color.Red(" [- VRFY allowed -]")
			}
			fmt.Print("Testing VRFY postmaster: ")
			vd := c.Verify("postmaster")
			if vd != nil {
				//fmt.Println(" [- VRFY disallowed -]")
				color.Green(" [- VRFY disallowed -]")
			} else {
				//fmt.Println(" [- VRFY allowed -]")
				color.Red(" [- VRFY allowed -]")
			}
			//
			//
			//
			fmt.Print("Testing RCPT ENUM postmaster@" + domain + ": ")
			c.Mail("postmaster@" + domain)
			r := c.Rcpt("postmaster@" + domain)
			if r != nil {
				//fmt.Println(" [- RCPT enum disallowed -]")
				color.Green(" [- RCPT enum disallowed -]")
			} else {
				//fmt.Println(" [- RCPT enum allowed -]")
				color.Red(" [- RCPT enum allowed -]")
			}
			//
			//
			//
			fmt.Print("Testing RCPT ENUM postmaster: ")
			c.Mail("postmaster@" + domain)
			rd := c.Rcpt("postmaster")
			if rd != nil {
				//fmt.Println(" [- RCPT enum disallowed -]")
				color.Green(" [- RCPT enum disallowed -]")
			} else {
				//fmt.Println(" [- RCPT enum allowed -]")
				color.Red(" [- RCPT enum allowed -]")
			}

		}
		defer c.Close()
		c.Quit()
		fmt.Println("")
	}
}
