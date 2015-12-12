package messaging

import (
	"fmt"
	"regexp"
	"time"
)

var pubnub = Pubnub{
	subscribeKey: "demo",
}

var (
	channelsSingleChannel    = *NewSubscriptionEntity()
	channelsThreeChannels    = *NewSubscriptionEntity()
	channelsSingleCG         = *NewSubscriptionEntity()
	channelsThreeCG          = *NewSubscriptionEntity()
	channelsChannelAndGroupC = *NewSubscriptionEntity()
	channelsChannelAndGroupG = *NewSubscriptionEntity()

	pubnubSingleChannel = Pubnub{
		channels:     channelsSingleChannel,
		subscribeKey: "my_key",
		uuid:         "my_uuid",
	}

	pubnubThreeChannels = Pubnub{
		channels:     channelsThreeChannels,
		subscribeKey: "my_key",
		uuid:         "my_uuid",
	}

	pubnubSingleCG = Pubnub{
		groups:       channelsSingleCG,
		subscribeKey: "my_key",
		uuid:         "my_uuid",
	}

	pubnubThreeCG = Pubnub{
		groups:       channelsThreeCG,
		subscribeKey: "my_key",
		uuid:         "my_uuid",
	}

	pubnubChannelAndGroup = Pubnub{
		channels:     channelsChannelAndGroupC,
		groups:       channelsChannelAndGroupG,
		subscribeKey: "my_key",
		uuid:         "my_uuid",
	}

	successChannel = make(chan SuccessResponse)
	errorChannel   = make(chan ErrorResponse)
	eventChannel   = make(chan ConnectionEvent)
)

var signatureRegexp, _ = regexp.Compile("&signature=.*$")

func timestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func truncateSignature(input string) (output string) {
	return signatureRegexp.ReplaceAllString(input, "")
}
