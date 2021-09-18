package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/macmagic/technical-test-deporvillage/internal/adapter"
	"github.com/macmagic/technical-test-deporvillage/internal/application/config"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

var cfg *config.Config

func TestE2e(t *testing.T) {
	a := require.New(t)

	setup := func() {
		cfg = &config.Config{
			Host:                 "localhost",
			Port:                 "5000",
			ConnectionType:       "tcp",
			MaxClientConnections: 5,
			SkuLogPath:           "./test_sku.log",
		}

		go func() {
			adapter.Run(cfg)
		}()

		time.Sleep(2 * time.Second)
	}

	t.Run("Given the system is running", func(t *testing.T) {
		setup()

		t.Run("When 5 clients tries to connect to the server", func(t *testing.T) {
			client1 := createClient(cfg)
			client2 := createClient(cfg)
			client3 := createClient(cfg)
			client4 := createClient(cfg)
			client5 := createClient(cfg)

			t.Run("And writes SKUS", func(t *testing.T) {
				sendMessage(client1, "ABCD-2222")
				sendMessage(client2, "SSSS-3333")
				sendMessage(client3, "2222-5555")
				sendMessage(client4, "2sssss")
				sendMessage(client5, "4444-SDAS")

				t.Run("Then the log file must contains two elements", func(t *testing.T) {
					file, err := os.OpenFile(cfg.SkuLogPath, os.O_RDWR, 755)

					if err != nil {
						log.Fatalln("Cannot open the file", err.Error())
					}

					lines, err := lineCounter(file)

					if err != nil {
						log.Fatalln("Error when try to read the log", err.Error())
					}

					a.Equal(3, lines)
				})
			})
		})
	})
}

func createClient(appConfig *config.Config) net.Conn {
	conn, err := net.Dial(appConfig.ConnectionType, fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port))

	if err != nil {
		log.Fatalln("Cannot connect to the server", err.Error())
	}

	return conn
}

func sendMessage(conn net.Conn, message string) {
	_, err := fmt.Fprintln(conn, message)

	if err != nil {
		log.Println("Cannot write the message", err.Error())
	}
}

func lineCounter(r io.Reader) (int, error) {

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
