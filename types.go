// Janus types relating to the videoroom plugin

package videoroom

// We need both the map defined below, and the list of keys that
// follows...as golang does not guarantee the order of range when
// applied to a map.  Perhaps we should autoinitialize evkeys...
// and that can be done.  But this is faster, though prone to error
// if someone adds a type to evtypes and not to evkeys
var EventTypes = map[string]func() interface{}{
	"error":       func() interface{} { return &ErrorMsg{} },
	"publishers":  func() interface{} { return &PublishMsg{} },
	"unpublished": func() interface{} { return &UnpublishMsg{} },
	"joining":     func() interface{} { return &JoiningMsg{} },
	"leaving":     func() interface{} { return &LeavingMsg{} },
	// The following is a catch all
	"videoroom":   func() interface{} { return &UnknownMsg{} },
}

// Defines the order in which the above map is searched.
var EventKeys = []string{
	"error",
	"publishers",
	"unpublished",
	"joining",
	"leaving",
	"videoroom",
}

type Publisher struct {
	ID uint64 `json:"id"`
	Display string `json:"display"`
	AudioCodec string `json:"audio_codec"`
	VideoCodec string `json:"video_codec"`
	Simulcast bool `json:"simulcast"`
	Talking bool `json:"talking"`
}

type Participant struct {
	ID uint64 `json:"id"`
	Display string `json:"display"`
	Publisher bool `json:"publisher"`
	Talking bool `json:"talking"`
}

type MsgListReq struct {
	Request string `json:"request"`
}

type MsgListParticipantsReq struct {
	Request string `json:"request"`
	Room uint64 `json:"room"`
}

type MsgListParticipantsResp struct {
	VideoRoom string `json:"videoroom"`
	Room uint64 `json:"room"`
	Participants []Participant `json:"participants"`
}

type JoinReq struct {
	Request string `json:"request"`
	Ptype string `json:"ptype"`
	Room uint64 `json:"room"`
	ID uint64 `json:"id"`
	Display string `json:"display"`
	//	Token string `json:"token"`
}

type JoinResp struct {
	VideoRoom string `json:"videoroom"`
	Room uint64 `json:"room"`
	Description string `json:"description"`
	ID uint64 `json:"id"`
	PrivateID uint64 `json:"private_id"`
	Publishers []Publisher `json:"publishers"`
}

type ErrorMsg struct {
	VideoRoom string `json:"videoroom"`
	ErrorCode uint64 `json:"error_code"`
	Error string `json:"error"`
}

type PublishMsg struct {
	VideoRoom string `json:"videoroom"`
	Room uint64 `json:"room"`
	// unpublished will contain the id of the unpublisher
	Publishers []Publisher `json:"publishers"`
}

type UnpublishMsg struct {
	VideoRoom string `json:"videoroom"`
	Room uint64 `json:"room"`
	Unpublished uint64 `json:"unpublished"`
}

type JoiningMsg struct {
	VideoRoom string `json:"videoroom"`
	Room uint64 `json:"room"`
	Joining struct {
		ID uint64 `json:"id"`
		Display string `json:"display"`
	} `json:"room"`
}

type LeavingMsg struct {
	VideoRoom string `json:"videoroom"`
	Room uint64 `json:"room"`
	Leaving uint64 `json:"leaving"`
}

type UnknownMsg struct {
	VideoRoom string `json:"videoroom"`
}

type RTPForwardReq struct {
	Request string `json:"request"`
	Room uint64 `json:"room"`
	PublisherID uint64 `json:"publisher_id"`
	Host string `json:"host"`
	HostFamily string `json:"host_family"`
	AudioPort uint64 `json:"audio_port"`
	VideoPort uint64 `json:"video_port"`
	Secret string `json:"secret"`
}

type RTPStopReq struct {
	Request string `json:"request"`
	Room uint64 `json:"room"`
	PublisherID uint64 `json:"publisher_id"`
	StreamID uint64 `json:"stream_id"`
	Secret string `json:"secret"`
}

type RTPListReq struct {
	Request string `json:"request"`
	Room uint64 `json:"room"`
	Secret string `json:"secret"`
}

type RTPForwardersResp struct {
	VideoRoom string `json:"videoroom"`
	Room uint64 `json:"room"`
	Forwarders []PubForwarder `json:"rtp_forwarders"`
}

type PubForwarder struct {
	Display string `json:"display"`
	PublisherID uint64 `json:"publisher_id"`
	Forwarder []Forwarder `json:"rtp_forwarder"`
}

type Forwarder struct {
	AudioStreamID uint64 `json:"audio_stream_id"`
	VideoStreamID uint64 `json:"video_stream_id"`
	DataStreamID uint64 `json:"data_stream_id"`
	IP string `json:"ip"`
	Port uint64 `json:"port"`
	RTCPPort uint64 `json:"rtcp_port"`
	SSRC uint64 `json:"ssrc"`
	PT uint64 `json:"pt"`
	Substream uint64 `json:"substream"`
	srtp bool `json:"srtp"`
}

type SubscribeReq struct {
	Request string `json:"request"`
	Ptype string `json:"ptype"`
	Room uint64 `json:"room"`
	Feed uint64 `json:"feed"`
	//      PrivateID uit64 `json:"private_id"`
	//	Token string `json:"token"`
}

type SubscribeResp struct {
	VideoRoom string `json:"videoroom"`
	Room uint64 `json:"room"`
	Feed uint64 `json:"feed"`
	Display string `json:"display"`
}

type StartReq struct {
	Request string `json:"request"`
}

type SDPPayload struct {
	SDP string `json:"sdp"`
	Type string `json:"type"`
}

type SDPMessage struct {
	SDP SDPPayload `json:"sdp"`
	//	ICE string `json:"ice"`
}
