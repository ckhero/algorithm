/**
 *@Description
 *@ClassName merge_sort_test.go
 *@Date 2021/6/5 下午5:36
 *@Author ckhero
 */

package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
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
			if got := MergeSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}