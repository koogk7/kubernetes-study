package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

// 두번째 인자는 port 번호
var portNumber = "8080"
var addr = flag.String("addr", ":"+portNumber, "http service address")
var temple = template.Must(template.New("qr").Parse(templateStr)) // HTML 템플릿 제작

func main() {
	// HTTP 포트 세팅
	flag.Parse()

	// URL 바인딩
	http.Handle("/", http.HandlerFunc(QR))
	http.Handle("/products/stop", http.HandlerFunc(DownServer))
	log.Println("Products Server is running on portnumber : " + portNumber)
	//
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	// HTML 템플릿 렌더링
	_ = temple.Execute(w, req.FormValue("s"))
}

func DownServer(w http.ResponseWriter, req *http.Request) {
	log.Fatal("Down Request")
}

const templateStr = `
<html>
<head>
<title>Products QR Link Generator</title>
</head>
<body>
<h1> Products QR Link Generator</h1>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`
