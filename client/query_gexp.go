package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type QueryGexpParam struct {

	// The start time for the query. This can be a relative or absolute timestamp.
	// The data type can only be string, int, or int64.
	// The value is required with non-zero value of the target type.
	Start interface{} `json:"start"`

	// An end time for the query. If not supplied, the TSD will assume the local
	// system time on the server. This may be a relative or absolute timestamp.
	// The data type can only be string, or int64.
	// The value is optional.
	End interface{} `json:"end,omitempty"`

	Exp string `json:"exp"`
}

func (query *QueryGexpParam) String() string {

	buffer := bytes.NewBuffer(nil)
	content, _ := json.Marshal(query)
	buffer.WriteString(fmt.Sprintf("%s\n", string(content)))
	return buffer.String()
}

func (c *clientImpl) QueryGexp(param QueryGexpParam) (*QueryResponse, error) {

	queryEndpoint := fmt.Sprintf("%s%s?start=%s&exp=%s", c.tsdbEndpoint, QueryGexpPath, param.Start, param.Exp)
	queryResp := QueryResponse{}

	if err := c.sendRequest(GetMethod, queryEndpoint, "", &queryResp); err != nil {
		return nil, err
	}
	return &queryResp, nil
}
