/**
 *@Description
 *@ClassName insert_sort_test.go
 *@Date 2021/6/6 下午3:47
 *@Author ckhero
 */

package insert_sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name:"test",
			args:args{data:[]int{6,1,2,7,9,3,4,5,10,8}},
			want: []int{1,2,3,4,5,6,7,8,9,10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}