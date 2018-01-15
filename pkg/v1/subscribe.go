// MIT License
//
// Copyright (c) [2017-2018] [Demitri Swan]
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package v1

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/master"
)

type EventStream chan *master.Event

// Subscribe subscribes to events on the Mesos master. This method blocks, so
// you. likely want to call it in a go routine. Process each *master.Event on
// the EventStream by checking the type (you may call GetType() on the
// *master.Event), then processing the data as needed. See the test/cmd
// package for an example.
func (m *Master) Subscribe(ctx context.Context, es EventStream) (err error) {
	var callType master.Call_Type = master.Call_SUBSCRIBE
	var callMsg proto.Message = &master.Call{Type: &callType}
	var b []byte
	b, err = proto.Marshal(callMsg)
	if err != nil {
		return
	}
	var buf io.Reader = bytes.NewBuffer(b)
	var event *master.Event = &master.Event{}
	var httpResponse *http.Response
	httpResponse, err = m.client.doProtoWrapper(ctx, buf, nil)
	if err != nil {
		return
	}
	var reader *bufio.Reader = bufio.NewReader(httpResponse.Body)
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
			// Get size as string
			var sizeString string
			sizeString, err = reader.ReadString('\n')
			if err != nil {
				return
			}
			sizeString = strings.TrimSpace(sizeString)
			// Convert string to int64
			var sizeInt int64
			sizeInt, err = strconv.ParseInt(sizeString, 10, 64)
			if err != nil {
				return err
			}
			// Read data specified by the size
			var eventBytes = make([]byte, sizeInt)
			var remaining int
			remaining, err = reader.Read(eventBytes)
			if err != nil {
				return
			}
			for int64(remaining) < sizeInt {
				var moreEvent []byte = make([]byte, remaining)
				remaining, err = reader.Read(moreEvent)
				if err != nil {
					return
				}
				eventBytes = append(eventBytes, moreEvent...)
			}
			// Unmarshal data into a master.Event
			err = proto.Unmarshal(eventBytes, event)
			if err != nil {
				return
			}
			es <- event
		}
	}
}
