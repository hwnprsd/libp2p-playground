package squad

import (
	"fmt"
	"log"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/rules"
	"github.com/solace-labs/skeyn/utils"
	protob "google.golang.org/protobuf/proto"
)

func (s *Squad) RuleBookKey() string {
	return "RULES_" + s.ID
}

func (s *Squad) SetSpendingCap() error {
	return nil
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
		ruleBook = proto.RuleBook{WalletAddress: rule.WalletAddress, Rules: make([]*proto.AccessControlRule, 0)}
	} else {
		if err := protob.Unmarshal(data, &ruleBook); err != nil {
			log.Println("Error unmarshalling rulebook")
			return err
		}
	}

	// Check if the signature is valid
	sig := utils.HexToBytes(rule.Signature)

	// TODO: Check for collissions

	// Do not need to pass a sender for most cases
	isValid, err := s.scw.ValidateRuleAddition(rule.Rule.Bytes(), sig, common.ZeroAddr())
	if err != nil {
		log.Println("Error validating Rule - ", err.Error())
		return err
	}

	if !isValid {
		log.Println("Rule is invalid")
		return fmt.Errorf("Invalid Rule Submitted")
	}

	// Check if the rule collides with any existing rules
	ruleBook.Rules = append(ruleBook.Rules, rule.Rule)

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

	var spendingCap *proto.SpendingCap
	for _, r := range ruleBook.SpendingCaps {
		if r.Sender == tx.Sender.GetAddr() && r.TokenAddress == tx.GetToAddr() {
			spendingCap = r
		}
	}

	addr, _ := common.NewEthWalletAddressString(tx.Sender.Addr)

	_, err = rules.ValidateTx(tx, addr, ruleBook.Rules)
	if err != nil {
		return err
	}

	// Rules Pass
	// No spending cap set. Return
	if spendingCap == nil {
		return nil
	}

	val := int32(tx.GetValue())

	if spendingCap.CurrentValue+val > spendingCap.Cap {
		return fmt.Errorf("Transaction exceeds the cap rule val = %d, curr = %d, cap = %d", val, spendingCap.CurrentValue, spendingCap.Cap)
	}

	// At this point, we are all good - I think
	spendingCap.CurrentValue = spendingCap.CurrentValue + val

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
