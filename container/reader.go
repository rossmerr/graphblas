// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package container

// Reader for importing data into a container
type Reader interface {
	Read() (record []string, err error)
}
