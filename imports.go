package webrtcvad

import (
	_ "github.com/rcrai/go-webrtcvad/webrtc"
	_ "github.com/rcrai/go-webrtcvad/webrtc/common_audio/signal_processing"
	_ "github.com/rcrai/go-webrtcvad/webrtc/common_audio/signal_processing/include"
	_ "github.com/rcrai/go-webrtcvad/webrtc/common_audio/vad"
	_ "github.com/rcrai/go-webrtcvad/webrtc/common_audio/vad/include"
	_ "github.com/rcrai/go-webrtcvad/webrtc/system_wrappers/interface"
)
