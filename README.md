# mxmap

Mxmap is a simple smtp scan that looking for: 
Open relay;
VRFY ENUM;
RCPT ENUM.

All ENUM tests has 2 process. First one a test with domain and second one without. Basicly you don't need a dicionary file for it. mxmap use postmaster account for that =]

**PS: x86_64 binary**

~~~~
./mxmap --domain=domain.tld
or
./mxmap -domain=domain.tld
~~~~

ex:
~~~~
./mxmap.go --domain=superonlineshop.com.br

------------------------------------------------------
[- MXMAP by ncaio -]
>> caiogore _|_ gmail _|_ com
------------------------------------------------------

Searching for MX records...
Mx found: 2

Testing: mail.superonlineshop.com.br. [- UP -]
Cheking for open relay:  [- Access denied -]
Testing VRFY postmaster@superonlineshop.com.br:  [- VRFY disallowed -]
Testing VRFY postmaster:  [- VRFY disallowed -]
Testing RCPT ENUM postmaster@superonlineshop.com.br:  [- RCPT enum allowed -]
Testing RCPT ENUM postmaster:  [- RCPT enum disallowed -]

Testing: www.superonlineshop.com.br. [- UP -]
Cheking for open relay:  [- Access denied -]
Testing VRFY postmaster@superonlineshop.com.br:  [- VRFY disallowed -]
Testing VRFY postmaster:  [- VRFY disallowed -]
Testing RCPT ENUM postmaster@superonlineshop.com.br:  [- RCPT enum allowed -]
Testing RCPT ENUM postmaster:  [- RCPT enum disallowed -]
~~~~
