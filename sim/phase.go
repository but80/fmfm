package sim

import (
	"github.com/but80/fmfm/ymf/ymfdata"
)

type phaseGenerator struct {
	chip                 *Chip
	evb                  bool
	dvb                  int
	phaseFrac64          uint64
	phaseIncrementFrac64 uint64
}

func newPhaseGenerator(chip *Chip) *phaseGenerator {
	return &phaseGenerator{
		chip: chip,
	}
}

func (pg *phaseGenerator) setFrequency(fnum, block, bo, mult, dt int) {
	baseFrequency := float64(fnum<<uint(block+3-bo)) / (16.0 * ymfdata.FNUMCoef)

	ksn := block<<1 | fnum>>9
	operatorFrequency := baseFrequency + ymfdata.DTCoef[dt][ksn]
	operatorFrequency *= ymfdata.MultTable[mult]

	pg.phaseIncrementFrac64 = uint64(operatorFrequency / pg.chip.sampleRate * ymfdata.Pow64Of2)
}

func (pg *phaseGenerator) getPhase(vibratoIndex int) uint64 {
	if pg.evb {
		pg.phaseFrac64 += (pg.phaseIncrementFrac64 >> 32) * ymfdata.VibratoTableInt32Frac32[pg.dvb][vibratoIndex]
	} else {
		pg.phaseFrac64 += pg.phaseIncrementFrac64
	}
	return pg.phaseFrac64
}

func (pg *phaseGenerator) keyOn() {
	pg.phaseFrac64 = 0
}