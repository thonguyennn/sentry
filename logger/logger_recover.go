/* !!
 * File: logger_sentry.go
 * File Created: Friday, 9th August 2019 9:47:11 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Friday, 9th August 2019 9:47:11 am
 * Modified By: KimEricko™ (phamkim.pr@gmail.com>)
 * -----
 * Copyright 2018 - 2019 mySoha Platform, VCCorp
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Developer: NhokCrazy199 (phamkim.pr@gmail.com)
 */

package logger

import (
	"errors"
	"fmt"
	"time"

	sentry "github.com/getsentry/sentry-go"
)

// SendErrorToSentry func
// Send Error log to screen and sentry system
// CaptureError formats and delivers an error to the Sentry server using the default *Client.
// Adds a stacktrace to the packet, excluding the call to this method.
func SendErrorToSentry(err error) {
	go func() {
		sentry.CaptureException(err)
		sentry.Flush(5 * time.Second)
	}()
}

func SentryRecover() {
	sentry.Recover()
}

// SendErrorTextToSentry func
// Send Error log to screen and sentry system
// CaptureError formats and delivers an error to the Sentry server using the default *Client.
// Adds a stacktrace to the packet, excluding the call to this method.
func SendErrorTextToSentry(errorText string) {
	go func() {
		err := errors.New(errorText)
		sentry.CaptureException(err)
		sentry.Flush(5 * time.Second)
	}()
}

// SendArrayArgsToSentry func
// Send Error log to screen and sentry system
// CaptureError formats and delivers an error to the Sentry server using the default *Client.
// Adds a stacktrace to the packet, excluding the call to this method.
func SendArrayArgsToSentry(format string, args ...interface{}) {
	go func() {
		err := fmt.Errorf(format, args...)
		sentry.CaptureException(err)
		sentry.Flush(5 * time.Second)
	}()
}

func init() {
	sentry.Init(sentry.ClientOptions{
		Dsn:   "https://b813f0e6efc149c7808da4969969da91@msh-st.sohatv.vn/13",
		Debug: true,
	})
}
