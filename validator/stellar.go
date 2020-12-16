package validator

// def get_account_sequence(chain, address) do
// url = "#{chain.config["endpoint"]}/accounts/#{address}"

// chain
// |> get(url)
// |> Map.fetch!("sequence")
// |> String.to_integer()
// end

// def pubkey_to_address(pubkey) do
// # checksum = prefix + pubkey |> crc |> int_to_bytes |> reverse
// # prefix + pubkey +checksum |> encode32 = address
// payload = Binary.append(@prefix, pubkey)

// checksum =
//   :crc_16_xmodem
//   |> CRC.crc(payload)
//   |> Serializer.int_to_bytes()
//   |> Binary.reverse()

// payload
// |> Binary.append(checksum)
// |> Base.encode32()
// end
