//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Each test suite can be run parallel, so it should not exist dependency
// between test different suite. However, inisde the suite itself, it can
// have dependency between Test1 and Test2, example:
// BTCStakingTestSuite has Test2SubmitCovenantSignature which
// depends on Test1CreateFinalityProviderAndDelegation for creating the
// finality provider used to query fp delegations and do its check.

// IBCTransferTestSuite tests IBC transfer end-to-end
func TestIBCTranferTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(IBCTransferTestSuite))
}

// TestBTCTimestampingTestSuite tests BTC timestamping protocol end-to-end
func TestBTCTimestampingTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BTCTimestampingTestSuite))
}

// TestBTCTimestampingPhase2HermesTestSuite tests BTC timestamping phase 2 protocol end-to-end,
// with the Hermes relayer
func TestBTCTimestampingPhase2HermesTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BTCTimestampingPhase2HermesTestSuite))
}

// TestBTCTimestampingPhase2RlyTestSuite tests BTC timestamping phase 2 protocol end-to-end,
// with the Go relayer
func TestBTCTimestampingPhase2RlyTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BTCTimestampingPhase2RlyTestSuite))
}

// TestBTCStakingTestSuite tests BTC staking protocol end-to-end
func TestBTCStakingTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BTCStakingTestSuite))
}

// TestSoftwareUpgradeTestSuite tests software upgrade protocol end-to-end
func TestSoftwareUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(SoftwareUpgradeVanillaTestSuite))
}
