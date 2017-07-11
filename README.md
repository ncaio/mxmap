# mxmap

Mxmap is a simple smtp scan for tasks like:
SPF test;
Dmarc test;
Google DKIM Selector test;
Open relay test;
VRFY ENUM test;
RCPT ENUM test;
Internal spoof attack.

All ENUM tests has 2 process. First, test with domain and second without. Basicly you don't need a dicionary file for it, mxmap use postmaster account by default. From the domain address, mxmap is able to find all mx records. If RCPT ENUM is possible then mxmap will try to do an internal spoof attack, sending a email with BODY "[- MXMAP SPOOFING TEST -]" from postmaster(+domain) to postmaster(+domain). But if you want, use --user=user flag to change it.

**PS: x86_64 binary**

**Basic usage** 

Basic usage 1 - [ Simple scan ]
helo flag is recommended in all cases

~~~~
./mxmap --domain=domain.tld --helo=domain.tld
~~~~

Basic usage 2 - [ Changing user name ]

~~~~
./mxmap --domain=domain.tld --helo=domain.tld --user=abuse
~~~~

Basic usage 3 - [ Changing a different open relay domain (and user) ]

~~~~
./mxmap --domain=domain.tld --helo=domain.tld --odomain=example.com
or
./mxmap -domain=domain.tld --helo=domain.tld -odomain=example.com --user=root
~~~~

Basic usage 3 - [ Saying yes to a spoofing attack ]

~~~~
./mxmap --domain=domain.tld --helo=domain.tld --odomain=example.com --spoof=on
~~~~

Basic usage 4 - [ Smtp banner ]

~~~~
./mxmap --domain=domain.tld --helo=domain.tld --odomain=example.com --banner=on
~~~~

**OUTPUT**

Ex1: ./mxmap --domain=zeplan.br.com --helo=mail.zeplan.br.com --banner=on

~~~~
----------------------------------------------------------------------
[- MXMAP by ncaio -]
>> caiogore _|_ gmail _|_ com
----------------------------------------------------------------------

Searching for MX records...
Mx found: 1
Dns txt records: [v=spf1 ip4:62.210.164.224/28 ip4:195.154.149.64/27 ip4:62.210.118.192/27 ip4:154.44.178.0/23 include:cmailsys.com ~all]

----------------------------------------------------------------------

SPF test: [- SPF Flag Found -]
* [- Sender-ID Result: SOFTFAIL -]

----------------------------------------------------------------------

DMARC test: [- Dmarc TXT found -]
Dns txt records: [v=DMARC1; p=none; rua=mailto:dmarc@zeplan.br.com]
* [- DMARC 'p' flag is none -]

----------------------------------------------------------------------

Google DKIM Selector test: [- DKIM TXT not found -]

----------------------------------------------------------------------

Testing: caloga-pub.caloga.com. -> 195.154.149.90 [- UP -]

Banner:
220 caloga-pub.caloga.com ESMTP Exim 4.84_2 Sun, 09 Jul 2017 20:17:38 +0200

Exim Vulnerability Statistics - https://www.cvedetails.com/product/19563/Exim-Exim.html?vendor_id=10919

Cheking for open relay:  [- Access denied -]
Testing VRFY postmaster :  [- VRFY disallowed -]
Testing VRFY postmaster@zeplan.br.com :  [- VRFY disallowed -]
Testing RCPT ENUM postmaster:  [- RCPT enum disallowed -]
Testing RCPT ENUM postmaster@zeplan.br.com:  [- RCPT enum allowed -]

----------------------------------------------------------------------
~~~~


References and regards:

*  Anti-Spam Recommendations for SMTP MTAs -> https://tools.ietf.org/html/rfc2505
*  Go -> https://golang.org
*  Color -> github.com/fatih/color
*  Resolution of the Sender Policy Framework (SPF) and Sender ID Experiments -> https://tools.ietf.org/html/rfc6686
*  DomainKeys Identified Mail (DKIM) Signatures -> https://tools.ietf.org/html/rfc6376
*  Vitaly Salnikov - How to send fake emails -> http://hackanddefense.com/blog/how-to-send-fake-emails/index.html
