package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ErrNotFoundFile = errors.New("файл не найден")
var ErrFileNotJSON = errors.New("файл JSON не найден")

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("файл не найден")
			return nil, ErrNotFoundFile
		}
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func ReadJSONFile(name string) ([]byte, error) {
	if strings.ToLower(filepath.Ext(name)) != ".json" {
		return nil, ErrFileNotJSON
	}

	data, err := os.ReadFile(name)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("файл не найден")
			return nil, ErrNotFoundFile
		}
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string){
	file, err := os.Create(name)
	if err != nil{
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}