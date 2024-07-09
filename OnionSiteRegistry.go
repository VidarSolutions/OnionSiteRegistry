// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OnlineSiteRegistry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OnionSiteRegistrySite is an auto generated low-level Go binding around an user-defined struct.
type OnionSiteRegistrySite struct {
	Title        string
	Description  string
	OnionUrl     string
	PublicPgpKey string
	Category     string
	Owner        common.Address
}

// OnionSiteRegistryMetaData contains all meta data concerning the OnionSiteRegistry contract.
var OnionSiteRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"CategoryApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"CategoryRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"onionUrl\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SiteRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"onionUrl\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SiteUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"adjustNewCategoryRequestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"adjustRegistrationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"adjustUpdateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allSites\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"onionUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicPgpKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"approveCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approvedCategories\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"categoryRequests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSites\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"onionUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicPgpKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structOnionSiteRegistry.Site[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"}],\"name\":\"getSiteByTitle\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"onionUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicPgpKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structOnionSiteRegistry.Site\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_category\",\"type\":\"string\"}],\"name\":\"getSitesByCategory\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"onionUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicPgpKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structOnionSiteRegistry.Site[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"isCategoryApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"isCategoryRequested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newCategoryRequestFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_onionUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_publicPgpKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_category\",\"type\":\"string\"}],\"name\":\"registerSite\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"requestNewCategory\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"sites\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"onionUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publicPgpKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_onionUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_publicPgpKey\",\"type\":\"string\"}],\"name\":\"updateSite\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"50f68c82": "adjustNewCategoryRequestFee(uint256)",
		"9da680f3": "adjustRegistrationFee(uint256)",
		"10470e53": "adjustUpdateFee(uint256)",
		"01d3fc41": "allSites(uint256)",
		"28e629ac": "approveCategory(string)",
		"81d07c77": "approvedCategories(uint256)",
		"73ac4822": "categoryRequests(string)",
		"02c28607": "getAllSites()",
		"5ff68490": "getSiteByTitle(string)",
		"41cc7530": "getSitesByCategory(string)",
		"32bb5b1a": "isCategoryApproved(string)",
		"f71da03c": "isCategoryRequested(string)",
		"06550320": "newCategoryRequestFee()",
		"8da5cb5b": "owner()",
		"319b85ca": "registerSite(string,string,string,string,string)",
		"14c44e09": "registrationFee()",
		"aff5063c": "requestNewCategory(string)",
		"72884128": "sites(string)",
		"758cdcf0": "updateFee()",
		"9b7a07db": "updateSite(string,string,string,string)",
		"2e1a7d4d": "withdraw(uint256)",
	},
	Bin: "0x6080604052670de0b6b3a7640000600555662386f26fc10000600655662386f26fc1000060075534801561003257600080fd5b50600480546001600160a01b0319163317905561259b806100546000396000f3fe60806040526004361061012a5760003560e01c806350f68c82116100ab57806381d07c771161006f57806381d07c771461033d5780638da5cb5b1461036a5780639b7a07db146103a25780639da680f3146103b5578063aff5063c146103d5578063f71da03c146103e857600080fd5b806350f68c821461028b5780635ff68490146102ab57806372884128146102d857806373ac4822146102f8578063758cdcf01461032757600080fd5b806328e629ac116100f257806328e629ac146101e85780632e1a7d4d14610208578063319b85ca1461022857806332bb5b1a1461023b57806341cc75301461026b57600080fd5b806301d3fc411461012f57806302c286071461016a578063065503201461018c57806310470e53146101b057806314c44e09146101d2575b600080fd5b34801561013b57600080fd5b5061014f61014a366004611e61565b610408565b60405161016196959493929190611eca565b60405180910390f35b34801561017657600080fd5b5061017f610705565b6040516101619190611fe2565b34801561019857600080fd5b506101a260075481565b604051908152602001610161565b3480156101bc57600080fd5b506101d06101cb366004611e61565b610a54565b005b3480156101de57600080fd5b506101a260055481565b3480156101f457600080fd5b506101d06102033660046120e7565b610a8c565b34801561021457600080fd5b506101d0610223366004611e61565b610bcd565b6101d0610236366004612124565b610c35565b34801561024757600080fd5b5061025b6102563660046120e7565b610f43565b6040519015158152602001610161565b34801561027757600080fd5b5061017f6102863660046120e7565b610fb4565b34801561029757600080fd5b506101d06102a6366004611e61565b611476565b3480156102b757600080fd5b506102cb6102c63660046120e7565b6114a5565b60405161016191906121f6565b3480156102e457600080fd5b5061014f6102f33660046120e7565b61186e565b34801561030457600080fd5b506103186103133660046120e7565b611894565b60405161016193929190612210565b34801561033357600080fd5b506101a260065481565b34801561034957600080fd5b5061035d610358366004611e61565b611959565b6040516101619190612243565b34801561037657600080fd5b5060045461038a906001600160a01b031681565b6040516001600160a01b039091168152602001610161565b6101d06103b0366004612256565b611a05565b3480156103c157600080fd5b506101d06103d0366004611e61565b611be9565b6101d06103e33660046120e7565b611c18565b3480156103f457600080fd5b5061025b6104033660046120e7565b611de2565b6003818154811061041857600080fd5b906000526020600020906006020160009150905080600001805461043b90612303565b80601f016020809104026020016040519081016040528092919081815260200182805461046790612303565b80156104b45780601f10610489576101008083540402835291602001916104b4565b820191906000526020600020905b81548152906001019060200180831161049757829003601f168201915b5050505050908060010180546104c990612303565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590612303565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b50505050509080600201805461055790612303565b80601f016020809104026020016040519081016040528092919081815260200182805461058390612303565b80156105d05780601f106105a5576101008083540402835291602001916105d0565b820191906000526020600020905b8154815290600101906020018083116105b357829003601f168201915b5050505050908060030180546105e590612303565b80601f016020809104026020016040519081016040528092919081815260200182805461061190612303565b801561065e5780601f106106335761010080835404028352916020019161065e565b820191906000526020600020905b81548152906001019060200180831161064157829003601f168201915b50505050509080600401805461067390612303565b80601f016020809104026020016040519081016040528092919081815260200182805461069f90612303565b80156106ec5780601f106106c1576101008083540402835291602001916106ec565b820191906000526020600020905b8154815290600101906020018083116106cf57829003601f168201915b505050600590930154919250506001600160a01b031686565b60606003805480602002602001604051908101604052809291908181526020016000905b82821015610a4b57838290600052602060002090600602016040518060c001604052908160008201805461075c90612303565b80601f016020809104026020016040519081016040528092919081815260200182805461078890612303565b80156107d55780601f106107aa576101008083540402835291602001916107d5565b820191906000526020600020905b8154815290600101906020018083116107b857829003601f168201915b505050505081526020016001820180546107ee90612303565b80601f016020809104026020016040519081016040528092919081815260200182805461081a90612303565b80156108675780601f1061083c57610100808354040283529160200191610867565b820191906000526020600020905b81548152906001019060200180831161084a57829003601f168201915b5050505050815260200160028201805461088090612303565b80601f01602080910402602001604051908101604052809291908181526020018280546108ac90612303565b80156108f95780601f106108ce576101008083540402835291602001916108f9565b820191906000526020600020905b8154815290600101906020018083116108dc57829003601f168201915b5050505050815260200160038201805461091290612303565b80601f016020809104026020016040519081016040528092919081815260200182805461093e90612303565b801561098b5780601f106109605761010080835404028352916020019161098b565b820191906000526020600020905b81548152906001019060200180831161096e57829003601f168201915b505050505081526020016004820180546109a490612303565b80601f01602080910402602001604051908101604052809291908181526020018280546109d090612303565b8015610a1d5780601f106109f257610100808354040283529160200191610a1d565b820191906000526020600020905b815481529060010190602001808311610a0057829003601f168201915b5050509183525050600591909101546001600160a01b03166020918201529082526001929092019101610729565b50505050905090565b6004546001600160a01b03163314610a875760405162461bcd60e51b8152600401610a7e9061233d565b60405180910390fd5b600655565b6004546001600160a01b03163314610ab65760405162461bcd60e51b8152600401610a7e9061233d565b610abf81611de2565b610b045760405162461bcd60e51b815260206004820152601660248201527510d85d1959dbdc9e481b9bdd081c995c5d595cdd195960521b6044820152606401610a7e565b60018082604051610b159190612387565b9081526040519081900360200190206001018054911515600160a01b0260ff60a01b19909216919091179055610b4a81610f43565b610b8c57600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01610b8a82826123f2565b505b80604051610b9a9190612387565b604051908190038120907f0b856c9432bec0142862f9f6d46c77a08a8da3e664b791fdaa538d36b19c66fe90600090a250565b6004546001600160a01b03163314610bf75760405162461bcd60e51b8152600401610a7e9061233d565b6004546040516001600160a01b039091169082156108fc029083906000818181858888f19350505050158015610c31573d6000803e3d6000fd5b5050565b600554341015610c875760405162461bcd60e51b815260206004820152601d60248201527f496e73756666696369656e7420726567697374726174696f6e206665650000006044820152606401610a7e565b600085604051610c979190612387565b90815260200160405180910390206002018054610cb390612303565b159050610d025760405162461bcd60e51b815260206004820152601c60248201527f4f6e696f6e2055524c20616c72656164792072656769737465726564000000006044820152606401610a7e565b610d0b81610f43565b610d4f5760405162461bcd60e51b815260206004820152601560248201527410d85d1959dbdc9e481b9bdd08185c1c1c9bdd9959605a1b6044820152606401610a7e565b60006040518060c00160405280868152602001858152602001878152602001848152602001838152602001336001600160a01b0316815250905080600087604051610d9a9190612387565b90815260405190819003602001902081518190610db790826123f2565b5060208201516001820190610dcc90826123f2565b5060408201516002820190610de190826123f2565b5060608201516003820190610df690826123f2565b5060808201516004820190610e0b90826123f2565b5060a09190910151600590910180546001600160a01b039092166001600160a01b031990921691909117905560038054600181018255600091909152815182916006027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01908190610e7d90826123f2565b5060208201516001820190610e9290826123f2565b5060408201516002820190610ea790826123f2565b5060608201516003820190610ebc90826123f2565b5060808201516004820190610ed190826123f2565b5060a09190910151600590910180546001600160a01b0319166001600160a01b039092169190911790556040513390610f0b908890612387565b604051908190038120907fc019112aab4e3340de1a9b784c5fc86fef4866e1392f7a63d2436c39eeedf3ba90600090a3505050505050565b6000805b600254811015610fab57828051906020012060028281548110610f6c57610f6c6124b2565b90600052602060002001604051610f8391906124c8565b604051809103902003610f995750600192915050565b80610fa38161253e565b915050610f47565b50600092915050565b60035460609060009067ffffffffffffffff811115610fd557610fd5612044565b60405190808252806020026020018201604052801561100e57816020015b610ffb611e22565b815260200190600190039081610ff35790505b5090506000805b6003548110156113c05784805190602001206003828154811061103a5761103a6124b2565b906000526020600020906006020160040160405161105891906124c8565b6040518091039020036113ae5760038181548110611078576110786124b2565b90600052602060002090600602016040518060c00160405290816000820180546110a190612303565b80601f01602080910402602001604051908101604052809291908181526020018280546110cd90612303565b801561111a5780601f106110ef5761010080835404028352916020019161111a565b820191906000526020600020905b8154815290600101906020018083116110fd57829003601f168201915b5050505050815260200160018201805461113390612303565b80601f016020809104026020016040519081016040528092919081815260200182805461115f90612303565b80156111ac5780601f10611181576101008083540402835291602001916111ac565b820191906000526020600020905b81548152906001019060200180831161118f57829003601f168201915b505050505081526020016002820180546111c590612303565b80601f01602080910402602001604051908101604052809291908181526020018280546111f190612303565b801561123e5780601f106112135761010080835404028352916020019161123e565b820191906000526020600020905b81548152906001019060200180831161122157829003601f168201915b5050505050815260200160038201805461125790612303565b80601f016020809104026020016040519081016040528092919081815260200182805461128390612303565b80156112d05780601f106112a5576101008083540402835291602001916112d0565b820191906000526020600020905b8154815290600101906020018083116112b357829003601f168201915b505050505081526020016004820180546112e990612303565b80601f016020809104026020016040519081016040528092919081815260200182805461131590612303565b80156113625780601f1061133757610100808354040283529160200191611362565b820191906000526020600020905b81548152906001019060200180831161134557829003601f168201915b5050509183525050600591909101546001600160a01b03166020909101528351849084908110611394576113946124b2565b602002602001018190525081806113aa9061253e565b9250505b806113b88161253e565b915050611015565b5060008167ffffffffffffffff8111156113dc576113dc612044565b60405190808252806020026020018201604052801561141557816020015b611402611e22565b8152602001906001900390816113fa5790505b50905060005b8281101561146d57838181518110611435576114356124b2565b602002602001015182828151811061144f5761144f6124b2565b602002602001018190525080806114659061253e565b91505061141b565b50949350505050565b6004546001600160a01b031633146114a05760405162461bcd60e51b8152600401610a7e9061233d565b600755565b6114ad611e22565b60005b600354811015611834578280519060200120600382815481106114d5576114d56124b2565b90600052602060002090600602016000016040516114f391906124c8565b6040518091039020036118225760038181548110611513576115136124b2565b90600052602060002090600602016040518060c001604052908160008201805461153c90612303565b80601f016020809104026020016040519081016040528092919081815260200182805461156890612303565b80156115b55780601f1061158a576101008083540402835291602001916115b5565b820191906000526020600020905b81548152906001019060200180831161159857829003601f168201915b505050505081526020016001820180546115ce90612303565b80601f01602080910402602001604051908101604052809291908181526020018280546115fa90612303565b80156116475780601f1061161c57610100808354040283529160200191611647565b820191906000526020600020905b81548152906001019060200180831161162a57829003601f168201915b5050505050815260200160028201805461166090612303565b80601f016020809104026020016040519081016040528092919081815260200182805461168c90612303565b80156116d95780601f106116ae576101008083540402835291602001916116d9565b820191906000526020600020905b8154815290600101906020018083116116bc57829003601f168201915b505050505081526020016003820180546116f290612303565b80601f016020809104026020016040519081016040528092919081815260200182805461171e90612303565b801561176b5780601f106117405761010080835404028352916020019161176b565b820191906000526020600020905b81548152906001019060200180831161174e57829003601f168201915b5050505050815260200160048201805461178490612303565b80601f01602080910402602001604051908101604052809291908181526020018280546117b090612303565b80156117fd5780601f106117d2576101008083540402835291602001916117fd565b820191906000526020600020905b8154815290600101906020018083116117e057829003601f168201915b5050509183525050600591909101546001600160a01b03166020909101529392505050565b8061182c8161253e565b9150506114b0565b5060405162461bcd60e51b815260206004820152600e60248201526d14da5d19481b9bdd08199bdd5b9960921b6044820152606401610a7e565b805160208183018101805160008252928201919093012091528054819061043b90612303565b80516020818301810180516001825292820191909301209152805481906118ba90612303565b80601f01602080910402602001604051908101604052809291908181526020018280546118e690612303565b80156119335780601f1061190857610100808354040283529160200191611933565b820191906000526020600020905b81548152906001019060200180831161191657829003601f168201915b505050600190930154919250506001600160a01b0381169060ff600160a01b9091041683565b6002818154811061196957600080fd5b90600052602060002001600091509050805461198490612303565b80601f01602080910402602001604051908101604052809291908181526020018280546119b090612303565b80156119fd5780601f106119d2576101008083540402835291602001916119fd565b820191906000526020600020905b8154815290600101906020018083116119e057829003601f168201915b505050505081565b600654341015611a575760405162461bcd60e51b815260206004820152601760248201527f496e73756666696369656e7420757064617465206665650000000000000000006044820152606401610a7e565b60008085604051611a689190612387565b90815260200160405180910390206002018054611a8490612303565b905011611ad35760405162461bcd60e51b815260206004820152601860248201527f4f6e696f6e2055524c206e6f74207265676973746572656400000000000000006044820152606401610a7e565b336001600160a01b0316600085604051611aed9190612387565b908152604051908190036020019020600501546001600160a01b031614611b565760405162461bcd60e51b815260206004820152601c60248201527f43616c6c6572206973206e6f74207468652073697465206f776e6572000000006044820152606401610a7e565b60008085604051611b679190612387565b908152604051908190036020019020905080611b8385826123f2565b5060018101611b9284826123f2565b5060038101611ba183826123f2565b506040513390611bb2908790612387565b604051908190038120907f579c44a201681f238068b663974d7345f78d33477588efaecba28057c975e7ca90600090a35050505050565b6004546001600160a01b03163314611c135760405162461bcd60e51b8152600401610a7e9061233d565b600555565b600754341015611c785760405162461bcd60e51b815260206004820152602560248201527f496e73756666696369656e74206e65772063617465676f727920726571756573604482015264742066656560d81b6064820152608401610a7e565b6000815111611cc95760405162461bcd60e51b815260206004820152601d60248201527f43617465676f7279206e616d652063616e6e6f7420626520656d7074790000006044820152606401610a7e565b611cd281611de2565b15611d1f5760405162461bcd60e51b815260206004820152601a60248201527f43617465676f727920616c7265616479207265717565737465640000000000006044820152606401610a7e565b604080516060810182528281523360208201526000818301529051600190611d48908490612387565b90815260405190819003602001902081518190611d6590826123f2565b506020820151600190910180546040938401511515600160a01b026001600160a81b03199091166001600160a01b0390931692909217919091179055513390611daf908390612387565b604051908190038120907f432f747420677650a70bef34a756e6d420d2236c04b38f0337cedfc0dd8b15b090600090a350565b6000806001600160a01b0316600183604051611dfe9190612387565b908152604051908190036020019020600101546001600160a01b0316141592915050565b6040518060c00160405280606081526020016060815260200160608152602001606081526020016060815260200160006001600160a01b031681525090565b600060208284031215611e7357600080fd5b5035919050565b60005b83811015611e95578181015183820152602001611e7d565b50506000910152565b60008151808452611eb6816020860160208601611e7a565b601f01601f19169290920160200192915050565b60c081526000611edd60c0830189611e9e565b8281036020840152611eef8189611e9e565b90508281036040840152611f038188611e9e565b90508281036060840152611f178187611e9e565b90508281036080840152611f2b8186611e9e565b91505060018060a01b03831660a0830152979650505050505050565b6000815160c08452611f5c60c0850182611e9e565b905060208301518482036020860152611f758282611e9e565b91505060408301518482036040860152611f8f8282611e9e565b91505060608301518482036060860152611fa98282611e9e565b91505060808301518482036080860152611fc38282611e9e565b60a0948501516001600160a01b03169590940194909452509092915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561203757603f19888603018452612025858351611f47565b94509285019290850190600101612009565b5092979650505050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261206b57600080fd5b813567ffffffffffffffff8082111561208657612086612044565b604051601f8301601f19908116603f011681019082821181831017156120ae576120ae612044565b816040528381528660208588010111156120c757600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156120f957600080fd5b813567ffffffffffffffff81111561211057600080fd5b61211c8482850161205a565b949350505050565b600080600080600060a0868803121561213c57600080fd5b853567ffffffffffffffff8082111561215457600080fd5b61216089838a0161205a565b9650602088013591508082111561217657600080fd5b61218289838a0161205a565b9550604088013591508082111561219857600080fd5b6121a489838a0161205a565b945060608801359150808211156121ba57600080fd5b6121c689838a0161205a565b935060808801359150808211156121dc57600080fd5b506121e98882890161205a565b9150509295509295909350565b6020815260006122096020830184611f47565b9392505050565b6060815260006122236060830186611e9e565b6001600160a01b0394909416602083015250901515604090910152919050565b6020815260006122096020830184611e9e565b6000806000806080858703121561226c57600080fd5b843567ffffffffffffffff8082111561228457600080fd5b6122908883890161205a565b955060208701359150808211156122a657600080fd5b6122b28883890161205a565b945060408701359150808211156122c857600080fd5b6122d48883890161205a565b935060608701359150808211156122ea57600080fd5b506122f78782880161205a565b91505092959194509250565b600181811c9082168061231757607f821691505b60208210810361233757634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602a908201527f4f6e6c7920636f6e7472616374206f776e65722063616e2063616c6c207468696040820152693990333ab731ba34b7b760b11b606082015260800190565b60008251612399818460208701611e7a565b9190910192915050565b601f8211156123ed57600081815260208120601f850160051c810160208610156123ca5750805b601f850160051c820191505b818110156123e9578281556001016123d6565b5050505b505050565b815167ffffffffffffffff81111561240c5761240c612044565b6124208161241a8454612303565b846123a3565b602080601f831160018114612455576000841561243d5750858301515b600019600386901b1c1916600185901b1785556123e9565b600085815260208120601f198616915b8281101561248457888601518255948401946001909101908401612465565b50858210156124a25787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60008083546124d681612303565b600182811680156124ee576001811461250357612532565b60ff1984168752821515830287019450612532565b8760005260208060002060005b858110156125295781548a820152908401908201612510565b50505082870194505b50929695505050505050565b60006001820161255e57634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220e7629825e2020c4f49e1b68c2b23e5ef6cfe1e28257000139818acc57f580be464736f6c63430008120033",
}

// OnionSiteRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use OnionSiteRegistryMetaData.ABI instead.
var OnionSiteRegistryABI = OnionSiteRegistryMetaData.ABI

// Deprecated: Use OnionSiteRegistryMetaData.Sigs instead.
// OnionSiteRegistryFuncSigs maps the 4-byte function signature to its string representation.
var OnionSiteRegistryFuncSigs = OnionSiteRegistryMetaData.Sigs

// OnionSiteRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OnionSiteRegistryMetaData.Bin instead.
var OnionSiteRegistryBin = OnionSiteRegistryMetaData.Bin

// DeployOnionSiteRegistry deploys a new Ethereum contract, binding an instance of OnionSiteRegistry to it.
func DeployOnionSiteRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OnionSiteRegistry, error) {
	parsed, err := OnionSiteRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnionSiteRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnionSiteRegistry{OnionSiteRegistryCaller: OnionSiteRegistryCaller{contract: contract}, OnionSiteRegistryTransactor: OnionSiteRegistryTransactor{contract: contract}, OnionSiteRegistryFilterer: OnionSiteRegistryFilterer{contract: contract}}, nil
}

// OnionSiteRegistry is an auto generated Go binding around an Ethereum contract.
type OnionSiteRegistry struct {
	OnionSiteRegistryCaller     // Read-only binding to the contract
	OnionSiteRegistryTransactor // Write-only binding to the contract
	OnionSiteRegistryFilterer   // Log filterer for contract events
}

// OnionSiteRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnionSiteRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnionSiteRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnionSiteRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnionSiteRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnionSiteRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnionSiteRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnionSiteRegistrySession struct {
	Contract     *OnionSiteRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OnionSiteRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnionSiteRegistryCallerSession struct {
	Contract *OnionSiteRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OnionSiteRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnionSiteRegistryTransactorSession struct {
	Contract     *OnionSiteRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OnionSiteRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnionSiteRegistryRaw struct {
	Contract *OnionSiteRegistry // Generic contract binding to access the raw methods on
}

// OnionSiteRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnionSiteRegistryCallerRaw struct {
	Contract *OnionSiteRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// OnionSiteRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnionSiteRegistryTransactorRaw struct {
	Contract *OnionSiteRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnionSiteRegistry creates a new instance of OnionSiteRegistry, bound to a specific deployed contract.
func NewOnionSiteRegistry(address common.Address, backend bind.ContractBackend) (*OnionSiteRegistry, error) {
	contract, err := bindOnionSiteRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnionSiteRegistry{OnionSiteRegistryCaller: OnionSiteRegistryCaller{contract: contract}, OnionSiteRegistryTransactor: OnionSiteRegistryTransactor{contract: contract}, OnionSiteRegistryFilterer: OnionSiteRegistryFilterer{contract: contract}}, nil
}

// NewOnionSiteRegistryCaller creates a new read-only instance of OnionSiteRegistry, bound to a specific deployed contract.
func NewOnionSiteRegistryCaller(address common.Address, caller bind.ContractCaller) (*OnionSiteRegistryCaller, error) {
	contract, err := bindOnionSiteRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnionSiteRegistryCaller{contract: contract}, nil
}

// NewOnionSiteRegistryTransactor creates a new write-only instance of OnionSiteRegistry, bound to a specific deployed contract.
func NewOnionSiteRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*OnionSiteRegistryTransactor, error) {
	contract, err := bindOnionSiteRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnionSiteRegistryTransactor{contract: contract}, nil
}

// NewOnionSiteRegistryFilterer creates a new log filterer instance of OnionSiteRegistry, bound to a specific deployed contract.
func NewOnionSiteRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*OnionSiteRegistryFilterer, error) {
	contract, err := bindOnionSiteRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnionSiteRegistryFilterer{contract: contract}, nil
}

