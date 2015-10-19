# Simple GUI for Carina by Rackspace

This project uses [libcarina](https://godoc.org/github.com/rackerlabs/libcarina) and [andlabs/ui](https://github.com/andlabs/ui) to create a simple GUI for managing clusters on Carina by Rackspace.  The GUI looks like this:

![carina gui](ui.png)

You can download the binary from the release page.

## Building

See [andlabs/ui](https://github.com/andlabs/ui) for the requirements for your platform to build the compiled binary.

```
GOOS=linux go build -a -installsuffix cgo -o carina-gui .
```

Then, once you build the binary, you have to do `chmod +x rcs-manager` for it to be executable
