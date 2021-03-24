package rpc

import (
	"context"
	"strings"

	pb "github.com/phaesoo/pigeonhole/gen/go/rpc/proto"
	"github.com/phaesoo/pigeonhole/models"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

// Send implements helloworld.GreeterServer
func (s *Server) Send(ctx context.Context, in *pb.NotificationRequest) (*pb.NotificationReply, error) {
	badge := int(in.Badge)
	notification := pb.PushNotification{
		Platform:         int(in.Platform),
		Tokens:           in.Tokens,
		Message:          in.Message,
		Title:            in.Title,
		Topic:            in.Topic,
		APIKey:           in.Key,
		Category:         in.Category,
		Sound:            in.Sound,
		ContentAvailable: in.ContentAvailable,
		ThreadID:         in.ThreadID,
		MutableContent:   in.MutableContent,
		Image:            in.Image,
		Priority:         strings.ToLower(in.GetPriority().String()),
	}

	if badge > 0 {
		notification.Badge = &badge
	}

	if in.Topic != "" && in.Platform == models.PlatFormAndroid {
		notification.To = in.Topic
	}

	if in.Alert != nil {
		notification.Alert = models.Alert{
			Title:        in.Alert.Title,
			Body:         in.Alert.Body,
			Subtitle:     in.Alert.Subtitle,
			Action:       in.Alert.Action,
			ActionLocKey: in.Alert.Action,
			LaunchImage:  in.Alert.LaunchImage,
			LocArgs:      in.Alert.LocArgs,
			LocKey:       in.Alert.LocKey,
			TitleLocArgs: in.Alert.TitleLocArgs,
			TitleLocKey:  in.Alert.TitleLocKey,
		}
	}

	if in.Data != nil {
		notification.Data = map[string]interface{}{}
		for k, v := range in.Data.Fields {
			notification.Data[k] = v
		}
	}

	//go gorush.SendNotification(ctx, notification)

	return &pb.NotificationReply{
		Success: true,
		Counts:  int32(len(notification.Tokens)),
	}, nil
}
