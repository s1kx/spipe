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

SHA256: `43bb616498de58cd499644d2439f2872c462a245dae27b8bfb2eb6134547f3f9`


**Windows 64-bit**

Download: http://storage.googleapis.com/spipe/bin/spipe-windows-amd64.exe

SHA256: `3ae70aed3c8d550bf2a75dd0289a0c363bd689ff29cbddafbb019c692125fe26`
