// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/company"
)

type CompanyCondition = company.Condition

// Company returns a company.ConditionStart for CompanyModel
func Company() company.ConditionStart {
	return company.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}
