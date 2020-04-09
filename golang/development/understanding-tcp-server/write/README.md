### import "net"


### Listen()

Listen announces on the local network address.

The network must be **"tcp"**, **"tcp4"**, **"tcp6"**, **"unix"** or **"unixpacket"**.

For TCP networks, if the host in the address parameter is empty or a literal unspecified IP address, Listen listens on all available unicast and anycast IP addresses of the local system. To only use IPv4, use network "tcp4". The address can use a host name, but this is not recommended, because it will create a listener for at most one of the host's IP addresses. If the port in the address parameter is empty or "0", as in "127.0.0.1:" or "[::1]:0", a port number is automatically chosen. The Addr method of Listener can be used to discover the chosen port.

### Accept()

    func (l *TCPListener) Accept() (Conn, error)

Accept implements the **Accept method** in the Listener interface; it waits for the next call and returns a generic **Conn**.

### WriteString()

    func WriteString(w Writer, s string) (n int, err error)

**WriteString** writes the contents of the string **s** to **w**, which accepts a slice of bytes. If **w** implements **StringWriter**, its WriteString method is invoked directly. Otherwise, w.Write is called exactly once.


### How to Run :

In one terminal run this program

    $ go run main.go

Another terminal run **telnet**

    $ telnet 127.0.0.1 8090
    
    Trying 127.0.0.1...
    Connected to localhost.
    Escape character is '^]'.

    What are you doing ?
    Connection closed by foreign host.