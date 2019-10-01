package main
import(
	"gopkg.in/gomail.v2"
	"fmt"
)
func main(){
	mail:=gomail.NewMessage()
	mail.SetHeader("From", "/*sender's address*/")
	mail.SetHeader("To", "/*receiver address*/")
	mail.SetHeader("Subject", "/*type your subject */")
	mail.SetBody("text/plain", "/*Type your message*/")
	dialer:=gomail.NewPlainDialer("smtp.gmail.com", 587, "/*sender's email address*/", "/*type your password*/")
	if err:=dialer.DialAndSend(mail);err!=nil{
		panic(err)
	}
	fmt.Println("Email sent")
}
