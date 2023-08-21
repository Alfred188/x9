//write x9 program to read argument and get urls from arguments in both arg mode and pip mode and seperate parameters.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// -h => help
// --help => help
// -u => url
var Url string

// -ul => url list
var UrlList []string

// -w => wordlist
var Wordlist []string

// -v => Value
var Values []string

// -c => chunk
var Chunk int

// -gs => generate strategy {normal,ignore,combine,all}
var GenerateStrategy string

// -vs => value strategy {replace,suffix}
var ValueStrategy string

// -o => output
var Output string

// -s, --silent => silent
var Silent bool

// result list
var ResultList []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	argsWithoutProg := os.Args[1:]

	GenerateStrategy = "all"
	ValueStrategy = "suffix"
	Chunk = 15
	Silent = false

	index := 1
	for _, arg := range argsWithoutProg {
		switch arg {
		case "-h", "--help":
			help()
			os.Exit(0)
		case "-u", "--url":
			GetUrl(argsWithoutProg[index])
		case "-ul", "--url-list":
			GetUrlList(argsWithoutProg[index])
		case "-w", "--wordlist":
			GetWordlist(argsWithoutProg[index])
		case "-v", "--value":
			GetValue(argsWithoutProg[index])
		case "-c", "--chunk":
			GetChunk(argsWithoutProg[index])
		case "-gs":
			GetGenerateStrategy(argsWithoutProg[index])
		case "-vs":
			GetValueStrategy(argsWithoutProg[index])
		case "-s", "--silent":
			Silent = true
		case "-o", "--output":
			GetOutput(argsWithoutProg[index])
		default:
			// fmt.Println("Unknown Argument")
		}

		// fmt.Println(index, ":", arg)
		index++
	}

	if Url != "" {
		// url mode
		Prosess(Url)
	} else if len(UrlList) > 0 {
		// url list mode
		for _, url := range UrlList {
			Prosess(url)
		}
	} else {
		fmt.Println("No url or url list provided")
	}

	if !Silent {
		logo()
	}

	Out()
}

func logo() {
	fmt.Println("         ________")
	fmt.Println("___  ___/   __   \\")
	fmt.Println("\\  \\/  /\\____    /")
	fmt.Println(" >    <    /    / ")
	fmt.Println("/__/\\__\\  /____/  ")
	fmt.Println("               ")
	fmt.Println(" v1.0.0 ")
	fmt.Println("Author: Alfred(MJ)")
	fmt.Println("E-Mail: Mohsen.Mahmoudjanlou@gmail.com")
	fmt.Println("Starting to process ... ")
}

func help() {
	fmt.Println("-h, --help			Show this help message and exit")
	fmt.Println("-u, --url			Target url")
	fmt.Println("-ul, --url-list			Target url list")
	fmt.Println("-w, --wordlist			Wordlist")
	fmt.Println("-v, --value			Value to insert")
	fmt.Println("-c, --chunk			Chunk, Default: 15")
	fmt.Println("-gs				Generate strategy {normal,ignore,combine,all}, Default: all \r\n\t\t\t\t\t Select the mode strategy from the available options: \r\n\t\t\t\t\t   normal: Remove all parameters and put the worldlist \r\n\t\t\t\t\t   ignore: Don't toch the URL and put the wordlist \r\n\t\t\t\t\t   combine: pitchfork combine on the existing parameters \r\n\t\t\t\t\t   all: all the above methods")
	fmt.Println("-vs				Value strategy {replace,suffix}, Default: suffix \r\n\t\t\t\t\tSelect the mode strategy from the available options: \r\n\t\t\t\t\t  replace: replace the value of the parameter with the gathered value \r\n\t\t\t\t\t  suffix: append the value to the end of the parameters")
	fmt.Println("-s, --silent			Silent")
	fmt.Println("-o, --output			Output results")
}

func GetUrl(url string) {
	Url = url
	// fmt.Println("URL:", Url)
}

func GetUrlList(path string) {
	//read file and get urls per line and add to UrlList
	dat, err := os.ReadFile(path)
	check(err)

	lins := strings.Split(string(dat), "\n")
	for _, line := range lins {
		if line != "" {
			UrlList = append(UrlList, line)
		}
	}

	// fmt.Println("URL List:", UrlList)
}

func GetWordlist(path string) {
	dat, err := os.ReadFile(path)
	check(err)

	lins := strings.Split(string(dat), "\n")
	for _, line := range lins {
		if line != "" {
			Wordlist = append(Wordlist, line)
		}
	}

	// fmt.Println("Wordlist:", Wordlist)
}

func GetValue(value string) {
	if value != "" {
		Values = strings.Split(value, ",")
	} else {
		fmt.Println("Value (-v, --value) is required")
		os.Exit(0)
	}
}

func GetChunk(chunk string) {
	ck, err := strconv.Atoi(chunk)
	check(err)
	Chunk = ck
}

func GetGenerateStrategy(generateStrategy string) {
	GenerateStrategy = strings.ToLower(generateStrategy)
}

func GetValueStrategy(valueStrategy string) {
	ValueStrategy = strings.ToLower(valueStrategy)
}

func GetOutput(output string) {
	Output = output
}

