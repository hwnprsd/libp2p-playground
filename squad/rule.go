package squad

import (
	"fmt"
	"log"

	"github.com/solace-labs/skeyn/proto"
	protob "google.golang.org/protobuf/proto"
)

func (s *Squad) RuleBookKey() string {
	return "RULES_" + s.ID
}

func (s *Squad) CreateRule(rule *proto.CreateRuleData) error {
	// TODO: VALIDATION
	// TODO: Verify with the owner
	data, err := s.db.Get([]byte(s.RuleBookKey()))
	if err != nil {
		return err
	}

	var ruleBook proto.RuleBook

	// Unmarshall and Marshal the new rulebook
	// TODO: Some sort of validation such that the rules don't clash

	if data == nil {
		ruleBook = proto.RuleBook{WalletAddress: rule.WalletAddress, Rules: make([]*proto.Rule, 0)}
	} else {
		if err := protob.Unmarshal(data, &ruleBook); err != nil {
			log.Println("Error unmarshalling rulebook")
			return err
		}
	}

	ruleBook.SpendingCap = append(ruleBook.SpendingCap, rule.SpendingCap)

	if err != nil {
		log.Println("Error marshalling rule to bytes")
		return err
	}

	data, err = protob.Marshal(&ruleBook)
	if err != nil {
		return fmt.Errorf("Error marshalling rulebook back to DB - %e", err)
	}

	err = s.db.Set([]byte(s.RuleBookKey()), data)
	if err != nil {
		return fmt.Errorf("Error writing rulebook back to DB - %e", err)
	}

	log.Printf("New Spending Cap Set - \nCurr = %d\nCap = %d\n", rule.SpendingCap.CurrentValue, rule.SpendingCap.Cap)

	return nil
}

func (s *Squad) validateTx(tx *proto.SolaceTx) error {
	data, err := s.db.Get([]byte(s.RuleBookKey()))
	if err != nil {
		return err
	}

	if err != nil {
		log.Println("Error opening rulebook")
		return err
	}

	if data == nil {
		log.Println("Rulebook is empty")
		return fmt.Errorf("Rulebook is empty")
	}

	var ruleBook proto.RuleBook
	if err := protob.Unmarshal(data, &ruleBook); err != nil {
		log.Println("Error unmarshalling rulebook")
		return err
	}

	var rule *proto.SpendingCap
	for _, r := range ruleBook.SpendingCap {
		if r.Sender == tx.Sender.GetAddr() && r.TokenAddress == tx.GetToAddr() {
			rule = r
		}
	}

	if rule == nil {
		return fmt.Errorf("No rule exists for the given sender")
	}

	val := int32(tx.GetValue())

	if rule.CurrentValue+val > rule.Cap {
		return fmt.Errorf("Transaction exceeds the cap rule val = %d, curr = %d, cap = %d", val, rule.CurrentValue, rule.Cap)
	}

	// At this point, we are all good - I think
	rule.CurrentValue = rule.CurrentValue + val

	data, err = protob.Marshal(&ruleBook)
	if err != nil {
		return fmt.Errorf("Error marshalling rulebook back to DB - %e", err)
	}

	err = s.db.Set([]byte(s.RuleBookKey()), data)
	if err != nil {
		return fmt.Errorf("Error writing rulebook back to DB - %e", err)
	}

	return nil

}
