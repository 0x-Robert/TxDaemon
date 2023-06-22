# Go 언어 첫 번째 강의

## 강의 링크 : https://www.udemy.com/course/google-go-web-development/

### 【한글자막】 Google Go 프로그래밍 언어로 웹 개발하기

1. 강사 : Todd McLeod
2. 강사 경력 : 프레즈노 시티 컬리지의 컴퓨터 정보 기술학과 종신 교수이자, 캘리포니아 주립 대학 프레즈노 캠퍼스의 컴퓨터 공학과 겸임 교수로 재직 중
3. 22년 이상 학생들을 가르친 경력이 있다.

---

**강의 내용**

**[강의에서 다룰 내용]**

다음과 같은 주제를 포함해 **그 이상**을 배울 수 있습니다.

**아키텍처**

- 네트워킹 아키텍처
- 클라이언트/서버 아키텍처
- 요청/ 응답 패턴
- IETF가 정의한 RFC 표준
- 클라이언트 측 요청 및 서버 측 응답의 형식

**템플릿**

- 서버 측 프로그래밍에서 템플릿의 역할
- Go의 표준 라이브러리에서 템플릿으로 작업하는 방법
- 템플릿으로 제대로 작업하기 위한 데이터 구조의 변경

**서버**

- TCP와 HTTP 간의 관계
- HTTP 요청에 응답하는 TCP 서버를 구축하는 방법
- 메모리 내 데이터베이스의 역할을 하는 TCP 서버를 생성하는 방법
- 다양한 라우트와 메서드를 처리하는 RESTful TCP 서버를 생성하는 법
- 웹 서버, 서브먹스, 멀티플렉서, 먹스 간의 차이
- Julien Schmidt 라우터와 같은 서드 파티 라우터를 사용하는 방법
- HTTP 메서드 및 상태 코드의 중요성

**net/http 패키지**

- net/http 패키지를 이용해 웹 개발을 간소화하는 방법
- net/http 패키지 간의 차이
- 핸들러 인터페이스
- http.ListenAndServe
- 고유한 서브먹스 만들기
- 디폴트 서브먹스 사용하기
- http.Handle 및 http.Handler
- http.Handlefunc, func(ResponseWriter, \*Request), 및 http.HandlerFunc
- http.ServeContent, http.ServeFile, http.FileServer
- http.StripPrefix
- http.NotFoundHandler

**상태 및 세션**

- UUID, 쿠키, URL에서의 값, 보안의 상태를 만드는 방법
- 로그인, 권한, 로그아웃 세션을 만드는 방법
- 세션을 만료시키는 방법

**배포**

- 도메인을 구매하는 방법
- 애플리케이션을 **Google Cloud**에 배포하는 방법

**Amazon Web Services**

- Amazon Web Services(AWS)를 사용하는 방법
- AWS EC2(Elastic Compute Cloud)에 Linux 가상 머신을 생성하는 방법
- SSH(Secure Shell)를 이용해 가상 머신을 관리하는 방법
- SCP(Secure Copy)를 이용해 가상 머신으로 파일을 전송하는 방법
- 로드 밸런서의 정의 및 AWS에서 사용하는 방법

**MySQL**

- AWS에서 MySQL을 사용하는 방법
- MySQL Workbench를 AWS로 연결하는 방법

**MongoDB**

- CRUD 이해하기
- MongoDB와 Go를 사용하는 방법

**MVC(모델-뷰-컨트롤러) 설계 패턴**

- MVC 설계 패턴 이해하기
- MVC 설계 패턴 활용하기

**Docker**

- 가상 머신vs 컨테이너 비교
- Docker의 장점 이해하기
- Docker 이미지, Docker 컨테이너, Docker 레지스트리
- Docker 및 Go 구현하기
- Docker 및Go 배포하기

**Google Cloud**

- Google Cloud Storage
- Google Cloud NoSQL Datastore
- Google Cloud Memcache
- Google Cloud PAAS App Engine

**웹 개발 툴킷**

- AJAX
- JSON
- json.Marhsal 및 json.Unmarshal
- json.Encode 및 json.Decode
- HMAC(해시 메시지 인증 코드)
- Base64 인코딩
- 웹 저장소
- 컨텍스트
- TLS 및 HTTPS
- 태그를 이용한 Go언어 JSON 작업

---

# 아키텍처