func Prosess(url string) {
	if len(Values) > 0 {
		// has value
		for _, Value := range Values {
			if strings.Contains(url, "?") {
				// has parameters
				parts := strings.Split(url, "?")

				if GenerateStrategy == "normal" {
					// normal mode
					if len(Wordlist) > 0 {
						// has wordlist
						for i := 0; i < len(Wordlist); i += Chunk {
							urls := ""
							for j := i; j < i+Chunk; j++ {
								if j < len(Wordlist) {
									if j > i {
										urls += "&"
									}
									urls += Wordlist[j] + "=" + Value
								}
							}
							ResultList = append(ResultList, parts[0]+"?"+urls)
						}
					} else {
						// no wordlist
						fmt.Println("Wordlist (-w) is required")
					}
				} else if GenerateStrategy == "ignore" {
					// ignore mode
					if len(Wordlist) > 0 {
						// has wordlist
						for i := 0; i < len(Wordlist); i += Chunk {
							urls := ""
							for j := i; j < i+Chunk; j++ {
								if j < len(Wordlist) {
									if j > i {
										urls += "&"
									}
									urls += Wordlist[j] + "=" + Value
								}
							}
							ResultList = append(ResultList, url+"&"+urls)
						}
					} else {
						// no wordlist
						fmt.Println("Wordlist (-w) is required")
					}

				} else if GenerateStrategy == "combine" {
					if strings.Contains(parts[1], "&") {
						parameters := strings.Split(parts[1], "&")

						len := len(parameters)
						for i := 0; i < len; i++ {
							c := 0
							strs := ""
							for _, parameter := range parameters {
								if c > 0 {
									strs += "&"
								}

								if i == c {
									if ValueStrategy == "replace" {
										// replace value
										param := strings.Split(parameter, "=")
										strs += param[0] + "=" + Value
									} else {
										// suffix value
										strs += parts[1] + Value
									}
								} else {
									strs += parameter
								}
								c++
							}
							ResultList = append(ResultList, parts[0]+"?"+strs)
						}
					} else {
						// only one parameter

						if ValueStrategy == "replace" {
							// replace value
							parameter := strings.Split(parts[1], "=")
							ResultList = append(ResultList, parts[0]+"?"+parameter[0]+"="+Value)
						} else {
							// suffix value
							ResultList = append(ResultList, parts[0]+"?"+parts[1]+Value)
						}
					}
				} else if GenerateStrategy == "all" {
					//combine
					if strings.Contains(parts[1], "&") {
						parameters := strings.Split(parts[1], "&")

						len := len(parameters)
						for i := 0; i < len; i++ {
							c := 0
							strs := ""
							for _, parameter := range parameters {
								if c > 0 {
									strs += "&"
								}

								if i == c {
									if ValueStrategy == "replace" {
										// replace value
										param := strings.Split(parameter, "=")
										strs += param[0] + "=" + Value
									} else {
										// suffix value
										strs += parts[1] + Value
									}
								} else {
									strs += parameter
								}
								c++
							}
							ResultList = append(ResultList, parts[0]+"?"+strs)
						}
					} else {
						// only one parameter

						if ValueStrategy == "replace" {
							// replace value
							parameter := strings.Split(parts[1], "=")
							ResultList = append(ResultList, parts[0]+"?"+parameter[0]+"="+Value)
						} else {
							// suffix value
							ResultList = append(ResultList, parts[0]+"?"+parts[1]+Value)
						}
					}
					// -- combine --
					// normal mode
					if len(Wordlist) > 0 {
						// has wordlist
						for i := 0; i < len(Wordlist); i += Chunk {
							urls := ""
							for j := i; j < i+Chunk; j++ {
								if j < len(Wordlist) {
									if j > i {
										urls += "&"
									}
									urls += Wordlist[j] + "=" + Value
								}
							}
							ResultList = append(ResultList, parts[0]+"?"+urls)
						}
					}
					// ignore mode
					if len(Wordlist) > 0 {
						// has wordlist
						for i := 0; i < len(Wordlist); i += Chunk {
							urls := ""
							for j := i; j < i+Chunk; j++ {
								if j < len(Wordlist) {
									if j > i {
										urls += "&"
									}
									urls += Wordlist[j] + "=" + Value
								}
							}
							ResultList = append(ResultList, url+"&"+urls)
						}
					}

				}

			} else {
				// no parameters
				if len(Wordlist) > 0 {
					// has wordlist
					for i := 0; i < len(Wordlist); i += Chunk {
						urls := ""
						for j := i; j < i+Chunk; j++ {
							if j < len(Wordlist) {
								if j > i {
									urls += "&"
								}
								urls += Wordlist[j] + "=" + Value
							}
						}
						ResultList = append(ResultList, url+"?"+urls)
					}
				} else {
					// no wordlist
					fmt.Println("Wordlist (-w, --wordlist) is required")
					os.Exit(0)
				}
			}
		}
	} else {
		fmt.Println("Value (-v, --value) is required")
		os.Exit(0)
	}
}

func Out() {
	if Output != "" {
		// write to file
		f, err := os.Create(Output)
		check(err)
		defer f.Close()

		for _, result := range ResultList {
			_, err := f.WriteString(result + "\n")
			check(err)
		}

		f.Sync()
	} else {
		// write to stdout
		for _, result := range ResultList {
			fmt.Println(result)
		}
	}
}
