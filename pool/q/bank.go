// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/bank"
)

type BankCondition = bank.Condition

// Bank returns a bank.ConditionStart for BankModel
func Bank() bank.ConditionStart {
	return bank.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}