// bindOnionSiteRegistry binds a generic wrapper to an already deployed contract.
func bindOnionSiteRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnionSiteRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnionSiteRegistry *OnionSiteRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnionSiteRegistry.Contract.OnionSiteRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnionSiteRegistry *OnionSiteRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.OnionSiteRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnionSiteRegistry *OnionSiteRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.OnionSiteRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OnionSiteRegistry *OnionSiteRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnionSiteRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OnionSiteRegistry *OnionSiteRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OnionSiteRegistry *OnionSiteRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllSites is a free data retrieval call binding the contract method 0x01d3fc41.
//
// Solidity: function allSites(uint256 ) view returns(string title, string description, string onionUrl, string publicPgpKey, string category, address owner)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) AllSites(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Title        string
	Description  string
	OnionUrl     string
	PublicPgpKey string
	Category     string
	Owner        common.Address
}, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "allSites", arg0)

	outstruct := new(struct {
		Title        string
		Description  string
		OnionUrl     string
		PublicPgpKey string
		Category     string
		Owner        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OnionUrl = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PublicPgpKey = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Category = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AllSites is a free data retrieval call binding the contract method 0x01d3fc41.
//
// Solidity: function allSites(uint256 ) view returns(string title, string description, string onionUrl, string publicPgpKey, string category, address owner)
func (_OnionSiteRegistry *OnionSiteRegistrySession) AllSites(arg0 *big.Int) (struct {
	Title        string
	Description  string
	OnionUrl     string
	PublicPgpKey string
	Category     string
	Owner        common.Address
}, error) {
	return _OnionSiteRegistry.Contract.AllSites(&_OnionSiteRegistry.CallOpts, arg0)
}

// AllSites is a free data retrieval call binding the contract method 0x01d3fc41.
//
// Solidity: function allSites(uint256 ) view returns(string title, string description, string onionUrl, string publicPgpKey, string category, address owner)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) AllSites(arg0 *big.Int) (struct {
	Title        string
	Description  string
	OnionUrl     string
	PublicPgpKey string
	Category     string
	Owner        common.Address
}, error) {
	return _OnionSiteRegistry.Contract.AllSites(&_OnionSiteRegistry.CallOpts, arg0)
}

// ApprovedCategories is a free data retrieval call binding the contract method 0x81d07c77.
//
// Solidity: function approvedCategories(uint256 ) view returns(string)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) ApprovedCategories(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "approvedCategories", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApprovedCategories is a free data retrieval call binding the contract method 0x81d07c77.
//
// Solidity: function approvedCategories(uint256 ) view returns(string)
func (_OnionSiteRegistry *OnionSiteRegistrySession) ApprovedCategories(arg0 *big.Int) (string, error) {
	return _OnionSiteRegistry.Contract.ApprovedCategories(&_OnionSiteRegistry.CallOpts, arg0)
}

// ApprovedCategories is a free data retrieval call binding the contract method 0x81d07c77.
//
// Solidity: function approvedCategories(uint256 ) view returns(string)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) ApprovedCategories(arg0 *big.Int) (string, error) {
	return _OnionSiteRegistry.Contract.ApprovedCategories(&_OnionSiteRegistry.CallOpts, arg0)
}

