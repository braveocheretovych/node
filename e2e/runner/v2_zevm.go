package runner

import (
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	testgatewayzevmcaller "github.com/zeta-chain/node/pkg/contracts/testgatewayzevmcaller"
	"github.com/zeta-chain/protocol-contracts/v2/pkg/gatewayzevm.sol"
)

var gasLimit = big.NewInt(1000000)

// V2ETHWithdraw calls Withdraw of Gateway with gas token on ZEVM
func (r *E2ERunner) V2ETHWithdraw(
	receiver ethcommon.Address,
	amount *big.Int,
	revertOptions gatewayzevm.RevertOptions,
) *ethtypes.Transaction {
	tx, err := r.GatewayZEVM.Withdraw(
		r.ZEVMAuth,
		receiver.Bytes(),
		amount,
		r.ETHZRC20Addr,
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}

// V2ETHWithdrawAndCall calls WithdrawAndCall of Gateway with gas token on ZEVM
func (r *E2ERunner) V2ETHWithdrawAndCall(
	receiver ethcommon.Address,
	amount *big.Int,
	payload []byte,
	revertOptions gatewayzevm.RevertOptions,
) *ethtypes.Transaction {
	tx, err := r.GatewayZEVM.WithdrawAndCall(
		r.ZEVMAuth,
		receiver.Bytes(),
		amount,
		r.ETHZRC20Addr,
		payload,
		gasLimit,
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}

// V2ETHWithdrawAndCall calls WithdrawAndCall of Gateway with gas token on ZEVM using authenticated call
func (r *E2ERunner) V2ETHWithdrawAndAuthenticatedCall(
	receiver ethcommon.Address,
	amount *big.Int,
	payload []byte,
	revertOptions gatewayzevm.RevertOptions,
) *ethtypes.Transaction {
	tx, err := r.GatewayZEVM.WithdrawAndCall2(
		r.ZEVMAuth,
		receiver.Bytes(),
		amount,
		r.ETHZRC20Addr,
		payload,
		gatewayzevm.CallOptions{
			IsArbitraryCall: false,
			GasLimit:        gasLimit,
		},
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}

// V2ETHWithdrawAndCall calls WithdrawAndCall of Gateway with gas token on ZEVM using authenticated call
// through contract
func (r *E2ERunner) V2ETHWithdrawAndAuthenticatedCallThroughContract(
	gatewayZEVMCaller *testgatewayzevmcaller.TestGatewayZEVMCaller,
	receiver ethcommon.Address,
	amount *big.Int,
	payload []byte,
	revertOptions testgatewayzevmcaller.RevertOptions,
) *ethtypes.Transaction {
	tx, err := gatewayZEVMCaller.WithdrawAndCallGatewayZEVM(
		r.ZEVMAuth,
		receiver.Bytes(),
		amount,
		r.ETHZRC20Addr,
		payload,
		testgatewayzevmcaller.CallOptions{
			IsArbitraryCall: false,
			GasLimit:        gasLimit,
		},
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}

// V2ERC20Withdraw calls Withdraw of Gateway with erc20 token on ZEVM
func (r *E2ERunner) V2ERC20Withdraw(
	receiver ethcommon.Address,
	amount *big.Int,
	revertOptions gatewayzevm.RevertOptions,
) *ethtypes.Transaction {
	tx, err := r.GatewayZEVM.Withdraw(
		r.ZEVMAuth,
		receiver.Bytes(),
		amount,
		r.ERC20ZRC20Addr,
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}

// V2ERC20WithdrawAndCall calls WithdrawAndCall of Gateway with erc20 token on ZEVM
func (r *E2ERunner) V2ERC20WithdrawAndCall(
	receiver ethcommon.Address,
	amount *big.Int,
	payload []byte,
	revertOptions gatewayzevm.RevertOptions,
) *ethtypes.Transaction {
	// this function take more gas than default 500k
	// so we need to increase the gas limit
	previousGasLimit := r.ZEVMAuth.GasLimit
	r.ZEVMAuth.GasLimit = 10000000
	defer func() {
		r.ZEVMAuth.GasLimit = previousGasLimit
	}()

	tx, err := r.GatewayZEVM.WithdrawAndCall(
		r.ZEVMAuth,
		receiver.Bytes(),
		amount,
		r.ERC20ZRC20Addr,
		payload,
		gasLimit,
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}

// V2ZEVMToEMVCall calls Call of Gateway on ZEVM
func (r *E2ERunner) V2ZEVMToEMVCall(
	receiver ethcommon.Address,
	payload []byte,
	revertOptions gatewayzevm.RevertOptions,
) *ethtypes.Transaction {
	tx, err := r.GatewayZEVM.Call0(
		r.ZEVMAuth,
		receiver.Bytes(),
		r.ETHZRC20Addr,
		payload,
		gasLimit,
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}

// V2ZEVMToEMVCall calls authenticated Call of Gateway on ZEVM
func (r *E2ERunner) V2ZEVMToEMVAuthenticatedCall(
	receiver ethcommon.Address,
	payload []byte,
	revertOptions gatewayzevm.RevertOptions,
) *ethtypes.Transaction {
	tx, err := r.GatewayZEVM.Call(
		r.ZEVMAuth,
		receiver.Bytes(),
		r.ETHZRC20Addr,
		payload,
		gatewayzevm.CallOptions{
			GasLimit:        gasLimit,
			IsArbitraryCall: false,
		},
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}

// V2ZEVMToEMVCall calls authenticated Call of Gateway on ZEVM through contract
func (r *E2ERunner) V2ZEVMToEMVAuthenticatedCallThroughContract(
	gatewayZEVMCaller *testgatewayzevmcaller.TestGatewayZEVMCaller,
	receiver ethcommon.Address,
	payload []byte,
	revertOptions testgatewayzevmcaller.RevertOptions,
) *ethtypes.Transaction {
	tx, err := gatewayZEVMCaller.CallGatewayZEVM(
		r.ZEVMAuth,
		receiver.Bytes(),
		r.ETHZRC20Addr,
		payload,
		testgatewayzevmcaller.CallOptions{
			GasLimit:        gasLimit,
			IsArbitraryCall: false,
		},
		revertOptions,
	)
	require.NoError(r, err)

	return tx
}
