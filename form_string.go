// Code generated by "stringer -type=Form"; DO NOT EDIT

package apd

import "fmt"

const _Form_name = "FiniteInfiniteNaNSignalingNaN"

var _Form_index = [...]uint8{0, 6, 14, 26, 29}

func (i Form) String() string {
	if i < 0 || i >= Form(len(_Form_index)-1) {
		return fmt.Sprintf("Form(%d)", i)
	}
	return _Form_name[_Form_index[i]:_Form_index[i+1]]
}
