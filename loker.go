package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var tipe []string
var nomor []string
var lowest_index int

func init_loker(){
	for len(tipe) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := strings.Fields(scanner.Text())

		if len(command)==2 && strings.ToLower(command[0]) == "init" {
			loker_allocated, err := strconv.Atoi(command[1])
			if err != nil {
				fmt.Println("Inisialisasi loker: init [angka]")
				continue
			}
			if loker_allocated < 1 {
				fmt.Println("Inisialisasi loker harus lebih dari 0")
			} else {
				for i:=0;i<loker_allocated;i++{
					tipe = append(tipe,"kosong")
					nomor = append(nomor, "")
				}
				fmt.Println("Berhasil membuat loker dengan jumlah", command[1])
			}
		} else if len(command)==1 && strings.ToLower(command[0]) == "exit"{
			os.Exit(1)
		}else {
			fmt.Println("Inisialisasi loker: init [angka]")
		}

	}
}

func get_status(){
	fmt.Println("No Loker Tipe Identitas\tNo Identitas")
	for i := range tipe {
		fmt.Println(strconv.Itoa(i+1), "\t", tipe[i], "\t\t", nomor[i])
	}
}

func get_lowest_index(){
	updated := false
	for i:=lowest_index; i<len(tipe);i++ {
		if nomor[i] == ""{
			updated = true
			lowest_index = i
			break
		}
	}
	if !updated {
		lowest_index = -1
	}
}

func input_data(input_tipe string, input_nomor string){
	if lowest_index>=0 {
		tipe[lowest_index] = input_tipe
		nomor[lowest_index] = input_nomor
		fmt.Println("Kartu identitas tersimpan di loker nomor", strconv.Itoa(lowest_index+1))
		get_lowest_index()
	} else {
		fmt.Println("Maaf loker sudah penuh")
	}
}

func delete_data(no_loker string){
	int_no_loker, err := strconv.Atoi(no_loker)
	if err == nil && int_no_loker>0 {
		if int_no_loker>len(tipe){
			fmt.Println("Tidak ada loker dengan nomor", no_loker)
		} else if nomor[int_no_loker-1]==""{
			fmt.Println("Loker sudah kosong")
		} else {
			tipe[int_no_loker-1] = "kosong"
			nomor[int_no_loker-1] = ""
			if lowest_index == -1 || int_no_loker-1 < lowest_index {
				lowest_index = int_no_loker-1
			}
			fmt.Println("Loker nomor", no_loker, "berhasil dikosongkan")
		}
	} else {
		fmt.Println("Nomor loker harus berupa angka positif")
	}
}

func find_data(no_identitas string){
	var found[] string
	for i:= range nomor {
		if nomor[i] == no_identitas {
			found = append(found, strconv.Itoa(i+1))
		}
	}
	if len(found) == 0 {
		fmt.Println("Nomor identitas tidak ditemukan")
	} else {
		fmt.Println("Kartu identitas tersebut berada di loker nomor", strings.Join(found, ","))
	}
}

func main() {
	lowest_index = 0
	init_loker()

	running := true
	for running {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := strings.Fields(scanner.Text())
		if len(command)<1 {
			fmt.Println("Available commands: status, input, leave, find, exit")
			continue
		}
		switch strings.ToLower(command[0]) {
			case "status":
				if len(command)>1{
					fmt.Println("usage: status")
				} else{
					get_status()
				}
			case "input":
				if len(command) == 3{
					input_tipe := command[1]
					input_nomor := command[2]
					input_data(input_tipe, input_nomor)
				} else {
					fmt.Println("usage: input [tipe identitas] [nomor identitas]")
				}
			case "leave":
				if len(command) == 2 {
					no_loker := command[1]
					delete_data(no_loker)
				} else {
					fmt.Println("usage: leave [nomor loker]")
				}
			case "find":
				if len(command) == 2 {
					no_identitas := command[1]
					find_data(no_identitas)
				} else {
					fmt.Println("usage: find [no. identitas]")
				}
			case "init":
				fmt.Println("Loker sudah diinisialisasi")
			case "exit":
				fmt.Println("Closing...")
				running = false
			default:
				fmt.Println("Available commands: status, input, leave, find, exit")
		}
	}
}