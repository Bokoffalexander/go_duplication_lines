/*
Печатает количество джублируемых строк, которые появились больше 1 раза.
Что в input:  
Считывает стандартный ввод Stdin или
считыввает из файлов (в аргумент передаем имя файла.txt).
Пример: ./duplication file.txt 
*/
package main

import (
	"bufio" // input := bufio.NewScanner(f)
	"fmt"
	"os" // os.Args[1:] os.Open(arg) 
)

func main() {
	fmt.Println("It counts duplications of lines") // Описание

	counts := make(map[string]int) // делаем МЕЙК МАП[ключ]значение
	files := os.Args[1:] // инициализируем массив (там строковые аргументы, если указали аргументы)
	fmt.Printf("INFO of args: the type is %T and the length is %v .\n", files, len(files)) // Печатаем тип массива и длину
	
	if len(files) == 0 { // Если НЕ указывали аргумент file.txt 
		fmt.Println("Type some lines below.")
		fmt.Println("For exit: type 0")
		countLines(os.Stdin, counts) // Передаем функции countLines: стандартный ввод и пустую МАП
	} else { // Если УКАЗЫВАЛИ аргументы file1.txt file2.txt ..
		for _, arg := range files { // Запустим цикл по аргументам. Первый file1.txt
			f, err := os.Open(arg) // Открываем файл назовем f
			if err != nil {
				fmt.Fprintf(os.Stderr, "duplication: %v\n", err) // Передаем в Стандартный Интерфейс Ошибки саму ошибку
				continue // идем на следующий аргумент
			}
			countLines(f, counts) // Отправляем функции countLines файл1 и пустую МАП
			f.Close() // Закрываем файл f
		}
	}

	printDuplications(counts) // Передаем функции печати параметр: заполненную МАП
}

func countLines(f *os.File, counts map[string]int) { // Считает сколько раз каждая строка была
	
	input := bufio.NewScanner(f) // Открываем Новый Сканер с переданным аргументом f
	for input.Scan() { // Считываем из f: либо стандартный ввод os.Stdin либо из file.txt
		counts[input.Text()]++
		if input.Text()=="0" {
			break
		}
	}
	// Игнорируем потенциальные ошибки из input.Err()
}

func printDuplications(counts map[string]int) { // Печатаем повторения
	fmt.Println("List of lines, more that once:")

	for line, n := range counts { // Создаем цикл проходов по МАП
		if n > 1 { // Если строка появилась больше одного раза, то на ПЕЧАТЬ
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

