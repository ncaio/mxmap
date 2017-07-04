# mxmap

Mxmap is a simple smtp scan for tasks like:
SPF test;
Open relay test;
VRFY ENUM test;
RCPT ENUM test;
Internal spoof attack.

All ENUM tests has 2 process. First, test with domain and second without. Basicly you don't need a dicionary file for it, mxmap use postmaster account by default. From the domain address, mxmap is able to find all mx records. If RCPT ENUM is possible then mxmap will try to do an internal spoof attack, sending a email with BODY "[- MXMAP SPOOFING TEST -]" from postmaster(+domain) to postmaster(+domain). But if you want, use --user=user flag to change it.

**PS: x86_64 binary**

**SPF Google Softail**

Normaly, domain with TXT spf include:_spf.google.com has a possibility to send emails to gmail.com accounts. That is possible because Google SPF record uses Softail. Mr. Vitaly Salnikov wrote about it in http://hackanddefense.com/blog/how-to-send-fake-emails/index.html


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

Ex1: ./mxmap --domain=nf-e.top --helo=nf-e.top --banner=on

~~~~
----------------------------------------------------------------------
[- MXMAP by ncaio -]
>> caiogore _|_ gmail _|_ com
----------------------------------------------------------------------

Searching for MX records...
Mx found: 1
Dns txt records: [v=spf1 a mx ip4:212.237.0.0/16 ip4:94.177.190.0/24 ip4:93.186.0.0/16 ip4:188.213.0.0/16 ~all]

----------------------------------------------------------------------

Testing: nf-e.top. -> 93.186.253.37 [- UP -]

Banner:
220 nf-e.top ESMTP

Cheking for open relay:  [- Access denied -]
Testing VRFY postmaster :  [- VRFY disallowed -]
Testing VRFY postmaster@nf-e.top :  [- VRFY disallowed -]
Testing RCPT ENUM postmaster:  [- RCPT enum allowed -]
... Maybe a Spoofing attack is possible
... When a enumeration is detected you can try a spoofing attack sending a e-mail from postmaster to postmaster. to do this, do you need --spoof=on flag. by default is off
Testing RCPT ENUM postmaster@nf-e.top:  [- RCPT enum allowed -]
... Maybe a Spoofing attack is possible
... When a enumeration is detected you can try a spoofing attack sending a e-mail from postmaster@nf-e.top to postmaster@nf-e.top. to do this, do you need --spoof=on flag. by default is off

----------------------------------------------------------------------
~~~~
