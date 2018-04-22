// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package table_test

import (
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/container/table"
)

func TestTable_Read(t *testing.T) {
	tests := []struct {
		name string
		in   string
		t    func(string) table.Table
	}{
		{
			name: "Explode Table",
			in: `log_id src_ip server_ip
001 128.0.0.1 208.29.69.138
002 192.168.1.2 157.166.255.18
003 128.0.0.1 74.125.224.72`,
			t: func(in string) table.Table {
				table, _ := table.NewTableFromReader(3, 5, strings.NewReader(in))
				return table
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			table := tt.t(tt.in)

			r1 := "log_id|001"
			r2 := "log_id|002"
			r3 := "log_id|003"
			c1 := "src_ip|128.0.0.1"
			c2 := "src_ip|192.168.1.2"
			c3 := "server_ip|157.166.255.18"
			c4 := "server_ip|208.29.69.138"
			c5 := "server_ip|74.125.224.72"

			if v := table.GetFloat64(r1, c1); v != 1 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 1)
			}

			if v := table.GetFloat64(r1, c2); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r1, c3); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r1, c4); v != 1 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 1)
			}

			if v := table.GetFloat64(r1, c5); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r2, c1); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r2, c2); v != 1 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 1)
			}

			if v := table.GetFloat64(r2, c3); v != 1 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 1)
			}

			if v := table.GetFloat64(r2, c4); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r2, c5); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r3, c1); v != 1 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 1)
			}

			if v := table.GetFloat64(r3, c2); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r3, c3); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r3, c4); v != 0 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 0)
			}

			if v := table.GetFloat64(r3, c5); v != 1 {
				t.Errorf("%+v Value = %+v, want %+v", tt.name, v, 1)
			}
		})
	}
}
