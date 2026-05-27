// ABOUTME: Audio playback module with device selection support.
// ABOUTME: Uses malgo (miniaudio bindings) for cross-platform audio output.

package audio

import (
	"os"
	"sync"

	"github.com/gen2brain/malgo"
	"github.com/go-audio/audio"
)

// DeviceInfo represents an audio output device
type DeviceInfo struct {
	Name      string
	IsDefault bool
}

// Player plays audio on a specific device
type Player struct {
	ctx        *malgo.AllocatedContext
	deviceID   *malgo.DeviceID // Stored copy of device ID (nil = default device)
	deviceName string
	volume     float64
	mu         sync.Mutex
}

// ListDevices returns all available audio output devices
func ListDevices() ([]DeviceInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// NewPlayer creates a new audio player for the specified device
// If deviceName is empty, the system default device is used
func NewPlayer(deviceName string, volume float64) (*Player, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find the device by name if specified

// Copy the DeviceID to avoid dangling pointer after devices slice is freed

// Play plays an audio file
func (p *Player) Play(soundPath string) error { _ = "STUB: not implemented"; return nil }

// Check if player is closed

// Check if file exists

// Decode audio file

// Apply volume

// Convert to bytes

// Create device config with larger buffer to prevent crackling

// Set specific device if configured

// Playback state

// Data callback

// 2 bytes per sample (16-bit)

// Fill remaining with silence

// Signal done when finished

// Initialize device

// Start playback

// Wait for playback to complete or timeout

// Delay to let buffer drain completely

// Close releases resources
func (p *Player) Close() error { _ = "STUB: not implemented"; return nil }

// decodeAudio decodes an audio file and returns samples, sample rate, and channel count
func (p *Player) decodeAudio(soundPath string) ([]int16, uint32, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func (p *Player) decodeMP3(f *os.File) ([]int16, uint32, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func (p *Player) decodeWAV(f *os.File) ([]int16, uint32, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func (p *Player) decodeFLAC(f *os.File) ([]int16, uint32, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func (p *Player) decodeOGG(f *os.File) ([]int16, uint32, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func (p *Player) decodeAIFF(f *os.File) ([]int16, uint32, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

// Convert samples based on bit depth

// streamToSamples converts a beep streamer to int16 samples
func streamToSamples(streamer interface {
	Stream([][2]float64) (int, bool)
}, sampleRate int, numChannels int) ([]int16, uint32, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

// Left channel

// Right channel (if stereo)

// intBufferToSamples converts go-audio IntBuffer to int16 samples
// bitDepth specifies the source bit depth (8, 16, 24, 32) for proper scaling
func intBufferToSamples(buf *audio.IntBuffer, bitDepth int) []int16 {
	_ = "STUB: not implemented"
	return nil
}

// Calculate shift amount based on bit depth
// We need to convert from source bit depth to 16-bit

// 8-bit: shift left by 8

// 16-bit: no conversion needed

// 24-bit: shift right by 8 to get upper 16 bits

// 32-bit: shift right by 16 to get upper 16 bits

// Fallback: assume 16-bit

// samplesToBytes converts int16 samples to bytes (little-endian)
func samplesToBytes(samples []int16) []byte { _ = "STUB: not implemented"; return nil }
