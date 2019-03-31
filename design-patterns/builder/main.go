package main

import (
	"encoding/json"
	"fmt"
)

// Notification is our final product. It represents complex object that
// construction we would like to encapsulate. In most cases we want to
// keep it simple and clean from any additional initialization logic.
type Notification struct {
	Body   []byte
	Format string
}

// NotificationBuilder is an abstract interface that defines a template
// for the steps to construct our product. Each of specific builders
// should satisfy that interface. This is the place where we define
// signature of a 'processes' that happens during initialization.
type NotificationBuilder interface {
	SetRecipient(recipient string)
	SetBody(body string)
	Notification() (*Notification, error)
}

// JSONNotificationBuilder is our concrete builder implementation. It satisfies
// NotificationBuilder implicitly. It may consist different set of members which
// not need to be related closely to our final product. Implementation of
// the NotificationBuilder (methods) is the place, where complex building logic
// should be placed. As an example we can take the place where notification body
// becomes marshaled, or simple input length validation happen.
type JSONNotificationBuilder struct {
	recipient string
	body      string
}

// SetRecipient is the part of NotificationBuilder implementation.
func (b *JSONNotificationBuilder) SetRecipient(recipient string) {
	lengthValidator := func(s string) bool {
		return len(s) > 0
	}

	switch isValid := lengthValidator(recipient); isValid {
	case true:
		b.recipient = recipient
	case false:
		b.recipient = "Default recipient"
	}
}

// SetBody is the part of NotificationBuilder implementation.
func (b *JSONNotificationBuilder) SetBody(body string) {
	lengthValidator := func(s string) bool {
		return len(s) > 0
	}

	switch isValid := lengthValidator(body); isValid {
	case true:
		b.body = body
	case false:
		b.body = "Default body"
	}
}

// Notification is the part of NotificationBuilder implementation.
func (b *JSONNotificationBuilder) Notification() (*Notification, error) {
	result := make(map[string]string)
	result["recipient"] = b.recipient
	result["body"] = b.body

	body, err := json.Marshal(result)

	if err != nil {
		return nil, err
	}

	return &Notification{body, "JSON"}, nil
}

// BuildNotification is our 'director' function. Its main responsibility is to
// construct the object through the builder interface. Important to notice here
// is the type of the builder parameter. It clearly indicates that we expect here
// not a specific type, but any type that satisfies general NotificationBuilder
// interface. As a result, we get function that could be used to build different kind
// of notifications.
func BuildNotification(builder NotificationBuilder) (*Notification, error) {
	builder.SetRecipient("Recipient")
	builder.SetBody("Body")
	return builder.Notification()
}

func main() {

	jsonBuilder := JSONNotificationBuilder{}
	jsonNotification, err := BuildNotification(&jsonBuilder)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Notification format: %+v \n", jsonNotification.Format)
	fmt.Printf("Notification body:   %+v \n", string(jsonNotification.Body))
}
