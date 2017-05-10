# mxmap

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