// CategoryRequests is a free data retrieval call binding the contract method 0x73ac4822.
//
// Solidity: function categoryRequests(string ) view returns(string name, address requester, bool approved)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) CategoryRequests(opts *bind.CallOpts, arg0 string) (struct {
	Name      string
	Requester common.Address
	Approved  bool
}, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "categoryRequests", arg0)

	outstruct := new(struct {
		Name      string
		Requester common.Address
		Approved  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Requester = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Approved = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// CategoryRequests is a free data retrieval call binding the contract method 0x73ac4822.
//
// Solidity: function categoryRequests(string ) view returns(string name, address requester, bool approved)
func (_OnionSiteRegistry *OnionSiteRegistrySession) CategoryRequests(arg0 string) (struct {
	Name      string
	Requester common.Address
	Approved  bool
}, error) {
	return _OnionSiteRegistry.Contract.CategoryRequests(&_OnionSiteRegistry.CallOpts, arg0)
}

// CategoryRequests is a free data retrieval call binding the contract method 0x73ac4822.
//
// Solidity: function categoryRequests(string ) view returns(string name, address requester, bool approved)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) CategoryRequests(arg0 string) (struct {
	Name      string
	Requester common.Address
	Approved  bool
}, error) {
	return _OnionSiteRegistry.Contract.CategoryRequests(&_OnionSiteRegistry.CallOpts, arg0)
}

// GetAllSites is a free data retrieval call binding the contract method 0x02c28607.
//
// Solidity: function getAllSites() view returns((string,string,string,string,string,address)[])
func (_OnionSiteRegistry *OnionSiteRegistryCaller) GetAllSites(opts *bind.CallOpts) ([]OnionSiteRegistrySite, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "getAllSites")

	if err != nil {
		return *new([]OnionSiteRegistrySite), err
	}

	out0 := *abi.ConvertType(out[0], new([]OnionSiteRegistrySite)).(*[]OnionSiteRegistrySite)

	return out0, err

}

// GetAllSites is a free data retrieval call binding the contract method 0x02c28607.
//
// Solidity: function getAllSites() view returns((string,string,string,string,string,address)[])
func (_OnionSiteRegistry *OnionSiteRegistrySession) GetAllSites() ([]OnionSiteRegistrySite, error) {
	return _OnionSiteRegistry.Contract.GetAllSites(&_OnionSiteRegistry.CallOpts)
}

// GetAllSites is a free data retrieval call binding the contract method 0x02c28607.
//
// Solidity: function getAllSites() view returns((string,string,string,string,string,address)[])
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) GetAllSites() ([]OnionSiteRegistrySite, error) {
	return _OnionSiteRegistry.Contract.GetAllSites(&_OnionSiteRegistry.CallOpts)
}

// GetSiteByTitle is a free data retrieval call binding the contract method 0x5ff68490.
//
// Solidity: function getSiteByTitle(string _title) view returns((string,string,string,string,string,address))
func (_OnionSiteRegistry *OnionSiteRegistryCaller) GetSiteByTitle(opts *bind.CallOpts, _title string) (OnionSiteRegistrySite, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "getSiteByTitle", _title)

	if err != nil {
		return *new(OnionSiteRegistrySite), err
	}

	out0 := *abi.ConvertType(out[0], new(OnionSiteRegistrySite)).(*OnionSiteRegistrySite)

	return out0, err

}

// GetSiteByTitle is a free data retrieval call binding the contract method 0x5ff68490.
//
// Solidity: function getSiteByTitle(string _title) view returns((string,string,string,string,string,address))
func (_OnionSiteRegistry *OnionSiteRegistrySession) GetSiteByTitle(_title string) (OnionSiteRegistrySite, error) {
	return _OnionSiteRegistry.Contract.GetSiteByTitle(&_OnionSiteRegistry.CallOpts, _title)
}

// GetSiteByTitle is a free data retrieval call binding the contract method 0x5ff68490.
//
// Solidity: function getSiteByTitle(string _title) view returns((string,string,string,string,string,address))
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) GetSiteByTitle(_title string) (OnionSiteRegistrySite, error) {
	return _OnionSiteRegistry.Contract.GetSiteByTitle(&_OnionSiteRegistry.CallOpts, _title)
}

// GetSitesByCategory is a free data retrieval call binding the contract method 0x41cc7530.
//
// Solidity: function getSitesByCategory(string _category) view returns((string,string,string,string,string,address)[])
func (_OnionSiteRegistry *OnionSiteRegistryCaller) GetSitesByCategory(opts *bind.CallOpts, _category string) ([]OnionSiteRegistrySite, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "getSitesByCategory", _category)

	if err != nil {
		return *new([]OnionSiteRegistrySite), err
	}

	out0 := *abi.ConvertType(out[0], new([]OnionSiteRegistrySite)).(*[]OnionSiteRegistrySite)

	return out0, err

}

// GetSitesByCategory is a free data retrieval call binding the contract method 0x41cc7530.
//
// Solidity: function getSitesByCategory(string _category) view returns((string,string,string,string,string,address)[])
func (_OnionSiteRegistry *OnionSiteRegistrySession) GetSitesByCategory(_category string) ([]OnionSiteRegistrySite, error) {
	return _OnionSiteRegistry.Contract.GetSitesByCategory(&_OnionSiteRegistry.CallOpts, _category)
}

// GetSitesByCategory is a free data retrieval call binding the contract method 0x41cc7530.
//
// Solidity: function getSitesByCategory(string _category) view returns((string,string,string,string,string,address)[])
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) GetSitesByCategory(_category string) ([]OnionSiteRegistrySite, error) {
	return _OnionSiteRegistry.Contract.GetSitesByCategory(&_OnionSiteRegistry.CallOpts, _category)
}

// IsCategoryApproved is a free data retrieval call binding the contract method 0x32bb5b1a.
//
// Solidity: function isCategoryApproved(string _name) view returns(bool)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) IsCategoryApproved(opts *bind.CallOpts, _name string) (bool, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "isCategoryApproved", _name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCategoryApproved is a free data retrieval call binding the contract method 0x32bb5b1a.
//
// Solidity: function isCategoryApproved(string _name) view returns(bool)
func (_OnionSiteRegistry *OnionSiteRegistrySession) IsCategoryApproved(_name string) (bool, error) {
	return _OnionSiteRegistry.Contract.IsCategoryApproved(&_OnionSiteRegistry.CallOpts, _name)
}

// IsCategoryApproved is a free data retrieval call binding the contract method 0x32bb5b1a.
//
// Solidity: function isCategoryApproved(string _name) view returns(bool)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) IsCategoryApproved(_name string) (bool, error) {
	return _OnionSiteRegistry.Contract.IsCategoryApproved(&_OnionSiteRegistry.CallOpts, _name)
}

// IsCategoryRequested is a free data retrieval call binding the contract method 0xf71da03c.
//
// Solidity: function isCategoryRequested(string _name) view returns(bool)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) IsCategoryRequested(opts *bind.CallOpts, _name string) (bool, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "isCategoryRequested", _name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCategoryRequested is a free data retrieval call binding the contract method 0xf71da03c.
//
// Solidity: function isCategoryRequested(string _name) view returns(bool)
func (_OnionSiteRegistry *OnionSiteRegistrySession) IsCategoryRequested(_name string) (bool, error) {
	return _OnionSiteRegistry.Contract.IsCategoryRequested(&_OnionSiteRegistry.CallOpts, _name)
}

// IsCategoryRequested is a free data retrieval call binding the contract method 0xf71da03c.
//
// Solidity: function isCategoryRequested(string _name) view returns(bool)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) IsCategoryRequested(_name string) (bool, error) {
	return _OnionSiteRegistry.Contract.IsCategoryRequested(&_OnionSiteRegistry.CallOpts, _name)
}

