// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/lang"
)

type LangCondition = lang.Condition

// Lang returns a lang.ConditionStart for LangModel
func Lang() lang.ConditionStart {
	return lang.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}
