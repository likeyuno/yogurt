/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

package static

import (
	"errors"
)

const (
	Ident = "\x20\x20"
)

type Respond struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type StatusFunc func() (int, string)

func RespondJSON(status StatusFunc, errs error, data interface{}) *Respond {
	code, message := status()
	respond := Respond{
		Code:    code,
		Message: message,
		Data:    data,
	}
	if errs != nil {
		if _, ok := errs.(interface{ Unwrap() error }); ok {
			for err := errs; err != nil; err = errors.Unwrap(err) {
				respond.Errors = append(respond.Errors, err.Error())
			}
		} else {
			respond.Errors = append(respond.Errors, errs.Error())
		}
	}
	return &respond
}

func OK() (int, string) {
	return 0, "ok"
}

func Error() (int, string) {
	return 10000, "unknown err"
}
