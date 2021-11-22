package array

import "testing"

func TestTwoSum(t *testing.T) {

	nums := []int{1, 2, 3, 4}
	target := 7

	res := TwoSum(nums, target)
	if (res[0] == 2 && res[1] == 3) || (res[0] == 3 && res[1] == 2) {
		t.Log("passed")
	} else {
		t.Errorf("got(%v) expected(%v)", res, []int{2, 3})
	}

}