// NewCategoryRequestFee is a free data retrieval call binding the contract method 0x06550320.
//
// Solidity: function newCategoryRequestFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) NewCategoryRequestFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "newCategoryRequestFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewCategoryRequestFee is a free data retrieval call binding the contract method 0x06550320.
//
// Solidity: function newCategoryRequestFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistrySession) NewCategoryRequestFee() (*big.Int, error) {
	return _OnionSiteRegistry.Contract.NewCategoryRequestFee(&_OnionSiteRegistry.CallOpts)
}

// NewCategoryRequestFee is a free data retrieval call binding the contract method 0x06550320.
//
// Solidity: function newCategoryRequestFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) NewCategoryRequestFee() (*big.Int, error) {
	return _OnionSiteRegistry.Contract.NewCategoryRequestFee(&_OnionSiteRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OnionSiteRegistry *OnionSiteRegistrySession) Owner() (common.Address, error) {
	return _OnionSiteRegistry.Contract.Owner(&_OnionSiteRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) Owner() (common.Address, error) {
	return _OnionSiteRegistry.Contract.Owner(&_OnionSiteRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) RegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "registrationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistrySession) RegistrationFee() (*big.Int, error) {
	return _OnionSiteRegistry.Contract.RegistrationFee(&_OnionSiteRegistry.CallOpts)
}

// RegistrationFee is a free data retrieval call binding the contract method 0x14c44e09.
//
// Solidity: function registrationFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) RegistrationFee() (*big.Int, error) {
	return _OnionSiteRegistry.Contract.RegistrationFee(&_OnionSiteRegistry.CallOpts)
}

// Sites is a free data retrieval call binding the contract method 0x72884128.
//
// Solidity: function sites(string ) view returns(string title, string description, string onionUrl, string publicPgpKey, string category, address owner)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) Sites(opts *bind.CallOpts, arg0 string) (struct {
	Title        string
	Description  string
	OnionUrl     string
	PublicPgpKey string
	Category     string
	Owner        common.Address
}, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "sites", arg0)

	outstruct := new(struct {
		Title        string
		Description  string
		OnionUrl     string
		PublicPgpKey string
		Category     string
		Owner        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Title = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OnionUrl = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PublicPgpKey = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Category = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Sites is a free data retrieval call binding the contract method 0x72884128.
//
// Solidity: function sites(string ) view returns(string title, string description, string onionUrl, string publicPgpKey, string category, address owner)
func (_OnionSiteRegistry *OnionSiteRegistrySession) Sites(arg0 string) (struct {
	Title        string
	Description  string
	OnionUrl     string
	PublicPgpKey string
	Category     string
	Owner        common.Address
}, error) {
	return _OnionSiteRegistry.Contract.Sites(&_OnionSiteRegistry.CallOpts, arg0)
}

// Sites is a free data retrieval call binding the contract method 0x72884128.
//
// Solidity: function sites(string ) view returns(string title, string description, string onionUrl, string publicPgpKey, string category, address owner)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) Sites(arg0 string) (struct {
	Title        string
	Description  string
	OnionUrl     string
	PublicPgpKey string
	Category     string
	Owner        common.Address
}, error) {
	return _OnionSiteRegistry.Contract.Sites(&_OnionSiteRegistry.CallOpts, arg0)
}

// UpdateFee is a free data retrieval call binding the contract method 0x758cdcf0.
//
// Solidity: function updateFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistryCaller) UpdateFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OnionSiteRegistry.contract.Call(opts, &out, "updateFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateFee is a free data retrieval call binding the contract method 0x758cdcf0.
//
// Solidity: function updateFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistrySession) UpdateFee() (*big.Int, error) {
	return _OnionSiteRegistry.Contract.UpdateFee(&_OnionSiteRegistry.CallOpts)
}

// UpdateFee is a free data retrieval call binding the contract method 0x758cdcf0.
//
// Solidity: function updateFee() view returns(uint256)
func (_OnionSiteRegistry *OnionSiteRegistryCallerSession) UpdateFee() (*big.Int, error) {
	return _OnionSiteRegistry.Contract.UpdateFee(&_OnionSiteRegistry.CallOpts)
}

// AdjustNewCategoryRequestFee is a paid mutator transaction binding the contract method 0x50f68c82.
//
// Solidity: function adjustNewCategoryRequestFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactor) AdjustNewCategoryRequestFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.contract.Transact(opts, "adjustNewCategoryRequestFee", _newFee)
}

// AdjustNewCategoryRequestFee is a paid mutator transaction binding the contract method 0x50f68c82.
//
// Solidity: function adjustNewCategoryRequestFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistrySession) AdjustNewCategoryRequestFee(_newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.AdjustNewCategoryRequestFee(&_OnionSiteRegistry.TransactOpts, _newFee)
}

// AdjustNewCategoryRequestFee is a paid mutator transaction binding the contract method 0x50f68c82.
//
// Solidity: function adjustNewCategoryRequestFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactorSession) AdjustNewCategoryRequestFee(_newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.AdjustNewCategoryRequestFee(&_OnionSiteRegistry.TransactOpts, _newFee)
}

// AdjustRegistrationFee is a paid mutator transaction binding the contract method 0x9da680f3.
//
// Solidity: function adjustRegistrationFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactor) AdjustRegistrationFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.contract.Transact(opts, "adjustRegistrationFee", _newFee)
}

// AdjustRegistrationFee is a paid mutator transaction binding the contract method 0x9da680f3.
//
// Solidity: function adjustRegistrationFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistrySession) AdjustRegistrationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.AdjustRegistrationFee(&_OnionSiteRegistry.TransactOpts, _newFee)
}

// AdjustRegistrationFee is a paid mutator transaction binding the contract method 0x9da680f3.
//
// Solidity: function adjustRegistrationFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactorSession) AdjustRegistrationFee(_newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.AdjustRegistrationFee(&_OnionSiteRegistry.TransactOpts, _newFee)
}

// AdjustUpdateFee is a paid mutator transaction binding the contract method 0x10470e53.
//
// Solidity: function adjustUpdateFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactor) AdjustUpdateFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.contract.Transact(opts, "adjustUpdateFee", _newFee)
}

// AdjustUpdateFee is a paid mutator transaction binding the contract method 0x10470e53.
//
// Solidity: function adjustUpdateFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistrySession) AdjustUpdateFee(_newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.AdjustUpdateFee(&_OnionSiteRegistry.TransactOpts, _newFee)
}

// AdjustUpdateFee is a paid mutator transaction binding the contract method 0x10470e53.
//
// Solidity: function adjustUpdateFee(uint256 _newFee) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactorSession) AdjustUpdateFee(_newFee *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.AdjustUpdateFee(&_OnionSiteRegistry.TransactOpts, _newFee)
}

// ApproveCategory is a paid mutator transaction binding the contract method 0x28e629ac.
//
// Solidity: function approveCategory(string _name) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactor) ApproveCategory(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _OnionSiteRegistry.contract.Transact(opts, "approveCategory", _name)
}

// ApproveCategory is a paid mutator transaction binding the contract method 0x28e629ac.
//
// Solidity: function approveCategory(string _name) returns()
func (_OnionSiteRegistry *OnionSiteRegistrySession) ApproveCategory(_name string) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.ApproveCategory(&_OnionSiteRegistry.TransactOpts, _name)
}

// ApproveCategory is a paid mutator transaction binding the contract method 0x28e629ac.
//
// Solidity: function approveCategory(string _name) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactorSession) ApproveCategory(_name string) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.ApproveCategory(&_OnionSiteRegistry.TransactOpts, _name)
}

