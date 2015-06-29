package main

import (
    "fmt"
    "github.com/vonwenm/libsmtp"
    "net/smtp"
    "io"
    "os"
)

func main() {
    auth := smtp.PlainAuth(
        "",
        "name@domain.com",
        "password",
        "mail.domain.com:25",
    )
    f, _ := os.Open("/path/to/file/")
    fmt.Println(
        libsmtp.SendMailWithAttachments(
            "smtp.domain.com:25",
            &auth,
            "mvw@name.com",
            "Y u no talk?",
            []string{"name@domain.com"},
            []byte("Message!"),
            map[string]io.Reader{
                "filename": f,
            },
        ),
    )

}