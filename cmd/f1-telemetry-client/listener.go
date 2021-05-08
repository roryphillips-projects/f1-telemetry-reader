package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const (
	// PACKET_BUFFER_KB buffer size in kilobytes to read packets into
	PACKET_BUFFER_KB = 2
)

func main() {
	listenForData()
}

func listenForData() {
	fmt.Println("Creating data export directory")
	dirName := fmt.Sprintf("./data/%v", time.Now().Unix())
	err := os.Mkdir(dirName, os.ModeDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenPacket("udp", "127.0.0.1:20777")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 4096; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(fmt.Sprintf("Starting listener %v", n))
			buffer := make([]byte, 1024 * PACKET_BUFFER_KB)
			for err == nil {
				fmt.Println("Reading Buffer")
				n, _, err := conn.ReadFrom(buffer)
				if err != nil {
					err = fmt.Errorf("failed to read connection: %v", err)
					continue
				}

				fileBuffer := make([]byte, n)
				copy(fileBuffer, buffer[:n])
				go func(timestamp int64) {
					_ = saveData(dirName, buffer[:n], timestamp)
				}(time.Now().UnixNano())
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	os.Exit(1)
}

func saveData(dirName string, data []byte, timestamp int64) error {
	fmt.Println("Saving data")
	file, err := os.Create(fmt.Sprintf("%v/%v.data", dirName, timestamp))
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data: %v", err)
	}
	return nil
}