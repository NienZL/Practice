package main

import (
	"fmt"
	"practice/fileWrite/model"
)

var (
	ChooseInputChan chan int
	file            model.File
	userInput       string
	fileContent     string
)

func init() {
	ChooseInputChan = make(chan int, 3)
	file = model.File{}
}

func main() {
	Start()
}

func Start() {
	var chooseInput int
	fmt.Println("選擇操作行為 1.建立新檔案  2.開啟舊檔案  3.寫入  4. 內容  5. 保存:")
	fmt.Scan(&chooseInput)
	ChooseInputChan <- chooseInput

	switch input := <-ChooseInputChan; input {
	case 1:
		create()
	case 2:
		open()
	case 3:
		write()
	case 4:
		read()
	case 5:
		save()
	default:
		fmt.Println("輸入錯誤! 請重新輸入數字選項!")
	}
	Start()
}

func create() {
	fmt.Println("請輸入新檔案之路徑(包含檔名與副檔名) : ")
	fmt.Scan(&userInput)

	file.Create(userInput)
}

func open() {
	fmt.Println("請輸入檔案之路徑(包含檔名與副檔名) : ")
	fmt.Scan(&userInput)
	file.Open(userInput)
}

func write() {
	fmt.Print("請輸入檔案之內容 : ")
	fmt.Scan(&fileContent)

	result := model.Calc(fileContent)
	file.Write(fileContent, result)
	fmt.Printf("%s = %s \n", fileContent, result)
}

func read() {
	fmt.Println("目前檔案名稱：" + file.Name + "." + file.SubName)
	fmt.Println("目前檔案內容如下：")
	fmt.Print(file.Content)
}

func save() {
	if file.Path == "" || file.Name == "" {
		fmt.Println("保存失敗，請確認路徑和檔案名稱是否正確。")
		fmt.Printf("檔案路徑 : %s\n檔案名稱 : %s.\n", file.Path, file.Name+"."+file.SubName)
	} else {
		fmt.Println("保存成功!")
	}
	// if err := file.Save(); err != nil {
	// 	fmt.Println("保存失敗，請確認路徑和檔案名稱。")
	// 	fmt.Printf("檔案路徑 : %s\n檔案名稱 : %s.\n", file.Path, file.Name+"."+file.SubName)
	// 	return
	// }

}

// type AbsPath struct {
// 	path, title, subtitle, wholePath string
// }

// var (
// 	ChooseInputChan        chan int
// 	result                 int
// 	userInput, fileContent string
// 	file                   AbsPath
// )

// func init() {
// 	ChooseInputChan = make(chan int, 3)
// }

// func main() {
// 	Start()
// }

// func pathConverter(input string) string {
// 	var replacer = strings.NewReplacer("\\", "/")
// 	input = replacer.Replace(input)
// 	return input
// }

// func Start() {
// 	var chooseInput int
// 	fmt.Println("輸入 1.建立一個新檔案  2.打開一個舊檔案 : ")
// 	fmt.Scan(&chooseInput)

// 	ChooseInputChan <- chooseInput

// 	switch input := <-ChooseInputChan; input {
// 	case 1:
// 		create()
// 	case 2:
// 		read()
// 	default:
// 		fmt.Println("輸入錯誤! 請重新輸入數字選項!")
// 		Start()
// 	}
// }

// func create() {
// 	// result := 0
// 	// fmt.Println("請輸入檔案之路徑(包含檔名與副檔名) : ")
// 	// fmt.Scan(&userInput)

// 	// userInput = pathConverter(userInput)
// 	// fileInfoConverter(userInput)
// 	askForPath()

// 	newFile, creatErr := os.Create(file.wholePath)
// 	if creatErr != nil {
// 		log.Fatal(creatErr)
// 	}

// 	// fmt.Println("請輸入檔案之內容 : ")
// 	// fmt.Scan(&fileContent)
// 	askForContent()

// 	// numArray := strings.Split(fileContent, "+")

// 	// for i := 0; i < len(numArray); i++ {
// 	// 	num, err := strconv.Atoi(numArray[i])
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	result += num
// 	// }

// 	newFile.WriteString(fileContent + "=" + strconv.Itoa(calc(fileContent)))

// 	// fileBuffer, readErr := ioutil.ReadFile(file.wholePath)
// 	// if readErr != nil {
// 	// 	fmt.Fprintf(os.Stderr, "Read File Error: %s\n", readErr)
// 	// }

// 	// fmt.Print("目前檔案內容如下：" + string(fileBuffer))
// 	printFile()
// }

// func read() {
// 	// fmt.Println("請輸入檔案之路徑(包含檔名與副檔名): ")
// 	// fmt.Scan(&userInput)

// 	// userInput = pathConverter(userInput)
// 	// fileInfoConverter(userInput)
// 	askForPath()

// 	// fileBuffer, readErr := ioutil.ReadFile(file.wholePath)
// 	// if readErr != nil {
// 	// 	log.Fatal(readErr)
// 	// }

// 	// fmt.Println("目前檔案內容如下：")
// 	// fmt.Printf("%s\n", string(fileBuffer))
// 	printFile()

// 	// fmt.Println("請輸入要修改的內容：")
// 	// fmt.Scan(&fileContent)
// 	askForContent()
// 	calc(fileContent)
// 	// result2 := fileContent + "=" + strconv.Itoa(result)
// 	err := ioutil.WriteFile(file.wholePath, []byte(fileContent+"="+strconv.Itoa(result)), 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// fileBuffer2, readErr := ioutil.ReadFile(file.wholePath)
// 	// if readErr != nil {
// 	// 	fmt.Fprintf(os.Stderr, "Read File Error: %s\n", readErr)
// 	// }

// 	// fmt.Print("目前檔案內容如下：" + string(fileBuffer2))
// 	printFile()
// }

// func fileInfoConverter(input string) {
// 	file.path = path.Dir(strings.Split(input, ".")[0])
// 	file.title = path.Base(strings.Split(input, ".")[0])
// 	file.subtitle = strings.Split(input, ".")[1]
// 	file.wholePath = input
// }

// func askForPath() {

// 	fmt.Println("請輸入檔案之路徑(包含檔名與副檔名) : ")
// 	fmt.Scan(&userInput)
// 	userInput = pathConverter(userInput)
// 	fileInfoConverter(userInput)

// }

// func askForContent() {
// 	fmt.Println("請輸入檔案之內容 : ")
// 	fmt.Scan(&fileContent)
// }

// func calc(input string) int {
// 	result = 0
// 	numArray := strings.Split(fileContent, "+")
// 	for i := 0; i < len(numArray); i++ {
// 		num, err := strconv.Atoi(numArray[i])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		result += num
// 	}

// 	return result
// }

// func printFile() {
// 	fileBuffer, readErr := ioutil.ReadFile(file.wholePath)
// 	if readErr != nil {
// 		fmt.Fprintf(os.Stderr, "Read File Error: %s\n", readErr)
// 	}

// 	fmt.Println("目前檔案內容如下：" + string(fileBuffer))
// }
