package mailer

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"midkaGolang/models"
)

// SendEmail отправляет письмо с указанным заголовком и телом на указанный адрес.
func SendEmail(to, subject, body string) error {
	from := "gormhome@gmail.com"      // мой email
	password := "jrpz yrsp erlt aoyw" // мой пароль

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// Функция для обработки успешной регистрации
func HandleSuccessfulRegistration(user models.User) {
	subject := "Регистрация прошла успешно"
	body := fmt.Sprintf("Здравствуйте, %s!<br>Вы успешно зарегистрировались на нашем сайте.", user.Email)

	err := SendEmail(user.Email, subject, body)
	if err != nil {
		fmt.Println("Ошибка при отправке письма:", err)
	}
}
