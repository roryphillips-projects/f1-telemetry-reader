package main

import (
	"encoding/json"
	"fmt"
	"github.com/roryphillips/f1-telemetry-client/internal"
	"github.com/roryphillips/f1-telemetry-client/internal/common"
	"github.com/roryphillips/f1-telemetry-client/internal/motion"
	"github.com/roryphillips/f1-telemetry-client/internal/session"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	hand := handler{packetParser: internal.NewPacketParser()}
	err := hand.demo("./data/1620373100")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	os.Exit(0)
}

type handler struct {
	packetParser internal.PacketParser
}

type writeReq struct {
	path string
	data []byte
}

func (h *handler) demo(rootDir string) error {
	errc := make(chan error, 1)
	writes := make(chan writeReq, 32)
	done := make(chan bool)

	wg := sync.WaitGroup{}

	// Load in a number of filepaths to parse and convert
	files, err := ioutil.ReadDir(rootDir)
	fileChan := make(chan os.FileInfo, len(files))
	if err != nil {
		return fmt.Errorf("unable to read directory: %v", err)
	}
	for _, file := range files {
		fileChan <- file
	}
	close(fileChan)

	// Spawn a number of concurrent file parsers
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			for file := range fileChan {
				err := h.reformatFile(writes, fmt.Sprintf("%s/%s", rootDir, file.Name()))
				if err != nil {
					errc <- err
				}
			}
			wg.Done()
		}()
	}

	go func() {
		bytesWritten := 0
		startTime := time.Now()
		for req := range writes {
			_ = ensureWritable(req.path)
			if err != nil {
				errc <- fmt.Errorf("failed to ensure writable path: %v", err)
			}

			err = ioutil.WriteFile(req.path, req.data, os.ModePerm)
			if err != nil {
				errc <- fmt.Errorf("failed to write json: %v", err)
			}
			bytesWritten += len(req.data)
		}
		endTime := time.Now()

		dur := endTime.Sub(startTime)
		fmt.Println(fmt.Sprintf("Wrote %v bytes in %v seconds", bytesWritten, dur.Seconds()))
	}()

	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case err := <-errc:
		if err != nil {
			return fmt.Errorf("failed to read file: %v", err)
		}
	case <-done:
		return nil
	}
	return nil
}

func (h *handler) reformatFile(writes chan<- writeReq, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read all file bytes: %v", err)
	}

	t, data, err := h.convertToJSON(bytes)
	if err != nil {
		return fmt.Errorf("failed to parse bytes: %v", err)
	}
	if len(data) > 0 && t != "" {
		path := transformPath(t, fileName)

		writes <- writeReq{
			path: path,
			data: data,
		}
	}

	return nil
}

func ensureWritable(path string) error {
	parts := strings.Split(path, "/")
	running := ""
	// The last part of the path will have an extension
	max := len(parts) - 1
	for i, part := range parts {
		if running != "" {
			running += "/"
		}
		running += part
		if i < max {
			if _, err := os.Stat(running); os.IsNotExist(err) {
				return os.Mkdir(running, os.ModeDir)
			}
		}
	}

	return nil
}

func transformPath(t string, in string) string {
	return strings.ReplaceAll(strings.ReplaceAll(in, ".data", ".json"), "data", fmt.Sprintf("output/%s", t))
}

func (h *handler) convertToJSON(data []byte) (string, []byte, error) {
	var t string
	var out []byte
	var header common.Header

	packet := internal.NewPacket(data)
	err := h.packetParser.Parse(packet, &header)
	if err != nil {
		return t, out, fmt.Errorf("failed to parse header %v", err)
	}

	switch header.PacketID {
	case common.PacketIDMotion:
		var parsed motion.Packet
		parsed.Header = header
		return h.parseJSON("motion", packet, &parsed)
	case common.PacketIDSession:
		var parsed session.Packet
		parsed.Header = header
		return h.parseJSON("session", packet, &parsed)
	}

	return t, out, nil
}

func (h *handler) parseJSON(t string, packet internal.Packet, dest interface{}) (string, []byte, error) {
	var out []byte

	err := h.packetParser.Parse(packet, dest)
	if err != nil {
		return t, out, fmt.Errorf("failed to parse motion packet: %v", err)
	}
	out, err = json.Marshal(dest)
	if err != nil {
		return t, out, fmt.Errorf("failed to marshal packet to json: %v", err)
	}
	return t, out, nil
}
