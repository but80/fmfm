# fmFM

[![Build Status](https://travis-ci.org/but80/fmfm.core.svg?branch=master)](https://travis-ci.org/but80/fmfm.core)
[![Go Report Card](https://goreportcard.com/badge/gopkg.in/but80/fmfm.core.v1?)](https://goreportcard.com/report/gopkg.in/but80/fmfm.core.v1)
[![Godoc](https://godoc.org/gopkg.in/but80/fmfm.core.v1?status.svg)](https://godoc.org/gopkg.in/but80/fmfm.core.v1)

**WORK IN PROGRESS**

**fmFM** (Fake Mobile FM synth) is a YAMAHA MA-5 (YMU765) / YMF825 clone software FM synthesizer.

Most of this code is based on [doomjs/opl3](https://github.com/doomjs/opl3).

# Requirements

- macOS
- [Go >= 1.9](https://golang.org/)
- [PortMIDI](http://portmedia.sourceforge.net/portmidi/)
- [PortAudio](http://www.portaudio.com/)

# Installation

1. Install [Homebrew](https://brew.sh/)

   ```bash
   /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
   ```
2. Install dependencies

   ```bash
   brew install go
   brew install portmidi
   brew install portaudio
   ```
3. Install fmfm-cli

   ```bash
   go get -u gopkg.in/but80/fmfm.core.v1/cmd/fmfm-cli
   ```

# Usage

```
NAME:
   fmfm-cli list - List MIDI devices

USAGE:
   fmfm-cli list
```

```
NAME:
   fmfm-cli midi - Listen MIDI events

USAGE:
   fmfm-cli midi [command options] [<Input MIDI device>]

OPTIONS:
   --mono, -m                 Force mono mode in all MIDI channels except drum PC
   --mute-nopc, -z            Mute if program change is not found
   --level value, -l value    Total level in dB (default: -12)
   --limiter value, -c value  Limiter threshold in dB (default: -3)
   --ignore value, -n value   Ignore specified MIDI channel (default: 0)
   --solo value, -s value     Accept only specified MIDI channel (default: 0)
   --dump value, -d value     Dump MIDI channel (default: 0)
   --print, -p                Print status
```

- Voice libraries (`*.vm5.pb`) must be placed under `voice/` before running. They can be generated by [smaf825](https://github.com/but80/smaf825/tree/v2) (currently use `v2` branch for this feature). [More information (Japanese)](https://github.com/but80/smaf825/tree/v2#ymf825%E7%94%A8%E3%83%88%E3%83%BC%E3%83%B3%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E6%8A%BD%E5%87%BA)
- fmFM receives MIDI messages via the MIDI port specified by the 1st argument.

# Todo

- Analyze ATS-MA5 output
  - Waveform of DVB
  - MIDI vibrato resolution
  - Channel pan resolution
  - Channel pan and voice pan blending
- Reduce multiplications in envelope generator

# License

MIT License
