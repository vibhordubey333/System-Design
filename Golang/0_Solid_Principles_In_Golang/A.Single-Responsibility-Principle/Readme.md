#### S- Single Responsibility


> When we do not respect Single Responsibility. The Single Responsibility Principle (SRP) states that each software module should have 
one and only one reason to change.
                - Uncle Bob aka Robert C. Martin

```
type EmailService struct {
	db           *gorm.DB
	smtpHost     string
	smtpPassword string
	smtpPort     int
}

func NewEmailService(db *gorm.DB, smtpHost string, smtpPassword string, smtpPort int) *EmailService {
	return &EmailService{
		db:           db,
		smtpHost:     smtpHost,
		smtpPassword: smtpPassword,
		smtpPort:     smtpPort,
	}
}

func (s *EmailService) Send(from string, to string, subject string, message string) error {
	email := EmailGorm{
		From:    from,
		To:      to,
		Subject: subject,
		Message: message,
	}

	err := s.db.Create(&email).Error
	if err != nil {
		log.Println(err)
		return err
	}
	
	auth := smtp.PlainAuth("", from, s.smtpPassword, s.smtpHost)
	
	server := fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort)
	
	err = smtp.SendMail(server, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
```

Let us examine the code block from above. There we have one struct, EmailService, with only one method, Send. We use this service for sending emails. Even if it looks fine, we realize that this code breaks every aspect of SRP when we scratch the surface.

The responsibility of EmailService is not just to send emails but to store an email message into DB and send it via SMTP protocol.

Take a closer look at the sentence above. The word "and" is bold with purpose. Using such an expression does not look like the case where we describe a single responsibility.

> As soon as describing the responsibility of some code struct requires the usage of the word "and", it already breaks the Single Responsibility Principle.

---
##### References
    1.  https://levelup.gitconnected.com/practical-solid-in-golang-single-responsibility-principle-20afb8643483
    2.  