# Type Template

## [template.Template](https://godoc.org/text/template#Template)

```Go
template.Template
```

---

# Parsing templates

## [template.ParseFiles](https://godoc.org/text/template#ParseFiles)

```Go
func ParseFiles(filenames ...string) (*Template, error)
```

## [template.ParseGlob](https://godoc.org/text/template#ParseGlob)

```Go
func ParseGlob(pattern string) (*Template, error)
```

---

## [template.Parse](https://godoc.org/text/template#Template.Parse)

```Go
func (t *Template) Parse(text string) (*Template, error)
```

## [template.ParseFiles](https://godoc.org/text/template#Template.ParseFiles)

```Go
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```

## [template.ParseGlob](https://godoc.org/text/template#Template.ParseGlob)

```Go
func (t *Template) ParseGlob(pattern string) (*Template, error)
```

---

# Executing templates

## [template.Execute](https://godoc.org/text/template#Template.Execute)

```Go
func (t *Template) Execute(wr io.Writer, data interface{}) error
```

## [template.ExecuteTemplate](https://godoc.org/text/template#Template.ExecuteTemplate)

```Go
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```

---

# Helpful template functions

## [template.Must](https://godoc.org/text/template#Must)

```Go
func Must(t *Template, err error) *Template
```

## [template.New](https://godoc.org/text/template#New)

```Go
func New(name string) *Template
```

---

# The init function

## [The init function](https://golang.org/doc/effective_go.html#init)

```Go
func init()
```

# HTTP Server

HTTP uses TCP.

To create a server that works with HTTP, we just create a TCP server.

To configure our code to handle request/response in an HTTP fashion which works with browsers, we need to adhere to HTTP standards.

# TCP server essentials

## Listen

[net.Listen](https://godoc.org/net#Listen)

```Go
func Listen(net, laddr string) (Listener, error)
```

## Listener

[net.Listener](https://godoc.org/net#Listener)

```Go
type Listener interface {
    // Accept waits for and returns the next connection to the listener.
    Accept() (Conn, error)

    // Close closes the listener.
    // Any blocked Accept operations will be unblocked and return errors.
    Close() error

    // Addr returns the listener's network address.
    Addr() Addr
}
```

## Connection

[net.Conn](https://godoc.org/net#Conn)

```Go
type Conn interface {
    // Read reads data from the connection.
    Read(b []byte) (n int, err error)

    // Write writes data to the connection.
    Write(b []byte) (n int, err error)

    // Close closes the connection.
    // Any blocked Read or Write operations will be unblocked and return errors.
    Close() error

    // LocalAddr returns the local network address.
    LocalAddr() Addr

    // RemoteAddr returns the remote network address.
    RemoteAddr() Addr

    SetDeadline(t time.Time) error

    SetReadDeadline(t time.Time) error

    SetWriteDeadline(t time.Time) error
}
```

## Dial

[net.Dial](https://godoc.org/net#Dial)

```Go
func Dial(network, address string) (Conn, error)
```

---

# Write

[io.WriteString](https://godoc.org/io#WriteString)

```Go
func WriteString(w Writer, s string) (n int, err error)
```

[fmt.Fprintln](https://godoc.org/fmt#Fprintln)

```Go
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```

---

# Read

- [ioutil.ReadAll](https://godoc.org/io/ioutil#ReadAll)

```Go
func ReadAll(r io.Reader) ([]byte, error)
```

- [bufio.NewScanner](https://godoc.org/bufio#NewScanner)

```Go
func NewScanner(r io.Reader) *Scanner
```

- [bufio.Scan](https://godoc.org/bufio#Scanner.Scan)

```Go
func (s *Scanner) Scan() bool
```

- [bufio.Text](https://godoc.org/bufio#Scanner.Text)

```Go
func (s *Scanner) Text() string
```

---

# Read & Write

- [io.Copy](https://godoc.org/io#Copy)

```GO
func Copy(dst Writer, src Reader) (written int64, err error)
```

# Postgres

# create table

```
CREATE TABLE Account  (
   id INT PRIMARY KEY     NOT NULL,
   name            CHAR(50)   NOT NULL,
   password          CHAR(50)     NOT NULL,
   first       CHAR(50) NOT NULL,
   last         CHAR(50) NOT NULL,
   role			   CHAR(50) NOT NULL
);
```

## show tables in a database (list down)

```
\d
```

## show details of a table

```
\d <table name>
```

## drop a table

```
DROP TABLE <table name>;
```

## schema

Schemas allow us to organize our database and database code.

A schema is like a folder.

Into this folder, you can put tables, views, indexes, sequences, data types, operators, and functions.

Unlike folders, however, schemas can't be nested.

Schemas provide namespacing.

[Read more about schemas](https://www.tutorialspoint.com/postgresql/postgresql_schema.htm)

# users & privileges

## see current user

```
SELECT current_user;
```

## details of users

```
\du
```

## create user

```
CREATE USER james WITH PASSWORD 'password';
```

## grant privileges

privileges: SELECT, INSERT, UPDATE, DELETE, RULE, ALL

```
GRANT ALL PRIVILEGES ON DATABASE company to james;
```

## revoke privileges

```
REVOKE ALL PRIVILEGES ON DATABASE company from james;
```

## alter

```
ALTER USER james WITH SUPERUSER;
```

```
ALTER USER james WITH NOSUPERUSER;
```

## remove

```
DROP USER james;
```