## 네트워크 아키텍처

### OSI Model

강의에서 언급한 네트워크 아키텍처로는 "OSI Model"과 "TCP/IP"모델이 있었다.
그렇다면 각각 어떤 것인지 정의를 살펴보고 무슨 뜻인지, 차이는 무엇인지 알아보자.

먼저 가장 유명한 OSI모델의 정의는 다음과 같다.
OSI 모델은 The International Organization for Standardization (ISO)가 설계한 Open Systems Interconnection (OSI) Reference Model이다. 이것은 네트워크의 각 계층을 7계층으로 나눈 개념인데 각 계층은 그 계층의 프로토콜을 하나이상 사용한다. 그리고 이런 계층은 네트워크 문제를 식별하는 데 도움을 준다. 7 계층의 모델 개념은 허니웰 인포메이션 서비스의 찰스 바크만의 노고에 의해서
탄생했다. 이 모델의 목적은 프로토콜을 기능별로 나눈 것이며 각 계층은 하위 계층의 기능만을 이용하고 상위 계층에는 기능을 제공한다.

그럼 이걸 왜 만들었나? 상이한 컴퓨터 시스템끼리 서로 통신할 수 있는 "표준"을 제공하기위해 만들었다고 보면 된다.
그렇다면 7계층을 하나씩 잘 살펴보자.

| Layer No. | Layer Name   | Description                                                                                               |
| --------- | ------------ | --------------------------------------------------------------------------------------------------------- |
| 7         | Application  | Consists of standard communication services and applications that everyone can use.                       |
| 6         | Presentation | Ensures that information is delivered to the receiving machine in a form that the machine can understand. |
| 5         | Session      | Manages the connections and terminations between cooperating computers.                                   |
| 4         | Transport    | Manages the transfer of data. Also assures that the received data are identical to the transmitted data.  |
| 3         | Network      | Manages data addressing and delivery between networks.                                                    |
| 2         | Data Link    | Handles the transfer of data across the network media.                                                    |
| 1         | Physical     | Defines the characteristics of the network hardware.                                                      |

#### 7 응용프로그램(어플리케이션) 계층 (HTTP, SMTP)

    이 계층은 사용자와 직접 상호작용하는 계층입니다. 주로 우리가 사용하는 웹 브라우저와 이메일 클라이언트가 이 응용계층에 의지합니다. 그렇다고 이 프로그램들이 응용프로그램 계층의 일부가 아닙니다. 그렇다면 이 계층에 속하는 프로토콜은 무엇일까요?
    그것은 HTTP와 SMTP 프로토콜입니다. HTTP란 Hyper Text Transfer Protocol입니다. SMTP란 Send Mail Transfer Protocol입니다.
    각각 "하이퍼 텍스트 전송 프로토콜", "이메일 전송 프로토콜"이라고 생각할 수 있겠네요.

#### 6 프레젠테이션 계층 (변환, 암호화, 압축)

     이 계층은 주로 데이터의 변환, 암호화, 압축을 담당합니다. 먼저 서로 통신하는 두 개의 통신장치는 서로 다른 인코딩 방법을 사용할 수 있으므로 수신장치의 어플리케이션 계층이 이해할 수 있는 구문으로 수신데이터를 "변환"하는 일을 합니다.

    두 장치가 암호화된 연결을 통해 통신할 경우 최종 수신자에게 암호화를 디코딩하여 암호화되지 않은 데이터로 애플리케이션을 제시할 수 있도록 하는 역할을 합니다.

    프레젠테이션 계층은 애플리케이션 계층에서 수신한 데이터를 세션계층으로 전송하기 전에 압축하는 일도 담당합니다. 이로써 통신 속도와 효율을 증가시킵니다.

#### 5. 세션 계층 (체크포인트, 세션 개념)

     두 기기 사이의 통신을 시작하고 종료하는 일을 담당하는 계층이고 통신이 시작되고 종료될때까지의 시간을 "세션"이라고 합니다.
     그리고 세션계층은 데이터 전송을 체크포인트와 동기화할 수 있다. 만약 100MB를 전송할 때 52MB에서 끊어지고 체크포인트가 5MB라면
     남은 50MB만 재전송하면 전송을 마무리할 수 있다.

