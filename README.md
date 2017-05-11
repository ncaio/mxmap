# mxmap

Mxmap is a simple smtp scan for tasks like:
Open relay test;
VRFY ENUM test;
RCPT ENUM test;
Internal spoof attack.

All ENUM tests has 2 process. First one a test with domain and second one without. Basicly you don't need a dicionary file for it. mxmap use postmaster account for that =]. From the domain address, mxmap is able to find all mx records. If RCPT ENUM is possible then mxmap will try to do a Internal spoof attack, sending a email with BODY "[- MXMAP SPOOFING TEST -]" from postmaster(+domain) to postmaster(+domain).

**PS: x86_64 binary**

~~~~
./mxmap --domain=domain.tld
or
./mxmap -domain=domain.tld
~~~~

ex:
~~~~
[hostname]# ./mxmap.go --domain=mikrocen***.com.br

------------------------------------------------------
[- MXMAP by ncaio -]
>> caiogore _|_ gmail _|_ com
------------------------------------------------------

Searching for MX records...
Mx found: 3

Testing: mx.mikrocen***.com.br. [- UP -]
Cheking for open relay:  [- Access denied -]
Testing VRFY postmaster@mikrocen***.com.br:  [- VRFY disallowed -]
Testing VRFY postmaster:  [- VRFY disallowed -]
Testing RCPT ENUM postmaster@mikrocen***.com.br:  [- RCPT enum disallowed -]
Testing RCPT ENUM postmaster:  [- RCPT enum allowed -]
... Spoofing is possible
Spoofing: sending mail from postmaster@mikrocen***.com.br to postmaster@mikrocen***.com.br [- Email Sended -]

Testing: mx1.mikrocen***.com.br. [- UP -]
Cheking for open relay:  [- Access denied -]
Testing VRFY postmaster@mikrocen***.com.br:  [- VRFY disallowed -]
Testing VRFY postmaster:  [- VRFY disallowed -]
Testing RCPT ENUM postmaster@mikrocen***.com.br:  [- RCPT enum disallowed -]
Testing RCPT ENUM postmaster:  [- RCPT enum allowed -]
... Spoofing is possible
Spoofing: sending mail from postmaster@mikrocen***.com.br to postmaster@mikrocen***.com.br [- Email Sended -]

Testing: mx2.mikrocen***.com.br. [- UP -]
Cheking for open relay:  [- Access denied -]
Testing VRFY postmaster@mikrocen***.com.br:  [- VRFY disallowed -]
Testing VRFY postmaster:  [- VRFY disallowed -]
Testing RCPT ENUM postmaster@mikrocen***.com.br:  [- RCPT enum disallowed -]
Testing RCPT ENUM postmaster:  [- RCPT enum allowed -]
... Spoofing is possible
Spoofing: sending mail from postmaster@mikrocen***.com.br to postmaster@mikrocen***.com.br [- Email Sended -]
~~~~