// RegisterSite is a paid mutator transaction binding the contract method 0x319b85ca.
//
// Solidity: function registerSite(string _onionUrl, string _title, string _description, string _publicPgpKey, string _category) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactor) RegisterSite(opts *bind.TransactOpts, _onionUrl string, _title string, _description string, _publicPgpKey string, _category string) (*types.Transaction, error) {
	return _OnionSiteRegistry.contract.Transact(opts, "registerSite", _onionUrl, _title, _description, _publicPgpKey, _category)
}

// RegisterSite is a paid mutator transaction binding the contract method 0x319b85ca.
//
// Solidity: function registerSite(string _onionUrl, string _title, string _description, string _publicPgpKey, string _category) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistrySession) RegisterSite(_onionUrl string, _title string, _description string, _publicPgpKey string, _category string) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.RegisterSite(&_OnionSiteRegistry.TransactOpts, _onionUrl, _title, _description, _publicPgpKey, _category)
}

// RegisterSite is a paid mutator transaction binding the contract method 0x319b85ca.
//
// Solidity: function registerSite(string _onionUrl, string _title, string _description, string _publicPgpKey, string _category) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactorSession) RegisterSite(_onionUrl string, _title string, _description string, _publicPgpKey string, _category string) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.RegisterSite(&_OnionSiteRegistry.TransactOpts, _onionUrl, _title, _description, _publicPgpKey, _category)
}

// RequestNewCategory is a paid mutator transaction binding the contract method 0xaff5063c.
//
// Solidity: function requestNewCategory(string _name) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactor) RequestNewCategory(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _OnionSiteRegistry.contract.Transact(opts, "requestNewCategory", _name)
}

// RequestNewCategory is a paid mutator transaction binding the contract method 0xaff5063c.
//
// Solidity: function requestNewCategory(string _name) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistrySession) RequestNewCategory(_name string) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.RequestNewCategory(&_OnionSiteRegistry.TransactOpts, _name)
}

// RequestNewCategory is a paid mutator transaction binding the contract method 0xaff5063c.
//
// Solidity: function requestNewCategory(string _name) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactorSession) RequestNewCategory(_name string) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.RequestNewCategory(&_OnionSiteRegistry.TransactOpts, _name)
}

// UpdateSite is a paid mutator transaction binding the contract method 0x9b7a07db.
//
// Solidity: function updateSite(string _onionUrl, string _title, string _description, string _publicPgpKey) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactor) UpdateSite(opts *bind.TransactOpts, _onionUrl string, _title string, _description string, _publicPgpKey string) (*types.Transaction, error) {
	return _OnionSiteRegistry.contract.Transact(opts, "updateSite", _onionUrl, _title, _description, _publicPgpKey)
}

// UpdateSite is a paid mutator transaction binding the contract method 0x9b7a07db.
//
// Solidity: function updateSite(string _onionUrl, string _title, string _description, string _publicPgpKey) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistrySession) UpdateSite(_onionUrl string, _title string, _description string, _publicPgpKey string) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.UpdateSite(&_OnionSiteRegistry.TransactOpts, _onionUrl, _title, _description, _publicPgpKey)
}

// UpdateSite is a paid mutator transaction binding the contract method 0x9b7a07db.
//
// Solidity: function updateSite(string _onionUrl, string _title, string _description, string _publicPgpKey) payable returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactorSession) UpdateSite(_onionUrl string, _title string, _description string, _publicPgpKey string) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.UpdateSite(&_OnionSiteRegistry.TransactOpts, _onionUrl, _title, _description, _publicPgpKey)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OnionSiteRegistry *OnionSiteRegistrySession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.Withdraw(&_OnionSiteRegistry.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OnionSiteRegistry *OnionSiteRegistryTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _OnionSiteRegistry.Contract.Withdraw(&_OnionSiteRegistry.TransactOpts, amount)
}

