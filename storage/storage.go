package storage

import (
	"fmt"
	"os"
)

func AppendMessage(topic string, msg string) error {
	err := os.MkdirAll("storage", 0755)
	if err != nil {
		return err
	}

	file := fmt.Sprintf("storage/%s.log", topic)

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(msg + "\n")
	return err
}
