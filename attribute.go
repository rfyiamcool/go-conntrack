package conntrack

import (
	"fmt"

	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

// ConnAttr represents the type and value of a attribute of a connection
type ConnAttr struct {
	Type ConnAttrType
	Data []byte
}

// ConnAttrType specifies the attribute of a connection
type ConnAttrType uint16

// Attributes of a connection
// based on libnetfilter_conntrack.h
const (
	AttrOrigIPv4Src             ConnAttrType = iota /* u32 bits */
	AttrOrigIPv4Dst             ConnAttrType = iota /* u32 bits */
	AttrReplIPv4Src             ConnAttrType = iota /* u32 bits */
	AttrReplIPv4Dst             ConnAttrType = iota /* u32 bits */
	AttrOrigIPv6Src             ConnAttrType = iota /* u128 bits */
	AttrOrigIPv6Dst             ConnAttrType = iota /* u128 bits */
	AttrReplIPv6Src             ConnAttrType = iota /* u128 bits */
	AttrReplIPv6Dst             ConnAttrType = iota /* u128 bits */
	AttrOrigPortSrc             ConnAttrType = iota /* u16 bits */
	AttrOrigPortDst             ConnAttrType = iota /* u16 bits */
	AttrReplPortSrc             ConnAttrType = iota /* u16 bits */
	AttrReplPortDst             ConnAttrType = iota /* u16 bits */
	AttrIcmpType                ConnAttrType = iota /* u8 bits */
	AttrIcmpCode                ConnAttrType = iota /* u8 bits */
	AttrIcmpID                  ConnAttrType = iota /* u16 bits */
	AttrOrigL3Proto             ConnAttrType = iota /* u8 bits */
	AttrReplL3Proto             ConnAttrType = iota /* u8 bits */
	AttrOrigL4Proto             ConnAttrType = iota /* u8 bits */
	AttrReplL4Proto             ConnAttrType = iota /* u8 bits */
	AttrTCPState                ConnAttrType = iota /* u8 bits */
	AttrSNatIPv4                ConnAttrType = iota /* u32 bits */
	AttrDNatIPv4                ConnAttrType = iota /* u32 bits */
	AttrSNatPort                ConnAttrType = iota /* u16 bits */
	AttrDNatPort                ConnAttrType = iota /* u16 bits */
	AttrTimeout                 ConnAttrType = iota /* u32 bits */
	AttrMark                    ConnAttrType = iota /* u32 bits */
	AttrOrigCounterPackets      ConnAttrType = iota /* u64 bits */
	AttrReplCounterPackets      ConnAttrType = iota /* u64 bits */
	AttrOrigCounterBytes        ConnAttrType = iota /* u64 bits */
	AttrReplCounterBytes        ConnAttrType = iota /* u64 bits */
	AttrUse                     ConnAttrType = iota /* u32 bits */
	AttrID                      ConnAttrType = iota /* u32 bits */
	AttrStatus                  ConnAttrType = iota /* u32 bits  */
	AttrTCPFlagsOrig            ConnAttrType = iota /* u8 bits */
	AttrTCPFlagsRepl            ConnAttrType = iota /* u8 bits */
	AttrTCPMaskOrig             ConnAttrType = iota /* u8 bits */
	AttrTCPMaskRepl             ConnAttrType = iota /* u8 bits */
	AttrMasterIPv4Src           ConnAttrType = iota /* u32 bits */
	AttrMasterIPv4Dst           ConnAttrType = iota /* u32 bits */
	AttrMasterIPv6Src           ConnAttrType = iota /* u128 bits */
	AttrMasterIPv6Dst           ConnAttrType = iota /* u128 bits */
	AttrMasterPortSrc           ConnAttrType = iota /* u16 bits */
	AttrMasterPortDst           ConnAttrType = iota /* u16 bits */
	AttrMasterL3Proto           ConnAttrType = iota /* u8 bits */
	AttrMasterL4Proto           ConnAttrType = iota /* u8 bits */
	AttrSecmark                 ConnAttrType = iota /* u32 bits */
	AttrOrigNatSeqCorrectionPos ConnAttrType = iota /* u32 bits */
	AttrOrigNatSeqOffsetBefore  ConnAttrType = iota /* u32 bits */
	AttrOrigNatSeqOffsetAfter   ConnAttrType = iota /* u32 bits */
	AttrReplNatSeqCorrectionPos ConnAttrType = iota /* u32 bits */
	AttrReplNatSeqOffsetBefore  ConnAttrType = iota /* u32 bits */
	AttrReplNatSeqOffsetAfter   ConnAttrType = iota /* u32 bits */
	AttrSctpState               ConnAttrType = iota /* u8 bits */
	AttrSctpVtagOrig            ConnAttrType = iota /* u32 bits */
	AttrSctpVtagRepl            ConnAttrType = iota /* u32 bits */
	AttrHelperName              ConnAttrType = iota /* string (30 bytes max) */
	AttrDccpState               ConnAttrType = iota /* u8 bits */
	AttrDccpRole                ConnAttrType = iota /* u8 bits */
	AttrDccpHandshakeSeq        ConnAttrType = iota /* u64 bits */
	AttrTCPWScaleOrig           ConnAttrType = iota /* u8 bits */
	AttrTCPWScaleRepl           ConnAttrType = iota /* u8 bits */
	AttrZone                    ConnAttrType = iota /* u16 bits */
	AttrSecCtx                  ConnAttrType = iota /* string */
	AttrTimestampStart          ConnAttrType = iota /* u64 bits linux >= 2.6.38 */
	AttrTimestampStop           ConnAttrType = iota /* u64 bits linux >= 2.6.38 */
	AttrHelperInfo              ConnAttrType = iota /* variable length */
	AttrConnlabels              ConnAttrType = iota /* variable length */
	AttrConnlabelsMask          ConnAttrType = iota /* variable length */
	AttrOrigzone                ConnAttrType = iota /* u16 bits */
	AttrReplzone                ConnAttrType = iota /* u16 bits */
	AttrSNatIPv6                ConnAttrType = iota /* u128 bits */
	AttrDNatIPv6                ConnAttrType = iota /* u128 bits */
	AttrMax                     ConnAttrType = iota
)

const (
	ctaUnspec        = iota
	ctaTupleOrig     = iota
	ctaTupleReply    = iota
	ctaStatus        = iota
	ctaProtoinfo     = iota
	ctaHelp          = iota
	ctaNatSrc        = iota
	ctaTimeout       = iota
	ctaMark          = iota
	ctaCountersOrig  = iota
	ctaCountersReply = iota
	ctaUse           = iota
	ctaID            = iota
	ctaNatDst        = iota
	ctaTupleMaster   = iota
	ctaSeqAdjOrig    = iota
	ctaSeqAdjRepl    = iota
	ctaSecmark       = iota
	ctaZone          = iota
	ctaSecCtx        = iota
	ctaTimestamp     = iota
	ctaMarkMask      = iota
	ctaLables        = iota
	ctaLablesMask    = iota
)

const (
	ctaIPv4Src = 1
	ctaIPv4Dst = 2
	ctaIPv6Src = 3
	ctaIPv6Dst = 4
)

const (
	ctaProtoNum        = 1
	ctaProtoSrcPort    = 2
	ctaProtoDstPort    = 3
	ctaProtoIcmpID     = 4
	ctaProtoIcmpType   = 5
	ctaProtoIcmpCode   = 6
	ctaProtoIcmpv6ID   = 7
	ctaProtoIcmpv6Type = 8
	ctaProtoIcmpv6Code = 9
)

const (
	ctaProtoinfoTCPState      = 1
	ctaProtoinfoTCPWScaleOrig = 2
	ctaProtoinfoTCPWScaleRepl = 3
	ctaProtoinfoTCPFlagsOrig  = 4
	ctaProtoinfoTCPFlagsRepl  = 5
)

const (
	ctaCounterPackets   = 1
	ctaCounterBytes     = 2
	ctaCounter32Packets = 3
	ctaCounter32Bytes   = 4
)

const (
	ctaTimestampStart = 1
	ctaTimestampStop  = 2
)

const nlafNested = (1 << 15)

func nestAttributes(filters []ConnAttr) ([]byte, error) {
	var attrs []netlink.Attribute

	for _, filter := range filters {
		switch filter.Type {
		case AttrMark:
			if len(filter.Data) != 4 {
				return nil, fmt.Errorf("Length of data for type %d has to be 4", filter.Type)
			}
			attrs = append(attrs, netlink.Attribute{Type: ctaMark, Data: filter.Data})
		default:
			return nil, fmt.Errorf("Type %d not yet implemented", filter.Type)
		}
	}

	return netlink.MarshalAttributes(attrs)
}

func checkHeader(data []byte) int {
	if (data[0] == unix.AF_INET || data[0] == unix.AF_INET6) && data[1] == unix.NFNETLINK_V0 {
		return 4
	}
	return 0
}

func extractTCPTuple(data []byte) ([]ConnAttr, error) {
	var connAttr []ConnAttr
	attributes, err := netlink.UnmarshalAttributes(data)
	if err != nil {
		return nil, err
	}
	for _, attr := range attributes {
		switch attr.Type & 0XFF {
		case ctaProtoinfoTCPState:
			connAttr = append(connAttr, ConnAttr{Type: AttrTCPState, Data: attr.Data})
		case ctaProtoinfoTCPWScaleOrig:
			connAttr = append(connAttr, ConnAttr{Type: AttrTCPWScaleOrig, Data: attr.Data})
		case ctaProtoinfoTCPWScaleRepl:
			connAttr = append(connAttr, ConnAttr{Type: AttrTCPWScaleRepl, Data: attr.Data})
		case ctaProtoinfoTCPFlagsOrig:
			connAttr = append(connAttr, ConnAttr{Type: AttrTCPFlagsOrig, Data: attr.Data})
		case ctaProtoinfoTCPFlagsRepl:
			connAttr = append(connAttr, ConnAttr{Type: AttrTCPFlagsRepl, Data: attr.Data})
		}
	}
	return connAttr, nil
}

func extractProtocolTuple(dir int, data []byte) ([]ConnAttr, int, error) {
	var connAttr []ConnAttr
	var protocol int
	attributes, err := netlink.UnmarshalAttributes(data)
	if err != nil {
		return nil, protocol, err
	}
	for _, attr := range attributes {
		switch attr.Type & 0XFF {
		case ctaProtoNum:
			protocol = int(attr.Data[0])
			if dir == -1 {
				connAttr = append(connAttr, ConnAttr{Type: AttrOrigL4Proto, Data: attr.Data})

			} else {
				connAttr = append(connAttr, ConnAttr{Type: AttrReplL4Proto, Data: attr.Data})

			}
		case ctaProtoSrcPort:
			connAttr = append(connAttr, ConnAttr{Type: ConnAttrType(ctaProtoSrcPort + dir + 8), Data: attr.Data})
		case ctaProtoDstPort:
			connAttr = append(connAttr, ConnAttr{Type: ConnAttrType(ctaProtoDstPort + dir + 8), Data: attr.Data})
		case ctaProtoIcmpID:
			connAttr = append(connAttr, ConnAttr{Type: AttrIcmpID, Data: attr.Data})
		case ctaProtoIcmpType:
			connAttr = append(connAttr, ConnAttr{Type: AttrIcmpType, Data: attr.Data})
		case ctaProtoIcmpCode:
			connAttr = append(connAttr, ConnAttr{Type: AttrIcmpCode, Data: attr.Data})
		case ctaProtoIcmpv6ID:
			connAttr = append(connAttr, ConnAttr{Type: AttrIcmpID, Data: attr.Data})
		case ctaProtoIcmpv6Type:
			connAttr = append(connAttr, ConnAttr{Type: AttrIcmpType, Data: attr.Data})
		case ctaProtoIcmpv6Code:
			connAttr = append(connAttr, ConnAttr{Type: AttrIcmpCode, Data: attr.Data})
		default:
			return nil, protocol, fmt.Errorf("Unexpected Protocol Tuple Attribute: %d", attr.Type&0xFF)
		}
	}
	return connAttr, protocol, nil
}

func extractIPTuple(dir int, data []byte) ([]ConnAttr, int, error) {
	var connAttr []ConnAttr
	var protocol int
	attributes, err := netlink.UnmarshalAttributes(data)
	if err != nil {
		return nil, protocol, err
	}
	for _, attr := range attributes {
		if attr.Type&nlafNested == nlafNested {
			tuple, proto, err := extractProtocolTuple(dir, attr.Data)
			if err != nil {
				return nil, protocol, err
			}
			protocol = proto
			connAttr = append(connAttr, tuple...)
			continue
		}
		switch attr.Type & 0XFF {
		case ctaIPv4Src:
			connAttr = append(connAttr, ConnAttr{Type: ConnAttrType(ctaIPv4Src + dir), Data: attr.Data})
		case ctaIPv4Dst:
			connAttr = append(connAttr, ConnAttr{Type: ConnAttrType(ctaIPv4Dst + dir), Data: attr.Data})
		case ctaIPv6Src:
			connAttr = append(connAttr, ConnAttr{Type: ConnAttrType(ctaIPv6Src + dir + 2), Data: attr.Data})
		case ctaIPv6Dst:
			connAttr = append(connAttr, ConnAttr{Type: ConnAttrType(ctaIPv6Dst + dir + 2), Data: attr.Data})
		default:
			return nil, protocol, fmt.Errorf("Unexpected IP Tuple Attribute: %d", attr.Type&0xFF)
		}
	}
	return connAttr, protocol, nil
}

func extractCounterTuple(dir int, data []byte) ([]ConnAttr, error) {
	var connAttr []ConnAttr
	attributes, err := netlink.UnmarshalAttributes(data)
	if err != nil {
		return nil, err
	}
	for _, attr := range attributes {
		switch attr.Type & 0XFF {
		case ctaCounter32Packets:
			fallthrough
		case ctaCounterPackets:
			if dir == -1 {
				connAttr = append(connAttr, ConnAttr{Type: AttrOrigCounterPackets, Data: attr.Data})
			} else {
				connAttr = append(connAttr, ConnAttr{Type: AttrReplCounterPackets, Data: attr.Data})
			}
		case ctaCounter32Bytes:
			fallthrough
		case ctaCounterBytes:
			if dir == -1 {
				connAttr = append(connAttr, ConnAttr{Type: AttrOrigCounterBytes, Data: attr.Data})
			} else {
				connAttr = append(connAttr, ConnAttr{Type: AttrReplCounterBytes, Data: attr.Data})
			}
		}
	}
	return connAttr, nil
}

func extractTimestampTuple(data []byte) ([]ConnAttr, error) {
	var connAttr []ConnAttr
	attributes, err := netlink.UnmarshalAttributes(data)
	if err != nil {
		return nil, err
	}
	for _, attr := range attributes {
		switch attr.Type & 0XFF {
		case ctaTimestampStart:
			connAttr = append(connAttr, ConnAttr{Type: AttrTimestampStart, Data: attr.Data})
		case ctaTimestampStop:
			connAttr = append(connAttr, ConnAttr{Type: AttrTimestampStop, Data: attr.Data})
		}
	}
	return connAttr, nil
}

func extractAttribute(data []byte) ([]ConnAttr, error) {
	var connAttr []ConnAttr
	var protocol int
	attributes, err := netlink.UnmarshalAttributes(data)
	if err != nil {
		return nil, err
	}

	for _, attr := range attributes {
		switch attr.Type & 0xFF {
		case ctaTupleOrig:
			tuple, proto, err := extractIPTuple(-1, attr.Data[4:])
			if err != nil {
				return nil, err
			}
			connAttr = append(connAttr, tuple...)
			protocol = proto
		case ctaTupleReply:
			tuple, proto, err := extractIPTuple(1, attr.Data[4:])
			if err != nil {
				return nil, err
			}
			connAttr = append(connAttr, tuple...)
			protocol = proto
		case ctaProtoinfo:
			if protocol == 6 {
				tuple, err := extractTCPTuple(attr.Data[4:])
				if err != nil {
					return nil, err
				}
				connAttr = append(connAttr, tuple...)
			}
		case ctaCountersOrig:
			tuple, err := extractCounterTuple(-1, attr.Data)
			if err != nil {
				return nil, err
			}
			connAttr = append(connAttr, tuple...)
		case ctaCountersReply:
			tuple, err := extractCounterTuple(1, attr.Data)
			if err != nil {
				return nil, err
			}
			connAttr = append(connAttr, tuple...)
		case ctaTimestamp:
			tuple, err := extractTimestampTuple(attr.Data)
			if err != nil {
				return nil, err
			}
			connAttr = append(connAttr, tuple...)
		case ctaTimeout:
			connAttr = append(connAttr, ConnAttr{Type: AttrTimeout, Data: attr.Data})
		case ctaID:
			connAttr = append(connAttr, ConnAttr{Type: AttrID, Data: attr.Data})
		case ctaUse:
			connAttr = append(connAttr, ConnAttr{Type: AttrUse, Data: attr.Data})
		case ctaStatus:
			connAttr = append(connAttr, ConnAttr{Type: AttrStatus, Data: attr.Data})
		case ctaMark:
			connAttr = append(connAttr, ConnAttr{Type: AttrMark, Data: attr.Data})
		case ctaSecCtx:
			connAttr = append(connAttr, ConnAttr{Type: AttrSecCtx, Data: attr.Data})
		default:
			fmt.Println(attr.Type&0xFF, "\t", attr.Length, "\t", attr.Data)
		}
	}
	return connAttr, nil
}

func extractAttributes(msg []byte) (*Conn, error) {
	var conn Conn

	offset := checkHeader(msg[:2])
	attr, err := extractAttribute(msg[offset:])
	if err != nil {
		return nil, err
	}
	conn.attr = attr
	return &conn, nil
}
