// Copyright 2018, The GoPacket Authors, All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.
//
//******************************************************************************

package layers

import (
	"github.com/graniet/gopacket"
)

func init()  {
	RegisterTCPPortLayerType(TCPPort(502), LayerTypeModbusPDU)
}

//******************************************************************************

// ModbusTCP Type
// --------
// Type ModbusTCP implements the DecodingLayer interface. Each ModbusTCP object
// represents in a structured form the MODBUS Application Protocol header (MBAP) record present as the TCP
// payload in an ModbusTCP TCP packet.
//
type ModbusPDU struct {
	BaseLayer // Stores the packet bytes and payload (Modbus PDU) bytes .

	TransactionIdentifier uint16         // Identification of a MODBUS Request/Response transaction
	ProtocolIdentifier    ModbusProtocol // It is used for intra-system multiplexing
	Length                uint16         // Number of following bytes (includes 1 byte for UnitIdentifier + Modbus data length
	UnitIdentifier        uint8          // Identification of a remote slave connected on a serial line or on other buses
}

//******************************************************************************

// LayerType returns the layer type of the ModbusTCP object, which is LayerTypeModbusTCP.
func (d *ModbusPDU) LayerType() gopacket.LayerType {
	return LayerTypeModbusPDU
}

//******************************************************************************

// decodeModbusTCP analyses a byte slice and attempts to decode it as an ModbusTCP
// record of a TCP packet.
//
// If it succeeds, it loads p with information about the packet and returns nil.
// If it fails, it returns an error (non nil).
//
// This function is employed in layertypes.go to register the ModbusTCP layer.
func decodeModbusPDU(data []byte, p gopacket.PacketBuilder) error {

	// Attempt to decode the byte slice.
	d := &ModbusPDU{}
	err := d.DecodeFromBytes(data, p)
	if err != nil {
		return err
	}
	// If the decoding worked, add the layer to the packet and set it
	// as the application layer too, if there isn't already one.
	p.AddLayer(d)
	p.SetApplicationLayer(d)

	return p.NextDecoder(d.NextLayerType())

}

//******************************************************************************

// DecodeFromBytes analyses a byte slice and attempts to decode it as an ModbusTCP
// record of a TCP packet.
//
// Upon succeeds, it loads the ModbusTCP object with information about the packet
// and returns nil.
// Upon failure, it returns an error (non nil).
func (d *ModbusPDU) DecodeFromBytes(data []byte, df gopacket.DecodeFeedback) error {
	return nil
}

//******************************************************************************

// NextLayerType returns the layer type of the ModbusTCP payload, which is LayerTypePayload.
func (d *ModbusPDU) NextLayerType() gopacket.LayerType {
	return gopacket.LayerTypePayload
}

//******************************************************************************

// Payload returns Modbus Protocol Data Unit (PDU) composed by Function Code and Data, it is carried within ModbusTCP packets
func (d *ModbusPDU) Payload() []byte {
	return d.BaseLayer.Payload
}

// CanDecode returns the set of layer types that this DecodingLayer can decode
func (s *ModbusPDU) CanDecode() gopacket.LayerClass {
	return LayerTypeModbusPDU
}
