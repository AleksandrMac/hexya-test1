// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/partner_industry"
)

type PartnerIndustryCondition = partner_industry.Condition

// PartnerIndustry returns a partner_industry.ConditionStart for PartnerIndustryModel
func PartnerIndustry() partner_industry.ConditionStart {
	return partner_industry.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}