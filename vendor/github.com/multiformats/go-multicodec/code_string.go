// Code generated by "stringer -type=Code -linecomment"; DO NOT EDIT.

package multicodec

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Identity-0]
	_ = x[Cidv1-1]
	_ = x[Cidv2-2]
	_ = x[Cidv3-3]
	_ = x[Ip4-4]
	_ = x[Tcp-6]
	_ = x[Sha1-17]
	_ = x[Sha2_256-18]
	_ = x[Sha2_512-19]
	_ = x[Sha3_512-20]
	_ = x[Sha3_384-21]
	_ = x[Sha3_256-22]
	_ = x[Sha3_224-23]
	_ = x[Shake128-24]
	_ = x[Shake256-25]
	_ = x[Keccak224-26]
	_ = x[Keccak256-27]
	_ = x[Keccak384-28]
	_ = x[Keccak512-29]
	_ = x[Blake3-30]
	_ = x[Sha2_384-32]
	_ = x[Dccp-33]
	_ = x[Murmur3X64_64-34]
	_ = x[Murmur3_32-35]
	_ = x[Ip6-41]
	_ = x[Ip6zone-42]
	_ = x[Ipcidr-43]
	_ = x[Path-47]
	_ = x[Multicodec-48]
	_ = x[Multihash-49]
	_ = x[Multiaddr-50]
	_ = x[Multibase-51]
	_ = x[Dns-53]
	_ = x[Dns4-54]
	_ = x[Dns6-55]
	_ = x[Dnsaddr-56]
	_ = x[Protobuf-80]
	_ = x[Cbor-81]
	_ = x[Raw-85]
	_ = x[DblSha2_256-86]
	_ = x[Rlp-96]
	_ = x[Bencode-99]
	_ = x[DagPb-112]
	_ = x[DagCbor-113]
	_ = x[Libp2pKey-114]
	_ = x[GitRaw-120]
	_ = x[TorrentInfo-123]
	_ = x[TorrentFile-124]
	_ = x[LeofcoinBlock-129]
	_ = x[LeofcoinTx-130]
	_ = x[LeofcoinPr-131]
	_ = x[Sctp-132]
	_ = x[DagJose-133]
	_ = x[DagCose-134]
	_ = x[EthBlock-144]
	_ = x[EthBlockList-145]
	_ = x[EthTxTrie-146]
	_ = x[EthTx-147]
	_ = x[EthTxReceiptTrie-148]
	_ = x[EthTxReceipt-149]
	_ = x[EthStateTrie-150]
	_ = x[EthAccountSnapshot-151]
	_ = x[EthStorageTrie-152]
	_ = x[EthReceiptLogTrie-153]
	_ = x[EthRecieptLog-154]
	_ = x[Aes128-160]
	_ = x[Aes192-161]
	_ = x[Aes256-162]
	_ = x[Chacha128-163]
	_ = x[Chacha256-164]
	_ = x[BitcoinBlock-176]
	_ = x[BitcoinTx-177]
	_ = x[BitcoinWitnessCommitment-178]
	_ = x[ZcashBlock-192]
	_ = x[ZcashTx-193]
	_ = x[Caip50-202]
	_ = x[Streamid-206]
	_ = x[StellarBlock-208]
	_ = x[StellarTx-209]
	_ = x[Md4-212]
	_ = x[Md5-213]
	_ = x[DecredBlock-224]
	_ = x[DecredTx-225]
	_ = x[Ipld-226]
	_ = x[Ipfs-227]
	_ = x[Swarm-228]
	_ = x[Ipns-229]
	_ = x[Zeronet-230]
	_ = x[Secp256k1Pub-231]
	_ = x[Dnslink-232]
	_ = x[Bls12_381G1Pub-234]
	_ = x[Bls12_381G2Pub-235]
	_ = x[X25519Pub-236]
	_ = x[Ed25519Pub-237]
	_ = x[Bls12_381G1g2Pub-238]
	_ = x[Sr25519Pub-239]
	_ = x[DashBlock-240]
	_ = x[DashTx-241]
	_ = x[SwarmManifest-250]
	_ = x[SwarmFeed-251]
	_ = x[Beeson-252]
	_ = x[Udp-273]
	_ = x[P2pWebrtcStar-275]
	_ = x[P2pWebrtcDirect-276]
	_ = x[P2pStardust-277]
	_ = x[WebrtcDirect-280]
	_ = x[Webrtc-281]
	_ = x[P2pCircuit-290]
	_ = x[DagJson-297]
	_ = x[Udt-301]
	_ = x[Utp-302]
	_ = x[Crc32-306]
	_ = x[Crc64Ecma-356]
	_ = x[Unix-400]
	_ = x[Thread-406]
	_ = x[P2p-421]
	_ = x[Https-443]
	_ = x[Onion-444]
	_ = x[Onion3-445]
	_ = x[Garlic64-446]
	_ = x[Garlic32-447]
	_ = x[Tls-448]
	_ = x[Sni-449]
	_ = x[Noise-454]
	_ = x[Quic-460]
	_ = x[QuicV1-461]
	_ = x[Webtransport-465]
	_ = x[Certhash-466]
	_ = x[Ws-477]
	_ = x[Wss-478]
	_ = x[P2pWebsocketStar-479]
	_ = x[Http-480]
	_ = x[Swhid1Snp-496]
	_ = x[Json-512]
	_ = x[Messagepack-513]
	_ = x[Car-514]
	_ = x[IpnsRecord-768]
	_ = x[Libp2pPeerRecord-769]
	_ = x[Libp2pRelayRsvp-770]
	_ = x[Memorytransport-777]
	_ = x[CarIndexSorted-1024]
	_ = x[CarMultihashIndexSorted-1025]
	_ = x[TransportBitswap-2304]
	_ = x[TransportGraphsyncFilecoinv1-2320]
	_ = x[TransportIpfsGatewayHttp-2336]
	_ = x[Multidid-3357]
	_ = x[Sha2_256Trunc254Padded-4114]
	_ = x[Sha2_224-4115]
	_ = x[Sha2_512_224-4116]
	_ = x[Sha2_512_256-4117]
	_ = x[Murmur3X64_128-4130]
	_ = x[Ripemd128-4178]
	_ = x[Ripemd160-4179]
	_ = x[Ripemd256-4180]
	_ = x[Ripemd320-4181]
	_ = x[X11-4352]
	_ = x[P256Pub-4608]
	_ = x[P384Pub-4609]
	_ = x[P521Pub-4610]
	_ = x[Ed448Pub-4611]
	_ = x[X448Pub-4612]
	_ = x[RsaPub-4613]
	_ = x[Sm2Pub-4614]
	_ = x[Ed25519Priv-4864]
	_ = x[Secp256k1Priv-4865]
	_ = x[X25519Priv-4866]
	_ = x[Sr25519Priv-4867]
	_ = x[RsaPriv-4869]
	_ = x[P256Priv-4870]
	_ = x[P384Priv-4871]
	_ = x[P521Priv-4872]
	_ = x[Kangarootwelve-7425]
	_ = x[AesGcm256-8192]
	_ = x[Silverpine-16194]
	_ = x[Sm3_256-21325]
	_ = x[Blake2b8-45569]
	_ = x[Blake2b16-45570]
	_ = x[Blake2b24-45571]
	_ = x[Blake2b32-45572]
	_ = x[Blake2b40-45573]
	_ = x[Blake2b48-45574]
	_ = x[Blake2b56-45575]
	_ = x[Blake2b64-45576]
	_ = x[Blake2b72-45577]
	_ = x[Blake2b80-45578]
	_ = x[Blake2b88-45579]
	_ = x[Blake2b96-45580]
	_ = x[Blake2b104-45581]
	_ = x[Blake2b112-45582]
	_ = x[Blake2b120-45583]
	_ = x[Blake2b128-45584]
	_ = x[Blake2b136-45585]
	_ = x[Blake2b144-45586]
	_ = x[Blake2b152-45587]
	_ = x[Blake2b160-45588]
	_ = x[Blake2b168-45589]
	_ = x[Blake2b176-45590]
	_ = x[Blake2b184-45591]
	_ = x[Blake2b192-45592]
	_ = x[Blake2b200-45593]
	_ = x[Blake2b208-45594]
	_ = x[Blake2b216-45595]
	_ = x[Blake2b224-45596]
	_ = x[Blake2b232-45597]
	_ = x[Blake2b240-45598]
	_ = x[Blake2b248-45599]
	_ = x[Blake2b256-45600]
	_ = x[Blake2b264-45601]
	_ = x[Blake2b272-45602]
	_ = x[Blake2b280-45603]
	_ = x[Blake2b288-45604]
	_ = x[Blake2b296-45605]
	_ = x[Blake2b304-45606]
	_ = x[Blake2b312-45607]
	_ = x[Blake2b320-45608]
	_ = x[Blake2b328-45609]
	_ = x[Blake2b336-45610]
	_ = x[Blake2b344-45611]
	_ = x[Blake2b352-45612]
	_ = x[Blake2b360-45613]
	_ = x[Blake2b368-45614]
	_ = x[Blake2b376-45615]
	_ = x[Blake2b384-45616]
	_ = x[Blake2b392-45617]
	_ = x[Blake2b400-45618]
	_ = x[Blake2b408-45619]
	_ = x[Blake2b416-45620]
	_ = x[Blake2b424-45621]
	_ = x[Blake2b432-45622]
	_ = x[Blake2b440-45623]
	_ = x[Blake2b448-45624]
	_ = x[Blake2b456-45625]
	_ = x[Blake2b464-45626]
	_ = x[Blake2b472-45627]
	_ = x[Blake2b480-45628]
	_ = x[Blake2b488-45629]
	_ = x[Blake2b496-45630]
	_ = x[Blake2b504-45631]
	_ = x[Blake2b512-45632]
	_ = x[Blake2s8-45633]
	_ = x[Blake2s16-45634]
	_ = x[Blake2s24-45635]
	_ = x[Blake2s32-45636]
	_ = x[Blake2s40-45637]
	_ = x[Blake2s48-45638]
	_ = x[Blake2s56-45639]
	_ = x[Blake2s64-45640]
	_ = x[Blake2s72-45641]
	_ = x[Blake2s80-45642]
	_ = x[Blake2s88-45643]
	_ = x[Blake2s96-45644]
	_ = x[Blake2s104-45645]
	_ = x[Blake2s112-45646]
	_ = x[Blake2s120-45647]
	_ = x[Blake2s128-45648]
	_ = x[Blake2s136-45649]
	_ = x[Blake2s144-45650]
	_ = x[Blake2s152-45651]
	_ = x[Blake2s160-45652]
	_ = x[Blake2s168-45653]
	_ = x[Blake2s176-45654]
	_ = x[Blake2s184-45655]
	_ = x[Blake2s192-45656]
	_ = x[Blake2s200-45657]
	_ = x[Blake2s208-45658]
	_ = x[Blake2s216-45659]
	_ = x[Blake2s224-45660]
	_ = x[Blake2s232-45661]
	_ = x[Blake2s240-45662]
	_ = x[Blake2s248-45663]
	_ = x[Blake2s256-45664]
	_ = x[Skein256_8-45825]
	_ = x[Skein256_16-45826]
	_ = x[Skein256_24-45827]
	_ = x[Skein256_32-45828]
	_ = x[Skein256_40-45829]
	_ = x[Skein256_48-45830]
	_ = x[Skein256_56-45831]
	_ = x[Skein256_64-45832]
	_ = x[Skein256_72-45833]
	_ = x[Skein256_80-45834]
	_ = x[Skein256_88-45835]
	_ = x[Skein256_96-45836]
	_ = x[Skein256_104-45837]
	_ = x[Skein256_112-45838]
	_ = x[Skein256_120-45839]
	_ = x[Skein256_128-45840]
	_ = x[Skein256_136-45841]
	_ = x[Skein256_144-45842]
	_ = x[Skein256_152-45843]
	_ = x[Skein256_160-45844]
	_ = x[Skein256_168-45845]
	_ = x[Skein256_176-45846]
	_ = x[Skein256_184-45847]
	_ = x[Skein256_192-45848]
	_ = x[Skein256_200-45849]
	_ = x[Skein256_208-45850]
	_ = x[Skein256_216-45851]
	_ = x[Skein256_224-45852]
	_ = x[Skein256_232-45853]
	_ = x[Skein256_240-45854]
	_ = x[Skein256_248-45855]
	_ = x[Skein256_256-45856]
	_ = x[Skein512_8-45857]
	_ = x[Skein512_16-45858]
	_ = x[Skein512_24-45859]
	_ = x[Skein512_32-45860]
	_ = x[Skein512_40-45861]
	_ = x[Skein512_48-45862]
	_ = x[Skein512_56-45863]
	_ = x[Skein512_64-45864]
	_ = x[Skein512_72-45865]
	_ = x[Skein512_80-45866]
	_ = x[Skein512_88-45867]
	_ = x[Skein512_96-45868]
	_ = x[Skein512_104-45869]
	_ = x[Skein512_112-45870]
	_ = x[Skein512_120-45871]
	_ = x[Skein512_128-45872]
	_ = x[Skein512_136-45873]
	_ = x[Skein512_144-45874]
	_ = x[Skein512_152-45875]
	_ = x[Skein512_160-45876]
	_ = x[Skein512_168-45877]
	_ = x[Skein512_176-45878]
	_ = x[Skein512_184-45879]
	_ = x[Skein512_192-45880]
	_ = x[Skein512_200-45881]
	_ = x[Skein512_208-45882]
	_ = x[Skein512_216-45883]
	_ = x[Skein512_224-45884]
	_ = x[Skein512_232-45885]
	_ = x[Skein512_240-45886]
	_ = x[Skein512_248-45887]
	_ = x[Skein512_256-45888]
	_ = x[Skein512_264-45889]
	_ = x[Skein512_272-45890]
	_ = x[Skein512_280-45891]
	_ = x[Skein512_288-45892]
	_ = x[Skein512_296-45893]
	_ = x[Skein512_304-45894]
	_ = x[Skein512_312-45895]
	_ = x[Skein512_320-45896]
	_ = x[Skein512_328-45897]
	_ = x[Skein512_336-45898]
	_ = x[Skein512_344-45899]
	_ = x[Skein512_352-45900]
	_ = x[Skein512_360-45901]
	_ = x[Skein512_368-45902]
	_ = x[Skein512_376-45903]
	_ = x[Skein512_384-45904]
	_ = x[Skein512_392-45905]
	_ = x[Skein512_400-45906]
	_ = x[Skein512_408-45907]
	_ = x[Skein512_416-45908]
	_ = x[Skein512_424-45909]
	_ = x[Skein512_432-45910]
	_ = x[Skein512_440-45911]
	_ = x[Skein512_448-45912]
	_ = x[Skein512_456-45913]
	_ = x[Skein512_464-45914]
	_ = x[Skein512_472-45915]
	_ = x[Skein512_480-45916]
	_ = x[Skein512_488-45917]
	_ = x[Skein512_496-45918]
	_ = x[Skein512_504-45919]
	_ = x[Skein512_512-45920]
	_ = x[Skein1024_8-45921]
	_ = x[Skein1024_16-45922]
	_ = x[Skein1024_24-45923]
	_ = x[Skein1024_32-45924]
	_ = x[Skein1024_40-45925]
	_ = x[Skein1024_48-45926]
	_ = x[Skein1024_56-45927]
	_ = x[Skein1024_64-45928]
	_ = x[Skein1024_72-45929]
	_ = x[Skein1024_80-45930]
	_ = x[Skein1024_88-45931]
	_ = x[Skein1024_96-45932]
	_ = x[Skein1024_104-45933]
	_ = x[Skein1024_112-45934]
	_ = x[Skein1024_120-45935]
	_ = x[Skein1024_128-45936]
	_ = x[Skein1024_136-45937]
	_ = x[Skein1024_144-45938]
	_ = x[Skein1024_152-45939]
	_ = x[Skein1024_160-45940]
	_ = x[Skein1024_168-45941]
	_ = x[Skein1024_176-45942]
	_ = x[Skein1024_184-45943]
	_ = x[Skein1024_192-45944]
	_ = x[Skein1024_200-45945]
	_ = x[Skein1024_208-45946]
	_ = x[Skein1024_216-45947]
	_ = x[Skein1024_224-45948]
	_ = x[Skein1024_232-45949]
	_ = x[Skein1024_240-45950]
	_ = x[Skein1024_248-45951]
	_ = x[Skein1024_256-45952]
	_ = x[Skein1024_264-45953]
	_ = x[Skein1024_272-45954]
	_ = x[Skein1024_280-45955]
	_ = x[Skein1024_288-45956]
	_ = x[Skein1024_296-45957]
	_ = x[Skein1024_304-45958]
	_ = x[Skein1024_312-45959]
	_ = x[Skein1024_320-45960]
	_ = x[Skein1024_328-45961]
	_ = x[Skein1024_336-45962]
	_ = x[Skein1024_344-45963]
	_ = x[Skein1024_352-45964]
	_ = x[Skein1024_360-45965]
	_ = x[Skein1024_368-45966]
	_ = x[Skein1024_376-45967]
	_ = x[Skein1024_384-45968]
	_ = x[Skein1024_392-45969]
	_ = x[Skein1024_400-45970]
	_ = x[Skein1024_408-45971]
	_ = x[Skein1024_416-45972]
	_ = x[Skein1024_424-45973]
	_ = x[Skein1024_432-45974]
	_ = x[Skein1024_440-45975]
	_ = x[Skein1024_448-45976]
	_ = x[Skein1024_456-45977]
	_ = x[Skein1024_464-45978]
	_ = x[Skein1024_472-45979]
	_ = x[Skein1024_480-45980]
	_ = x[Skein1024_488-45981]
	_ = x[Skein1024_496-45982]
	_ = x[Skein1024_504-45983]
	_ = x[Skein1024_512-45984]
	_ = x[Skein1024_520-45985]
	_ = x[Skein1024_528-45986]
	_ = x[Skein1024_536-45987]
	_ = x[Skein1024_544-45988]
	_ = x[Skein1024_552-45989]
	_ = x[Skein1024_560-45990]
	_ = x[Skein1024_568-45991]
	_ = x[Skein1024_576-45992]
	_ = x[Skein1024_584-45993]
	_ = x[Skein1024_592-45994]
	_ = x[Skein1024_600-45995]
	_ = x[Skein1024_608-45996]
	_ = x[Skein1024_616-45997]
	_ = x[Skein1024_624-45998]
	_ = x[Skein1024_632-45999]
	_ = x[Skein1024_640-46000]
	_ = x[Skein1024_648-46001]
	_ = x[Skein1024_656-46002]
	_ = x[Skein1024_664-46003]
	_ = x[Skein1024_672-46004]
	_ = x[Skein1024_680-46005]
	_ = x[Skein1024_688-46006]
	_ = x[Skein1024_696-46007]
	_ = x[Skein1024_704-46008]
	_ = x[Skein1024_712-46009]
	_ = x[Skein1024_720-46010]
	_ = x[Skein1024_728-46011]
	_ = x[Skein1024_736-46012]
	_ = x[Skein1024_744-46013]
	_ = x[Skein1024_752-46014]
	_ = x[Skein1024_760-46015]
	_ = x[Skein1024_768-46016]
	_ = x[Skein1024_776-46017]
	_ = x[Skein1024_784-46018]
	_ = x[Skein1024_792-46019]
	_ = x[Skein1024_800-46020]
	_ = x[Skein1024_808-46021]
	_ = x[Skein1024_816-46022]
	_ = x[Skein1024_824-46023]
	_ = x[Skein1024_832-46024]
	_ = x[Skein1024_840-46025]
	_ = x[Skein1024_848-46026]
	_ = x[Skein1024_856-46027]
	_ = x[Skein1024_864-46028]
	_ = x[Skein1024_872-46029]
	_ = x[Skein1024_880-46030]
	_ = x[Skein1024_888-46031]
	_ = x[Skein1024_896-46032]
	_ = x[Skein1024_904-46033]
	_ = x[Skein1024_912-46034]
	_ = x[Skein1024_920-46035]
	_ = x[Skein1024_928-46036]
	_ = x[Skein1024_936-46037]
	_ = x[Skein1024_944-46038]
	_ = x[Skein1024_952-46039]
	_ = x[Skein1024_960-46040]
	_ = x[Skein1024_968-46041]
	_ = x[Skein1024_976-46042]
	_ = x[Skein1024_984-46043]
	_ = x[Skein1024_992-46044]
	_ = x[Skein1024_1000-46045]
	_ = x[Skein1024_1008-46046]
	_ = x[Skein1024_1016-46047]
	_ = x[Skein1024_1024-46048]
	_ = x[Xxh32-46049]
	_ = x[Xxh64-46050]
	_ = x[Xxh3_64-46051]
	_ = x[Xxh3_128-46052]
	_ = x[PoseidonBls12_381A2Fc1-46081]
	_ = x[PoseidonBls12_381A2Fc1Sc-46082]
	_ = x[Urdca2015Canon-46083]
	_ = x[Ssz-46337]
	_ = x[SszSha2_256Bmt-46338]
	_ = x[JsonJcs-46593]
	_ = x[Iscc-52225]
	_ = x[ZeroxcertImprint256-52753]
	_ = x[Varsig-53248]
	_ = x[Es256k-53479]
	_ = x[Bls12381G1Sig-53482]
	_ = x[Bls12381G2Sig-53483]
	_ = x[Eddsa-53485]
	_ = x[Eip191-53649]
	_ = x[Jwk_jcsPub-60241]
	_ = x[FilCommitmentUnsealed-61697]
	_ = x[FilCommitmentSealed-61698]
	_ = x[Plaintextv2-7367777]
	_ = x[HolochainAdrV0-8417572]
	_ = x[HolochainAdrV1-8483108]
	_ = x[HolochainKeyV0-9728292]
	_ = x[HolochainKeyV1-9793828]
	_ = x[HolochainSigV0-10645796]
	_ = x[HolochainSigV1-10711332]
	_ = x[SkynetNs-11639056]
	_ = x[ArweaveNs-11704592]
	_ = x[SubspaceNs-11770128]
	_ = x[KumandraNs-11835664]
	_ = x[Es256-13636096]
	_ = x[Es284-13636097]
	_ = x[Es512-13636098]
	_ = x[Rs256-13636101]
}

