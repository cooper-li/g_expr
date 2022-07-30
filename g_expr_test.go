package g_expr

import (
	"testing"
)

type rule struct {
	data map[string]interface{}
	rule string
	res  bool
}

func TestMatch(t *testing.T) {

	ruleList := []rule{
		{
			data: map[string]interface{}{
				"name": "cooper",
				"age":  11,
				"sex":  2,
			},
			rule: `(age > 10 && in_data("name", "cooper|jack")) || sex == 1`,
			res:  true,
		},
		{
			data: map[string]interface{}{
				"name": "codee",
				"age":  11,
				"addr": "aaa",
			},
			rule: `addr == "aaa" && sex == 1`,
			res:  false,
		},
	}

	for _, ca := range ruleList {
		res, err := Match(ca.rule, ca.data)
		if err != nil && res != ca.res {
			t.Logf("res: %v, err: %v", ca.res, err)
		}
	}
}
