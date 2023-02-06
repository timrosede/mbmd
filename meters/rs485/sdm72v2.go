package rs485

import . "github.com/volkszaehler/mbmd/meters"

func init() {
	Register("SDM72-V2", NewSDM72V2Producer)
}

type SDM72V2Producer struct {
	Opcodes
}

func NewSDM72V2Producer() Producer {
	/**
	 * Opcodes as defined by Eastron SDM72.
	 * See https://data.stromzähler.eu/eastron/SDM72DM-V2-manual.pdf
	 */
	ops := Opcodes{
		VoltageL1: 0x0000,
		VoltageL2: 0x0002,
		VoltageL3: 0x0004,
		CurrentL1: 0x0006,
		CurrentL2: 0x0008,
		CurrentL3: 0x000A,

		PowerL1: 0x000C,
		PowerL2: 0x000E,
		PowerL3: 0x0010,
		Power:   0x0034,

		ApparentPowerL1: 0x0012, // apparent power l1
		ApparentPowerL2: 0x0014, // apparent power l1
		ApparentPowerL3: 0x0016, // apparent power l1

		ReactivePowerL1: 0x0018, // reactive power l1
		ReactivePowerL2: 0x001A, // reactive power l2
		ReactivePowerL3: 0x001C, // reactive power l3

		Import: 0x0048,
		Export: 0x004a,
		Sum:    0x0156,

		CosphiL1:  0x001e, //      230
		CosphiL2:  0x0020,
		CosphiL3:  0x0022,
		Cosphi:    0x003e,
		Frequency: 0x0046, //      230
	}
	return &SDM72V2Producer{Opcodes: ops}
}

func (p *SDM72V2Producer) Description() string {
	return "Eastron SDM72-V2"
}

func (p *SDM72V2Producer) snip(iec Measurement) Operation {
	operation := Operation{
		FuncCode:  ReadInputReg,
		OpCode:    p.Opcode(iec),
		ReadLen:   2,
		IEC61850:  iec,
		Transform: RTUIeee754ToFloat64,
	}
	return operation
}

// This device does not provide voltage data
// so it is not possible to automatically detect the device
func (p *SDM72V2Producer) Probe() Operation {
	return Operation{}
}

func (p *SDM72V2Producer) Produce() (res []Operation) {
	for op := range p.Opcodes {
		res = append(res, p.snip(op))
	}

	return res
}
