package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Recording is successful")
	file.Close()
}
