package main

//if you love it please donate 1CbE3SsUcvJWZ2YNaDwUj9AQtT8k4AGmLe
//you can sell it or do what ever you want
//I figure it out 1 year very long time

import (
    "crypto/tls"
    "log"
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/adaptor/v2"
    "golang.org/x/crypto/acme/autocert"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to the home page!")
    })

    app.Get("/login", func(c *fiber.Ctx) error {
        return c.SendString("Login Page")
    })

    app.Get("/register", func(c *fiber.Ctx) error {
        return c.SendString("Register Page")
    })

    certManager := autocert.Manager{
        Prompt: autocert.AcceptTOS,
        // ระบุโดเมนทั้งสองใน HostWhitelist don't forget add your ip address map with you website to me i use readyidc and windows vps for hosting with disable some firewall rule
        //[let empty] A 3600 [your host IP like 123.231.234.12]
        //www A 3600 [your IP]
        HostPolicy: autocert.HostWhitelist("www.yourdomainname.com", "YourSameNameDomain.com"),
        Cache:      autocert.DirCache("certs"),
    }

    server := &http.Server{
        Addr: ":443",
        Handler: adaptor.FiberApp(app),
        TLSConfig: &tls.Config{
            GetCertificate: certManager.GetCertificate,
            MinVersion:     tls.VersionTLS12,
        },
    }

    go func() {
        h := certManager.HTTPHandler(nil)
        log.Fatal(http.ListenAndServe(":80", h))
    }()

    log.Fatal(server.ListenAndServeTLS("", "")) 
}

/*
note dont delete it
server := &http.Server{
    Addr: ":443",
    Handler: logRequest(adaptor.FiberApp(app)), // ใช้ middleware สำหรับ log รายละเอียดของ request
    TLSConfig: &tls.Config{
        GetCertificate: certManager.GetCertificate,
        MinVersion:     tls.VersionTLS12,
    },
}

func logRequest(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Received request for %s from %s", r.Host, r.RemoteAddr)
        handler.ServeHTTP(w, r)
    })
}
*/

/*
if error so use this mod (go.mod)

module ChangeNameToYourRealProject

go 1.22.0

require (
	github.com/gofiber/adaptor/v2 v2.2.1
	github.com/gofiber/fiber/v2 v2.52.2
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.14.0
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.7
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)


*/
