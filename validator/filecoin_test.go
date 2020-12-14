package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilecoinValidateAddress(t *testing.T) {
	validator := &Filecoin{}

	var mainnetCases = map[string]*Result{
		"f00":      {true, FilID, ""},
		"f01":      {true, FilID, ""},
		"f010":     {true, FilID, ""},
		"f0150":    {true, FilID, ""},
		"f0499":    {true, FilID, ""},
		"f01024":   {true, FilID, ""},
		"f01729":   {true, FilID, ""},
		"f0999999": {true, FilID, ""},
		"f15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq":                                              {true, FilSecp256k1, ""},
		"f12fiakbhe2gwd5cnmrenekasyn6v5tnaxaqizq6a":                                              {true, FilSecp256k1, ""},
		"f1wbxhu3ypkuo6eyp6hjx6davuelxaxrvwb2kuwva":                                              {true, FilSecp256k1, ""},
		"f1xtwapqc6nh4si2hcwpr3656iotzmlwumogqbuaa":                                              {true, FilSecp256k1, ""},
		"f1xcbgdhkgkwht3hrrnui3jdopeejsoatkzmoltqy":                                              {true, FilSecp256k1, ""},
		"f17uoq6tp427uzv7fztkbsnn64iwotfrristwpryy":                                              {true, FilSecp256k1, ""},
		"f24vg6ut43yw2h2jqydgbg2xq7x6f4kub3bg6as6i":                                              {true, FilActor, ""},
		"f25nml2cfbljvn4goqtclhifepvfnicv6g7mfmmvq":                                              {true, FilActor, ""},
		"f2nuqrg7vuysaue2pistjjnt3fadsdzvyuatqtfei":                                              {true, FilActor, ""},
		"f24dd4ox4c2vpf5vk5wkadgyyn6qtuvgcpxxon64a":                                              {true, FilActor, ""},
		"f2gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr23y":                                              {true, FilActor, ""},
		"f3vvmn62lofvhjd2ugzca6sof2j2ubwok6cj4xxbfzz4yuxfkgobpihhd2thlanmsh3w2ptld2gqkn2jvlss4a": {true, FilBLS, ""},
		"f3wmuu6crofhqmm3v4enos73okk2l366ck6yc4owxwbdtkmpk42ohkqxfitcpa57pjdcftql4tojda2poeruwa": {true, FilBLS, ""},
		"f3s2q2hzhkpiknjgmf4zq3ejab2rh62qbndueslmsdzervrhapxr7dftie4kpnpdiv2n6tvkr743ndhrsw6d3a": {true, FilBLS, ""},
		"f3q22fijmmlckhl56rn5nkyamkph3mcfu5ed6dheq53c244hfmnq2i7efdma3cj5voxenwiummf2ajlsbxc65a": {true, FilBLS, ""},
		"f3u5zgwa4ael3vuocgc5mfgygo4yuqocrntuuhcklf4xzg5tcaqwbyfabxetwtj4tsam3pbhnwghyhijr5mixa": {true, FilBLS, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"t00":      {true, FilID, ""},
		"t01":      {true, FilID, ""},
		"t010":     {true, FilID, ""},
		"t0150":    {true, FilID, ""},
		"t0499":    {true, FilID, ""},
		"t01024":   {true, FilID, ""},
		"t01729":   {true, FilID, ""},
		"t0999999": {true, FilID, ""},
		"t15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq":                                              {true, FilSecp256k1, ""},
		"t12fiakbhe2gwd5cnmrenekasyn6v5tnaxaqizq6a":                                              {true, FilSecp256k1, ""},
		"t1wbxhu3ypkuo6eyp6hjx6davuelxaxrvwb2kuwva":                                              {true, FilSecp256k1, ""},
		"t1xtwapqc6nh4si2hcwpr3656iotzmlwumogqbuaa":                                              {true, FilSecp256k1, ""},
		"t1xcbgdhkgkwht3hrrnui3jdopeejsoatkzmoltqy":                                              {true, FilSecp256k1, ""},
		"t17uoq6tp427uzv7fztkbsnn64iwotfrristwpryy":                                              {true, FilSecp256k1, ""},
		"t24vg6ut43yw2h2jqydgbg2xq7x6f4kub3bg6as6i":                                              {true, FilActor, ""},
		"t25nml2cfbljvn4goqtclhifepvfnicv6g7mfmmvq":                                              {true, FilActor, ""},
		"t2nuqrg7vuysaue2pistjjnt3fadsdzvyuatqtfei":                                              {true, FilActor, ""},
		"t24dd4ox4c2vpf5vk5wkadgyyn6qtuvgcpxxon64a":                                              {true, FilActor, ""},
		"t2gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr23y":                                              {true, FilActor, ""},
		"t3vvmn62lofvhjd2ugzca6sof2j2ubwok6cj4xxbfzz4yuxfkgobpihhd2thlanmsh3w2ptld2gqkn2jvlss4a": {true, FilBLS, ""},
		"t3wmuu6crofhqmm3v4enos73okk2l366ck6yc4owxwbdtkmpk42ohkqxfitcpa57pjdcftql4tojda2poeruwa": {true, FilBLS, ""},
		"t3s2q2hzhkpiknjgmf4zq3ejab2rh62qbndueslmsdzervrhapxr7dftie4kpnpdiv2n6tvkr743ndhrsw6d3a": {true, FilBLS, ""},
		"t3q22fijmmlckhl56rn5nkyamkph3mcfu5ed6dheq53c244hfmnq2i7efdma3cj5voxenwiummf2ajlsbxc65a": {true, FilBLS, ""},
		"t3u5zgwa4ael3vuocgc5mfgygo4yuqocrntuuhcklf4xzg5tcaqwbyfabxetwtj4tsam3pbhnwghyhijr5mixa": {true, FilBLS, ""},
	}

	for addr, result := range testnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"Q2gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr23y": {false, Unknown, ""},
		"t4gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr23y": {false, Unknown, ""},
		"t2gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr24y": {false, Unknown, ""},
		"t0banananananannnnnnnnn":                   {false, Unknown, ""},
		"t0banananananannnnnnnn ":                   {false, Unknown, ""},
		"t2gfvuyh7v2sx3patm1k23wdzmhyhtmqctasbr24y": {false, Unknown, ""},
		"t2gfvuyh7v2sx3paTm1k23wdzmhyhtmqctasbr24y": {false, Unknown, ""},
		"t2":    {false, Unknown, ""},
		"t1234": {false, Unknown, ""},
		"f4gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr23y": {false, Unknown, ""},
		"f2gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr24y": {false, Unknown, ""},
		"f0banananananannnnnnnnn":                   {false, Unknown, ""},
		"f0banananananannnnnnnn ":                   {false, Unknown, ""},
		"f2gfvuyh7v2sx3patm1k23wdzmhyhtmqctasbr24y": {false, Unknown, ""},
		"f2gfvuyh7v2sx3paTm1k23wdzmhyhtmqctasbr24y": {false, Unknown, ""},
		"f2":    {false, Unknown, ""},
		"f1234": {false, Unknown, ""},
		"abcde": {false, Unknown, ""},
		"":      {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
