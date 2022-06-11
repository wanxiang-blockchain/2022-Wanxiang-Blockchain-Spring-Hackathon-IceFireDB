/*
 *
 *  * Licensed to the Apache Software Foundation (ASF) under one or more
 *  * contributor license agreements.  See the NOTICE file distributed with
 *  * this work for additional information regarding copyright ownership.
 *  * The ASF licenses this file to You under the Apache License, Version 2.0
 *  * (the "License"); you may not use this file except in compliance with
 *  * the License.  You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

const (
	tab = "    "
)

type Record struct {
	Name string
	File string
	Line int
}

func (r *Record) String() string {
	if r == nil {
		return "[nil-record]"
	}
	return fmt.Sprintf("%s:%d %s", r.File, r.Line, r.Name)
}

type Stack []*Record

func Trace() Stack {
	return TraceN(1, 32)
}

func (s Stack) String() string {
	return s.StringWithIndent(0)
}

func (s Stack) StringWithIndent(indent int) string {
	var b bytes.Buffer
	for i, r := range s {
		for j := 0; j < indent; j++ {
			fmt.Fprint(&b, tab)
		}
		fmt.Fprintf(&b, "%-3d %s:%d\n", len(s)-i-1, r.File, r.Line)
		for j := 0; j < indent; j++ {
			fmt.Fprint(&b, tab)
		}
		fmt.Fprint(&b, tab, tab)
		fmt.Fprint(&b, r.Name, "\n")
	}
	if len(s) != 0 {
		for j := 0; j < indent; j++ {
			fmt.Fprint(&b, tab)
		}
		fmt.Fprint(&b, tab, "... ...\n")
	}
	return b.String()
}

func TraceN(skip, depth int) Stack {
	s := make([]*Record, 0, depth)
	for i := 0; i < depth; i++ {
		r := Caller(skip + i + 1)
		if r == nil {
			break
		}
		s = append(s, r)
	}
	return s
}

func Caller(skip int) *Record {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return nil
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil || strings.HasPrefix(fn.Name(), "runtime.") {
		return nil
	}
	return &Record{
		Name: fn.Name(),
		File: file,
		Line: line,
	}
}
