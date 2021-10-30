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
// the main Ethereum network.
var MainnetBootnodes = []string{
	// FandomChain Bootnodes Mainnet
	"enode://5b1254aa4a7a91d5a95a3aa67aed2a2b8a12e2b686f6f48de2856055a71123a32b666974f2f273cdf28924228ace6dda424e54a7ca2ef061a308869a6f02116b@152.70.246.150:10303",
	"enode://ea3b2ebd55f4c36a69bdc9e689a67d881aad9e30e05d24e97c4d8039a080b7765185d44969206ceedc51ff348b43fa304cb87418eeb9ab4dd5a10b8d91aa0a62@152.70.93.73:10303",
	"enode://68927910ce58e6fef623988c78ff2ecd1b7001ba4f0f2afe8f4fb6180b52c7568b4d97b4cc606b0bb481f186458558dd78340b1b2680813dc509d063b5e7496f@132.226.171.251:10303",
	"enode://a7022180f186a2e6f05c29404e7026bc8c2c92dab62dda13512314846129bac7eeb012ee14b27cb8f7484ff04c84ee0b8183efb233fd776883c1579bbd4c24c8@152.70.237.249:10303",
	"enode://7800410d2383bf3f826be98189ed9df7b3e19cd1005d6f09feca35da19895cdc1defe9d8c01e8a5035e84de1eaff0df88fbce47b41ba58c5d69cd0dd46225c6b@146.56.171.183:10303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	// FandomChain Bootnodes Testnet
	"enode://426d2092178adc391eb9fe99e78e985bf9e2dea03efb0aca4858495ea2f03847a3796ef51b760a3e240d892b274570c78de3871a642000e8783f958213306dbb@140.238.14.234:10303",
	"enode://e5d7dc26402d2d79b330edc720fe07423cbd078216c6446ec3b3c4f099b7f4eacc3027086ccf89aedb8788ec00d05d4c9c566e20c4daaf3f03b159b92c9a8456@140.238.14.234:20303",
	"enode://df1e3afb70f3e99cc64ad63f57c39c5bd478ed9e74a4d96f0468157f2b1e6903052ed9ca3d8de84502812e2cb76e9f472fd9c444924edf5ec7b450a169f32f3a@140.238.14.234:30303",
	"enode://4dbbefa4092c07abbf1204c0992a16ca9aeba4072a353e60d0b0a6cd6b998d757ddfbd6fab26ba2e23c8f99b0d06300d25c143368a51c6f0e6973b155dac6447@140.238.14.234:40303",
	"enode://0411b5250fc60b503f405dfa3188cdf7c468bdbdacd8a48077a37bf167651613a804a8823815f929564f2ee950a322ff332bc0bfdb2a28f4d81309704b375fcc@140.238.14.234:50303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
