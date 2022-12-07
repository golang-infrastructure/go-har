package har

import (
	"encoding/json"
	"time"
)

// ParseHar 解析HAR格式的文件
func ParseHar(harFileBytes []byte) (*Har, error) {
	har := new(Har)
	err := json.Unmarshal(harFileBytes, har)
	if err != nil {
		return nil, err
	}
	return har, nil
}

// ------------------------------------------------- --------------------------------------------------------------------

type Har struct {
	Log Log `json:"log"`
}

type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type PageTimings struct {
	OnContentLoad float64 `json:"onContentLoad"`
	OnLoad        float64 `json:"onLoad"`
	Comment       string  `json:"comment"`
}

type Pages struct {
	StartedDateTime time.Time   `json:"startedDateTime"`
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	PageTimings     PageTimings `json:"pageTimings"`
}

type Headers struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Request struct {
	Method      string    `json:"method"`
	URL         string    `json:"url"`
	HTTPVersion string    `json:"httpVersion"`
	Cookies     []Cookie  `json:"cookies"`
	Headers     []Headers `json:"headers"`
	QueryString []any     `json:"queryString"`
	HeadersSize int       `json:"headersSize"`
	BodySize    int       `json:"bodySize"`
}

type Cookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path"`
	Domain   string    `json:"domain"`
	Expires  time.Time `json:"expires"`
	HTTPOnly bool      `json:"httpOnly"`
	Secure   bool      `json:"secure"`
	SameSite string    `json:"sameSite"`
}

type Content struct {
	Size     int    `json:"size"`
	MimeType string `json:"mimeType"`
}

type Response struct {
	Status       int       `json:"status"`
	StatusText   string    `json:"statusText"`
	HTTPVersion  string    `json:"httpVersion"`
	Cookies      []Cookie  `json:"cookies"`
	Headers      []Headers `json:"headers"`
	RedirectURL  string    `json:"redirectURL"`
	HeadersSize  int       `json:"headersSize"`
	BodySize     int       `json:"bodySize"`
	Content      Content   `json:"content"`
	TransferSize int       `json:"_transferSize"`
	Error        any       `json:"_error"`
}

type BeforeRequest struct {
}

type AfterRequest struct {
}

type Cache struct {
	BeforeRequest BeforeRequest `json:"beforeRequest"`
	AfterRequest  AfterRequest  `json:"afterRequest"`
	Comment       string        `json:"comment"`
}

type Timings struct {
	Blocked         float64     `json:"blocked"`
	DNS             float64     `json:"dns"`
	Connect         float64     `json:"connect"`
	Send            float64     `json:"send"`
	Wait            float64     `json:"wait"`
	Receive         float64     `json:"receive"`
	Ssl             float64     `json:"ssl"`
	BlockedQueueing float64 `json:"_blocked_queueing"`
	BlockedProxy    float64 `json:"_blocked_proxy"`
}

type Entries struct {
	StartedDateTime time.Time `json:"startedDateTime"`
	Time            float64       `json:"time"`
	Request         Request   `json:"request"`
	Response        Response  `json:"response"`
	Cache           Cache     `json:"cache"`
	Timings         Timings   `json:"timings"`
	Pageref         string    `json:"pageref"`
	Initiator       Initiator `json:"_initiator"`
	Priority        string    `json:"_priority"`
	ResourceType    string    `json:"_resourceType"`
	Connection      string    `json:"connection"`
	ServerIPAddress string    `json:"serverIPAddress"`
}

type Initiator struct {
	Type       string `json:"type"`
	URL        string `json:"url"`
	LineNumber int    `json:"lineNumber"`
	Stack      Stack  `json:"stack"`
}

type Stack struct {
	CallFrames []CallFrame `json:"callFrames"`
	Parent     Parent      `json:"parent"`
}

type Parent struct {
	Parent      *Parent     `json:"parent"`
	Description string      `json:"description"`
	CallFrames  []CallFrame `json:"callFrames"`
	ParentID    ParentID    `json:"parentId"`
}

type ParentID struct {
	ID         string `json:"id"`
	DebuggerID string `json:"debuggerId"`
}

type CallFrame struct {
	FunctionName string `json:"functionName"`
	ScriptID     string `json:"scriptId"`
	URL          string `json:"url"`
	LineNumber   int    `json:"lineNumber"`
	ColumnNumber int    `json:"columnNumber"`
}

type Log struct {
	Version string    `json:"version"`
	Creator Creator   `json:"creator"`
	Pages   []Pages   `json:"pages"`
	Entries []Entries `json:"entries"`
}

// ------------------------------------------------- --------------------------------------------------------------------
