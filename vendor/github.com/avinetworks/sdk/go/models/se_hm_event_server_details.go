package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// SeHmEventServerDetails se hm event server details
// swagger:model SeHmEventServerDetails
type SeHmEventServerDetails struct {

	// Placeholder for description of property app_info of obj type SeHmEventServerDetails field type str  type object
	AppInfo []*AppInfo `json:"app_info,omitempty"`

	// Healthmonitor Failure code. Enum options - ARP_UNRESOLVED, CONNECTION_REFUSED, CONNECTION_TIMEOUT, RESPONSE_CODE_MISMATCH, PAYLOAD_CONTENT_MISMATCH, SERVER_UNREACHABLE, CONNECTION_RESET, CONNECTION_ERROR, HOST_ERROR, ADDRESS_ERROR, NO_PORT, PAYLOAD_TIMEOUT, NO_RESPONSE, NO_RESOURCES, SSL_ERROR, SSL_CERT_ERROR, PORT_UNREACHABLE, SCRIPT_ERROR, OTHER_ERROR, SERVER_DISABLED, REMOTE_STATE, MAINTENANCE_RESPONSE_CODE_MATCH, MAINTENANCE_PAYLOAD_CONTENT_MATCH, CHUNKED_RESPONSE_PAYLOAD_NOT_FOUND, GSLB_POOL_MEMBER_DOWN, GSLB_POOL_MEMBER_DISABLED, GSLB_POOL_MEMBER_STATE_UNKNOWN, INSUFFICIENT_HEALTH_MONITORS_UP, GSLB_POOL_MEMBER_REMOTE_STATE_UNKNOWN, RESPONSE_BUFFER_OVERFLOW.
	FailureCode *string `json:"failure_code,omitempty"`

	// Host name or VM name or DNS name for the server.
	Hostname *string `json:"hostname,omitempty"`

	// IP address of server.
	// Required: true
	IP *IPAddr `json:"ip"`

	// Port override form the pool port. If server port is not specified, the pool port is used.
	Port *int32 `json:"port,omitempty"`

	// Placeholder for description of property shm of obj type SeHmEventServerDetails field type str  type object
	Shm []*SeHmEventShmDetails `json:"shm,omitempty"`

	//  Enum options - ADF_CLIENT_CONN_SETUP_REFUSED, ADF_SERVER_CONN_SETUP_REFUSED, ADF_CLIENT_CONN_SETUP_TIMEDOUT, ADF_SERVER_CONN_SETUP_TIMEDOUT, ADF_CLIENT_CONN_SETUP_FAILED_INTERNAL, ADF_SERVER_CONN_SETUP_FAILED_INTERNAL, ADF_CLIENT_CONN_SETUP_FAILED_BAD_PACKET, ADF_UDP_CONN_SETUP_FAILED_INTERNAL, ADF_UDP_SERVER_CONN_SETUP_FAILED_INTERNAL, ADF_CLIENT_SENT_RESET, ADF_SERVER_SENT_RESET, ADF_CLIENT_CONN_TIMEDOUT, ADF_SERVER_CONN_TIMEDOUT, ADF_USER_DELETE_OPERATION, ADF_CLIENT_REQUEST_TIMEOUT, ADF_CLIENT_CONN_ABORTED, ADF_CLIENT_SSL_HANDSHAKE_FAILURE, ADF_CLIENT_CONN_FAILED, ADF_SERVER_CERTIFICATE_VERIFICATION_FAILED, ADF_SERVER_SIDE_SSL_HANDSHAKE_FAILED, ADF_IDLE_TIMEDOUT, ADF_CLIENT_CONNECTION_CLOSED_BEFORE_REQUEST, ADF_CLIENT_HIGH_TIMEOUT_RETRANSMITS, ADF_SERVER_HIGH_TIMEOUT_RETRANSMITS, ADF_CLIENT_HIGH_RX_ZERO_WINDOW_SIZE_EVENTS, ADF_SERVER_HIGH_RX_ZERO_WINDOW_SIZE_EVENTS, ADF_CLIENT_RTT_ABOVE_SEC, ADF_SERVER_RTT_ABOVE_500MS, ADF_CLIENT_HIGH_TOTAL_RETRANSMITS, ADF_SERVER_HIGH_TOTAL_RETRANSMITS, ADF_CLIENT_HIGH_OUT_OF_ORDERS, ADF_SERVER_HIGH_OUT_OF_ORDERS, ADF_CLIENT_HIGH_TX_ZERO_WINDOW_SIZE_EVENTS, ADF_SERVER_HIGH_TX_ZERO_WINDOW_SIZE_EVENTS, ADF_CLIENT_POSSIBLE_WINDOW_STUCK, ADF_SERVER_POSSIBLE_WINDOW_STUCK, ADF_SERVER_UNANSWERED_SYNS, ADF_CLIENT_CLOSE_CONNECTION_ON_VS_UPDATE, ADF_RESPONSE_CODE_4XX, ADF_RESPONSE_CODE_5XX, ADF_LOAD_BALANCING_FAILED, ADF_DATASCRIPT_EXECUTION_FAILED, ADF_REQUEST_NO_POOL, ADF_RATE_LIMIT_DROP_CLIENT_IP, ADF_RATE_LIMIT_DROP_URI, ADF_RATE_LIMIT_DROP_CLIENT_IP_URI, ADF_RATE_LIMIT_DROP_UNKNOWN_URI, ADF_RATE_LIMIT_DROP_BAD_URI, ADF_REQUEST_VIRTUAL_HOSTING_APP_SELECT_FAILED, ADF_RATE_LIMIT_DROP_UNKNOWN_CIP, ADF_RATE_LIMIT_DROP_BAD_CIP, ADF_RATE_LIMIT_DROP_CLIENT_IP_BAD, ADF_RATE_LIMIT_DROP_URI_BAD, ADF_RATE_LIMIT_DROP_CLIENT_IP_URI_BAD, ADF_RATE_LIMIT_DROP_REQ, ADF_RATE_LIMIT_DROP_CLIENT_IP_CONN, ADF_RATE_LIMIT_DROP_CONN, ADF_RATE_LIMIT_DROP_HEADER, ADF_RATE_LIMIT_DROP_CUSTOM, ADF_HTTP_VERSION_LT_1_0, ADF_CLIENT_HIGH_RESPONSE_TIME, ADF_SERVER_HIGH_RESPONSE_TIME, ADF_PERSISTENT_SERVER_CHANGE, ADF_DOS_SERVER_BAD_GATEWAY, ADF_DOS_SERVER_GATEWAY_TIMEOUT, ADF_DOS_CLIENT_SENT_RESET, ADF_DOS_CLIENT_CONN_TIMEOUT, ADF_DOS_CLIENT_REQUEST_TIMEOUT, ADF_DOS_CLIENT_CONN_ABORTED, ADF_DOS_CLIENT_BAD_REQUEST, ADF_DOS_CLIENT_REQUEST_ENTITY_TOO_LARGE, ADF_DOS_CLIENT_REQUEST_URI_TOO_LARGE, ADF_DOS_CLIENT_REQUEST_HEADER_TOO_LARGE, ADF_DOS_CLIENT_CLOSED_REQUEST, ADF_DOS_SSL_ERROR, ADF_REQUEST_MEMORY_LIMIT_EXCEEDED, ADF_X509_CLIENT_CERTIFICATE_VERIFICATION_FAILED, ADF_X509_CLIENT_CERTIFICATE_NOT_YET_VALID, ADF_X509_CLIENT_CERTIFICATE_EXPIRED, ADF_X509_CLIENT_CERTIFICATE_REVOKED, ADF_X509_CLIENT_CERTIFICATE_INVALID_CA, ADF_X509_CLIENT_CERTIFICATE_CRL_NOT_PRESENT, ADF_X509_CLIENT_CERTIFICATE_CRL_NOT_YET_VALID, ADF_X509_CLIENT_CERTIFICATE_CRL_EXPIRED, ADF_X509_CLIENT_CERTIFICATE_CRL_ERROR, ADF_X509_CLIENT_CERTIFICATE_CHAINING_ERROR, ADF_X509_CLIENT_CERTIFICATE_INTERNAL_ERROR, ADF_X509_CLIENT_CERTIFICATE_FORMAT_ERROR, ADF_UDP_PORT_NOT_REACHABLE, ADF_UDP_CONN_TIMEOUT, ADF_X509_SERVER_CERTIFICATE_VERIFICATION_FAILED, ADF_X509_SERVER_CERTIFICATE_NOT_YET_VALID, ADF_X509_SERVER_CERTIFICATE_EXPIRED, ADF_X509_SERVER_CERTIFICATE_REVOKED, ADF_X509_SERVER_CERTIFICATE_INVALID_CA, ADF_X509_SERVER_CERTIFICATE_CRL_NOT_PRESENT, ADF_X509_SERVER_CERTIFICATE_CRL_NOT_YET_VALID, ADF_X509_SERVER_CERTIFICATE_CRL_EXPIRED, ADF_X509_SERVER_CERTIFICATE_CRL_ERROR, ADF_X509_SERVER_CERTIFICATE_CHAINING_ERROR, ADF_X509_SERVER_CERTIFICATE_INTERNAL_ERROR, ADF_X509_SERVER_CERTIFICATE_FORMAT_ERROR, ADF_X509_SERVER_CERTIFICATE_HOSTNAME_ERROR, ADF_SSL_R_BAD_CHANGE_CIPHER_SPEC, ADF_SSL_R_BLOCK_CIPHER_PAD_IS_WRONG, ADF_SSL_R_DIGEST_CHECK_FAILED, ADF_SSL_R_ERROR_IN_RECEIVED_CIPHER_LIST, ADF_SSL_R_EXCESSIVE_MESSAGE_SIZE, ADF_SSL_R_LENGTH_MISMATCH, ADF_SSL_R_NO_CIPHERS_PASSED, ADF_SSL_R_NO_CIPHERS_SPECIFIED, ADF_SSL_R_NO_COMPRESSION_SPECIFIED, ADF_SSL_R_NO_SHARED_CIPHER, ADF_SSL_R_RECORD_LENGTH_MISMATCH, ADF_SSL_R_PARSE_TLSEXT, ADF_SSL_R_UNEXPECTED_MESSAGE, ADF_SSL_R_UNEXPECTED_RECORD, ADF_SSL_R_UNKNOWN_ALERT_TYPE, ADF_SSL_R_UNKNOWN_PROTOCOL, ADF_SSL_R_WRONG_VERSION_NUMBER, ADF_SSL_R_DECRYPTION_FAILED_OR_BAD_RECORD_MAC, ADF_SSL_R_RENEGOTIATE_EXT_TOO_LONG, ADF_SSL_R_RENEGOTIATION_ENCODING_ERR, ADF_SSL_R_RENEGOTIATION_MISMATCH, ADF_SSL_R_UNSAFE_LEGACY_RENEGOTIATION_DISABLED, ADF_SSL_R_SCSV_RECEIVED_WHEN_RENEGOTIATING, ADF_SSL_R_INAPPROPRIATE_FALLBACK, ADF_SSL_R_SSLV3_ALERT_UNEXPECTED_MESSAGE, ADF_SSL_R_SSLV3_ALERT_BAD_RECORD_MAC, ADF_SSL_R_TLSV1_ALERT_DECRYPTION_FAILED, ADF_SSL_R_TLSV1_ALERT_RECORD_OVERFLOW, ADF_SSL_R_SSLV3_ALERT_DECOMPRESSION_FAILURE, ADF_SSL_R_SSLV3_ALERT_HANDSHAKE_FAILURE, ADF_SSL_R_SSLV3_ALERT_NO_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_BAD_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_UNSUPPORTED_CERTIFICATE, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_REVOKED, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_EXPIRED, ADF_SSL_R_SSLV3_ALERT_CERTIFICATE_UNKNOWN, ADF_SSL_R_SSLV3_ALERT_ILLEGAL_PARAMETER, ADF_SSL_R_TLSV1_ALERT_UNKNOWN_CA, ADF_SSL_R_TLSV1_ALERT_ACCESS_DENIED, ADF_SSL_R_TLSV1_ALERT_DECODE_ERROR, ADF_SSL_R_TLSV1_ALERT_DECRYPT_ERROR, ADF_SSL_R_TLSV1_ALERT_EXPORT_RESTRICTION, ADF_SSL_R_TLSV1_ALERT_PROTOCOL_VERSION, ADF_SSL_R_TLSV1_ALERT_INSUFFICIENT_SECURITY, ADF_SSL_R_TLSV1_ALERT_INTERNAL_ERROR, ADF_SSL_R_TLSV1_ALERT_USER_CANCELLED, ADF_SSL_R_TLSV1_ALERT_NO_RENEGOTIATION, ADF_CLIENT_AUTH_UNKNOWN_USER, ADF_CLIENT_AUTH_LOGIN_FAILED, ADF_CLIENT_AUTH_MISSING_CREDENTIALS, ADF_CLIENT_AUTH_SERVER_CONN_ERROR, ADF_CLIENT_AUTH_USER_NOT_AUTHORIZED, ADF_CLIENT_AUTH_TIMED_OUT, ADF_CLIENT_AUTH_UNKNOWN_ERROR, ADF_CLIENT_DNS_FAILED_INVALID_QUERY, ADF_CLIENT_DNS_FAILED_INVALID_DOMAIN, ADF_CLIENT_DNS_FAILED_NO_SERVICE, ADF_CLIENT_DNS_FAILED_GS_DOWN, ADF_CLIENT_DNS_FAILED_NO_VALID_GS_MEMBER, ADF_SERVER_DNS_ERROR_RESPONSE, ADF_CLIENT_DNS_FAILED_UNSUPPORTED_QUERY, ADF_MEMORY_EXHAUSTED, ADF_CLIENT_DNS_POLICY_DROP, ADF_CLIENT_DNS_RL_POLICY_HIT, ADF_WAF_MATCH, ADF_HTTP2_CLIENT_TIMEDOUT, ADF_HTTP2_PROXY_PROTOCOL_ERROR, ADF_HTTP2_INVALID_CONNECTION_PREFACE, ADF_HTTP2_CLIENT_INVALID_DATA_FRAME_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_PADDED_DATA_FRAME_WITH_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_VIOLATED_CONN_FLOW_CONTROL, ADF_HTTP2_CLIENT_VIOLATED_STREAM_FLOW_CONTROL, ADF_HTTP2_CLIENT_DATA_FRAME_HALF_CLOSED_STREAM, ADF_HTTP2_CLIENT_HEADERS_FRAME_WITH_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_HEADERS_FRAME_WITH_EMPTY_HEADER_BLOCK, ADF_HTTP2_CLIENT_PADDED_HEADERS_FRAME_WITH_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_HEADERS_FRAME_INCORRECT_IDENTIFIER, ADF_HTTP2_CLIENT_HEADERS_FRAME_STREAM_INCORRECT_DEPENDENCY, ADF_HTTP2_CONCURRENT_STREAMS_EXCEEDED, ADF_HTTP2_CLIENT_STREAM_DATA_BEFORE_ACK_SETTINGS, ADF_HTTP2_CLIENT_HEADER_BLOCK_TOO_LONG_SIZE_UPDATE, ADF_HTTP2_CLIENT_HEADER_BLOCK_TOO_LONG_HEADER_INDEX, ADF_HTTP2_CLIENT_HEADER_BLOCK_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_INVALID_HPACK_TABLE_INDEX, ADF_HTTP2_CLIENT_OUT_OF_BOUND_HPACK_TABLE_INDEX, ADF_HTTP2_CLIENT_INVALID_TABLE_SIZE_UPDATE, ADF_HTTP2_CLIENT_HEADER_FIELD_TOO_LONG_LENGTH_VALUE, ADF_HTTP2_CLIENT_EXCEEDED_HTTP2_MAX_FIELD_SIZE_LIMIT, ADF_HTTP2_CLIENT_INVALID_ENCODED_HEADER_FIELD, ADF_HTTP2_CLIENT_EXCEEDED_HTTP2_MAX_HEADER_SIZE_LIMIT, ADF_HTTP2_CLIENT_INVALID_HEADER_NAME, ADF_HTTP2_CLIENT_HEADER_WITH_INVALID_VALUE, ADF_HTTP2_CLIENT_UNKNOWN_PSEUDO_HEADER, ADF_HTTP2_CLIENT_DUPLICATE_PATH_HEADER, ADF_HTTP2_CLIENT_EMPTY_PATH_HEADER, ADF_HTTP2_CLIENT_INVALID_PATH_HEADER, ADF_HTTP2_CLIENT_DUPLICATE_METHOD_HEADER, ADF_HTTP2_CLIENT_EMPTY_METHOD_HEADER, ADF_HTTP2_CLIENT_INVALID_METHOD_HEADER, ADF_HTTP2_CLIENT_DUPLICATE_SCHEME_HEADER, ADF_HTTP2_CLIENT_EMPTY_SCHEME_HEADER, ADF_HTTP2_CLIENT_NO_METHOD_HEADER, ADF_HTTP2_CLIENT_NO_SCHEME_HEADER, ADF_HTTP2_CLIENT_NO_PATH_HEADER, ADF_HTTP2_CLIENT_PREMATURELY_CLOSED_STREAM, ADF_HTTP2_CLIENT_PREMATURELY_CLOSED_CONNECTION, ADF_HTTP2_CLIENT_LARGER_DATA_BODY_THAN_DECLARED, ADF_HTTP2_CLIENT_LARGE_CHUNKED_BODY, ADF_HTTP2_NEGATIVE_WINDOW_UPDATE, ADF_HTTP2_SEND_WINDOW_FLOW_CONTROL_ERROR, ADF_HTTP2_CLIENT_UNEXPECTED_CONTINUATION_FRAME, ADF_HTTP2_CLIENT_WINDOW_UPDATE_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_WINDOW_UPDATE_FRAME_INCORRECT_INCREMENT, ADF_HTTP2_CLIENT_WINDOW_UPDATE_FRAME_INCREMENT_NOT_ALLOWED_FOR_WINDOW, ADF_HTTP2_CLIENT_GOAWAY_FRAME_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_PING_FRAME_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_PUSH_PROMISE, ADF_HTTP2_CLIENT_SETTINGS_FRAME_INCORRECT_MAX_FRAME_SIZE, ADF_HTTP2_CLIENT_SETTINGS_FRAME_INCORRECT_INIIAL_WINDOW_SIZE, ADF_HTTP2_CLIENT_SETTINGS_FRAME_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_SETTINGS_FRAME_ACK_FLAG_NONZERO_LENGTH, ADF_HTTP2_CLIENT_RST_STREAM_FRAME_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_RST_STREAM_FRAME_INCORRECT_IDENTIFIER, ADF_HTTP2_CLIENT_PRIORITY_FRAME_INCORRECT_DEPENDENCY, ADF_HTTP2_CLIENT_PRIORITY_FRAME_INCORRECT_IDENTIFIER, ADF_HTTP2_CLIENT_PRIORITY_FRAME_INCORRECT_LENGTH, ADF_HTTP2_CLIENT_CONTINUATION_FRAME_INCORRECT_IDENTIFIER, ADF_HTTP2_CLIENT_CONTINUATION_FRAME_EXPECTED_INAPPROPRIATE_FRAME, ADF_HTTP2_CLIENT_INVALID_HEADER, ADF_USER_DELETE_OPERATION_DATASCRIPT_RESET_CONN, ADF_USER_DELETE_OPERATION_HTTP_RULE_SECURITY_ACTION_CLOSE_CONN, ADF_USER_DELETE_OPERATION_HTTP_RULE_SECURITY_RATE_LIMIT_ACTION_CLOSE_CONN, ADF_USER_DELETE_OPERATION_HTTP_RULE_MISSING_TOKEN_ACTION_CLOSE_CONN, ADF_HTTP_BAD_REQUEST_INVALID_HOST_IN_REQUEST_LINE, ADF_HTTP_BAD_REQUEST_RECEIVED_VERSION_LESS_THAN_10, ADF_HTTP_NOT_ALLOWED_DATASCRIPT_RESPONSE_RETURNED_4XX, ADF_HTTP_NOT_ALLOWED_RUM_FLAGGED_INVALID_METHOD, ADF_HTTP_NOT_ALLOWED_UNSUPPORTED_TRACE_METHOD, ADF_HTTP_REQUEST_TIMEOUT_WAITING_FOR_CLIENT, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_INVALID_CONTENT_LENGTH, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_HTTP11_WITHOUT_HOST_HDR, ADF_HTTP_BAD_REQUEST_FAILED_TO_PARSE_URI, ADF_HTTP_BAD_REQUEST_INVALID_HEADER_LINE, ADF_HTTP_BAD_REQUEST_ERROR_WHILE_READING_CLIENT_HEADERS, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_DUPLICATE_HEADER, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_INVALID_HOST_HEADER, ADF_HTTP_NOT_IMPLEMENTED_CLIENT_SENT_UNKNOWN_TRANSFER_ENCODING, ADF_HTTP_BAD_REQUEST_REQUESTED_SERVER_NAME_DIFFERS, ADF_HTTP_BAD_REQUEST_CLIENT_SENT_INVALID_CHUNKED_BODY, ADF_HTTP_BAD_REQUEST_INVALID_HEADER_IN_SPDY, ADF_HTTP_BAD_REQUEST_INVALID_HEADER_BLOCK_IN_SPDY, ADF_HTTP_BAD_REQUEST_DATA_ERROR_IN_SPDY, ADF_HTTP_BAD_REQUEST_NO_METHOD_URI_OR_PROT_IN_REQ_CREATE_SPDY, ADF_HTTP_BAD_REQUEST_CLIENT_PREMATURELY_CLOSED_SPDY_STREAM, ADF_HTTP_BAD_REQUEST_DATA_ERROR_IN_SPDY_READ_REQ_BODY, ADF_HTTP_BAD_REQUEST_CERT_ERROR, ADF_HTTP_BAD_REQUEST_PLAIN_HTTP_REQUEST_SENT_ON_HTTPS_PORT, ADF_HTTP_BAD_REQUEST_NO_CERT_ERROR, ADF_HTTP_BAD_REQUEST_HEADER_TOO_LARGE, ADF_SERVER_HIGH_RESPONSE_TIME_L7, ADF_SERVER_HIGH_RESPONSE_TIME_L4, ADF_COOKIE_SIZE_GREATER_THAN_MAX, ADF_COOKIE_SIZE_LESS_THAN_MIN_COOKIE_LEN, ADF_PERSISTENCE_PROFILE_KEYS_NOT_CONFIGURED, ADF_PERSISTENCE_COOKIE_VERSION_MISMATCH, ADF_COOKIE_ABSENT_FROM_KEYS_IN_PERSISTENCE_PROFILE, ADF_GSLB_SITE_PERSISTENCE_REMOTE_SITE_DOWN, ADF_HTTP_NOT_ALLOWED_DATASCRIPT_RESPONSE_RETURNED_5XX, ADF_SERVER_UPSTREAM_TIMEOUT, ADF_SERVER_UPSTREAM_READ_ERROR, ADF_SERVER_UPSTREAM_RESOLVER_ERROR, ADF_SIP_INVALID_MESSAGE_FROM_CLIENT, ADF_SIP_MESSAGE_UPDATE_FAILED, ADF_SIP_SERVER_UNKNOWN_CALLID, ADF_SIP_REQUEST_FAILED, ADF_SIP_REQUEST_TIMEDOUT, ADF_SIP_CONN_IDLE_TIMEDOUT, ADF_SIP_TRANSACTION_TIMEDOUT, ADF_SIP_SVR_UDP_PORT_NOT_REACHABLE, ADF_SIP_CLT_UDP_PORT_NOT_REACHABLE, ADF_SIP_INVALID_MESSAGE_FROM_SERVER, ADF_L4_DATASCRIPT_DROP, ADF_L4_DATASCRIPT_SIGNIFICANCE, ADF_SAML_COOKIE_VERSION_MISMATCH, ADF_SAML_COOKIE_KEYS_NOT_CONFIGURED, ADF_SAML_COOKIE_ABSENT_FROM_KEYS_IN_SAML_AUTH_POLICY, ADF_SAML_COOKIE_INVALID, ADF_SAML_COOKIE_DECRYPTION_ERROR, ADF_SAML_COOKIE_ENCRYPTION_ERROR, ADF_SAML_COOKIE_DECODE_ERROR, ADF_SAML_COOKIE_SESSION_COOKIE_GREATER_THAN_MAX, ADF_SAML_ASSERTION_DOES_NOT_MATCH_REQUEST_ID, ADF_SAML_AUTHENTICATION_UNSUPPORTED_METHOD, ADF_SAML_COOKIE_SESSION_COOKIE_TIMEOUT, ADF_SAML_ACS_URL_MISMATCH, ADF_SAML_ASSERTION_NO_BODY, ADF_SAML_ASSERTION_INVALID, ADF_SAML_ASSERTION_ATTRIBUTE_ERROR, ADF_SAML_LOGIN_ERROR, ADF_HTTP_SERVER_RESELECT_OCCURRENCE, ADF_HTTP_RULE_SECURITY_RATE_LIMIT_ACTION_REPORT, ADF_HTTP_RULE_SECURITY_RATE_LIMIT_ACTION_REDIRECT, ADF_HTTP_RULE_SECURITY_RATE_LIMIT_ACTION_RESPONSE, ADF_HTTP2_SERVER_SENT_UNEXPECTED_FRAME, ADF_HTTP2_SERVER_SENT_FRAME_UNKNOWN_STREAM, ADF_HTTP2_SERVER_REJECTED_REQUEST_WITH_ERROR, ADF_HTTP2_SERVER_SENT_GOAWAY_WITH_ERROR, ADF_HTTP2_SERVER_SENT_UNEXPECTED_PUSH_PROMISE, ADF_HTTP2_SERVER_SENT_INVALID_HEADER, ADF_HTTP2_SERVER_SENT_DUP_STATUS_HEADER, ADF_HTTP2_SERVER_SENT_INVALID_STATUS_HEADER, ADF_HTTP2_SERVER_SENT_LARGE_HEADER_NAME_LEN, ADF_HTTP2_SERVER_NO_STATUS_HEADER, ADF_HTTP2_SERVER_SENT_DATA_FOR_UNKNOWN_STREAM, ADF_HTTP2_SERVER_STREAM_FLOW_CONTROL_VIOLATION, ADF_HTTP2_SERVER_CONN_FLOW_CONTROL_VIOLATION, ADF_HTTP2_SERVER_SENT_INVALID_TRAILER, ADF_HTTP2_SERVER_SENT_TRAILER_NO_ENDSTREAM_FLAG, ADF_HTTP2_SERVER_SENT_SHORT_FRAME, ADF_HTTP2_SERVER_SENT_FRAME_LONG_PADDING, ADF_HTTP2_SERVER_SENT_LARGE_FRAME, ADF_HTTP2_SERVER_SENT_FRAME_INVALID_LENGTH, ADF_HTTP2_SERVER_SENT_TRUNCATED_HEADER, ADF_HTTP2_SERVER_SENT_INVALID_TABLE_INDEX, ADF_HTTP2_SERVER_SENT_INVALID_TABLE_SIZE_UPDATE, ADF_HTTP2_SERVER_SENT_TABLE_INDEX_CONT_FLAG, ADF_HTTP2_SERVER_SENT_ZERO_HEADER_NAME_LEN, ADF_HTTP2_SERVER_SENT_INVALID_ENCODED_HEADER, ADF_HTTP2_SERVER_SENT_RST_INVALID_LENGTH, ADF_HTTP2_SERVER_SENT_GOAWAY_NONZERO_ID, ADF_HTTP2_SERVER_SENT_GOAWAY_INVALID_LEN, ADF_HTTP2_SERVER_SENT_WIN_UPDATE_INVALID_LEN, ADF_HTTP2_SERVER_SENT_WIN_UPDATE_LARGE_LEN, ADF_HTTP2_SERVER_SENT_SETTINGS_NONZERO_ID, ADF_HTTP2_SERVER_SENT_SETTINGS_ACK_NONZERO_ID, ADF_HTTP2_SERVER_SENT_SETTINGS_INVALID_LEN, ADF_HTTP2_SERVER_SENT_SETTING_LARGE_INI_WIN_SIZE, ADF_HTTP2_SERVER_SENT_PING_NONZERO_ID, ADF_HTTP2_SERVER_SENT_PING_INVALID_LEN, ADF_HTTP2_SERVER_SENT_PING_ACK, ADF_HTTP2_SERVER_NO_UPSTREAM_KEEPALIVE, ADF_HTTP2_CLIENT_SENT_TRAILER.
	SslErrorCode *string `json:"ssl_error_code,omitempty"`
}
