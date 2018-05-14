# fmFM

[![Build Status](https://travis-ci.org/but80/fmfm.core.svg?branch=master)](https://travis-ci.org/but80/fmfm.core)
[![Go Report Card](https://goreportcard.com/badge/github.com/but80/fmfm.core)](https://goreportcard.com/report/github.com/but80/fmfm.core)
[![Godoc](https://godoc.org/github.com/but80/fmfm.core?status.svg)](http://godoc.org/github.com/but80/fmfm.core)

**WORK IN PROGRESS**

**fmFM** (Fake Mobile FM synth) is a YAMAHA MA-5 (YMU765) / YMF825 clone software FM synthesizer.

Most of this code is based on [doomjs/opl3](https://github.com/doomjs/opl3).

# Requirements

- macOS
- [PortMIDI](http://portmedia.sourceforge.net/portmidi/)

  ```
  # On macOS
  brew install portmidi
  ```
- [PortAudio](http://www.portaudio.com/)

  ```
  # On macOS
  brew install portaudio
  ```

# Usage

```
go run cmd/fmfm-cli/main.go
```

- Voice libraries (`*.vm5.pb`) must be placed under `voice/` before running. They can be generated by [smaf825](https://github.com/but80/smaf825/tree/v2) (currently use `v2` branch for this feature). [More information (Japanese)](https://github.com/but80/smaf825/tree/v2#ymf825%E7%94%A8%E3%83%88%E3%83%BC%E3%83%B3%E3%83%87%E3%83%BC%E3%82%BF%E3%81%AE%E6%8A%BD%E5%87%BA)
- The IAC virtual MIDI port named `IAC YAMAHA Virtual MIDI Device 0` must be created before running.
  fmFM receives the MIDI messages via this port.

# Todo

- Analyze ATS-MA5 output
  - Waveform of DAM
  - Waveform of DVB
  - LPF character at feedback
  - AR shifting by KSR
  - Channel pan resolution
  - Channel pan and voice pan blending
  - MIDI vibrato resolution
  - LFO reset timing

# License

GPL 3.0
