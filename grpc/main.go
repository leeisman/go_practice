package main

import (
	"context"
	"gitlab.silkrode.com.tw/team_golang/jys/infra/pb/nc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:8090"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	ncClient := nc.NewNCServiceClient(conn)
	create(ncClient)
	//log.Print("sleep")
	//time.Sleep(time.Second * 20)
	//log.Print("wake up")
	//update(ncClient)
}

func create(ncClient nc.NCServiceClient) {
	publishAt := time.Now()
	in := nc.CreateNotificationReq{
		Title:       "go_practice title",
		Content:     "go_practice content",
		Type:        nc.NotificationType_Feedback,
		PublishedAt: &publishAt,
		Users: []*nc.User{
			&nc.User{
				ID:           1,
				UserDeviceID: "FrankieTest",
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := ncClient.CreateNotification(ctx, &in)
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}

	log.Print("create resp: ", resp)
}

func update(ncClient nc.NCServiceClient) {
	in := nc.UpdateNotificationReq{
		WhereCond: nc.WhereNotificationCondition{
			Notification: nc.Notification{
				ID: 14,
			},
		},
		Notification: &nc.Notification{
			Title:   "update title",
			Content: "update content",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := ncClient.UpdateNotification(ctx, &in)
	if err != nil {
		log.Fatalf("could not update: %v", err)
	}
	log.Print("update resp: ", resp)
}
