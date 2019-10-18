//Расчет суммы и количество платежей в файле реестра, направляемого в адрес ООО РМК
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
        "strconv"
)


func loadCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", path, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '|'            //разделитель полей в файле scv
	r.Comment = '#'          // символ комментария
	r.LazyQuotes = true      // разрешить ковычки в полях
	rows, err := r.ReadAll() //прочитать весь файл в массив [][]string
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}

var dN [][]string
var numbers []float64

func main() {
    log.SetPrefix("grep: ")
    log.SetFlags(0) // no extra info in log messages

    if len(os.Args) != 2 {
        fmt.Printf("Usage: %v FILE\n", os.Args[0])
        return
    }

    dN = loadCSV(os.Args[1])
    var iline int
    var sSumm float64

	for _, line := range dN {
 		//		vline = strings.Join(line, ";") // объединим поля в строку с применением разделителя ;
		//	if strings.Contains(vline, key){ // если найдено в строке ключевое слово
		iline = iline + 1
                // Fio            Address                                                          Ls     Summ    Period Peny   Xz1      Xz2
                //Кулагин Г. И.|Липецкая обл., Данковский район, с. Ягодное, ул. Молодежная, д. 4|1652796|123.45|092019|0.00||22481755|18.09.2019
		//fmt.Printf("%s\n",line[3])
                i, err := strconv.ParseFloat(line[3], 64)
                if err == nil {
                 numbers = append(numbers, i)
	         sSumm = sSumm + i
                }
	}

        fmt.Println(numbers)    
        fmt.Println(len(numbers))    
	fmt.Printf("Количетсво записей в реестре %s  %10d шт.\n",os.Args[1],iline)
	fmt.Printf("Сумма по реестру             %s  %10.2f \n",os.Args[1], sSumm)


}
