// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package triple_test

import (
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/container/table"
	"github.com/RossMerr/Caudex.GraphBLAS/container/triple"
)

func TestNewTripleFromTable(t *testing.T) {

	type args struct {
		t  func(string) *triple.Store
		r  []string
		c  []string
		in string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Table Triples",
			args: args{
				in: `log_id src_ip server_ip
001 128.0.0.1 208.29.69.138
002 192.168.1.2 157.166.255.18
003 128.0.0.1 74.125.224.72`,
				t: func(in string) *triple.Store {
					table, _ := table.NewTableFromReader(3, 5, strings.NewReader(in))
					return triple.NewTripleStoreFromTable(table)
				},
				r: []string{"log_id|001", "log_id|001", "log_id|002", "log_id|002", "log_id|003", "log_id|003"},
				c: []string{"src_ip|128.0.0.1", "server_ip|208.29.69.138", "src_ip|192.168.1.2", "server_ip|157.166.255.18", "src_ip|128.0.0.1", "server_ip|74.125.224.72"},
			},
		},
		{
			name: "Table Transpose Triples",
			args: args{
				in: `log_id src_ip server_ip
001 128.0.0.1 208.29.69.138
002 192.168.1.2 157.166.255.18
003 128.0.0.1 74.125.224.72`,
				t: func(in string) *triple.Store {
					table, _ := table.NewTableFromReader(3, 5, strings.NewReader(in))
					return triple.NewTripleStoreFromTable(table).Transpose()
				},
				r: []string{"src_ip|128.0.0.1", "server_ip|208.29.69.138", "src_ip|192.168.1.2", "server_ip|157.166.255.18", "src_ip|128.0.0.1", "server_ip|74.125.224.72"},
				c: []string{"log_id|001", "log_id|001", "log_id|002", "log_id|002", "log_id|003", "log_id|003"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := tt.args.t(tt.args.in)
			for i, triple := range store.Triples {
				if tt.args.r[i] != triple.Row {
					t.Errorf("%+v got %+v, want %+v", tt.name, triple.Row, tt.args.r[i])
				}

				if tt.args.c[i] != triple.Column {
					t.Errorf("%+v got %+v, want %+v", tt.name, triple.Column, tt.args.c[i])
				}
			}
		})
	}
}
