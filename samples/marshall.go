package samples

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func MarshallSample() error {
	type User struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
		Password string `json:"-"`
		IsAdmin  bool   `json:"is_admin"`
	}

	const (
		filePathJson = "data/output/user.json"
	)

	user := User{
		Name:    "Ivan",
		Age:     32,
		Email:   "",
		IsAdmin: true,
	}

	jsonData, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		return err
	}

	// fileName := filepath.Base(filePathJson)
	dir := filepath.Dir(filePathJson)

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePathJson)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