#### 4. 전송 계층 (세그먼트, 오류제어, 흐름제어)

     전송계층은 두 기기간의 종단 간 통신을 담당한다. 세션계층에서 데이터를 가져와서 계층3으로 보내기 전에 "세그먼트"라고 하는 조각으로 분할하는 일도 포함된다. 또한 흐름 제어 및 오류 제어의 기능도 하는데 송신자가 연결속도가 너무 빠르고 수신자는 연결속도가 느릴 때 최적의 전송속도로 전송할 수 있게 결정하고 제어하는 역할을 한다. 그리고 수신된 데이터가 완료됐는지 확인하고 완료되지 않으면 재전송을 요청하는 오류제어기능도 같이 수행한다.

#### 3. 네트워크 계층 (패킷, 라우팅)

     서로 다른 두 네트워크 간에 데이터 전송을 용이하게 하는 역할을 한다. 서로 통신하는 두 장치가 동일한 네트워크에 있을 때는 이 네트워크 계층은 필요하지 않다. 즉 다른 네트워크에 있을 때 필요하다고 보면 될 것 같다. 그리고 네트워크 계층은 전송 계층의 세그먼트를 송신자 장치에서 "패킷"이라고 불리는 더 작은 단위로 세분화해서 수신장치에서 다시 조립한다. 그리고 네트워크 계층은 데이터가 표적에 도달하기 위한 최상의 물리적 경로를 찾는데 이를 "라우팅"이라고 한다. 서버 백엔드에서 Rest API의 경로를 설정해주는 작업도 "라우팅"이라는 말을 많이 쓰는데 서버단 공부를 할 때 보면 이해할 수 있을 것 이다. 여기서 멀티플렉서라는 개념도 떠올릴 수도 있다.

#### 2. 데이터 연결 계층 (프레임, 이더넷, 프레임에 주소부여( MAC{Media Access Control} 물리적 주소 ) )

     네트워크 계층과 비슷하며 동일한 네트워크에 있는 두 개의 장치간 데이터 전송을 더 용이하게 한다. 데이터 연결 계층은 네트워크 계층에서 패킷을 가져와서 프레임이라고 불리는 더 작은 조각으로 세분화한다. 네트워크 계층과 마찬가지로 데이터 연결 계층도
     인트라 네트워크 통신에서 흐름제어와 오류제어를 담당한다.

#### 1. 물리적 계층 (케이블, 스위치, 와이파이)

     이 계층에서는 케이블, 스위치 등 데이터 전송과 관련된 물리적 장비가 포함된다. 이 계층에서는 1과 0의 문자열인 bitstream으로 변환되는 계층이다. 그리고 두 장치의 물리적계층은 신호계층에 의해서 두 장치의 1이 0과 구별되야한다.

#### OSI 모델을 통해 이메일통신에서 보는 7단계 흐름

- 쿠퍼가 파머한테 이메일을 보낸다. 송신자가 수신자에게 데이터를 보내는 작업 (계층 7>6>5>4>3>2>1) 단계로 진행된다.

1. 이메일 애플리케이션에서 메시지 작성 후 발송을 누른다.
2. 이메일 애플리케이션이 이메일 메시지를 애플리케이션 계층으로 넘기면 애플리케이션 레이어는 SMTP(Send Mail Transfer Protocol)을 선택하고 프레젠테이션 계층으로 넘긴다.
3. 프레젠테이션 계층에서는 데이터를 압축해서 세션계층에 전달한다.
4. 세션계층은 세션을 시작한다. 이후 전송 계층으로 데이터를 넘긴다.
5. 전송계층에서 데이터는 세그먼트로 나눠진다.
6. 네트워크 계층에서 패킷으로 나눠지고 다시 데이터 연결 계층으로 넘긴다.
7. 데이터 연결계층에서는 패킷을 프레임으로 나눈 뒤 물리적 계층으로 전달한다.
8. 물리적 계층에서는 데이터를 0과 1의 비트스트림으로 변환하고 물리적매체인 케이블(예시)을 통해 전송한다.

- 파머는 쿠퍼한테 받은 데이터를 위와 반대로 받게 된다.

