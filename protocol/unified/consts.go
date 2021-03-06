package unified

const (
	CMD_BIND_REQ      uint32 = 0x0001
	CMD_SUBMIT_REQ    uint32 = 0x0002
	CMD_DELIVER_REQ   uint32 = 0x0003
	CMD_REPORT_REQ    uint32 = 0x0004
	CMD_HEARTBEAT_REQ uint32 = 0x0005

	CMD_BIND_ACK      uint32 = 0x9001
	CMD_SUBMIT_ACK    uint32 = 0x9002
	CMD_DELIVER_ACK   uint32 = 0x9003
	CMD_REPORT_ACK    uint32 = 0x9004
	CMD_HEARTBEAT_ACK uint32 = 0x9005
)
