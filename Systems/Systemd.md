# Systemd


## Use of '-' in assignment
The argument passed should be an absolute filename or wildcard expression, optionally prefixed with "-", which indicates that if the file does not exist, it will not be read and no error or warning message is logged.

If any of those commands (not prefixed with -) fail, the rest are not executed and the unit is considered failed.
