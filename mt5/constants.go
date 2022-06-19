package mt5

const (
	// Prefixes
	PREFIX_API    = "MT5WEBAPI%04x%04x"
	PACKET_FORMAT = "%04x%04x"

	// Configuration
	API_VERSION           = 3000
	META_SIZE             = 9
	PARAM_RETURN_CODE     = "RET_CODE"
	PARAM_SRV_RAND        = "SRV_RAND"
	PARAM_SRV_RAND_ANSWER = "SRV_RAND_ANSWER"
	PARAM_CLI_RAND        = "CLI_RAND"
	MAX_COMMANDS          = 16383

	// Authorization commands
	CMD_AUTH_START  = "AUTH_START"
	CMD_AUTH_ANSWER = "AUTH_ANSWER"

	// Misc
	WORD_API     = "WebAPI"
	WORD_MANAGER = "MANAGER"
	CRYPT_METHOD = "AES256OFB"
)
