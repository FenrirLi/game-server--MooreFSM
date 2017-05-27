package machine

import (
	"fmt"
)

var TableRulesGroup = map[string][]TableRules{
	"TABLE_RULE_READY":{},
	"TABLE_RULE_DEAL":{},
	"TABLE_RULE_STEP":{&LiuJuRule{}},
	"TABLE_RULE_END":{},
	"TABLE_RULE_SETTLE_FOR_ROUND":{&SettleForRoundRule{}},
}

//===========================TableRulesManager===========================
type TableRulesManager struct {}
func (this *TableRulesManager) condition( table Table, rule_group string ) bool {
	//依据检验的组对规则进行遍历
	if rules_array, ok := TableRulesGroup[rule_group]; ok {
		for _,rule := range rules_array {
			//满足规则则进行处理
			if rule.Condition( table ) {
				rule.Action( table )
			}
		}
	} else {
		fmt.Println("Manager : rule_group Not Found")
	}

	table.Machine.CurrentStatus.NextStatus( table )
	return false
}