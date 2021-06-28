# Goa Demo Service

HTTP service created with the [Goa v3](https://goa.design/) framework.
It can be used as an example and a starting point
for writing microservices with [Goa](https://goa.design/).

## goagen.sh

Goa code generation doesn't work when you use a vendor directory.
This means that if you explicitly build your project with vendor
support enabled, `goa gen` will fail.

The `goagen.sh` script is a work-around that limitation. 
It alters GOFLAGS so that the goa code generation doesn't look for
dependencies in the `vendor` directory.

> If you don't use the vendor mode, then you may 
> not need to use this script.