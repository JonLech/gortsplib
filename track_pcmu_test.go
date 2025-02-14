package gortsplib //nolint:dupl

import (
	"testing"

	psdp "github.com/pion/sdp/v3"
	"github.com/stretchr/testify/require"
)

func TestTrackPCMUAttributes(t *testing.T) {
	track := &TrackPCMU{}
	require.Equal(t, 8000, track.ClockRate())
	require.Equal(t, "", track.GetControl())
}

func TestTrackPCMUClone(t *testing.T) {
	track := &TrackPCMU{}

	clone := track.clone()
	require.NotSame(t, track, clone)
	require.Equal(t, track, clone)
}

func TestTrackPCMUMediaDescription(t *testing.T) {
	track := &TrackPCMU{}

	require.Equal(t, &psdp.MediaDescription{
		MediaName: psdp.MediaName{
			Media:   "audio",
			Protos:  []string{"RTP", "AVP"},
			Formats: []string{"0"},
		},
		Attributes: []psdp.Attribute{
			{
				Key:   "rtpmap",
				Value: "0 PCMU/8000",
			},
			{
				Key:   "control",
				Value: "",
			},
		},
	}, track.MediaDescription())
}
