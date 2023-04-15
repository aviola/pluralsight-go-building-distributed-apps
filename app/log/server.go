package log

import (
	"io/ioutil"
	stlog "log" // Redefine the standard log package to avoid name clash with the current package
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

// fileLog implements the io.Writer interface.
// Write() opens the log file, writes the log and closes the file.
func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(data)
}

// Run() takes the input 'destination' string, converts it to a fileLog (that is an io.Writer) and instantiates the logger.
func Run(destination string) {
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

// Register HTTP endpoints
func RegisterHandlers() {
	// Register the HTTP endpoint for the /log route, providing an handler function
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		// The handler function reads the content of the request body...
		msg, err := ioutil.ReadAll(r.Body)
		// ...and in case of error while reading or empty message...
		if err != nil || len(msg) == 0 {
			// ...reply with an HTTP response with error code 400: Bad Request
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// Otherwise, convert the message into a string and write it to the log file
		write(string(msg))
	})
}

// write() takes the message as input and writes to the log file using our custom logger.
func write(message string) {
	log.Printf("%v\n", message)
}