const _Code_name = "identitycidv1cidv2cidv3ip4tcpsha1sha2-256sha2-512sha3-512sha3-384sha3-256sha3-224shake-128shake-256keccak-224keccak-256keccak-384keccak-512blake3sha2-384dccpmurmur3-x64-64murmur3-32ip6ip6zoneipcidrpathmulticodecmultihashmultiaddrmultibasednsdns4dns6dnsaddrprotobufcborrawdbl-sha2-256rlpbencodedag-pbdag-cborlibp2p-keygit-rawtorrent-infotorrent-fileleofcoin-blockleofcoin-txleofcoin-prsctpdag-josedag-coseeth-blocketh-block-listeth-tx-trieeth-txeth-tx-receipt-trieeth-tx-receipteth-state-trieeth-account-snapshoteth-storage-trieeth-receipt-log-trieeth-reciept-logaes-128aes-192aes-256chacha-128chacha-256bitcoin-blockbitcoin-txbitcoin-witness-commitmentzcash-blockzcash-txcaip-50streamidstellar-blockstellar-txmd4md5decred-blockdecred-txipldipfsswarmipnszeronetsecp256k1-pubdnslinkbls12_381-g1-pubbls12_381-g2-pubx25519-pubed25519-pubbls12_381-g1g2-pubsr25519-pubdash-blockdash-txswarm-manifestswarm-feedbeesonudpp2p-webrtc-starp2p-webrtc-directp2p-stardustwebrtc-directwebrtcp2p-circuitdag-jsonudtutpcrc32crc64-ecmaunixthreadp2phttpsoniononion3garlic64garlic32tlssninoisequicquic-v1webtransportcerthashwswssp2p-websocket-starhttpswhid-1-snpjsonmessagepackcaripns-recordlibp2p-peer-recordlibp2p-relay-rsvpmemorytransportcar-index-sortedcar-multihash-index-sortedtransport-bitswaptransport-graphsync-filecoinv1transport-ipfs-gateway-httpmultididsha2-256-trunc254-paddedsha2-224sha2-512-224sha2-512-256murmur3-x64-128ripemd-128ripemd-160ripemd-256ripemd-320x11p256-pubp384-pubp521-pubed448-pubx448-pubrsa-pubsm2-pubed25519-privsecp256k1-privx25519-privsr25519-privrsa-privp256-privp384-privp521-privkangarootwelveaes-gcm-256silverpinesm3-256blake2b-8blake2b-16blake2b-24blake2b-32blake2b-40blake2b-48blake2b-56blake2b-64blake2b-72blake2b-80blake2b-88blake2b-96blake2b-104blake2b-112blake2b-120blake2b-128blake2b-136blake2b-144blake2b-152blake2b-160blake2b-168blake2b-176blake2b-184blake2b-192blake2b-200blake2b-208blake2b-216blake2b-224blake2b-232blake2b-240blake2b-248blake2b-256blake2b-264blake2b-272blake2b-280blake2b-288blake2b-296blake2b-304blake2b-312blake2b-320blake2b-328blake2b-336blake2b-344blake2b-352blake2b-360blake2b-368blake2b-376blake2b-384blake2b-392blake2b-400blake2b-408blake2b-416blake2b-424blake2b-432blake2b-440blake2b-448blake2b-456blake2b-464blake2b-472blake2b-480blake2b-488blake2b-496blake2b-504blake2b-512blake2s-8blake2s-16blake2s-24blake2s-32blake2s-40blake2s-48blake2s-56blake2s-64blake2s-72blake2s-80blake2s-88blake2s-96blake2s-104blake2s-112blake2s-120blake2s-128blake2s-136blake2s-144blake2s-152blake2s-160blake2s-168blake2s-176blake2s-184blake2s-192blake2s-200blake2s-208blake2s-216blake2s-224blake2s-232blake2s-240blake2s-248blake2s-256skein256-8skein256-16skein256-24skein256-32skein256-40skein256-48skein256-56skein256-64skein256-72skein256-80skein256-88skein256-96skein256-104skein256-112skein256-120skein256-128skein256-136skein256-144skein256-152skein256-160skein256-168skein256-176skein256-184skein256-192skein256-200skein256-208skein256-216skein256-224skein256-232skein256-240skein256-248skein256-256skein512-8skein512-16skein512-24skein512-32skein512-40skein512-48skein512-56skein512-64skein512-72skein512-80skein512-88skein512-96skein512-104skein512-112skein512-120skein512-128skein512-136skein512-144skein512-152skein512-160skein512-168skein512-176skein512-184skein512-192skein512-200skein512-208skein512-216skein512-224skein512-232skein512-240skein512-248skein512-256skein512-264skein512-272skein512-280skein512-288skein512-296skein512-304skein512-312skein512-320skein512-328skein512-336skein512-344skein512-352skein512-360skein512-368skein512-376skein512-384skein512-392skein512-400skein512-408skein512-416skein512-424skein512-432skein512-440skein512-448skein512-456skein512-464skein512-472skein512-480skein512-488skein512-496skein512-504skein512-512skein1024-8skein1024-16skein1024-24skein1024-32skein1024-40skein1024-48skein1024-56skein1024-64skein1024-72skein1024-80skein1024-88skein1024-96skein1024-104skein1024-112skein1024-120skein1024-128skein1024-136skein1024-144skein1024-152skein1024-160skein1024-168skein1024-176skein1024-184skein1024-192skein1024-200skein1024-208skein1024-216skein1024-224skein1024-232skein1024-240skein1024-248skein1024-256skein1024-264skein1024-272skein1024-280skein1024-288skein1024-296skein1024-304skein1024-312skein1024-320skein1024-328skein1024-336skein1024-344skein1024-352skein1024-360skein1024-368skein1024-376skein1024-384skein1024-392skein1024-400skein1024-408skein1024-416skein1024-424skein1024-432skein1024-440skein1024-448skein1024-456skein1024-464skein1024-472skein1024-480skein1024-488skein1024-496skein1024-504skein1024-512skein1024-520skein1024-528skein1024-536skein1024-544skein1024-552skein1024-560skein1024-568skein1024-576skein1024-584skein1024-592skein1024-600skein1024-608skein1024-616skein1024-624skein1024-632skein1024-640skein1024-648skein1024-656skein1024-664skein1024-672skein1024-680skein1024-688skein1024-696skein1024-704skein1024-712skein1024-720skein1024-728skein1024-736skein1024-744skein1024-752skein1024-760skein1024-768skein1024-776skein1024-784skein1024-792skein1024-800skein1024-808skein1024-816skein1024-824skein1024-832skein1024-840skein1024-848skein1024-856skein1024-864skein1024-872skein1024-880skein1024-888skein1024-896skein1024-904skein1024-912skein1024-920skein1024-928skein1024-936skein1024-944skein1024-952skein1024-960skein1024-968skein1024-976skein1024-984skein1024-992skein1024-1000skein1024-1008skein1024-1016skein1024-1024xxh-32xxh-64xxh3-64xxh3-128poseidon-bls12_381-a2-fc1poseidon-bls12_381-a2-fc1-scurdca-2015-canonsszssz-sha2-256-bmtjson-jcsiscczeroxcert-imprint-256varsiges256kbls-12381-g1-sigbls-12381-g2-sigeddsaeip-191jwk_jcs-pubfil-commitment-unsealedfil-commitment-sealedplaintextv2holochain-adr-v0holochain-adr-v1holochain-key-v0holochain-key-v1holochain-sig-v0holochain-sig-v1skynet-nsarweave-nssubspace-nskumandra-nses256es284es512rs256"

