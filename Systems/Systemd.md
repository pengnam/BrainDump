# Systemd


## Use of '-' in assignment
The argument passed should be an absolute filename or wildcard expression, optionally prefixed with "-", which indicates that if the file does not exist, it will not be read and no error or warning message is logged.

If any of those commands (not prefixed with -) fail, the rest are not executed and the unit is considered failed.

## Specifying process properties

There are soft and hard limits for executed processes. setrlimit can be used for resource limit. When set to *infinity*, it implies that there are no limits.


## Configuring resource limits

There are typically two ways to configure resource limits:


1. ulimit:

Used to change resource limits on a temporary basis. 

2. /etc/security/limits.conf

Persistent limits can be set for a particular user by editing the `/etc/security/limits.conf`.



