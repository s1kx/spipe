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
