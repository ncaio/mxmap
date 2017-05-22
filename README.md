# mxmap

Mxmap is a simple smtp scan for tasks like:
Open relay test;
VRFY ENUM test;
RCPT ENUM test;
Internal spoof attack.

All ENUM tests has 2 process. First one a test with domain and second one without. Basicly you don't need a dicionary file for it, mxmap use postmaster account for default. From the domain address, mxmap is able to find all mx records. If RCPT ENUM is possible then mxmap will try to do an internal spoof attack, sending a email with BODY "[- MXMAP SPOOFING TEST -]" from postmaster(+domain) to postmaster(+domain). But if you want, use --user=user flag to change it.

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
/mxmap --domain=domain.tld --helo=domain.tld --odomain=example.com --spoof=on
~~~~

**OUTPUT**

Ex1:

~~~~
----------------------------------------------------------------------
[- MXMAP by ncaio -]
>> caiogore _|_ gmail _|_ com
----------------------------------------------------------------------

Searching for MX records...
Mx found: 1

----------------------------------------------------------------------

Testing: nf-e.top. [- UP -]

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
