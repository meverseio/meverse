package app

import (
	"fmt"

	"github.com/fletaio/fleta_v2/common"
	"github.com/fletaio/fleta_v2/common/amount"
	"github.com/fletaio/fleta_v2/common/bin"
	"github.com/fletaio/fleta_v2/contract/formulator"
	"github.com/fletaio/fleta_v2/contract/gateway"
	"github.com/fletaio/fleta_v2/contract/token"
	"github.com/fletaio/fleta_v2/core/types"
)

func Genesis() *types.ContextData {
	adminAddress := common.HexToAddress("0x24D5da1B198c5049016c513099916498ceE9ccf5")
	generators := []common.Address{
		common.HexToAddress("0x7abD922630E41765d6674784C65D794015c1676B"),
		common.HexToAddress("0xDDB9c2B2198daA32f70A4E50B817C0f4bb0BA6b2"),
		common.HexToAddress("0x3E229C2165f3ce8fA3712a779c366603753Fcd1C"),
		common.HexToAddress("0x8cC1b16aAd2d9baAefd03110b5E76B89E66CF8cC"),
		common.HexToAddress("0x6aDed46ff1dfb0263C2b8c8ec6618Ebec19682CD"),
		common.HexToAddress("0x0CBb5713E066c2c432A588Fb1C37BB907E99cd7B"),
		common.HexToAddress("0x81779750Fc6EdCbeBd7c0CCE1F4DEAabb8774A7e"),
		common.HexToAddress("0x60F53D40C32ec5f09A5e12f4F2e6d464bE9F91ba"),
		common.HexToAddress("0x3c6E849cA24b0c8A62697e1870Eb74F5190735b5"),
		common.HexToAddress("0x36104d32d97DE4e11E3590A1c0bDfb715521F256"),
	}
	supplyAddress := common.HexToAddress("0x24D5da1B198c5049016c513099916498ceE9ccf5")

	alphaOwners := genesisAlphas()

	sigmaOwners := genesisSigmas()

	omegaOwners := genesisOmegas()

	stakingMap := map[common.Address]map[common.Address]*amount.Amount{}

	ClassMap := map[string]uint64{}
	if true {
		ClassID, err := types.RegisterContractType(&token.TokenContract{})
		if err != nil {
			panic(err)
		}
		ClassMap["Token"] = ClassID
	}
	if true {
		ClassID, err := types.RegisterContractType(&formulator.FormulatorContract{})
		if err != nil {
			panic(err)
		}
		ClassMap["Formulator"] = ClassID
	}
	if true {
		ClassID, err := types.RegisterContractType(&gateway.GatewayContract{})
		if err != nil {
			panic(err)
		}
		ClassMap["Gateway"] = ClassID
	}

	genesis := types.NewEmptyContext()
	var tokenAddress common.Address
	var formulatorAddress common.Address
	var gatewayAddress common.Address
	if true {
		if err := genesis.SetAdmin(adminAddress, true); err != nil {
			panic(err)
		}
		for _, v := range generators {
			if err := genesis.SetGenerator(v, true); err != nil {
				panic(err)
			}
		}
	}
	if true {
		arg := &token.TokenContractConstruction{
			Name:   "MEVerse",
			Symbol: "MEV",
			InitialSupplyMap: map[common.Address]*amount.Amount{
				supplyAddress: amount.MustParseAmount("1706092169.9918"),
				common.HexToAddress("0x01133FB825b385f6158A598Ca5136f4385fDd722"): amount.MustParseAmount("32754.741713057648982596"),
				common.HexToAddress("0x016c02F3A7527427a0ab236Bb1F3c31B7D9A2A8e"): amount.MustParseAmount("66505.275860816795243437"),
				common.HexToAddress("0x03e8F84C3C455B11518c9a357A3De4161ce078c2"): amount.MustParseAmount("23109.97148585"),
				common.HexToAddress("0x049C28F0655BEeCef1551EE322492b2325C53206"): amount.MustParseAmount("12945.70036099366912"),
				common.HexToAddress("0x04A5A1d9bEdf099a4Ca17414376561Fd20270Df0"): amount.MustParseAmount("99.203437408273234111"),
				common.HexToAddress("0x04b939CB4dB94B1d086269E7F46328F43D5BA7b5"): amount.MustParseAmount("1605.0472299988372"),
				common.HexToAddress("0x04d72a36318F89C84a68291AC881Dd87a59667Cf"): amount.MustParseAmount("181167.33396430356182205"),
				common.HexToAddress("0x055b08F4A7F48664ED5522B39429b7AcE00f62F5"): amount.MustParseAmount("29.2"),
				common.HexToAddress("0x05F32570781f2f28A4783C2bf29BDE1805154BfD"): amount.MustParseAmount("180"),
				common.HexToAddress("0x06E4d94e171A89B38AAf0cf096087a5342453060"): amount.MustParseAmount("65447.536867295239746109"),
				common.HexToAddress("0x0730973507cFE6E1EBDc371da28994D8e1Ff71A7"): amount.MustParseAmount("5691.56354583"),
				common.HexToAddress("0x083f01f1bDfEEdA6F65696FF378b4837f2c64897"): amount.MustParseAmount("33224.9409066116772"),
				common.HexToAddress("0x08dc123181ad3144e8e8f1884C67c8FEcc76dD68"): amount.MustParseAmount("55"),
				common.HexToAddress("0x092bCe1574E6df6f012BAd8413dB1B2C614E6303"): amount.MustParseAmount("2893.010806997364448"),
				common.HexToAddress("0x0B48FBa2aAf5622135BfB1Ab64A8F4A3e2CC9856"): amount.MustParseAmount("1000"),
				common.HexToAddress("0x0D7873a3Af7d78948aB5144f6b88854a742F2F14"): amount.MustParseAmount("110548.288129394720019948"),
				common.HexToAddress("0x0Ea063836e25C921bb782e3ad17BcD91fEa6B862"): amount.MustParseAmount("65268.912567072132"),
				common.HexToAddress("0x0a19415170b95A5dd349c8918F13b2ba16CCd3dd"): amount.MustParseAmount("3486.713691186065784"),
				common.HexToAddress("0x0cB80a1542007B5E3ebB9Fbd18C542dF88D76071"): amount.MustParseAmount("183.285085734249115473"),
				common.HexToAddress("0x0d61A733dDfaBb66e877f3C28486A0a863808cF2"): amount.MustParseAmount("1667.000931321050763911"),
				common.HexToAddress("0x0fB16e7674284858d4839ce836906e71ff05e1Dc"): amount.MustParseAmount("800000"),
				common.HexToAddress("0x1085DBd9835ee92151ec62CB5E2F86AeA1CE8ed7"): amount.MustParseAmount("700"),
				common.HexToAddress("0x10B0838a963b8304541e7eBFAcF7e385f130a873"): amount.MustParseAmount("5389.3"),
				common.HexToAddress("0x10dCBCF98a5C7CddFe68aC997910E8f6377c39e6"): amount.MustParseAmount("999.7"),
				common.HexToAddress("0x11602FD93Ef73d2fDA09d4D40493D8976DEC5dEa"): amount.MustParseAmount("1414.1258"),
				common.HexToAddress("0x11D840061bD904933B5A48555fcAC8d61D4827BD"): amount.MustParseAmount("3064.018389237314379686"),
				common.HexToAddress("0x13BF10E29EA90C2eb74fEb8E274C732Da97e53B8"): amount.MustParseAmount("69.999999999996617263"),
				common.HexToAddress("0x148e261C298836e9e03003FfbeD04A092463D565"): amount.MustParseAmount("464.4821565595276"),
				common.HexToAddress("0x154726CDdb9cD868c7d1E7B6C2c7b145ac281c8B"): amount.MustParseAmount("50654.710554499351718174"),
				common.HexToAddress("0x15669B764ffD45eDaFa79a85C09A1B021E02eDEE"): amount.MustParseAmount("33592.662449717270778893"),
				common.HexToAddress("0x15794c1E7661027567Ea95FD7AbDEbBBaCb26723"): amount.MustParseAmount("71183.430813897397907443"),
				common.HexToAddress("0x164289FC78d6D7500C8E2FA559f254FCC3Cbd1d7"): amount.MustParseAmount("1008.125881740637020824"),
				common.HexToAddress("0x165758bc476ef4a22CFebBC3ca53624d57FFe564"): amount.MustParseAmount("165226.022063317517072"),
				common.HexToAddress("0x1684C75BA970BbfF7b67Ce21a05835A023839eDD"): amount.MustParseAmount("203054.240675854537514429"),
				common.HexToAddress("0x16faa4E24A0dBE547C11eD4Db8671821AB37Bb64"): amount.MustParseAmount("3.70653666228272"),
				common.HexToAddress("0x177aA85B29eC87Fa6a072feE3f8721EfddeDa21F"): amount.MustParseAmount("226.933424803525661608"),
				common.HexToAddress("0x1795aF6Fc412e2977bb20De79380d54Fa3a42513"): amount.MustParseAmount("319.1588"),
				common.HexToAddress("0x19e3E749b40D96F78A53d4F3A1FBDad04384a2eb"): amount.MustParseAmount("971.4"),
				common.HexToAddress("0x1A07D4f64d046CD6bDee86B9217B1E6CBeF7CbA0"): amount.MustParseAmount("9070.06754196464668"),
				common.HexToAddress("0x1A0c935741af4F27b3eBd0364Fa4DA35cbDF12F0"): amount.MustParseAmount("3928670.451215677772"),
				common.HexToAddress("0x1F402ed33A5cf8545838d97523DAfDE55C50af79"): amount.MustParseAmount("101356.782958933132435808"),
				common.HexToAddress("0x1c521b991888F3bA2c5B9b0F6bEd0da87dBD54E6"): amount.MustParseAmount("100"),
				common.HexToAddress("0x1d90392F0591e0188b410fDfD3fF4e4e17ff7982"): amount.MustParseAmount("1000"),
				common.HexToAddress("0x1d958843d039B0c237E315F4EdC4EFf1D01Eb718"): amount.MustParseAmount("23279.85852751957483"),
				common.HexToAddress("0x1dD1e68ee4DF49E5a541737A44D87F0e1d244fdf"): amount.MustParseAmount("24030.764691217658944"),
				common.HexToAddress("0x1da830ae4f2F66554771f117F63C9fE5c3BEcB5B"): amount.MustParseAmount("200"),
				common.HexToAddress("0x200CEc6944b02ABa6CFDdb714EAD5Bf18da19a6e"): amount.MustParseAmount("4867.548903139418857115"),
				common.HexToAddress("0x20Ae2A4f7ea047251473FD774bc0CafcF4D1d742"): amount.MustParseAmount("3000"),
				common.HexToAddress("0x20C336A50939293286805A4d1eDA31dd3964D9Ff"): amount.MustParseAmount("2619083.500810451828"),
				common.HexToAddress("0x23485D0257a64A35ED2a41760e77b2f30f407497"): amount.MustParseAmount("20"),
				common.HexToAddress("0x236Dbcf2d888819B34b083e966D1781583104c46"): amount.MustParseAmount("53266.90149615"),
				common.HexToAddress("0x25F05dd5469876537eB5d3C566202D740BbDf8e5"): amount.MustParseAmount("0.499999999999999995"),
				common.HexToAddress("0x270bD8Ca688D48179B7Ab532C8f5F37e924dad54"): amount.MustParseAmount("10960.714101932569242377"),
				common.HexToAddress("0x273dEf47E1cDd54431EB5a1094615A99511d9e23"): amount.MustParseAmount("0.326533775155048"),
				common.HexToAddress("0x2761C7FF3556DF0f58dBFDC328Bc940EC3Eee652"): amount.MustParseAmount("994.962632554498583013"),
				common.HexToAddress("0x27b0af646384F705a70B354F259311107A86fe5D"): amount.MustParseAmount("500000"),
				common.HexToAddress("0x28D5042528635CdC420A3c6079176C777D5A47dc"): amount.MustParseAmount("0.205120628758592712"),
				common.HexToAddress("0x29b7d3B4fC4bD8aebE8ad4528EeC77E29E2c36eB"): amount.MustParseAmount("12626.98101509393712"),
				common.HexToAddress("0x2A3C2290E71bf659665Cee995926621c152fA591"): amount.MustParseAmount("700"),
				common.HexToAddress("0x2C02007FcAe567925C562D9aE03d983506e386da"): amount.MustParseAmount("47833.8275916408836696"),
				common.HexToAddress("0x2C5Af9110F7Ca24db3c01CB6CD7e71E3fc747111"): amount.MustParseAmount("48460.807833810388062531"),
				common.HexToAddress("0x2D08E6561FBAc164e6F137dBe87fe9B00BB3e122"): amount.MustParseAmount("66165.432914956653491293"),
				common.HexToAddress("0x2E99a0b5E7326B78B35d76e1df6DeEde16070194"): amount.MustParseAmount("500"),
				common.HexToAddress("0x2F412861891C6B54BD5C8f5D76E328D38b9c1a64"): amount.MustParseAmount("7179.8"),
				common.HexToAddress("0x2b1F84694B933261a5B3b95584DEC6aEec7d8B52"): amount.MustParseAmount("56529.998144924726354887"),
				common.HexToAddress("0x2f97eB2314AEdB2a5e798dA609eFFF793380803F"): amount.MustParseAmount("80000"),
				common.HexToAddress("0x308bc4f4F64a61B6Ee2b26E45A6B874D1979943D"): amount.MustParseAmount("33951.779432750968877291"),
				common.HexToAddress("0x316BCeece4bFAE14a5dada54341223C0673f89b9"): amount.MustParseAmount("11735.1256"),
				common.HexToAddress("0x31DedBc002c53C7DC6915b7ca5130eBDb84678C7"): amount.MustParseAmount("10000"),
				common.HexToAddress("0x323f2DB5237CA6D07E5F70704416e197049644B2"): amount.MustParseAmount("36136.28814788325968"),
				common.HexToAddress("0x3399A885E14578B4357bee6775b7C47DA52B661c"): amount.MustParseAmount("25648.019593620661334769"),
				common.HexToAddress("0x33bBdAEcf806d9F16305E9eBBA0Ed236432046Ad"): amount.MustParseAmount("700"),
				common.HexToAddress("0x34456D8364263Ce16b03A60bC70F1C9eDCf1FaCa"): amount.MustParseAmount("79.9"),
				common.HexToAddress("0x348e94330bb3e8e31430a35B3575dF69c0e19fD7"): amount.MustParseAmount("102114.500141076138765377"),
				common.HexToAddress("0x365164f5b42f4fc987750aF6540cAc36fBdBAba6"): amount.MustParseAmount("119.613725216085935"),
				common.HexToAddress("0x36b73Ce25a2FB4c5c9d178Affea7cD8776b9f4Ac"): amount.MustParseAmount("5886.042463739572034295"),
				common.HexToAddress("0x37C5b9c5CA65E3B31e2D0010B60fCD520A035fa8"): amount.MustParseAmount("29555.088107761669687632"),
				common.HexToAddress("0x38351261041b5E5bd040AdCB6FfEE7Ff71E8A868"): amount.MustParseAmount("153.113439011933704"),
				common.HexToAddress("0x3B6b3661D1851Bb90EF143a2aA0e01CE00961D9B"): amount.MustParseAmount("199.6"),
				common.HexToAddress("0x3BEabBD8C325842Be92C0f84615107ed8c5134DF"): amount.MustParseAmount("1000"),
				common.HexToAddress("0x3C9b35cEFf575b24b13A9477A1ca38B7C94861f7"): amount.MustParseAmount("700"),
				common.HexToAddress("0x3aA050753DbF087B84Cca2740E0F36eFCb83e321"): amount.MustParseAmount("161322.321413577105013201"),
				common.HexToAddress("0x3c5354E762BaD0c0FB21853D31aec621a39B8dF5"): amount.MustParseAmount("34999.9"),
				common.HexToAddress("0x3e570781422B9c5Ac8bB897386020E2A55668ad3"): amount.MustParseAmount("200000"),
				common.HexToAddress("0x4109c6D228DebA5632FefcF126c45Dc45C62949A"): amount.MustParseAmount("250000"),
				common.HexToAddress("0x418dbE88e2b2C33E59093bC0B2fCe547857E8A94"): amount.MustParseAmount("181.53572246"),
				common.HexToAddress("0x423D66E9C860f407EA6747c4AF67815Bc204Ad43"): amount.MustParseAmount("200963.563884416598672"),
				common.HexToAddress("0x42642ad8E524B8264ea5c61e7c1eeeb46bBF03C1"): amount.MustParseAmount("681.82787927"),
				common.HexToAddress("0x43c9566A27dC1e85AB2f8588B79988d72364126a"): amount.MustParseAmount("1.2"),
				common.HexToAddress("0x44B63f1DbDce97C417Aab3bB18384C2fC1f230c6"): amount.MustParseAmount("0.934077558448206828"),
				common.HexToAddress("0x4523261ADc33284DAECcf234D22827Beb65cFC5b"): amount.MustParseAmount("698"),
				common.HexToAddress("0x457d6ACB9225dd9f93d305c38C8a253BfF4816DB"): amount.MustParseAmount("7723.055068700033530071"),
				common.HexToAddress("0x45eE05C8c661a64790a44c99aB187277571e1a15"): amount.MustParseAmount("177901.24282351231"),
				common.HexToAddress("0x466Aa028f9b0e8DA1D39D96C79e8e8eca42c39c8"): amount.MustParseAmount("10000"),
				common.HexToAddress("0x4677DFAF4B04b1A68A8112fC8E6419e9e05Cb9a0"): amount.MustParseAmount("640.189840770250896"),
				common.HexToAddress("0x46b99Eab9B01200C9Edf3Ed7CDd4d9912a2C1aB0"): amount.MustParseAmount("100000"),
				common.HexToAddress("0x471Cb122A750C3e625Cc1d6d34dF190B25f6b0a7"): amount.MustParseAmount("369.46174250037952"),
				common.HexToAddress("0x47BEC4c227E26c7C352CC8d1Cd79D6A3589916Af"): amount.MustParseAmount("480.090019082917752097"),
				common.HexToAddress("0x48b6E6f9983aC1391753B305a7eDFC17f1BFCF02"): amount.MustParseAmount("0.957158463977977657"),
				common.HexToAddress("0x494a598d5653996a2d802c264AB82655938C3885"): amount.MustParseAmount("49021.546393950906242258"),
				common.HexToAddress("0x497523039F1ffb5ab215c4ED2618388CC5063601"): amount.MustParseAmount("12185.760589297661800026"),
				common.HexToAddress("0x498F6730E5A48413e5d2c8a8CB321051fA0Aa509"): amount.MustParseAmount("0.63630068854571122"),
				common.HexToAddress("0x4A4f2Dc6dA2D33080f784bD837518C90761bB15e"): amount.MustParseAmount("30772.428452222964496"),
				common.HexToAddress("0x4B0E707A4eE1596B3E6E8A175CD8B4b954c8015A"): amount.MustParseAmount("308027.705518301918325922"),
				common.HexToAddress("0x4a5F5088FEe3150a9CED11F7eD5666B8cB583f02"): amount.MustParseAmount("76189.449569230058467454"),
				common.HexToAddress("0x4b3f78301EE22fCDee40f8b9227D122676F2E1AF"): amount.MustParseAmount("370.27284396787532907"),
				common.HexToAddress("0x4bb55854407dB8edB457D530ebE03386fA0dC9dC"): amount.MustParseAmount("5001.34253553140219744"),
				common.HexToAddress("0x4c4726C25f471A7891dB46731657A7B05FDB8430"): amount.MustParseAmount("4621.144201421499704298"),
				common.HexToAddress("0x4c99F3EE43CCAbf7df61112bF297b7195da28908"): amount.MustParseAmount("311.276794087691856376"),
				common.HexToAddress("0x4d9B86aB150F6E67aC4Be5202571CCECF3550264"): amount.MustParseAmount("6003"),
				common.HexToAddress("0x4dF8dc4982c8471CC4F7B3F7bc0059697443D546"): amount.MustParseAmount("2091.75095615"),
				common.HexToAddress("0x4e7047087218082D57f4a0415cb58EEA6B8CeFAb"): amount.MustParseAmount("9413.918762356289376"),
				common.HexToAddress("0x4fd1968b2bD486748bb76164d02164Cb111F42b2"): amount.MustParseAmount("29990.545514341706822402"),
				common.HexToAddress("0x502d175dD588dddF38d3Ef371acB95CC638D0676"): amount.MustParseAmount("5449.74164168"),
				common.HexToAddress("0x5070091e68F939E81eF49c46FDC5728481fae5b7"): amount.MustParseAmount("71655.26160935912578"),
				common.HexToAddress("0x5103CE8D9C84F1060FE4b441835727a6Ab2A07c3"): amount.MustParseAmount("89560.819515594458896766"),
				common.HexToAddress("0x519A2BF558b871E2178Ae89BAFfBF9d874546d17"): amount.MustParseAmount("51564.234911589564504"),
				common.HexToAddress("0x520ce691F66af08748b296545f29a7f1c6D023e3"): amount.MustParseAmount("999.1"),
				common.HexToAddress("0x52B968c3a1FE701CEf7A8c36875808bf6abbDB7F"): amount.MustParseAmount("553.3659645"),
				common.HexToAddress("0x531747Dd441ECa1E91f7Ae7bB0586dDB01AF7aD9"): amount.MustParseAmount("11198.064893575356248"),
				common.HexToAddress("0x533828Bfb7D55425e24EE43359afa2F3d45a6cD8"): amount.MustParseAmount("400.786637137475032"),
				common.HexToAddress("0x5432587Ff095e87139Bd83f903352271c4C43505"): amount.MustParseAmount("2044.796438158402022326"),
				common.HexToAddress("0x558771Dd216c386431625F3Bacb20D7CA2d00497"): amount.MustParseAmount("0.000795859328438133"),
				common.HexToAddress("0x55F654F873B794bcF5ec04355FD4631Cd6156AE4"): amount.MustParseAmount("700"),
				common.HexToAddress("0x565A348611Bc6dc55F6a65D76A2Ca510047a8B60"): amount.MustParseAmount("700"),
				common.HexToAddress("0x56D7a66b5e91E5e40225E9090c5f4590B055D454"): amount.MustParseAmount("1451.2120734366912"),
				common.HexToAddress("0x57c8f95ac0708010e6D3b4029ccfe482e99f7384"): amount.MustParseAmount("179.6"),
				common.HexToAddress("0x5Cab786a9A06aa2d6c24cE74Af14314184aD80Af"): amount.MustParseAmount("1595.03015873776657064"),
				common.HexToAddress("0x5Ce764a19051C0B9403Be07063D8696523986F33"): amount.MustParseAmount("90227.950392281758407431"),
				common.HexToAddress("0x5D857C9D5b03780AcF9e77EaAF163c2C2fD8704c"): amount.MustParseAmount("31859.938009456498856476"),
				common.HexToAddress("0x5DBf953cc612D1e21285DC0B8397D8FC28878bA3"): amount.MustParseAmount("707.36808061202536"),
				common.HexToAddress("0x5Ee9905846b0060B98d2dE5B0CD40614CE9cbEde"): amount.MustParseAmount("95052.03487154"),
				common.HexToAddress("0x5b0f6Ac8cF7aa71A9d330366407cd7aBB03cfFBd"): amount.MustParseAmount("740.882528160529149309"),
				common.HexToAddress("0x5b85756bebBC79d35302c280E7B73272549918E7"): amount.MustParseAmount("4999.7"),
				common.HexToAddress("0x5d0f83D85F0474E4089921d722f34A5fad66E327"): amount.MustParseAmount("20"),
				common.HexToAddress("0x5d80Ad2d783323B83Ff55dd257089Cd1fb02D60a"): amount.MustParseAmount("4589.4810536726908"),
				common.HexToAddress("0x5e2441bAFA193EA83F666155037359033Be9C496"): amount.MustParseAmount("94302.12098792047283"),
				common.HexToAddress("0x5f1CCa47d2D0d028e96aaaFc8277FAEbC96ac558"): amount.MustParseAmount("661.809436803267165174"),
				common.HexToAddress("0x607A78BFf8FB2Acd4838F073eed969B9F08c673F"): amount.MustParseAmount("8817.889131949939837179"),
				common.HexToAddress("0x615aB70ac137A85E2149b073369d03ee89FCD5dF"): amount.MustParseAmount("1000"),
				common.HexToAddress("0x63a26fedfF6Cf6984481F9e699713f653cFbBeA5"): amount.MustParseAmount("128712.19559994337465472"),
				common.HexToAddress("0x6569939e4E474F3Dc03eEFC4870E2B634082601d"): amount.MustParseAmount("2037.23155461152316648"),
				common.HexToAddress("0x65Ac312A7320beE1db4Fd194082265f8d5C81a12"): amount.MustParseAmount("1248046.918730300277913283"),
				common.HexToAddress("0x6632141ba3473A2399B83a7593ce58cD95E4bbCf"): amount.MustParseAmount("656.02781607"),
				common.HexToAddress("0x6633F9e05a5541Ed91eaa5fcB057bE843345a814"): amount.MustParseAmount("353.408013116239792"),
				common.HexToAddress("0x665E49280ad62f0A5f84fa158461D44ede4b4A8e"): amount.MustParseAmount("11310.54815551"),
				common.HexToAddress("0x673BC80826d7cD466Bb616d7A47A0AF92d5A8c73"): amount.MustParseAmount("8016.8913273809968"),
				common.HexToAddress("0x676E77E8282a79b5FfDCeD34994B66a76bb08C62"): amount.MustParseAmount("1010.172247002019839949"),
				common.HexToAddress("0x67720D3cef8797a8d3A20318f08E74A2F784a285"): amount.MustParseAmount("7885.472006447245296"),
				common.HexToAddress("0x67fF403d97a3E96fC4F7F8007e9d2b4f5930d10A"): amount.MustParseAmount("40"),
				common.HexToAddress("0x69288a3Ea9D824dC08DCA00F17C6Fee8e51908CC"): amount.MustParseAmount("846.1734951814288"),
				common.HexToAddress("0x696eF71Fa303220E550f1941BB04D1C179547731"): amount.MustParseAmount("320.86185911290005428"),
				common.HexToAddress("0x699a1A44cd62D0794CFa7f6EC53918577BfA1bBe"): amount.MustParseAmount("99.8"),
				common.HexToAddress("0x69CE1511DCA7D8dd4F8072346538AceFaa718649"): amount.MustParseAmount("3817.96063489507704"),
				common.HexToAddress("0x6AA70c22475a2040F976A010B6D8f74607875979"): amount.MustParseAmount("3134.32577767"),
				common.HexToAddress("0x6B1977a9D24D8cF43D09dA2c697f464d19f3a3DD"): amount.MustParseAmount("700"),
				common.HexToAddress("0x6D65C73385bb35674057059d966423f64d8fFD63"): amount.MustParseAmount("9990.000000000000003767"),
				common.HexToAddress("0x6F1D04326324994EfE7702933b47a2Bb1BBa3D55"): amount.MustParseAmount("14594.91057457"),
				common.HexToAddress("0x6aa0Df04929906056cC2F2192F874841Bb621B9e"): amount.MustParseAmount("1185.02514941968928"),
				common.HexToAddress("0x6af2c0c7a67B6B2C2be9c47Da88dCD1fC4F88859"): amount.MustParseAmount("164638.185696502881241175"),
				common.HexToAddress("0x6b659F8f516eb1b2A9C5D2FbF36a363F3f082D63"): amount.MustParseAmount("2007.565587818165393775"),
				common.HexToAddress("0x70084E9Ae9F54ba31663861542D73fE718bc27f7"): amount.MustParseAmount("949.723133858886948263"),
				common.HexToAddress("0x703db9c9F0BC2F58DB92C04f12E7D63632A1Cf62"): amount.MustParseAmount("700"),
				common.HexToAddress("0x71cDd63A4fca12a64Fcd71342fe7EeF468C6F4Bb"): amount.MustParseAmount("1429.2178"),
				common.HexToAddress("0x71d7a97c09446bfE3B7abf56ad733E6745ec2482"): amount.MustParseAmount("82735.649711772638932782"),
				common.HexToAddress("0x71e33129c522E71BF3fc3700b584984d7C819F6C"): amount.MustParseAmount("26974.419614340571185577"),
				common.HexToAddress("0x7217D4A72dAEEd42AbAb9bd7A6B461BCA6C8Ef19"): amount.MustParseAmount("1177.281257671722751401"),
				common.HexToAddress("0x728c0001F46de39BA08CdD16d6023Fede987646D"): amount.MustParseAmount("8521.073482871569601723"),
				common.HexToAddress("0x72A832AE6f4025Ab21CA5265016eef34BD503E3b"): amount.MustParseAmount("49172.873664554602173178"),
				common.HexToAddress("0x72fAefED9bc0eD588797853F7a0027286ED0E72c"): amount.MustParseAmount("819383.468300846544951148"),
				common.HexToAddress("0x737C92222EaF4D1083278242Cc7D5252B88d1c1c"): amount.MustParseAmount("199.6"),
				common.HexToAddress("0x7400f026068eA0Ff984Be97e2b701dE669Ea6E3F"): amount.MustParseAmount("8157.19696573641796"),
				common.HexToAddress("0x74079F4654D971B6Ee19B9Eae2f2ea745e84F3a2"): amount.MustParseAmount("495.2631401"),
				common.HexToAddress("0x742C9E5B999F1940FD0c29dFB1E33717ee1D59bb"): amount.MustParseAmount("161608.708950501662559335"),
				common.HexToAddress("0x74D8E5092f0c2E8A829C8EEF0237d53Da2BF0A60"): amount.MustParseAmount("6478.474282800826073766"),
				common.HexToAddress("0x768F903ff09027B66F90F2212e8A2FeeE123395A"): amount.MustParseAmount("189191.937025543323102"),
				common.HexToAddress("0x7825a4eE7F4e49fc27B7DE961355B07389E5dAFF"): amount.MustParseAmount("5327.73276883592711392"),
				common.HexToAddress("0x78Be24C61B8A273BCB7F73543CD02f7A48810Dd5"): amount.MustParseAmount("193681.90542611217656"),
				common.HexToAddress("0x78d7794106A1365DE34d1ab5A8cAE6de96Ec995F"): amount.MustParseAmount("167.9718197"),
				common.HexToAddress("0x7972d210E29b7a6E615c81cA2f6d635D8E76a1FC"): amount.MustParseAmount("1205.30076112309799674"),
				common.HexToAddress("0x79e83d2108DD1EaEf51B8fD579cD9d03136Fc7f2"): amount.MustParseAmount("3949.01543955"),
				common.HexToAddress("0x7A562B2C61daB3C02c9930414736b6f5700fD6B0"): amount.MustParseAmount("10219.4328"),
				common.HexToAddress("0x7Ada445b4129bB2c78C216A0B728456e4Edf74D2"): amount.MustParseAmount("500000"),
				common.HexToAddress("0x7BC81c97C94Ade359733240Ff1fE8d7a303C3eab"): amount.MustParseAmount("299.5"),
				common.HexToAddress("0x7D1a1d1e62FBCC77D0086415d8844CDD243Afc7a"): amount.MustParseAmount("79824.796133713217142"),
				common.HexToAddress("0x7D9965D0e87001228ACfFd0981a4557b287D16f8"): amount.MustParseAmount("1526.965520844102841018"),
				common.HexToAddress("0x7F1e0271F0828bBf1D60440BcCc77e3a6753Dd55"): amount.MustParseAmount("8218.42690166031828492"),
				common.HexToAddress("0x7F3ed728E32257944F8156c6a91f15662Ca70f45"): amount.MustParseAmount("33970.73369879864879572"),
				common.HexToAddress("0x7c0E9b7858F7568f163Ef18F1E5e570432dDd976"): amount.MustParseAmount("358.839765010727566796"),
				common.HexToAddress("0x7c59B5Cc837e1400Ed27215b72aF4AAE94bA4cCC"): amount.MustParseAmount("990.6"),
				common.HexToAddress("0x7cd4F4B2b811cF35559CaA1D85F2097F38f882bb"): amount.MustParseAmount("4999.9"),
				common.HexToAddress("0x7d1c4A60a8599DD087e1edD970D7CF80d147F840"): amount.MustParseAmount("540"),
				common.HexToAddress("0x7e404300E7a2483E99A63047658bB931dB65Ed79"): amount.MustParseAmount("18.9"),
				common.HexToAddress("0x7e95fdbc0203ab6fD5Cf16129f0b42e4F2BD130d"): amount.MustParseAmount("999.6"),
				common.HexToAddress("0x7f9bEC0d93cE2B51E9247d3b41cD99841220c593"): amount.MustParseAmount("99.1"),
				common.HexToAddress("0x805CD293dF99ab01eaa3cFAB571B6Da02338571A"): amount.MustParseAmount("470"),
				common.HexToAddress("0x818a792cE79dc387254975D234e73D400A18a784"): amount.MustParseAmount("465.811939804084942948"),
				common.HexToAddress("0x81d37b002Db4dCa2907E12e15d532E736d0AEA75"): amount.MustParseAmount("13604.5"),
				common.HexToAddress("0x828b764DdcE9D93fA5928CFFF86b7a510c8f715e"): amount.MustParseAmount("1854.679878519547285126"),
				common.HexToAddress("0x83D70D2a9F88252818FfdcB387d1d69126856EeA"): amount.MustParseAmount("7802.60477792953256"),
				common.HexToAddress("0x84135Fa217B959A91A951E1223874E2536c909b3"): amount.MustParseAmount("36257.074990019742797512"),
				common.HexToAddress("0x84285c11b11312490aEBfef76E85C2d5542B085E"): amount.MustParseAmount("43.525003969974768"),
				common.HexToAddress("0x843eE5F04CA1187915fbCf09d5000f3d5AdC985C"): amount.MustParseAmount("32.6"),
				common.HexToAddress("0x84E9038E7721959A92F8C604d89Ccfaa783cd97b"): amount.MustParseAmount("4883.5417253267376"),
				common.HexToAddress("0x86CE5F1ca8F23355B4F86c326ce2009780eFa3F1"): amount.MustParseAmount("270.70469261251665956"),
				common.HexToAddress("0x877D8416ab6f3bDfDB08C44a3201dd3b1E0f7662"): amount.MustParseAmount("1792.824131108697524163"),
				common.HexToAddress("0x898abdd3BB93D47b7C56CDde7e1ce579CAaA2797"): amount.MustParseAmount("999.9"),
				common.HexToAddress("0x8FEbD2F8D7C91851C032924A0195002796C9D71D"): amount.MustParseAmount("21007.768046764125794835"),
				common.HexToAddress("0x8ac973F2b15f8B6779AA0417a6528Ac1f96a8232"): amount.MustParseAmount("2115.930165986860319986"),
				common.HexToAddress("0x8bb54C4c1c84D44F2F055605867624Ce28373EB5"): amount.MustParseAmount("0.986944780460894782"),
				common.HexToAddress("0x900E7021450106eB3043ed8d85cE9A1749eAc5AB"): amount.MustParseAmount("5395.66306998"),
				common.HexToAddress("0x905b1AE23FdEd3CA5a55A8e5e5E017D11D2F9A24"): amount.MustParseAmount("4999.1"),
				common.HexToAddress("0x90c418626ED552A2089b28e14c9634bC6f178c6a"): amount.MustParseAmount("101.945092627821317096"),
				common.HexToAddress("0x90c759b6255A252D9b091ad791f1d56B3eFc62f2"): amount.MustParseAmount("29.5089791923316"),
				common.HexToAddress("0x914529d6f0681261089dB03BC33cC929238969Ba"): amount.MustParseAmount("500000"),
				common.HexToAddress("0x9289bFd8642424af66944aEa7c22941d706fA82C"): amount.MustParseAmount("19.9"),
				common.HexToAddress("0x92AfF669F8c51F6eAFA7750C78e0E17cF78F77b4"): amount.MustParseAmount("5740.662930659556464"),
				common.HexToAddress("0x92F0683926D991Ee5eaB6DDC27e950c075DE9534"): amount.MustParseAmount("99.8"),
				common.HexToAddress("0x93132E8cf7556a66ae19F13a5A122e2034dAa675"): amount.MustParseAmount("1000"),
				common.HexToAddress("0x9365A03ce4c748C216891c1a5F8024514eC56d11"): amount.MustParseAmount("700"),
				common.HexToAddress("0x9400fFe5627D384c07DD60A35c4637F8C88a724E"): amount.MustParseAmount("2000"),
				common.HexToAddress("0x9472A5F51BdD3278f1C91Fc971c3e6A3B8c69A18"): amount.MustParseAmount("2781.25741721225152"),
				common.HexToAddress("0x948663C1cB42c7EA3694CEAf819489650b15A2E4"): amount.MustParseAmount("43050.44483247079530664"),
				common.HexToAddress("0x96e5f2118A031289588cF1ebDDEb0966ee4D2c82"): amount.MustParseAmount("30000"),
				common.HexToAddress("0x9717B077e0F192b6251cfB1C07700dA6B0cC6c07"): amount.MustParseAmount("0.673356129620574033"),
				common.HexToAddress("0x9779076eF9f09eDF4eCc5A7ea17aC502fCdDaEa9"): amount.MustParseAmount("203.84767627"),
				common.HexToAddress("0x980dbe6B26799566344bD9a729e215b90cd12eE4"): amount.MustParseAmount("20661.246287767536525932"),
				common.HexToAddress("0x983e41b155c26531A5dadD3BB416131d56CC930b"): amount.MustParseAmount("183.1166453"),
				common.HexToAddress("0x987B23e03d7Fa2F1893252F5E8394Bd142116925"): amount.MustParseAmount("249.9"),
				common.HexToAddress("0x988eFeCb764ca5d1503eBb6EC32EA1EC4033d5D1"): amount.MustParseAmount("3000"),
				common.HexToAddress("0x99F8B99009779CA938a81ACedF933796577c46Da"): amount.MustParseAmount("51064.39121232248568"),
				common.HexToAddress("0x9A35c15f630FAABe2C4ba42Af284CFcfFCEAA27f"): amount.MustParseAmount("2557505.328144284108066831"),
				common.HexToAddress("0x9AD3FAb3014696DDeE4D84c9a67463791F17D108"): amount.MustParseAmount("227.14064086"),
				common.HexToAddress("0x9Bb38896B1135be5FAA45d2954390A93f48eef5a"): amount.MustParseAmount("8998.422428478300343795"),
				common.HexToAddress("0x9a021031CbD7cb5bdDFB10f4C632CB57bD35C0E7"): amount.MustParseAmount("1153.0050641091496"),
				common.HexToAddress("0x9c730c8914dC2eE07c10E9a81DE8A83752A059A5"): amount.MustParseAmount("100.899999999999999991"),
				common.HexToAddress("0xA0D45F76DeF589dcBB217Df1bFAB3d3ef84283c3"): amount.MustParseAmount("700"),
				common.HexToAddress("0xA151E37CaC28F2ED94a247f8Da38D7eba9CA9911"): amount.MustParseAmount("0.880406666444130525"),
				common.HexToAddress("0xA2D57c430209B448250fFd1280AA269a2631e31e"): amount.MustParseAmount("135839.62721"),
				common.HexToAddress("0xA3a076D795Ba216AeF674E326eBD3946CBD61e97"): amount.MustParseAmount("1790496.60042103838098"),
				common.HexToAddress("0xA46c8BC4A771aA97e314407d9626883B6e152035"): amount.MustParseAmount("563759.030383298840628067"),
				common.HexToAddress("0xA97A8049042e364751B45fACCCB33f67E7992941"): amount.MustParseAmount("20416.356898378245925548"),
				common.HexToAddress("0xAA5507F0d96eD27fad1e8f83b20dDb3627FCBaB4"): amount.MustParseAmount("64595.80186506150565328"),
				common.HexToAddress("0xAEb795e46B6C520635dAd1387609619b3f89DEbB"): amount.MustParseAmount("239.234033137818377864"),
				common.HexToAddress("0xAd57101D9810186686e7F042177AA369e645417f"): amount.MustParseAmount("12914.566188627673018475"),
				common.HexToAddress("0xB38fc1E79c680b899d7c69cFC8E0b0D0373F77a9"): amount.MustParseAmount("7.963396873051895391"),
				common.HexToAddress("0xB6C6F6Eb8ab960B450187642afC7d60B3ab65374"): amount.MustParseAmount("999.6"),
				common.HexToAddress("0xBEAac2179F26829A01CcC0Bc036E702b461B07A8"): amount.MustParseAmount("0.720011073879072551"),
				common.HexToAddress("0xBFFb8e79fCEF07575a24dfC023b97EF016F00205"): amount.MustParseAmount("0.000000000055130795"),
				common.HexToAddress("0xBfEdcFb5237c360fFc541Fa65934161708Faa108"): amount.MustParseAmount("1615269.98091519346016"),
				common.HexToAddress("0xC3C76c7637a11D4fc061a315762e993a6708cD88"): amount.MustParseAmount("31174.963262016268096246"),
				common.HexToAddress("0xC5A668483F060EAfA707cccB78CE50d681Cc6683"): amount.MustParseAmount("405.9278"),
				common.HexToAddress("0xC5Cf312681ae925714057D7F6570f742c751282a"): amount.MustParseAmount("87242.753951821144661637"),
				common.HexToAddress("0xC5D24B0a635447b51868648b774F7Ba283E8A351"): amount.MustParseAmount("200994.328983877027233275"),
				common.HexToAddress("0xC5eF67FcA705d5067619d7545fd9e1253C56CC8b"): amount.MustParseAmount("372.734097133739686164"),
				common.HexToAddress("0xC632C1DB8c8B2EdA5A224fbe62cD539929d6f104"): amount.MustParseAmount("93335.99838334"),
				common.HexToAddress("0xC73e632109Ba7F2aFd1928F62ecFbF90286bfDAF"): amount.MustParseAmount("1000"),
				common.HexToAddress("0xC98301EF0c63f468097eB4e0B6743d55928CBE64"): amount.MustParseAmount("80156.977035610510120652"),
				common.HexToAddress("0xC9b5FcCD173fF9F1E4634f259206C07f997d239b"): amount.MustParseAmount("498.8"),
				common.HexToAddress("0xCC74D2eC6f52eC6CE3a943D111F4d52e3eeC9f35"): amount.MustParseAmount("159095.720727723050082189"),
				common.HexToAddress("0xCc29a75F0297Af3F183CDda38Da05073724437De"): amount.MustParseAmount("399.56174250037952"),
				common.HexToAddress("0xD110eD592A0aD1f87a1320A3DcAEFbe58cde942c"): amount.MustParseAmount("13.90705632"),
				common.HexToAddress("0xD14DC00336dEc0FA79Cd652e87cB67ba314b8C57"): amount.MustParseAmount("169882.9"),
				common.HexToAddress("0xD2600295A2C3bBD4d33dde1A99b8b603047fd00E"): amount.MustParseAmount("3004354.93175310545"),
				common.HexToAddress("0xD3080C0A8aD922D7021D1bc0790d94BF8E20C39d"): amount.MustParseAmount("30.808"),
				common.HexToAddress("0xD6978DA07b4DB4034B1245BD898920f8f1B13499"): amount.MustParseAmount("358.839765010727566796"),
				common.HexToAddress("0xD7A38a6b486db1f32588c0aa5493905BBDAfb845"): amount.MustParseAmount("1000"),
				common.HexToAddress("0xD7FCF3A71309dC4a4bdd720de3Bb0B62CF1A31Df"): amount.MustParseAmount("0.25782343889016"),
				common.HexToAddress("0xD916a85529983D85853c2f3d9A4D5B4c56A0F8A4"): amount.MustParseAmount("2844.6"),
				common.HexToAddress("0xDA16deBbae3AE7916095d8779534FaBDd2cedb4a"): amount.MustParseAmount("124002.4016911242728"),
				common.HexToAddress("0xDA1b0cCfc7EbDA51C1C46f357E9bC05634baD636"): amount.MustParseAmount("0.555114550744"),
				common.HexToAddress("0xDC2264ECb69A09dacD2847dAb259C570A267Eef0"): amount.MustParseAmount("2870144.893794929735944914"),
				common.HexToAddress("0xDD2fa60caa7f0bbdaf6749418bD95Ca85f4a04B9"): amount.MustParseAmount("1156.89802054082224015"),
				common.HexToAddress("0xDaCDE68e10742D3576c70ae0324519C174705EA0"): amount.MustParseAmount("3.8387909"),
				common.HexToAddress("0xDf0f0F953326c64d2178d841Bf4912FCDCbFb9Ad"): amount.MustParseAmount("53481.524064031115188729"),
				common.HexToAddress("0xE663B7cD451e544aCd04968099cceB00e2E6EF04"): amount.MustParseAmount("474703.362605703168629981"),
				common.HexToAddress("0xE738E650eff9D5B369D8171c94ed4307C1732B98"): amount.MustParseAmount("700"),
				common.HexToAddress("0xE8454C8fe12E5E146cf2D8D70e3659Dea575186F"): amount.MustParseAmount("10506.70975714"),
				common.HexToAddress("0xEEcA2D5171CD6A33389573533b78E729c8F96BC9"): amount.MustParseAmount("144070.957997027509443677"),
				common.HexToAddress("0xEaf4a591dB67290821C4784FC2A2C5365963C93c"): amount.MustParseAmount("621.78107615529488"),
				common.HexToAddress("0xEb2f08Cb9E9043898eC2445Fb8Ef55F08A770a9a"): amount.MustParseAmount("66473.322981670479842218"),
				common.HexToAddress("0xEf8923DAA9ba2741C91246D38fd0dbebd6673f4c"): amount.MustParseAmount("38.924779153973688"),
				common.HexToAddress("0xF0d5D0Fa9ec74eb93448161585547eCa9aE25bB7"): amount.MustParseAmount("299.8"),
				common.HexToAddress("0xF58F3FC306DF98fB8ED1992D6a7851cB55f08a47"): amount.MustParseAmount("101019.63196652"),
				common.HexToAddress("0xF5D3608296f701a9c04f7f52023E4DB05E75CBD7"): amount.MustParseAmount("209271.624947529461615033"),
				common.HexToAddress("0xF699027a46B9AD39180849D5F0864EF689ADD09f"): amount.MustParseAmount("185406.949748577842376126"),
				common.HexToAddress("0xF9398803Bd128eBb2cAE729C26E22ED7f70DC6F1"): amount.MustParseAmount("0.899999999999999991"),
				common.HexToAddress("0xF93cEA58905F5699546108B16f06E61834680472"): amount.MustParseAmount("2907.115593505227115503"),
				common.HexToAddress("0xFBb86775C3326fbEA011B78Fc0F7E46DD8b370e5"): amount.MustParseAmount("654065.901274103673320881"),
				common.HexToAddress("0xFEB842E977381514860e66FcFd11FC7db332862e"): amount.MustParseAmount("1000"),
				common.HexToAddress("0xFaBBF15DCf7950D59E85EeF6efb56eB29B748461"): amount.MustParseAmount("0.285145721161774154"),
				common.HexToAddress("0xFc42248F0DFcf5E48192AeDD4C14B37562a00973"): amount.MustParseAmount("838.945580370459829102"),
				common.HexToAddress("0xFc73E8EE0C4c34EFA3BfA5Bc5Cce632211d2763f"): amount.MustParseAmount("700"),
				common.HexToAddress("0xa1904b9e57Aeefaec56117D05525E15B263b7ac4"): amount.MustParseAmount("1800000"),
				common.HexToAddress("0xa1EC05825785B18Db57D10d80B9E92B48191C46d"): amount.MustParseAmount("664.833742496317579005"),
				common.HexToAddress("0xa21eC062b30FA545448fA8492142E6d9F9467853"): amount.MustParseAmount("52.73963775058790679"),
				common.HexToAddress("0xa32B77fCad4cE0dFc85F50F3674cb8Db5A52f8e0"): amount.MustParseAmount("9489.9"),
				common.HexToAddress("0xa5338478E039ff27C03EE880788327303c44e231"): amount.MustParseAmount("199.2"),
				common.HexToAddress("0xa61627a61Becd5C6DA8A216E409206f1f61298AA"): amount.MustParseAmount("9.9"),
				common.HexToAddress("0xa9B465f2977d0bB6D914edaBDC55B3F597D8b9dC"): amount.MustParseAmount("1098.64740707858"),
				common.HexToAddress("0xaD8C21941D1DAA2C99BdD73e3184b552e3696287"): amount.MustParseAmount("700"),
				common.HexToAddress("0xad8D3aD50733C90171544111FA86639AFC23D308"): amount.MustParseAmount("9131.04483221"),
				common.HexToAddress("0xb07F9012eEf574072C5dD8795Ba0e2bde804c431"): amount.MustParseAmount("100000"),
				common.HexToAddress("0xb16845Bef7E536A9C0454d179387Bbc450AEA9Fb"): amount.MustParseAmount("82.246469826684568"),
				common.HexToAddress("0xb2BD59197ffa4ad521ef0E3E86A683f1393D91b1"): amount.MustParseAmount("21604.4561"),
				common.HexToAddress("0xb2b53a8bF273380eD1FBcdc46849176D3B321eE1"): amount.MustParseAmount("4654.93773624557224"),
				common.HexToAddress("0xb3E0226a91E54A8FB07188A6C6412A588911F87a"): amount.MustParseAmount("2129.349896810031024"),
				common.HexToAddress("0xb40Ad050012FAaB3061945490A196Df5Ffc22591"): amount.MustParseAmount("134.698976651803209716"),
				common.HexToAddress("0xb4AC84F260CE649bE580692E89ae18522c4d7579"): amount.MustParseAmount("86516.71792021"),
				common.HexToAddress("0xb4Feb654de411a4df250586ec67f09E30b2F0888"): amount.MustParseAmount("59836.945280936385313186"),
				common.HexToAddress("0xb67E062105E5090CFdB79C5A0Acd2dEe912D2d8B"): amount.MustParseAmount("20421.0213690660323602"),
				common.HexToAddress("0xb94e9590A9d71a5EC8F1844715f99752A849bC8e"): amount.MustParseAmount("86470.338268216047823155"),
				common.HexToAddress("0xbA00785E74746b61496B96a112B6550c88CA1422"): amount.MustParseAmount("39.20808847148608"),
				common.HexToAddress("0xbC4D11D7dBeD9969a14Ad0279698E681460061c7"): amount.MustParseAmount("529.1306"),
				common.HexToAddress("0xbF15ed505b363e9e0d9CD47867B5742aC7E0dc8C"): amount.MustParseAmount("19396.617088409208325139"),
				common.HexToAddress("0xbF21B3bD2c1F30f2031387f09d6d8d4e78c983A8"): amount.MustParseAmount("940.20625415725352"),
				common.HexToAddress("0xbc36795464C231860296bBc449349F028Cf8Ed97"): amount.MustParseAmount("600"),
				common.HexToAddress("0xbc4d35dAb3161B880fE0bfC88feA79A26eB3e3eF"): amount.MustParseAmount("19.4"),
				common.HexToAddress("0xbdc32957d16eE2cf11C99750eD789f54F418e8d4"): amount.MustParseAmount("15393.188958577696491107"),
				common.HexToAddress("0xbfF99071E3e975228165A8B5A4d6A929AB8c43e6"): amount.MustParseAmount("1"),
				common.HexToAddress("0xc13C0035Ec712D8e0282b131335739b924e4184D"): amount.MustParseAmount("109.769250966875238898"),
				common.HexToAddress("0xc45CbeEA1aaF686D1a195dd16AB5f872821Af0a3"): amount.MustParseAmount("1035.00966334009272"),
				common.HexToAddress("0xc7A0e55a80b38EB49d9BCA9918414fdd633C84A2"): amount.MustParseAmount("23932.421172229869936788"),
				common.HexToAddress("0xc7E628d39859bFC5c638DdAec6ED73AcaB5405D9"): amount.MustParseAmount("12048.977430994479462983"),
				common.HexToAddress("0xc8E98E3314a9777FFe378E84d68C3D77911dcFD0"): amount.MustParseAmount("1047.22949096589552"),
				common.HexToAddress("0xc98B48B9509F501aFc1b53dfd8896148EA617097"): amount.MustParseAmount("15968.413544497869943413"),
				common.HexToAddress("0xcA30A951530523cEa94bbB98AF692Ea7E696EAE6"): amount.MustParseAmount("6463.384149660339896"),
				common.HexToAddress("0xcA44461E6A343E27Ce5E87034272D0A0361FD3b9"): amount.MustParseAmount("24312.06587622896"),
				common.HexToAddress("0xcAF4217a1f40748ee3Ea9fab73dA060fFFfA2faA"): amount.MustParseAmount("0.003888986612215964"),
				common.HexToAddress("0xcAfCd28Dd170EA2DC41701E32a8AbEFdfDC221d8"): amount.MustParseAmount("70554.89569766380327866"),
				common.HexToAddress("0xcB68d98C427EA984Da540B1d2850fC35dDFD0336"): amount.MustParseAmount("2075.574592854287134941"),
				common.HexToAddress("0xcBEA83f29F3c3A1270bAc9F5cD0826E5813c2c18"): amount.MustParseAmount("0.000000000000000008"),
				common.HexToAddress("0xcDbF2470ef91cB4a9d18725b66eCce873C229DD5"): amount.MustParseAmount("252142.78420094435979"),
				common.HexToAddress("0xca4F2cec68e4f0A76D0F5C5311D3b4c07cc9C82A"): amount.MustParseAmount("485.622170429200145279"),
				common.HexToAddress("0xcbD42B7f4E01d2BF56453dF7dB23574286699a53"): amount.MustParseAmount("98.4"),
				common.HexToAddress("0xcf074C8bC53BD24B4Cce0e2e1e1C27366527d754"): amount.MustParseAmount("5311.84696990151808"),
				common.HexToAddress("0xd1990cC46658c24B653af93C7B9769C0955941cC"): amount.MustParseAmount("1990.34697378997648"),
				common.HexToAddress("0xd2aEe5Bce35d2e57C6483690F69e26A8FdE2AABC"): amount.MustParseAmount("3481.261407142408204607"),
				common.HexToAddress("0xd38013fbb464517AAd2f976576885052a0f5fEF6"): amount.MustParseAmount("0.216126841776819586"),
				common.HexToAddress("0xd741bA086f243178951c88D6deA37C0D6166f494"): amount.MustParseAmount("175564.523213199419791058"),
				common.HexToAddress("0xd7DB6fB25E1034C3775979bf51aFEB060D16aA48"): amount.MustParseAmount("84.89294368"),
				common.HexToAddress("0xd833EB915EBE79058B4F04141179F15AC55f6EB8"): amount.MustParseAmount("3344.937789247657643327"),
				common.HexToAddress("0xd930e2D8Ae342f83968C912dC05bF341430A049A"): amount.MustParseAmount("2742.107996523733256699"),
				common.HexToAddress("0xd98dbAA2E5e0Ea3875Ec552C963464cDe2089Bb4"): amount.MustParseAmount("149.649147116325283876"),
				common.HexToAddress("0xdD1c46929521bCECbccd484CcC3aC08B96221aD1"): amount.MustParseAmount("12745.445367493768256612"),
				common.HexToAddress("0xdc266DE43376eb3262B1C56980bCCe6af896bf6b"): amount.MustParseAmount("1000"),
				common.HexToAddress("0xdf4FE21e353d1BD3C4B9A45471703D8C2Aa22AF0"): amount.MustParseAmount("2004.723788101362305934"),
				common.HexToAddress("0xe0F83dC078ecaeDEf974264D3c4a998878134816"): amount.MustParseAmount("2.34"),
				common.HexToAddress("0xe3038bd7A4bf7E4A52E1b51848de2b7D6164a6c7"): amount.MustParseAmount("23902.224968676967699246"),
				common.HexToAddress("0xe636836bc387905983eB0a23c3E3FdB926289Fc7"): amount.MustParseAmount("12712.073936020037300131"),
				common.HexToAddress("0xe6729C5FB575077d85078fA00c21a39D456D7681"): amount.MustParseAmount("191382.358903178096648"),
				common.HexToAddress("0xe8339eEFe2E0f86c1755702818876CE84D09d757"): amount.MustParseAmount("100000"),
				common.HexToAddress("0xe854fC537ba3C880BaCdD76082532483cCeA495e"): amount.MustParseAmount("12383.053487385895237622"),
				common.HexToAddress("0xe900DbCeEF5C1c5Df0fD8D9dAd016bADA9dC3D5b"): amount.MustParseAmount("2984.244868466543169415"),
				common.HexToAddress("0xe904fD3ECBfE4f448334501e9e16EBe6c17d2eeD"): amount.MustParseAmount("13645.2544720653292"),
				common.HexToAddress("0xe90e03eD54461894DB3F26BedfFa364e2FA45d7B"): amount.MustParseAmount("1"),
				common.HexToAddress("0xeD5AA3De684126B468a4059B4323a9F4dB8Bcc78"): amount.MustParseAmount("11535.772334258117668784"),
				common.HexToAddress("0xeE06c26536A2e1E054cF0CaDC9394997EB5d2a77"): amount.MustParseAmount("20"),
				common.HexToAddress("0xeF00607B5c28BcC68A442b584aCa7438FAB5451C"): amount.MustParseAmount("487.5912525989012"),
				common.HexToAddress("0xef3624A4929Cd758Bb844dE71A3C4b9fafC885F5"): amount.MustParseAmount("30000"),
				common.HexToAddress("0xf0E25d6CB9e8C20f8fab342896e399bF51061403"): amount.MustParseAmount("580.8"),
				common.HexToAddress("0xf13beB84d612A97ab778B098c4354b91F8d5dbE5"): amount.MustParseAmount("119621.855305677232732"),
				common.HexToAddress("0xf54142ed63aBDFF6bE03371b79349652fd7BC85B"): amount.MustParseAmount("0.913389906579969456"),
				common.HexToAddress("0xf5FFfae0Fa335Df728a37188c87B47532e242238"): amount.MustParseAmount("2000"),
				common.HexToAddress("0xf672b570cFc99aB12494600A10b1a3311b4f3bBC"): amount.MustParseAmount("110016.319086198382967503"),
				common.HexToAddress("0xf6eb7EFaE6BcEA2913a0636a7F510B4D9060f530"): amount.MustParseAmount("288848.162974752895039568"),
				common.HexToAddress("0xfaf601dC4C4e7A413aD3A5fFF56Bf3071912467C"): amount.MustParseAmount("176152.513886125956079632"),
				common.HexToAddress("0xfc4aaF3791634E4DACc289e23eDdFadE0B2F32bB"): amount.MustParseAmount("999.9"),
				common.HexToAddress("0xfe58AE2f8CC09BaD1490F5b0aD613E1f279C657c"): amount.MustParseAmount("5810.00281456423504"),
				common.HexToAddress("0xff9F5b26878bcBE38D03799736303a549f09A7F8"): amount.MustParseAmount("0"),
			},
		}
		bs, _, err := bin.WriterToBytes(arg)
		if err != nil {
			panic(err)
		}
		cont, err := genesis.DeployContract(adminAddress, ClassMap["Token"], bs)
		if err != nil {
			panic(err)
		}
		tokenAddress = cont.Address()
		genesis.SetMainToken(tokenAddress)

		fmt.Println("Token", tokenAddress.String())
		// Token 0xeF3432F1D54eC559613f44879FEF8a7866dA3e07
	}
	if true {
		arg := &gateway.GatewayContractConstruction{
			TokenAddress: tokenAddress,
		}
		bs, _, err := bin.WriterToBytes(arg)
		if err != nil {
			panic(err)
		}
		v, err := genesis.DeployContract(adminAddress, ClassMap["Gateway"], bs)
		if err != nil {
			panic(err)
		}
		cont := v.(*gateway.GatewayContract)
		gatewayAddress = cont.Address()

		cc := genesis.ContractContext(cont, adminAddress)
		intr := types.NewInteractor(genesis, cont, cc)
		cc.Exec = intr.Exec

		PlatformList := map[string]*amount.Amount{
			"MEVerse":   amount.MustParseAmount("0.1"),
			"ETH":       amount.MustParseAmount("300"),
			"BSC":       amount.MustParseAmount("20"),
			"Klaytn":    amount.MustParseAmount("5"),
			"Polygon":   amount.MustParseAmount("5"),
			"Tomochain": amount.MustParseAmount("50"),
		}
		for platform, am := range PlatformList {
			if err := cont.AddPlatform(cc, platform, am); err != nil {
				panic(err)
			}
		}
		intr.Distroy()

		fmt.Println("Gateway", gatewayAddress.String())
		// Gateway 0x0eA20dAcf567eB840f1B2463F366650086fb995c
	}

	if true {
		arg := &formulator.FormulatorContractConstruction{
			TokenAddress: tokenAddress,
			FormulatorPolicy: formulator.FormulatorPolicy{
				AlphaAmount:    amount.NewAmount(200000, 0),
				SigmaCount:     4,
				SigmaBlocks:    5184000,
				OmegaCount:     2,
				OmegaBlocks:    5184000,
				HyperAmount:    amount.NewAmount(0, 0),
				MinStakeAmount: amount.NewAmount(100, 0),
			},
			RewardPolicy: formulator.RewardPolicy{
				RewardPerBlock:        amount.MustParseAmount("0.4756468797564688"),
				AlphaEfficiency1000:   1000,
				SigmaEfficiency1000:   1150,
				OmegaEfficiency1000:   1300,
				HyperEfficiency1000:   1300,
				StakingEfficiency1000: 700,
				CommissionRatio1000:   50,
				MiningFeeAddress:      adminAddress,
				MiningFee1000:         300,
			},
		}
		bs, _, err := bin.WriterToBytes(arg)
		if err != nil {
			panic(err)
		}
		v, err := genesis.DeployContract(adminAddress, ClassMap["Formulator"], bs)
		if err != nil {
			panic(err)
		}
		cont := v.(*formulator.FormulatorContract)
		formulatorAddress = cont.Address()

		if true {
			v, err := genesis.Contract(tokenAddress)
			if err != nil {
				panic(err)
			}
			cont := v.(*token.TokenContract)
			cc := genesis.ContractContext(cont, adminAddress)
			if err := cont.SetMinter(cc, formulatorAddress, true); err != nil {
				panic(err)
			}
		}

		cc := genesis.ContractContext(cont, cont.Address())
		intr := types.NewInteractor(genesis, cont, cc)
		cc.Exec = intr.Exec
		for _, addr := range alphaOwners {
			if _, err := cont.CreateGenesisAlpha(cc, addr); err != nil {
				panic(err)
			}
		}
		for _, addr := range sigmaOwners {
			if _, err := cont.CreateGenesisSigma(cc, addr); err != nil {
				panic(err)
			}
		}
		for _, addr := range omegaOwners {
			if _, err := cont.CreateGenesisOmega(cc, addr); err != nil {
				panic(err)
			}
		}
		for hyper, mp := range stakingMap {
			for addr, am := range mp {
				if err := cont.AddGenesisStakingAmount(cc, hyper, addr, am); err != nil {
					panic(err)
				}
			}
		}
		intr.Distroy()

		fmt.Println("formulatorAddress", formulatorAddress.String())
		// formulatorAddress 0x75A098f86bAe71039217a879f064d034c59C3766
	}
	return genesis.Top()
}
