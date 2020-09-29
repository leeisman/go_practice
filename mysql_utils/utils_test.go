package mysql_utils

import "testing"

func TestGetTopUpRecordsByOrderIDs(t *testing.T) {
	type args struct {
		orderIDs []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "records",
			args: args{
				orderIDs: []string{"3821233532538191872", "3821232783875899392"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTopUpRecordsByOrderIDs(tt.args.orderIDs)
		})
	}
}
