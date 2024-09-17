package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	notifactionpb "notifaction-service/protos/notifaction"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type NotificationServiceServer struct {
	notifactionpb.UnimplementedNotificationServiceServer
	Consumer *kafka.Consumer
}

func NewNotificationServer(cons *kafka.Consumer) *NotificationServiceServer {
	ns := &NotificationServiceServer{
		Consumer: cons,
	}
	return ns
}

func (s *NotificationServiceServer) ListenForKafkaMessages() {
	for {
		msg, err := s.Consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Olingan Kafka xabari: %s\n", string(msg.Value))

			var emailMsg notifactionpb.NotificationReq
			err := json.Unmarshal(msg.Value, &emailMsg)
			if err != nil {
				log.Printf("Xabarni parsing qilishda xatolik: %v", err)
				continue
			}

			err = SendEmail(emailMsg.UserEmail, emailMsg.Subject, emailMsg.Message)
			if err != nil {
				log.Printf("Email yuborishda xatolik: %v", err)
			}
		} else {
			log.Printf("Xabar o'qishda xatolik: %v\n", err)
		}
	}
}

func SendEmail(to, subject, body string) error {
	from := "dilshoddilmurodov112@gmail.com"
	pass := "xmxu rdhp pmdf pezk"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
