package samples

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func UnmarshallSample() error {
	type User struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
		Password string `json:"-"`
		IsAdmin  bool   `json:"is_admin"`
	}

	var (
		err          error
		importedUser User
	)

	const (
		userFilePath = "data/output/user.json"
	)

	file, err := os.Open(userFilePath)
	if err != nil {
		return err
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, &importedUser)

	fmt.Println("Строковое представление объекта данных JSON:")
	fmt.Println(string(fileData))

	fmt.Printf("importedUser: \n%v\n", importedUser)

	return nil
}
