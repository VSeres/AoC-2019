package intcode

func (p *Program) allocateMemory(size int) {
	newMemory := make([]int, len(p.memory)+size)
	copy(newMemory, p.memory)
	p.memory = newMemory
}

func (p *Program) putValue(dest int, value int, mode byte) {
	memorySize := len(p.memory)
	addr := p.getDestAddress(mode, dest)
	if memorySize <= addr {
		p.allocateMemory(addr + 1 - memorySize)
	}
	p.memory[addr] = value
}

func (p *Program) getValue(mode byte, param int) int {
	memorySize := len(p.memory)
	switch mode {
	case '0':
		if memorySize <= param {
			p.allocateMemory(param + 1 - memorySize)
		}
		return p.memory[param]
	case '1':
		return param
	case '2':
		if memorySize <= p.base+param {
			p.allocateMemory(p.base + param + 1 - memorySize)
		}
		return p.memory[p.base+param]
	default:
		return -100000000
	}
}

func (p *Program) getDestAddress(mode byte, param int) int {
	memorySize := len(p.memory)
	switch mode {
	case '0':
		if memorySize <= param {
			p.allocateMemory(param + 1 - memorySize)
		}
		return param
	case '2':
		if memorySize <= p.base+param {
			p.allocateMemory(p.base + param + 1 - memorySize)
		}
		return p.base + param
	default:
		return -100000000
	}
}
