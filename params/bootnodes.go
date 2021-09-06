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
	"enode://bd26321c6a8306a3e167220ad506a0ffa19f9e720f6d73ede120b313d56b092c85daec055eb1dc48ff2aff47a87b1c5a21e9a73f707af468c70595456f534463@140.238.26.150:30303",
	"enode://e0b20132d2cd37b0aa1b986ae7621ee8fb71863b7e636f568bad7a9329e626f4ce6677ebf3f7c30e9cc13a7623434ea123d86fbc025cacca6baac015591a9ed7@140.238.9.216:30303",
	"enode://4067d8875a0684a0bdc786c66c8a592627d15491455b57f8342298b50654e00e30303ffa6d057dbcf76ebe114f592771dc2e4fd064dcb6e06274c3490564d910@140.238.2.213:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	// FandomChain Bootnodes Testnet
	"enode://846ba8638e7433d23293a3e9a8f6474a7ad2b1515733d3a8c96073314351f195466de2b757a0c4f4dbc041b6898fffad5ad79fdd5a3659f170c70785f69c875b@146.56.179.153:10303",
	"enode://8de51d4bc9e3331d608bfb2e4ee21123f76a107805d13a7128a1e8ea2dd839b80945083529ac96dac4222c90b83b0ba77054822de13ac237413b8bb6b072144e@146.56.179.153:20303",
	"enode://7776e27a3c77a6ddd8feecbdf40756242c97999d94e2cf7041e67883eda74b66e42c1cf89ef2d2356683f899312f7959ad0b0f60bdee4a33d648c1ff6f849f51@146.56.179.153:30303",
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
