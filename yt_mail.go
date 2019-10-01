package main
import(
	"gopkg.in/gomail.v2"
	"fmt"
	"log"
	"net/http"
)
func main(){
	mail:=gomail.NewMessage()
	mail.SetHeader("From", "email")
	mail.SetHeader("To", "address")
	mail.SetHeader("Subject", "/*type your subject */")
	mail.SetBody("text/plain", "/*Type your message*/")
	dialer:=gomail.NewPlainDialer("smtp.gmail.com", 587, "/*sender's email address*/", "/*type your password*/")
	if err:=dialer.DialAndSend(mail);err!=nil{
		panic(err)
	}
	fmt.Println("Email sent")
}
