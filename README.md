# myzinx
Golang轻量级并发服务器框架zinx


##func ResolveTCPAddr
```go
func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
```
ResolveTCPAddr将addr作为TCP地址解析并返回。参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；net必须是"tcp"、"tcp4"或"tcp6"。

IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80"。

##func ListenTCP
```go
func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)
```
ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener，net参数必须是"tcp"、"tcp4"、"tcp6"，如果laddr的端口字段为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。

##func (*TCPListener) AcceptTCP
```go
func (l *TCPListener) AcceptTCP() (*TCPConn, error)
```
AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn。