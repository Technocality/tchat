package serverdomain

import (
	"tchat/internal/types"
	"tchat/server/serverdata"
	"time"
)

type ChannelService struct {
	userRepository *serverdata.UserRepository
	repository     *serverdata.ChannelRepository
}

func NewChannelService(userRepository *serverdata.UserRepository, channelRepository *serverdata.ChannelRepository) *ChannelService {
	return &ChannelService{
		userRepository: userRepository,
		repository:     channelRepository,
	}
}

func (cs *ChannelService) CreateChannel(userID, channelName string) error {
	c := types.Channel{
		Name:      channelName,
		Owner:     userID,
		CreatedAt: time.Now(),
	}

	return cs.repository.CreateChannel(c)
}

func (cs *ChannelService) GetAll() ([]types.Channel, error) {
	return cs.GetAll()
}
