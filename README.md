# YouTokenToGo

### This repo contains bindings for some of the functions provided by https://github.com/VKCOM/YouTokenToMe, to use with golang.

## Getting started

First build the shard lib with
```
make build # or just "make"
```

After that, you can use the library through golang by importing "C" and adding this on top of the import
```
#cgo LDFLAGS: -L${SRCDIR} -lbpewrapper
#include <stdlib.h>
#include <path/to/wrapper.h>
```
