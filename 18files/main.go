package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
f,err := os.Open("myfile.txt")

if err != nil {
    panic(err)
}

fileInfo,err :=  f.Stat()

if err != nil {
    panic(err)
}

fmt.Println("fileinfo name", fileInfo.Name())
fmt.Println("fileinfo size", fileInfo.Size())
fmt.Println("fileinfo permissions", fileInfo.Mode())
fmt.Println("fileinfo modified at", fileInfo.ModTime())
fmt.Println("fileinfo isDir?", fileInfo.IsDir())


defer f.Close() // close the file when we're done



// buff := make([]byte, 10)

// d, err := f.Read(buff)
// if err != nil {
// 	panic(err)
// }

// fmt.Printf("read %d bytes: %s\n", d, string(buff))

//read

//for small file read 
// file,err := os.ReadFile("myfile.tsx")

// if err != nil {
//     panic(err)
// }

// fmt.Println(string(file))


//read folders

dir , err := os.Open("../")
checkNilError(err)

defer dir.Close()

fileInformation , err := dir.ReadDir(-1)

checkNilError(err)

for _, fi := range fileInformation {
    fmt.Println(fi.Name())
}



//create file
crf , err := os.Create("new.txt")
checkNilError(err)

defer crf.Close()

// crf.WriteString("hello world")
// crf.WriteString("append data to file")

bytes :=[]byte("hello world")
crf.Write(bytes)


//read and write to another file ( streaming mode)
srcfile , err := os.Open("myfile.txt")

checkNilError(err)

defer srcfile.Close()


dstfile , err := os.Create("new2.txt")
checkNilError(err)

defer dstfile.Close()

reader :=bufio.NewReader(srcfile)
writer := bufio.NewWriter(dstfile)

for {
	b,err:= reader.ReadByte()
	if err == io.EOF {
        break
    }

	e := writer.WriteByte(b)
	if e != nil {
        panic(e)
    }


}

writer.Flush() // make sure everything is written to the file

fmt.Println("written to new file")


//delelte  the file

err = os.Remove("new.txt")
if err != nil{
	fmt.Println("failed to delete new.txt")
    
}

fmt.Println("deleted new.txt")


}


func checkNilError(err error) error{
	if err != nil {
        panic(err)
    }
    return nil
}