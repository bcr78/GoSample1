package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {

	input_file := flag.String("i", "input.csv", "an input file")
	output_file := flag.String("o", "output.csv", "an output file")
	input_separator := flag.String("ip_sep", ",", "an input separator")
	output_separator := flag.String("op_sep", "|", "an output separator")
	line := flag.String("line", "", "line break")
	flag.Parse()
	/*
		fmt.Println("input:", *input_file)
		fmt.Println("output:", *output_file)
		fmt.Println("input_separator:", *input_separator)
		fmt.Println("output_separator:", *output_separator)
		if *line != "" {
			fmt.Println("line break:", *line)
		}
	*/
	// var ip_sep = *input_separator

	csvFile, err := os.Open(*input_file) //("sample1.csv")
	if err != nil {
		fmt.Println("Input file name is missing ")
		return
	}
	reader := csv.NewReader(csvFile)
	reader.Comma, _ = utf8.DecodeLastRuneInString(*input_separator) //','

	//	info, _ := os.Stat(*input_file)
	//	mode := info.Mode()

	//	fmt.Println(csvFile, "mode is ", mode)

	csvfile2, err := os.Create(*output_file) //("sample3.csv")
	if err != nil {
		fmt.Println(err)
		opFilePath := *output_file
		fmt.Println("opFilePath=", opFilePath)
		fmt.Printf("type is %T\n", opFilePath)

		return
	}
	writer := csv.NewWriter(csvfile2)
	writer.Comma, _ = utf8.DecodeLastRuneInString(*output_separator) //output_separator //'|'
	//writer.UseCRLF = true

	//Read the file fully at a time
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error in reading file : ", err)
		return
	}
	for _, record := range csvData {
		if *line != "" {
			writer.UseCRLF = false
			lastField := record[len(record)-1]
			lastField = lastField + *line
			record[len(record)-1] = lastField
		} else {
			writer.UseCRLF = true
		}
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	writer.Flush()
	defer csvFile.Close()
	defer csvfile2.Close()

}
