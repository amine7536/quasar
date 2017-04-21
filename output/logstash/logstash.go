package outputlogstash

import (
	"encoding/json"
	"net"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/amine7536/quasar/conf"
	"github.com/amine7536/quasar/event"
)

// Name output name
const Name = "logstash"

// Output main type
type Output struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Conn net.Conn
}

// DefaultOutputConfig return default config struct
func DefaultOutputConfig() Output {
	return Output{}
}

// InitOutput initialize an output
func InitOutput(configRaw *interface{}, logger *logrus.Entry) (conf.OutputHandler, error) {
	config := DefaultOutputConfig()
	err := conf.ReflectConfig(configRaw, &config)
	if err != nil {
		return &config, err
	}

	if err := config.Connect(); err != nil {
		logger.Fatalf("Unable to connect to %s:%s", config.Host, config.Port)
	}

	return &config, err
}

// Connect to logstash
func (t *Output) Connect() error {
	service := t.Host + ":" + t.Port

	// Open TCP connection
	if _, err := net.ResolveTCPAddr("tcp4", service); err != nil {
		return err
	}

	// conn, err := net.DialTimeout("tcp", nil, tcpAddr)
	timeout := time.Duration(30 * time.Second)
	conn, errDial := net.DialTimeout("tcp", service, timeout)
	if errDial != nil {
		return errDial
	}

	t.Conn = conn

	return nil
}

// Send output
func (t *Output) Send(event event.Event) error {

	// defer t.Conn.Close()
	out, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if connected := isAlive(t.Conn, []byte{}); connected == false {
		err := t.Connect()
		if err != nil {
			return err
		}
	}

	// Event
	conn := t.Conn
	_, err = conn.Write(out)
	if err != nil {
		return err
	}

	// Send EOF
	_, errEOF := conn.Write([]byte("\r\n\r\n"))
	if errEOF != nil {
		return errEOF
	}

	return nil
}

func isAlive(c net.Conn, buffer []byte) bool {
	_, err := c.Write(buffer)
	if err != nil {
		c.Close()
		return false
	}
	return true
}
