package e2etests

import (
	"github.com/stretchr/testify/require"

	"github.com/zeta-chain/node/e2e/runner"
	"github.com/zeta-chain/node/e2e/utils"
	crosschaintypes "github.com/zeta-chain/node/x/crosschain/types"
)

func TestBitcoinDeposit(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 1)

	depositAmount := parseFloat(r, args[0])

	r.SetBtcAddress(r.Name, false)

	txHash := r.DepositBTCWithAmount(depositAmount)
	r.Logger.Print("🔍 waiting for the deposit to be mined: %s", txHash.String())

	// wait for the cctx to be mined
	cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, txHash.String(), r.CctxClient, r.Logger, r.CctxTimeout)
	r.Logger.CCTX(*cctx, "deposit")
	r.Logger.Print("✅ deposit cctx mined: %s", cctx.Index)
	utils.RequireCCTXStatus(r, cctx, crosschaintypes.CctxStatus_OutboundMined)
}