// OnionSiteRegistryCategoryApprovedIterator is returned from FilterCategoryApproved and is used to iterate over the raw logs and unpacked data for CategoryApproved events raised by the OnionSiteRegistry contract.
type OnionSiteRegistryCategoryApprovedIterator struct {
	Event *OnionSiteRegistryCategoryApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnionSiteRegistryCategoryApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnionSiteRegistryCategoryApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnionSiteRegistryCategoryApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnionSiteRegistryCategoryApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnionSiteRegistryCategoryApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnionSiteRegistryCategoryApproved represents a CategoryApproved event raised by the OnionSiteRegistry contract.
type OnionSiteRegistryCategoryApproved struct {
	Name common.Hash
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCategoryApproved is a free log retrieval operation binding the contract event 0x0b856c9432bec0142862f9f6d46c77a08a8da3e664b791fdaa538d36b19c66fe.
//
// Solidity: event CategoryApproved(string indexed name)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) FilterCategoryApproved(opts *bind.FilterOpts, name []string) (*OnionSiteRegistryCategoryApprovedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _OnionSiteRegistry.contract.FilterLogs(opts, "CategoryApproved", nameRule)
	if err != nil {
		return nil, err
	}
	return &OnionSiteRegistryCategoryApprovedIterator{contract: _OnionSiteRegistry.contract, event: "CategoryApproved", logs: logs, sub: sub}, nil
}

// WatchCategoryApproved is a free log subscription operation binding the contract event 0x0b856c9432bec0142862f9f6d46c77a08a8da3e664b791fdaa538d36b19c66fe.
//
// Solidity: event CategoryApproved(string indexed name)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) WatchCategoryApproved(opts *bind.WatchOpts, sink chan<- *OnionSiteRegistryCategoryApproved, name []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _OnionSiteRegistry.contract.WatchLogs(opts, "CategoryApproved", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnionSiteRegistryCategoryApproved)
				if err := _OnionSiteRegistry.contract.UnpackLog(event, "CategoryApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCategoryApproved is a log parse operation binding the contract event 0x0b856c9432bec0142862f9f6d46c77a08a8da3e664b791fdaa538d36b19c66fe.
//
// Solidity: event CategoryApproved(string indexed name)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) ParseCategoryApproved(log types.Log) (*OnionSiteRegistryCategoryApproved, error) {
	event := new(OnionSiteRegistryCategoryApproved)
	if err := _OnionSiteRegistry.contract.UnpackLog(event, "CategoryApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnionSiteRegistryCategoryRequestedIterator is returned from FilterCategoryRequested and is used to iterate over the raw logs and unpacked data for CategoryRequested events raised by the OnionSiteRegistry contract.
type OnionSiteRegistryCategoryRequestedIterator struct {
	Event *OnionSiteRegistryCategoryRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnionSiteRegistryCategoryRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnionSiteRegistryCategoryRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnionSiteRegistryCategoryRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnionSiteRegistryCategoryRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnionSiteRegistryCategoryRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnionSiteRegistryCategoryRequested represents a CategoryRequested event raised by the OnionSiteRegistry contract.
type OnionSiteRegistryCategoryRequested struct {
	Name      common.Hash
	Requester common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCategoryRequested is a free log retrieval operation binding the contract event 0x432f747420677650a70bef34a756e6d420d2236c04b38f0337cedfc0dd8b15b0.
//
// Solidity: event CategoryRequested(string indexed name, address indexed requester)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) FilterCategoryRequested(opts *bind.FilterOpts, name []string, requester []common.Address) (*OnionSiteRegistryCategoryRequestedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _OnionSiteRegistry.contract.FilterLogs(opts, "CategoryRequested", nameRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &OnionSiteRegistryCategoryRequestedIterator{contract: _OnionSiteRegistry.contract, event: "CategoryRequested", logs: logs, sub: sub}, nil
}

// WatchCategoryRequested is a free log subscription operation binding the contract event 0x432f747420677650a70bef34a756e6d420d2236c04b38f0337cedfc0dd8b15b0.
//
// Solidity: event CategoryRequested(string indexed name, address indexed requester)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) WatchCategoryRequested(opts *bind.WatchOpts, sink chan<- *OnionSiteRegistryCategoryRequested, name []string, requester []common.Address) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _OnionSiteRegistry.contract.WatchLogs(opts, "CategoryRequested", nameRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnionSiteRegistryCategoryRequested)
				if err := _OnionSiteRegistry.contract.UnpackLog(event, "CategoryRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCategoryRequested is a log parse operation binding the contract event 0x432f747420677650a70bef34a756e6d420d2236c04b38f0337cedfc0dd8b15b0.
//
// Solidity: event CategoryRequested(string indexed name, address indexed requester)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) ParseCategoryRequested(log types.Log) (*OnionSiteRegistryCategoryRequested, error) {
	event := new(OnionSiteRegistryCategoryRequested)
	if err := _OnionSiteRegistry.contract.UnpackLog(event, "CategoryRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnionSiteRegistrySiteRegisteredIterator is returned from FilterSiteRegistered and is used to iterate over the raw logs and unpacked data for SiteRegistered events raised by the OnionSiteRegistry contract.
type OnionSiteRegistrySiteRegisteredIterator struct {
	Event *OnionSiteRegistrySiteRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnionSiteRegistrySiteRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnionSiteRegistrySiteRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnionSiteRegistrySiteRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnionSiteRegistrySiteRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnionSiteRegistrySiteRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnionSiteRegistrySiteRegistered represents a SiteRegistered event raised by the OnionSiteRegistry contract.
type OnionSiteRegistrySiteRegistered struct {
	OnionUrl common.Hash
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSiteRegistered is a free log retrieval operation binding the contract event 0xc019112aab4e3340de1a9b784c5fc86fef4866e1392f7a63d2436c39eeedf3ba.
//
// Solidity: event SiteRegistered(string indexed onionUrl, address indexed owner)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) FilterSiteRegistered(opts *bind.FilterOpts, onionUrl []string, owner []common.Address) (*OnionSiteRegistrySiteRegisteredIterator, error) {

	var onionUrlRule []interface{}
	for _, onionUrlItem := range onionUrl {
		onionUrlRule = append(onionUrlRule, onionUrlItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _OnionSiteRegistry.contract.FilterLogs(opts, "SiteRegistered", onionUrlRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &OnionSiteRegistrySiteRegisteredIterator{contract: _OnionSiteRegistry.contract, event: "SiteRegistered", logs: logs, sub: sub}, nil
}

// WatchSiteRegistered is a free log subscription operation binding the contract event 0xc019112aab4e3340de1a9b784c5fc86fef4866e1392f7a63d2436c39eeedf3ba.
//
// Solidity: event SiteRegistered(string indexed onionUrl, address indexed owner)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) WatchSiteRegistered(opts *bind.WatchOpts, sink chan<- *OnionSiteRegistrySiteRegistered, onionUrl []string, owner []common.Address) (event.Subscription, error) {

	var onionUrlRule []interface{}
	for _, onionUrlItem := range onionUrl {
		onionUrlRule = append(onionUrlRule, onionUrlItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _OnionSiteRegistry.contract.WatchLogs(opts, "SiteRegistered", onionUrlRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnionSiteRegistrySiteRegistered)
				if err := _OnionSiteRegistry.contract.UnpackLog(event, "SiteRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSiteRegistered is a log parse operation binding the contract event 0xc019112aab4e3340de1a9b784c5fc86fef4866e1392f7a63d2436c39eeedf3ba.
//
// Solidity: event SiteRegistered(string indexed onionUrl, address indexed owner)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) ParseSiteRegistered(log types.Log) (*OnionSiteRegistrySiteRegistered, error) {
	event := new(OnionSiteRegistrySiteRegistered)
	if err := _OnionSiteRegistry.contract.UnpackLog(event, "SiteRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnionSiteRegistrySiteUpdatedIterator is returned from FilterSiteUpdated and is used to iterate over the raw logs and unpacked data for SiteUpdated events raised by the OnionSiteRegistry contract.
type OnionSiteRegistrySiteUpdatedIterator struct {
	Event *OnionSiteRegistrySiteUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnionSiteRegistrySiteUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnionSiteRegistrySiteUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnionSiteRegistrySiteUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnionSiteRegistrySiteUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnionSiteRegistrySiteUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnionSiteRegistrySiteUpdated represents a SiteUpdated event raised by the OnionSiteRegistry contract.
type OnionSiteRegistrySiteUpdated struct {
	OnionUrl common.Hash
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSiteUpdated is a free log retrieval operation binding the contract event 0x579c44a201681f238068b663974d7345f78d33477588efaecba28057c975e7ca.
//
// Solidity: event SiteUpdated(string indexed onionUrl, address indexed owner)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) FilterSiteUpdated(opts *bind.FilterOpts, onionUrl []string, owner []common.Address) (*OnionSiteRegistrySiteUpdatedIterator, error) {

	var onionUrlRule []interface{}
	for _, onionUrlItem := range onionUrl {
		onionUrlRule = append(onionUrlRule, onionUrlItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _OnionSiteRegistry.contract.FilterLogs(opts, "SiteUpdated", onionUrlRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &OnionSiteRegistrySiteUpdatedIterator{contract: _OnionSiteRegistry.contract, event: "SiteUpdated", logs: logs, sub: sub}, nil
}

// WatchSiteUpdated is a free log subscription operation binding the contract event 0x579c44a201681f238068b663974d7345f78d33477588efaecba28057c975e7ca.
//
// Solidity: event SiteUpdated(string indexed onionUrl, address indexed owner)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) WatchSiteUpdated(opts *bind.WatchOpts, sink chan<- *OnionSiteRegistrySiteUpdated, onionUrl []string, owner []common.Address) (event.Subscription, error) {

	var onionUrlRule []interface{}
	for _, onionUrlItem := range onionUrl {
		onionUrlRule = append(onionUrlRule, onionUrlItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _OnionSiteRegistry.contract.WatchLogs(opts, "SiteUpdated", onionUrlRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnionSiteRegistrySiteUpdated)
				if err := _OnionSiteRegistry.contract.UnpackLog(event, "SiteUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSiteUpdated is a log parse operation binding the contract event 0x579c44a201681f238068b663974d7345f78d33477588efaecba28057c975e7ca.
//
// Solidity: event SiteUpdated(string indexed onionUrl, address indexed owner)
func (_OnionSiteRegistry *OnionSiteRegistryFilterer) ParseSiteUpdated(log types.Log) (*OnionSiteRegistrySiteUpdated, error) {
	event := new(OnionSiteRegistrySiteUpdated)
	if err := _OnionSiteRegistry.contract.UnpackLog(event, "SiteUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
