// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main PlatON network.
var MainnetBootnodes = []string{
	"enode://81dd24640878badc06de82a82fdb0fe55f24d27877144261d81b1f39faa6686ac8f5d2489dbf97cd44d583b4b00976a8f92845378084d25c7a8bae671a543983@ms1.bfa6.platon.network:16789",
	"enode://32d628cfd32d3f464666792f4fa0bf097c723045f8fe415a8015f1c3cbd0a1bba23e7c76defac277967101dc73e3bd8fc255febb7a52d77c1018ed0cbf8d3ad4@ms2.6cc3.platon.network:16789",
	"enode://3b2ca03c94a2b8f36b88983d8666947dd08e15347980f95c395b36a4e69218c902894e9e2e92c5a2e0fe8b5c137732d2df40a118766245fdac88c480eb120c18@ms3.cd41.platon.network:16789",
	"enode://7b5323a73e9cbffd1e6d9178f0b1d55e92649aa71ebe55a0a9c577d374a9ae21ee4980aef2a3214b6e16aa9928ee48df65a382bd2d7ec19f7b87e6d993654d17@ms4.1fda.platon.network:16789",
	"enode://e6123b585a8e030b42d873d7d09b68847d1f3bba86fab84490fc29acf332a94682a8f8e1518ca857fc75391d62eaf2117703dfeed386b4e0926bf017b5cae445@ms5.ee7a.platon.network:16789",
	"enode://ab2f7cdf347d4ca26f4fdf5657d7b669464c5712cddc42609ad2060691226187815f0ce87f4dca2cac3ee618d4beeeba9618dbd31c54f97af21d16b7cbf0dccd@ms6.63a8.platon.network:16789",
	"enode://4c5a092156c43d5aa3dc71f9dc11d304d7631d393725b09e574577c583759e58ddc245e38f993cc0f32fe873cc782bfc1c62fbd49097eec3278b240de785800b@ms7.66dc.platon.network:16789",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://3fec5e5982a0b32a25168dae575c4705ab8509f266947cb8b16b62ac9eafb78d3e7efce2c31bac447edce3446a12b71383a41dcbdbe80fa856d8739b0214ff35@ts1.7757.platon.network:16789",
	"enode://da6e4c410df2128a79e13179e744b5d28a704c2ea8d729bf63aa3c6f75afed24ef01385380560a5d607f63bda23a9a1e6bf2848e7793f3ea9b59e12f35e662a1@ts2.c72f.platon.network:16789",
	"enode://bf4c161ec0d0854bc4161a4831f15312f9115bdeb0e262022d24c3f17ca06f693f0e36c09412d165155232b11c2d77b5e5c58a805fb4d4533f648b1babd51040@ts3.0464.platon.network:16789",
	"enode://9f3ebc1ddf763e47626d7a13d4f501fcd0d83db8cc4fa26b7a2658ad41e76d29ffa8bb6354ff71d5e22d079a570a07333a3312a59350a75b625850e5819fcbd0@ts4.31e0.platon.network:16789",
	"enode://18820d5df0684e7871643351573f095d99c5c6ecb26021627c60ff55c6d146f5b9984d43cc55f9efaaf043b40628a1c2a770e7e891878e1dc168a3acb7e05dc6@ts5.5eb0.platon.network:16789",
	"enode://991f577221be7a18800ab8ae8eb3214ef3e1fa0be4a161c9cd95100b56e95df677aa12dc624aaedc073c5b48367df0351a4955e91290f19fba8c37413cf5b5e0@ts6.8847.platon.network:16789",
	"enode://86fbd54d03786d924ee33d1fe15542c4ceb528a058c955e42327f0c528db7b2494aecc814b361b4f501950f46aca4c2926e255e6542db991546e0540d24a9d35@ts7.72f5.platon.network:16789",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}

// AlayanetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var AlayanetBootnodes = []string{
	"enode://6928c02520ff4e1784d49b4987eee9852dd7b1552f89836292de1002869a78697a9e697d96001e079b5c662135fd5234ae31b3a94e53283d56caf7652f6d6e90@seed1.a72b.alaya.network:16789",
	"enode://3fdefc6e19e46cb05700b947ff8261087706697fe6054ddf925a261af2780084252bf7b2f6cf652e1fec4d64d3e9539ca9408a1b5ead5fe82b555d95cf143fb2@seed2.afc7.alaya.network:16789",
	"enode://3fe92730eb9b53a2e58a9be11a1707c346432ee0c531c24a22d4bb2d0d4a9b4ef04e23988b4fa5d91a790c7f821e9983ae71b03903a3d75bfcce156b060cf99b@seed3.ccf5.alaya.network:16789",
	"enode://1972a5a7d75010e199eac231ab513e564cad5f0e88331a53606b7d55220803c1816d3b0d06ca9b0e10389264f4fade77c46814dd44df502599d3f0a286160498@seed4.5e92.alaya.network:16789",
	"enode://02dc695641f5cada2c685e3bf3dca0218a9dc7a5d5ce8165a2f5bee40d002d18ec6d899abaac1472d88b71e49691019766abd177b8d5d94f72f6f6dc842fded2@seed5.10a1.alaya.network:16789",
	"enode://49648f184dab8acf0927238452e1f7e8f0e86135dcd148baa0a2d22cd931ed9770f726fccf13a3976bbdb738a880a074dc21d811037190267cd9c8c1378a6043@seed6.8d4b.alaya.network:16789",
	"enode://5670e1b34fe39da46ebdd9c3377053c5214c7b3e6b371d31fcc381f788414a38d44cf844ad71305eb1d0d8afddee8eccafb4d30b33b54ca002db47c4864ba080@seed7.bdac.alaya.network:16789",
}

// AlayaTestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var AlayaTestnetBootnodes = []string{}
