package validation

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
)

const pathToStyle = "styles/standard.txt"

func CheckStyleModification() error {
	originalHash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	currentHash, err := HashFile(pathToStyle)
	if err != nil {
		return err
	}

	if originalHash != currentHash {
		err := errors.New("File has been modified.")
		return err
	}

	return nil
}

func HashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
