// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NpaPolicyResponseItem struct {
	RuleData *NpaPolicyRuleData `json:"rule_data,omitempty"`
	RuleID   *string            `json:"rule_id,omitempty"`
	RuleName *string            `json:"rule_name,omitempty"`
}

func (o *NpaPolicyResponseItem) GetRuleData() *NpaPolicyRuleData {
	if o == nil {
		return nil
	}
	return o.RuleData
}

func (o *NpaPolicyResponseItem) GetRuleID() *string {
	if o == nil {
		return nil
	}
	return o.RuleID
}

func (o *NpaPolicyResponseItem) GetRuleName() *string {
	if o == nil {
		return nil
	}
	return o.RuleName
}