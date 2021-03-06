// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cerror

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type errorTestSuite struct {
	suite.Suite
}

const (
	// EcodeUnknown is unknown error info
	EcodeUnknown = 10009999
	// EcodeNotFile errors for operate on dir but file is required
	EcodeNotFile = 10000001
	// EcodeNotDir errors for operate on file but dir is required
	EcodeNotDir = 10000002
	// EcodeNotExists errors for operate on target but doesn't exists
	EcodeNotExists = 10000003
	// EcodeExists errors for Add target but already exists
	EcodeExists = 10000004
	// EcodeDirNotEmpty errors for Remove directory but directory has child etc
	EcodeDirNotEmpty = 10000005
)

var templateError = map[int]string{
	EcodeUnknown:   "Unknown Error",
	EcodeNotFile:   "Target is Not File",
	EcodeNotDir:    "Target is Not Dir",
	EcodeNotExists: "Target is not exists",
	EcodeExists:    "Target is exists",
}

func (s *errorTestSuite) SetupTest() {
	errorsMessage = templateError
}

func (s *errorTestSuite) TearDownTest() {
	errorsMessage = map[int]string{}
}

func (s *errorTestSuite) TestNewError() {
	for k, v := range errorsMessage {
		e := NewError(k, v)
		s.Equal(k, e.ErrorCode)
		s.Equal(v, e.Message)
		s.Equal(v, e.Cause)
	}
}

func (s *errorTestSuite) TestNewErrorUnkownCode() {
	code := 0
	cause := "Unknown"

	e := NewError(code, cause)
	s.Equal(code, e.ErrorCode)
	s.Equal("", e.Message)
	s.Equal(cause, e.Cause)
}

func (s *errorTestSuite) TestJSONString() {
	e := NewError(EcodeNotDir, "TestJSONString")
	str := e.JSONString()

	str2, err := json.Marshal(e)
	s.NoError(err)
	s.Equal(string(str2), str)
}

func (s *errorTestSuite) TestJSONStringError() {
	marshal = func(interface{}) ([]byte, error) {
		return nil, errors.New("Error Marshal failed")
	}
	defer func() {
		marshal = json.Marshal
	}()

	e := NewError(EcodeNotDir, "TestJSONString")
	str := e.JSONString()

	str2, err2 := json.Marshal(e)
	s.NoError(err2)
	s.Equal(string(str2), str)
}

func (s *errorTestSuite) TestSetErrorMessageOK() {
	errorsMessage = map[int]string{}
	SetErrorsMessage(templateError)

	s.Equal(len(errorsMessage), len(templateError))
	for k, v := range templateError {
		v1, ok := errorsMessage[k]
		s.True(ok)
		s.Equal(v, v1)
	}
}

func (s *errorTestSuite) TestSetErrorMessageReplace() {
	errorsMessage = map[int]string{}
	SetErrorsMessage(templateError)

	otherMessage := map[int]string{
		100:         "100",
		200:         "200",
		EcodeNotDir: "EcodeNotDir",
	}
	SetErrorsMessage(otherMessage)

	s.Equal(len(templateError)+2, len(errorsMessage))
	for k, v := range templateError {
		v1, ok := errorsMessage[k]
		s.True(ok)

		_, ok = otherMessage[k]
		if !ok {
			s.Equal(v, v1)
		}
	}

	for k, v := range otherMessage {
		v1, ok := errorsMessage[k]
		s.True(ok)
		s.Equal(v, v1)
	}
}

func (s *errorTestSuite) TestIsOk() {
	type testCase struct {
		description string
		err         *Error
		errCode     int
		target      bool
	}
	testCases := []testCase{
		{
			description: "normal test ErrorNotFile",
			err:         NewError(EcodeNotFile, ""),
			errCode:     EcodeNotFile,
			target:      true,
		},
		{
			description: "normal test ErrorNotDir",
			err:         NewError(EcodeNotDir, ""),
			errCode:     EcodeNotDir,
			target:      true,
		},
	}
	for _, tc := range testCases {
		actual := Is(tc.err, tc.errCode)
		if actual != tc.target {
			s.Failf(tc.description, "expect %v, got %v", tc.target, actual)
		}
	}
}

func (s *errorTestSuite) TestIsFailed() {
	type testCase struct {
		description string
		err         error
		errCode     int
		target      bool
	}
	var err *Error
	testCases := []testCase{
		{
			description: "ErrorCode doesn't match failed",
			err:         NewError(EcodeNotFile, ""),
			errCode:     EcodeNotDir,
			target:      false,
		},
		{
			description: "nil error failed",
			err:         nil,
			errCode:     EcodeNotDir,
			target:      false,
		},
		{
			description: "nil error value failed",
			err:         err,
			errCode:     EcodeNotDir,
			target:      false,
		},
		{
			description: "error type match failed",
			err:         fmt.Errorf(""),
			errCode:     EcodeNotDir,
			target:      false,
		},
	}
	for _, tc := range testCases {
		actual := Is(tc.err, tc.errCode)
		if actual != tc.target {
			s.Failf(tc.description, "expect %v, got %v", tc.target, actual)
		}
	}
}

func (s *errorTestSuite) TestIsErrorOk() {
	type testCase struct {
		description string
		err         error
		target      bool
	}
	var err *Error
	testCases := []testCase{
		{
			description: "normal test",
			err:         NewError(EcodeDirNotEmpty, ""),
			target:      true,
		},
		{
			description: "nil error value test",
			err:         err,
			target:      true,
		},
	}
	for _, tc := range testCases {
		actual := IsError(tc.err)
		if actual != tc.target {
			s.Failf(tc.description, "expect %v, got %v", tc.target, actual)
		}
	}
}

func (s *errorTestSuite) TestIsErrorFailed() {
	type testCase struct {
		description string
		err         error
		target      bool
	}
	testCases := []testCase{
		{
			description: "nil error failed",
			err:         nil,
			target:      false,
		},
		{
			description: "error type match failed",
			err:         fmt.Errorf(""),
			target:      false,
		},
	}
	for _, tc := range testCases {
		actual := IsError(tc.err)
		if actual != tc.target {
			s.Failf(tc.description, "expect %v, got %v", tc.target, actual)
		}
	}
}

func TestErrorTestSuite(t *testing.T) {
	s := &errorTestSuite{}
	suite.Run(t, s)
}
