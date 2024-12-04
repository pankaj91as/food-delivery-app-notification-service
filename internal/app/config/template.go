package config

import (
	"fmt"
	"os"
)

func GetNotificationTemplate(QueueName, NotificationType string) (string, error) {
	fmt.Printf("Reading template from %s", "./template/"+QueueName+"/"+NotificationType+".tmpl\n")
	contents, err := os.ReadFile("./template/" + QueueName + "/" + NotificationType + ".tmpl")
	if err != nil {
		fmt.Println("File reading error", err)
		return "", err
	}
	return string(contents), nil
}
