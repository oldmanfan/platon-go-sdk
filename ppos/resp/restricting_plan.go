package resp

import "math/big"

type RestrictingPlan struct {

	/**
	 * 表示结算周期的倍数。与每个结算周期出块数的乘积表示在目标区块高度上释放锁定的资金。如果 account 是激励池地址，
	 * 那么 period 值是 120（即，30*4） 的倍数。另外，period * 每周期的区块数至少要大于最高不可逆区块高度
	 */
	Epoch *big.Int `json:"Epoch"`

	/**
	 * 表示目标区块上待释放的金额
	 */
	Amount *big.Int `json:"Amount"`
}

type RestrictingInfo struct {

	/**
	 * 释放区块高度
	 */
	BlockNumber *big.Int
	/**
	 * 释放金额
	 */
	Amount *big.Int
}

type RestrictingItem struct {
	/**
	 * 锁仓余额
	 */
	Balance *big.Int
	/**
	 * 质押/抵押金额
	 */
	Pledge *big.Int
	/**
	 * 欠释放金额
	 */
	Debt *big.Int
	/**
	 * 锁仓分录信息
	 */
	Info []RestrictingInfo `json:"plans"`
}
