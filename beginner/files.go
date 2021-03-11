package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	/**
	Создание файла
	*/
	//file, err := os.Create("test.txt")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Fatal(file)
	//file.Close()

	/**
	Очистка файла 100 байт
	*/
	//error := os.Truncate("test.txt", 100)
	//if error != nil {
	//	log.Fatal(error.Error())
	//}

	/**
	Информация о файле
	*/
	//fileInfo, err := os.Stat("text.txt")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Println(fileInfo.Name())
	//fmt.Println(fileInfo.Size())
	//fmt.Println(fileInfo.Mode())
	//fmt.Println(fileInfo.ModTime())
	//fmt.Println(fileInfo.IsDir())

	/**
	Переименовывание файла и перемещение
	*/
	//original := "test.txt"
	//newName := "test2.txt"
	//err := os.Rename(original, newName)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	file, err := os.Open("test2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	byteSlice := make([]byte, 160)
	file.Read(byteSlice)
	fmt.Println(string(byteSlice))

	data, err := ioutil.ReadFile("test2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(data))

}
