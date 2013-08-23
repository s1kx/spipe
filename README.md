# spipe

Superfast TCP/UDP port redirector, powered by Go


### Usage

```shell
spipe [options] <listen port> <remote host> <remote port>

  -i="0.0.0.0": Listening interface IP address
  -net="tcp": Network type (tcp, tcp4, tcp6)
  -s=0: Outbound source port number

Example: spipe -i 127.0.0.1 80 example.com 80
```

### Install

#### Source

You need to have a properly configured Go installation (including having $GOBIN in your $PATH) to build spipe from source.

```shell
go get github.com/s1kx/spipe
go install github.com/s1kx/spipe
```

#### Pre-compiled executables

**Linux 64-bit**

Download: http://storage.googleapis.com/spipe/bin/spipe-linux-amd64

SHA256: `b739226af1c651e15d9df3bfbae93294b4a34fbdc8cf6a17ec910a6d30f36273`


**Windows 64-bit**

Download: http://storage.googleapis.com/spipe/bin/spipe-windows-amd64.exe

SHA256: `7b9febfa43369395dec5aa00f5c145bf4faa5feb765446f9a6507196e041c2d4`
