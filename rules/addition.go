package rules

import "fmt"

// Validate that a new rule can be added to the ACL List
func ValidateRuleAddition(rules ACL, newRule ACLRule) error {
	// Generate IDs for all the existing rules
	ids := make(map[string]struct{}, 0)

	for _, rule := range rules {
		for _, id := range getRuleIds(rule) {
			ids[id] = struct{}{}
		}
	}

	for _, id := range getRuleIds(newRule) {
		_, exists := ids[id]
		if exists {
			return fmt.Errorf("Rule collision")
		}
	}

	// Check if any IDs collide
	// TODO: More advanced checks for VR
	return nil
}
