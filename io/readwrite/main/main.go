package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// 1. 復制文件
	originalFile, err := os.Open("test001.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopy, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopy)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

	// 2. 寫文件
	file, err := os.OpenFile(
		"test001.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSliceWritten := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSliceWritten)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

	// 3. 快速寫入
	err = ioutil.WriteFile("test001.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 4. ***緩存寫***(重要)
	// 4.1 先寫入緩存
	// 4.2 再寫入硬盤
	fileBufferWrite, err := os.OpenFile("test001.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileBufferWrite.Close()

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(fileBufferWrite)

	// 写字节到buffer
	bytesBufferWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesBufferWritten)

	// 写字符串到buffer
	// 也可以使用 WriteRune() 和 WriteByte()
	bytesBufferWritten, err = bufferedWriter.WriteString(
		"Buffered string\n",
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesBufferWritten)

	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// *写内存buffer到硬盘*
	bufferedWriter.Flush()

	// 丢弃还没有flush的缓存的内容，清除错误并把它的输出传给参数中的writer
	// 当你想将缓存传给另外一个writer时有用
	bufferedWriter.Reset(bufferedWriter)

	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	bufferedWriter = bufio.NewWriterSize(
		bufferedWriter,
		8000,
	)

	// resize后检查缓存的大小
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 5. 讀取文件
	fileRead, err := os.Open("test001.txt")
	if err != nil {
		log.Fatal(err)
	}

	// os.File.Read(), io.ReadFull() 和
	// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	dataReadAll, err := ioutil.ReadAll(fileRead)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as hex: %x\n", dataReadAll)

	// 6. 快讀到內存
	dataReadFile, err := ioutil.ReadFile("test001.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", dataReadFile)

	// 7. ***使用緩存讀***(重要)
	fileBufferRead, err := os.Open("test001.txt")
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(fileBufferRead)

	// 得到字节，当前指针不变
	byteSliceRead := make([]byte, 5)
	byteSliceRead, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSliceRead)

	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSliceRead)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSliceRead)

	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)

	// 读取到分隔符，包含分隔符，返回byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read bytes: %s\n", dataBytes)

	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read string: %s\n", dataString)

}