var _Code_map = map[Code]string{
	0:        _Code_name[0:8],
	1:        _Code_name[8:13],
	2:        _Code_name[13:18],
	3:        _Code_name[18:23],
	4:        _Code_name[23:26],
	6:        _Code_name[26:29],
	17:       _Code_name[29:33],
	18:       _Code_name[33:41],
	19:       _Code_name[41:49],
	20:       _Code_name[49:57],
	21:       _Code_name[57:65],
	22:       _Code_name[65:73],
	23:       _Code_name[73:81],
	24:       _Code_name[81:90],
	25:       _Code_name[90:99],
	26:       _Code_name[99:109],
	27:       _Code_name[109:119],
	28:       _Code_name[119:129],
	29:       _Code_name[129:139],
	30:       _Code_name[139:145],
	32:       _Code_name[145:153],
	33:       _Code_name[153:157],
	34:       _Code_name[157:171],
	35:       _Code_name[171:181],
	41:       _Code_name[181:184],
	42:       _Code_name[184:191],
	43:       _Code_name[191:197],
	47:       _Code_name[197:201],
	48:       _Code_name[201:211],
	49:       _Code_name[211:220],
	50:       _Code_name[220:229],
	51:       _Code_name[229:238],
	53:       _Code_name[238:241],
	54:       _Code_name[241:245],
	55:       _Code_name[245:249],
	56:       _Code_name[249:256],
	80:       _Code_name[256:264],
	81:       _Code_name[264:268],
	85:       _Code_name[268:271],
	86:       _Code_name[271:283],
	96:       _Code_name[283:286],
	99:       _Code_name[286:293],
	112:      _Code_name[293:299],
	113:      _Code_name[299:307],
	114:      _Code_name[307:317],
	120:      _Code_name[317:324],
	123:      _Code_name[324:336],
	124:      _Code_name[336:348],
	129:      _Code_name[348:362],
	130:      _Code_name[362:373],
	131:      _Code_name[373:384],
	132:      _Code_name[384:388],
	133:      _Code_name[388:396],
	134:      _Code_name[396:404],
	144:      _Code_name[404:413],
	145:      _Code_name[413:427],
	146:      _Code_name[427:438],
	147:      _Code_name[438:444],
	148:      _Code_name[444:463],
	149:      _Code_name[463:477],
	150:      _Code_name[477:491],
	151:      _Code_name[491:511],
	152:      _Code_name[511:527],
	153:      _Code_name[527:547],
	154:      _Code_name[547:562],
	160:      _Code_name[562:569],
	161:      _Code_name[569:576],
	162:      _Code_name[576:583],
	163:      _Code_name[583:593],
	164:      _Code_name[593:603],
	176:      _Code_name[603:616],
	177:      _Code_name[616:626],
	178:      _Code_name[626:652],
	192:      _Code_name[652:663],
	193:      _Code_name[663:671],
	202:      _Code_name[671:678],
	206:      _Code_name[678:686],
	208:      _Code_name[686:699],
	209:      _Code_name[699:709],
	212:      _Code_name[709:712],
	213:      _Code_name[712:715],
	224:      _Code_name[715:727],
	225:      _Code_name[727:736],
	226:      _Code_name[736:740],
	227:      _Code_name[740:744],
	228:      _Code_name[744:749],
	229:      _Code_name[749:753],
	230:      _Code_name[753:760],
	231:      _Code_name[760:773],
	232:      _Code_name[773:780],
	234:      _Code_name[780:796],
	235:      _Code_name[796:812],
	236:      _Code_name[812:822],
	237:      _Code_name[822:833],
	238:      _Code_name[833:851],
	239:      _Code_name[851:862],
	240:      _Code_name[862:872],
	241:      _Code_name[872:879],
	250:      _Code_name[879:893],
	251:      _Code_name[893:903],
	252:      _Code_name[903:909],
	273:      _Code_name[909:912],
	275:      _Code_name[912:927],
	276:      _Code_name[927:944],
	277:      _Code_name[944:956],
	280:      _Code_name[956:969],
	281:      _Code_name[969:975],
	290:      _Code_name[975:986],
	297:      _Code_name[986:994],
	301:      _Code_name[994:997],
	302:      _Code_name[997:1000],
	306:      _Code_name[1000:1005],
	356:      _Code_name[1005:1015],
	400:      _Code_name[1015:1019],
	406:      _Code_name[1019:1025],
	421:      _Code_name[1025:1028],
	443:      _Code_name[1028:1033],
	444:      _Code_name[1033:1038],
	445:      _Code_name[1038:1044],
	446:      _Code_name[1044:1052],
	447:      _Code_name[1052:1060],
	448:      _Code_name[1060:1063],
	449:      _Code_name[1063:1066],
	454:      _Code_name[1066:1071],
	460:      _Code_name[1071:1075],
	461:      _Code_name[1075:1082],
	465:      _Code_name[1082:1094],
	466:      _Code_name[1094:1102],
	477:      _Code_name[1102:1104],
	478:      _Code_name[1104:1107],
	479:      _Code_name[1107:1125],
	480:      _Code_name[1125:1129],
	496:      _Code_name[1129:1140],
	512:      _Code_name[1140:1144],
	513:      _Code_name[1144:1155],
	514:      _Code_name[1155:1158],
	768:      _Code_name[1158:1169],
	769:      _Code_name[1169:1187],
	770:      _Code_name[1187:1204],
	777:      _Code_name[1204:1219],
	1024:     _Code_name[1219:1235],
	1025:     _Code_name[1235:1261],
	2304:     _Code_name[1261:1278],
	2320:     _Code_name[1278:1308],
	2336:     _Code_name[1308:1335],
	3357:     _Code_name[1335:1343],
	4114:     _Code_name[1343:1367],
	4115:     _Code_name[1367:1375],
	4116:     _Code_name[1375:1387],
	4117:     _Code_name[1387:1399],
	4130:     _Code_name[1399:1414],
	4178:     _Code_name[1414:1424],
	4179:     _Code_name[1424:1434],
	4180:     _Code_name[1434:1444],
	4181:     _Code_name[1444:1454],
	4352:     _Code_name[1454:1457],
	4608:     _Code_name[1457:1465],
	4609:     _Code_name[1465:1473],
	4610:     _Code_name[1473:1481],
	4611:     _Code_name[1481:1490],
	4612:     _Code_name[1490:1498],
	4613:     _Code_name[1498:1505],
	4614:     _Code_name[1505:1512],
	4864:     _Code_name[1512:1524],
	4865:     _Code_name[1524:1538],
	4866:     _Code_name[1538:1549],
	4867:     _Code_name[1549:1561],
	4869:     _Code_name[1561:1569],
	4870:     _Code_name[1569:1578],
	4871:     _Code_name[1578:1587],
	4872:     _Code_name[1587:1596],
	7425:     _Code_name[1596:1610],
	8192:     _Code_name[1610:1621],
	16194:    _Code_name[1621:1631],
	21325:    _Code_name[1631:1638],
	45569:    _Code_name[1638:1647],
	45570:    _Code_name[1647:1657],
	45571:    _Code_name[1657:1667],
	45572:    _Code_name[1667:1677],
	45573:    _Code_name[1677:1687],
	45574:    _Code_name[1687:1697],
	45575:    _Code_name[1697:1707],
	45576:    _Code_name[1707:1717],
	45577:    _Code_name[1717:1727],
	45578:    _Code_name[1727:1737],
	45579:    _Code_name[1737:1747],
	45580:    _Code_name[1747:1757],
	45581:    _Code_name[1757:1768],
	45582:    _Code_name[1768:1779],
	45583:    _Code_name[1779:1790],
	45584:    _Code_name[1790:1801],
	45585:    _Code_name[1801:1812],
	45586:    _Code_name[1812:1823],
	45587:    _Code_name[1823:1834],
	45588:    _Code_name[1834:1845],
	45589:    _Code_name[1845:1856],
	45590:    _Code_name[1856:1867],
	45591:    _Code_name[1867:1878],
	45592:    _Code_name[1878:1889],
	45593:    _Code_name[1889:1900],
	45594:    _Code_name[1900:1911],
	45595:    _Code_name[1911:1922],
	45596:    _Code_name[1922:1933],
	45597:    _Code_name[1933:1944],
	45598:    _Code_name[1944:1955],
	45599:    _Code_name[1955:1966],
	45600:    _Code_name[1966:1977],
	45601:    _Code_name[1977:1988],
	45602:    _Code_name[1988:1999],
	45603:    _Code_name[1999:2010],
	45604:    _Code_name[2010:2021],
	45605:    _Code_name[2021:2032],
	45606:    _Code_name[2032:2043],
	45607:    _Code_name[2043:2054],
	45608:    _Code_name[2054:2065],
	45609:    _Code_name[2065:2076],
	45610:    _Code_name[2076:2087],
	45611:    _Code_name[2087:2098],
	45612:    _Code_name[2098:2109],
	45613:    _Code_name[2109:2120],
	45614:    _Code_name[2120:2131],
	45615:    _Code_name[2131:2142],
	45616:    _Code_name[2142:2153],
	45617:    _Code_name[2153:2164],
	45618:    _Code_name[2164:2175],
	45619:    _Code_name[2175:2186],
	45620:    _Code_name[2186:2197],
	45621:    _Code_name[2197:2208],
	45622:    _Code_name[2208:2219],
	45623:    _Code_name[2219:2230],
	45624:    _Code_name[2230:2241],
	45625:    _Code_name[2241:2252],
	45626:    _Code_name[2252:2263],
	45627:    _Code_name[2263:2274],
	45628:    _Code_name[2274:2285],
	45629:    _Code_name[2285:2296],
	45630:    _Code_name[2296:2307],
	45631:    _Code_name[2307:2318],
	45632:    _Code_name[2318:2329],
	45633:    _Code_name[2329:2338],
	45634:    _Code_name[2338:2348],
	45635:    _Code_name[2348:2358],
	45636:    _Code_name[2358:2368],
	45637:    _Code_name[2368:2378],
	45638:    _Code_name[2378:2388],
	45639:    _Code_name[2388:2398],
	45640:    _Code_name[2398:2408],
	45641:    _Code_name[2408:2418],
	45642:    _Code_name[2418:2428],
	45643:    _Code_name[2428:2438],
	45644:    _Code_name[2438:2448],
	45645:    _Code_name[2448:2459],
	45646:    _Code_name[2459:2470],
	45647:    _Code_name[2470:2481],
	45648:    _Code_name[2481:2492],
	45649:    _Code_name[2492:2503],
	45650:    _Code_name[2503:2514],
	45651:    _Code_name[2514:2525],
	45652:    _Code_name[2525:2536],
	45653:    _Code_name[2536:2547],
	45654:    _Code_name[2547:2558],
	45655:    _Code_name[2558:2569],
	45656:    _Code_name[2569:2580],
	45657:    _Code_name[2580:2591],
	45658:    _Code_name[2591:2602],
	45659:    _Code_name[2602:2613],
	45660:    _Code_name[2613:2624],
	45661:    _Code_name[2624:2635],
	45662:    _Code_name[2635:2646],
	45663:    _Code_name[2646:2657],
	45664:    _Code_name[2657:2668],
	45825:    _Code_name[2668:2678],
	45826:    _Code_name[2678:2689],
	45827:    _Code_name[2689:2700],
	45828:    _Code_name[2700:2711],
	45829:    _Code_name[2711:2722],
	45830:    _Code_name[2722:2733],
	45831:    _Code_name[2733:2744],
	45832:    _Code_name[2744:2755],
	45833:    _Code_name[2755:2766],
	45834:    _Code_name[2766:2777],
	45835:    _Code_name[2777:2788],
	45836:    _Code_name[2788:2799],
	45837:    _Code_name[2799:2811],
	45838:    _Code_name[2811:2823],
	45839:    _Code_name[2823:2835],
	45840:    _Code_name[2835:2847],
	45841:    _Code_name[2847:2859],
	45842:    _Code_name[2859:2871],
	45843:    _Code_name[2871:2883],
	45844:    _Code_name[2883:2895],
	45845:    _Code_name[2895:2907],
	45846:    _Code_name[2907:2919],
	45847:    _Code_name[2919:2931],
	45848:    _Code_name[2931:2943],
	45849:    _Code_name[2943:2955],
	45850:    _Code_name[2955:2967],
	45851:    _Code_name[2967:2979],
	45852:    _Code_name[2979:2991],
	45853:    _Code_name[2991:3003],
	45854:    _Code_name[3003:3015],
	45855:    _Code_name[3015:3027],
	45856:    _Code_name[3027:3039],
	45857:    _Code_name[3039:3049],
	45858:    _Code_name[3049:3060],
	45859:    _Code_name[3060:3071],
	45860:    _Code_name[3071:3082],
	45861:    _Code_name[3082:3093],
	45862:    _Code_name[3093:3104],
	45863:    _Code_name[3104:3115],
	45864:    _Code_name[3115:3126],
	45865:    _Code_name[3126:3137],
	45866:    _Code_name[3137:3148],
	45867:    _Code_name[3148:3159],
	45868:    _Code_name[3159:3170],
	45869:    _Code_name[3170:3182],
	45870:    _Code_name[3182:3194],
	45871:    _Code_name[3194:3206],
	45872:    _Code_name[3206:3218],
	45873:    _Code_name[3218:3230],
	45874:    _Code_name[3230:3242],
	45875:    _Code_name[3242:3254],
	45876:    _Code_name[3254:3266],
	45877:    _Code_name[3266:3278],
	45878:    _Code_name[3278:3290],
	45879:    _Code_name[3290:3302],
	45880:    _Code_name[3302:3314],
	45881:    _Code_name[3314:3326],
	45882:    _Code_name[3326:3338],
	45883:    _Code_name[3338:3350],
	45884:    _Code_name[3350:3362],
	45885:    _Code_name[3362:3374],
	45886:    _Code_name[3374:3386],
	45887:    _Code_name[3386:3398],
	45888:    _Code_name[3398:3410],
	45889:    _Code_name[3410:3422],
	45890:    _Code_name[3422:3434],
	45891:    _Code_name[3434:3446],
	45892:    _Code_name[3446:3458],
	45893:    _Code_name[3458:3470],
	45894:    _Code_name[3470:3482],
	45895:    _Code_name[3482:3494],
	45896:    _Code_name[3494:3506],
	45897:    _Code_name[3506:3518],
	45898:    _Code_name[3518:3530],
	45899:    _Code_name[3530:3542],
	45900:    _Code_name[3542:3554],
	45901:    _Code_name[3554:3566],
	45902:    _Code_name[3566:3578],
	45903:    _Code_name[3578:3590],
	45904:    _Code_name[3590:3602],
	45905:    _Code_name[3602:3614],
	45906:    _Code_name[3614:3626],
	45907:    _Code_name[3626:3638],
	45908:    _Code_name[3638:3650],
	45909:    _Code_name[3650:3662],
	45910:    _Code_name[3662:3674],
	45911:    _Code_name[3674:3686],
	45912:    _Code_name[3686:3698],
	45913:    _Code_name[3698:3710],
	45914:    _Code_name[3710:3722],
	45915:    _Code_name[3722:3734],
	45916:    _Code_name[3734:3746],
	45917:    _Code_name[3746:3758],
	45918:    _Code_name[3758:3770],
	45919:    _Code_name[3770:3782],
	45920:    _Code_name[3782:3794],
	45921:    _Code_name[3794:3805],
	45922:    _Code_name[3805:3817],
	45923:    _Code_name[3817:3829],
	45924:    _Code_name[3829:3841],
	45925:    _Code_name[3841:3853],
	45926:    _Code_name[3853:3865],
	45927:    _Code_name[3865:3877],
	45928:    _Code_name[3877:3889],
	45929:    _Code_name[3889:3901],
	45930:    _Code_name[3901:3913],
	45931:    _Code_name[3913:3925],
	45932:    _Code_name[3925:3937],
	45933:    _Code_name[3937:3950],
	45934:    _Code_name[3950:3963],
	45935:    _Code_name[3963:3976],
	45936:    _Code_name[3976:3989],
	45937:    _Code_name[3989:4002],
	45938:    _Code_name[4002:4015],
	45939:    _Code_name[4015:4028],
	45940:    _Code_name[4028:4041],
	45941:    _Code_name[4041:4054],
	45942:    _Code_name[4054:4067],
	45943:    _Code_name[4067:4080],
	45944:    _Code_name[4080:4093],
	45945:    _Code_name[4093:4106],
	45946:    _Code_name[4106:4119],
	45947:    _Code_name[4119:4132],
	45948:    _Code_name[4132:4145],
	45949:    _Code_name[4145:4158],
	45950:    _Code_name[4158:4171],
	45951:    _Code_name[4171:4184],
	45952:    _Code_name[4184:4197],
	45953:    _Code_name[4197:4210],
	45954:    _Code_name[4210:4223],
	45955:    _Code_name[4223:4236],
	45956:    _Code_name[4236:4249],
	45957:    _Code_name[4249:4262],
	45958:    _Code_name[4262:4275],
	45959:    _Code_name[4275:4288],
	45960:    _Code_name[4288:4301],
	45961:    _Code_name[4301:4314],
	45962:    _Code_name[4314:4327],
	45963:    _Code_name[4327:4340],
	45964:    _Code_name[4340:4353],
	45965:    _Code_name[4353:4366],
	45966:    _Code_name[4366:4379],
	45967:    _Code_name[4379:4392],
	45968:    _Code_name[4392:4405],
	45969:    _Code_name[4405:4418],
	45970:    _Code_name[4418:4431],
	45971:    _Code_name[4431:4444],
	45972:    _Code_name[4444:4457],
	45973:    _Code_name[4457:4470],
	45974:    _Code_name[4470:4483],
	45975:    _Code_name[4483:4496],
	45976:    _Code_name[4496:4509],
	45977:    _Code_name[4509:4522],
	45978:    _Code_name[4522:4535],
	45979:    _Code_name[4535:4548],
	45980:    _Code_name[4548:4561],
	45981:    _Code_name[4561:4574],
	45982:    _Code_name[4574:4587],
	45983:    _Code_name[4587:4600],
	45984:    _Code_name[4600:4613],
	45985:    _Code_name[4613:4626],
	45986:    _Code_name[4626:4639],
	45987:    _Code_name[4639:4652],
	45988:    _Code_name[4652:4665],
	45989:    _Code_name[4665:4678],
	45990:    _Code_name[4678:4691],
	45991:    _Code_name[4691:4704],
	45992:    _Code_name[4704:4717],
	45993:    _Code_name[4717:4730],
	45994:    _Code_name[4730:4743],
	45995:    _Code_name[4743:4756],
	45996:    _Code_name[4756:4769],
	45997:    _Code_name[4769:4782],
	45998:    _Code_name[4782:4795],
	45999:    _Code_name[4795:4808],
	46000:    _Code_name[4808:4821],
	46001:    _Code_name[4821:4834],
	46002:    _Code_name[4834:4847],
	46003:    _Code_name[4847:4860],
	46004:    _Code_name[4860:4873],
	46005:    _Code_name[4873:4886],
	46006:    _Code_name[4886:4899],
	46007:    _Code_name[4899:4912],
	46008:    _Code_name[4912:4925],
	46009:    _Code_name[4925:4938],
	46010:    _Code_name[4938:4951],
	46011:    _Code_name[4951:4964],
	46012:    _Code_name[4964:4977],
	46013:    _Code_name[4977:4990],
	46014:    _Code_name[4990:5003],
	46015:    _Code_name[5003:5016],
	46016:    _Code_name[5016:5029],
	46017:    _Code_name[5029:5042],
	46018:    _Code_name[5042:5055],
	46019:    _Code_name[5055:5068],
	46020:    _Code_name[5068:5081],
	46021:    _Code_name[5081:5094],
	46022:    _Code_name[5094:5107],
	46023:    _Code_name[5107:5120],
	46024:    _Code_name[5120:5133],
	46025:    _Code_name[5133:5146],
	46026:    _Code_name[5146:5159],
	46027:    _Code_name[5159:5172],
	46028:    _Code_name[5172:5185],
	46029:    _Code_name[5185:5198],
	46030:    _Code_name[5198:5211],
	46031:    _Code_name[5211:5224],
	46032:    _Code_name[5224:5237],
	46033:    _Code_name[5237:5250],
	46034:    _Code_name[5250:5263],
	46035:    _Code_name[5263:5276],
	46036:    _Code_name[5276:5289],
	46037:    _Code_name[5289:5302],
	46038:    _Code_name[5302:5315],
	46039:    _Code_name[5315:5328],
	46040:    _Code_name[5328:5341],
	46041:    _Code_name[5341:5354],
	46042:    _Code_name[5354:5367],
	46043:    _Code_name[5367:5380],
	46044:    _Code_name[5380:5393],
	46045:    _Code_name[5393:5407],
	46046:    _Code_name[5407:5421],
	46047:    _Code_name[5421:5435],
	46048:    _Code_name[5435:5449],
	46049:    _Code_name[5449:5455],
	46050:    _Code_name[5455:5461],
	46051:    _Code_name[5461:5468],
	46052:    _Code_name[5468:5476],
	46081:    _Code_name[5476:5501],
	46082:    _Code_name[5501:5529],
	46083:    _Code_name[5529:5545],
	46337:    _Code_name[5545:5548],
	46338:    _Code_name[5548:5564],
	46593:    _Code_name[5564:5572],
	52225:    _Code_name[5572:5576],
	52753:    _Code_name[5576:5597],
	53248:    _Code_name[5597:5603],
	53479:    _Code_name[5603:5609],
	53482:    _Code_name[5609:5625],
	53483:    _Code_name[5625:5641],
	53485:    _Code_name[5641:5646],
	53649:    _Code_name[5646:5653],
	60241:    _Code_name[5653:5664],
	61697:    _Code_name[5664:5687],
	61698:    _Code_name[5687:5708],
	7367777:  _Code_name[5708:5719],
	8417572:  _Code_name[5719:5735],
	8483108:  _Code_name[5735:5751],
	9728292:  _Code_name[5751:5767],
	9793828:  _Code_name[5767:5783],
	10645796: _Code_name[5783:5799],
	10711332: _Code_name[5799:5815],
	11639056: _Code_name[5815:5824],
	11704592: _Code_name[5824:5834],
	11770128: _Code_name[5834:5845],
	11835664: _Code_name[5845:5856],
	13636096: _Code_name[5856:5861],
	13636097: _Code_name[5861:5866],
	13636098: _Code_name[5866:5871],
	13636101: _Code_name[5871:5876],
}

func (i Code) String() string {
	if str, ok := _Code_map[i]; ok {
		return str
	}
	return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
}