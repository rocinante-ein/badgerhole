package entities

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// NewWWWLog create new wwwlog instance
func NewWWWLog(r *http.Request) *WWWLogImpl {

	// parse request parameters
	r.ParseForm()

	// create wwwlog instance from request parameter
	var wwwlog = &WWWLogImpl{
		RequestID:   uuid.New(),
		Source:      r.RemoteAddr,
		SourceIP:    strings.Split(r.RemoteAddr, ":")[0],
		SourcePort:  parsePort(r.RemoteAddr),
		Dist:        r.Host,
		DistIP:      strings.Split(r.Host, ":")[0],
		DistPort:    parsePort(r.Host),
		HTTPMethod:  r.Method,
		URL:         r.URL.String(),
		RequestURI:  r.RequestURI,
		Referer:     r.Referer(),
		UserAgent:   r.UserAgent(),
		Headers:     r.Header,
		RequestBody: r.PostForm,
		// Future Update:
		//    Multipart form data file registered quarantine directory(Write ONLY)
		//		MultipartForm: r.MultipartForm.File,
		Timestamp: time.Now(),
	}

	return wwwlog
}

// WWWLogImpl honey pod log format struct
type WWWLogImpl struct {
	RequestID   uuid.UUID           `json:"request_id"`
	Source      string              `json:"source"`
	SourceIP    string              `json:"source_ip"`
	SourcePort  int                 `json:"source_port"`
	Dist        string              `json:"dist"`
	DistIP      string              `json:"dist_ip"`
	DistPort    int                 `json:"dist_port"`
	HTTPMethod  string              `json:"http_method"`
	URL         string              `json:"url"`
	RequestURI  string              `json:"request_uri"`
	Referer     string              `json:"referer"`
	UserAgent   string              `json:"user_agent"`
	Headers     map[string][]string `json:"headers"`
	RequestBody map[string][]string `json:"request_body"`
	Timestamp   time.Time           `json:"timestamp"`
}

// ToJSON Cast JSON byte slice
func (wlf WWWLogImpl) ToJSON() []byte {
	jsonLog, _ := json.MarshalIndent(wlf, "", "   ")

	return jsonLog
}

// GetRequestID get unique request_id
func (wlf WWWLogImpl) GetRequestID() string {
	return wlf.RequestID.String()
}

// private module parse port
func parsePort(host string) int {

	splitHost := strings.Split(host, ":")

	if len(splitHost) < 2 {
		return 80
	}

	hostPort, _ := strconv.Atoi(splitHost[1])

	return hostPort
}
