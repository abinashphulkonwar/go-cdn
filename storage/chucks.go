package storage

import (
	"fmt"
	"io"
	"os"
)

func ReadChunks() {
	// Open the file (replace "file.txt" with the path to your file)
	file, err := os.Open("../temp/638e0509715a1b79a57b24eca1669694")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Define the buffer size for reading chunks
	var chunkSize int64 = 100 * 1024

	buffer := make([]byte, chunkSize)
	_, err = file.Seek(chunkSize, io.SeekStart)
	if err != nil {
		fmt.Println("Error seeking file:", err)
	}

	n, err := file.Read(buffer)
	println(n)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
	}

	processChunk(buffer[:n])

}

func processChunk(chunk []byte) {
	// Replace this function with your desired processing logic for each chunk
	fmt.Println("Processing chunk:", chunk[0], len(chunk))
}
