/**
 *@Description
 *@ClassName quick_sort_test
 *@Date 2021/6/5 下午3:32
 *@Author ckhero
 */

package quick

import (
	"fmt"
	"reflect"
	"testing"
)

type Test interface {
	Test()
	Test2()
	Test22()
}


func TestQuickSort(t *testing.T) {
	data := []int{6,1,2,7,9,3,4,5,10,8}
	QuickSort(data, 0, len(data) - 1)
	fmt.Println(data)
}

func TestQuickSort1(t *testing.T) {
	type args struct {
		data []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name:"test",
			args:args{data:[]int{6,1,2,7,9,3,4,5,10,8}, low:0, high:9},
			want: []int{1,2,3,4,5,6,7,8,9,10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.data; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}