1. 물리적 매체(와이파이, 예시)를 통해 비트스트림을 수신한다.
2. 물리적계층은 비트스트림을 1과 0에서 프레임으로 변환해서 데이터 연결 계층으로 넘긴다.
3. 데이터 연결 계층은 프레임을 패킷으로 재조립해 네트워크 계층으로 넘긴다.
4. 네트워크 계층은 패킷으로 세그먼트를 만들어 전송계층으로 넘긴다.
5. 전송 계층은 세그먼트를 재조립해서 하나의 데이터를 만든다.
6. 세션계층에서 데이터를 받은 뒤 프레젠테이션 계층으로 넘기면 통신 세션 종료된다.
7. 프레젠테이션 계층은 압축을 제거하고 원본 데이터를 애플리케이션 계층으로 넘긴다.
8. 애플리케이션 계층은 사람이 읽을 수 있는 데이터를 파머의 이메일 소프트웨어에 제공하고 파머씨는 이메일확인을 할 수 있게된다.

---

### TCP/IP 모델

| OSI Ref. Layer No. . | OSI Layer Equivalent               | TCP/IP Layer     | TCP/IP Protocol Examples                                                    |
| -------------------- | ---------------------------------- | ---------------- | --------------------------------------------------------------------------- |
| 5,6,7                | Application, session, presentation | Application      | NFS, NIS+, DNS, telnet, ftp, rlogin, rsh, rcp, RIP, RDISC, SNMP, and others |
| 4                    | Transport                          | Transport        | TCP, UDP                                                                    |
| 3                    | Network                            | Internet         | IP, ARP, ICMP                                                               |
| 2                    | Data Link                          | Data Link        | PPP, IEEE 802.2                                                             |
| 1                    | Physical                           | Physical Network | Ethernet (IEEE 802.3) Token Ring, RS-232, others                            |

# Template

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

### 참고 :

OSI Model

https://www.cloudflare.com/ko-kr/learning/ddos/glossary/open-systems-interconnection-model-osi/#:~:text=%EA%B0%9C%EB%B0%A9%ED%98%95%20%EC%8B%9C%EC%8A%A4%ED%85%9C%20%EC%83%81%ED%98%B8%20%EC%97%B0%EA%B2%B0(OSI,%EC%9E%88%EB%8A%94%20%ED%91%9C%EC%A4%80%EC%9D%84%20%EC%A0%9C%EA%B3%B5%ED%95%A9%EB%8B%88%EB%8B%A4.

위키피디아 OSI Model

https://ko.wikipedia.org/wiki/OSI_%EB%AA%A8%ED%98%95

패킷

https://www.cloudflare.com/ko-kr/learning/network-layer/what-is-a-packet/

라우터

https://www.cloudflare.com/ko-kr/learning/network-layer/what-is-routing/

전기전자공학자협회

IEEE
https://ko.wikipedia.org/wiki/%EC%A0%84%EA%B8%B0%EC%A0%84%EC%9E%90%EA%B3%B5%ED%95%99%EC%9E%90%ED%98%91%ED%9A%8C#:~:text=%EC%A0%84%EA%B8%B0%EC%A0%84%EC%9E%90%EA%B3%B5%ED%95%99%EC%9E%90%ED%98%91%ED%9A%8C(Institute,%EC%A0%84%EB%AC%B8%EA%B0%80%EB%93%A4%EC%9D%98%20%EA%B5%AD%EC%A0%9C%EC%A1%B0%EC%A7%81%EC%9D%B4%EB%8B%A4.&text=IEEE%EB%8A%94%20'I%2DTriple%2D,%EC%9D%98%20%EB%89%B4%EC%9A%95%EC%97%90%20%EC%9C%84%EC%B9%98%ED%95%98%EA%B3%A0%20%EC%9E%88%EB%8B%A4.

이더넷

1. LAN, WAN, WAN에서 가장 많이 활용되는 기술 규격
2. OSI 모델의 물리 계층에서 신호와 배선, 데이터 링크 계층에서 MAC 패킷과 프로토콜의 형식을 정의한다.
3. 이더넷은 네트워크에 연결된 각 기기들이 48비트 길이의 고유의 MAC 주소를 가지고 이 주소를 이용해 상호간에 데이터를 주고 받을 수 있도록 만들어졌다. 전송 매체로는 BNC 케이블 또는 UTP, STP 케이블을 사용하며, 각 기기를 상호 연결시키는 데에는 허브, 네트워크 스위치, 리피터 등의 장치를 이용한다. 4.대부분 IEEE 802.3규약으로 표준화 됨

https://ko.wikipedia.org/wiki/%EC%9D%B4%EB%8D%94%EB%84%B7
