package main

import (
	"bytes"
	"context"
	"log"
	"net"

	todo "github.com/shashank-priyadarshi/projects/tools/markdownapp/proto"
	"github.com/yuin/goldmark"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":8082"
)

type markdownApplicationService struct {
	todo.MarkdownApplicationServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	markdownAppServer := grpc.NewServer()
	todo.RegisterMarkdownApplicationServer(markdownAppServer, &markdownApplicationService{})
	if err = markdownAppServer.Serve(lis); err != nil {
		log.Fatalf("failed to start markdownAppServer: %v", err)
	}
}

// Fetches list from DB
func (m *markdownApplicationService) List(context.Context, *emptypb.Empty) (items *todo.Items, err error) {
	return
}

// Parses given markdown
func (m *markdownApplicationService) Parse(context.Context, *emptypb.Empty) (errors *todo.Error, err error) {
	return
}

func (m *markdownApplicationService) parse() error {
	// ../../designs/markdown-app/data-structure.md
	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		panic(err)
	}
}

// Adds todo to DB and markdown
func (m *markdownApplicationService) Add(context.Context, *todo.Items) (errors *todo.Error, err error) {
	return
}

// Updates todo entry in DB and markdown
func (m *markdownApplicationService) Edit(context.Context, *todo.Items) (errors *todo.Error, err error) {
	return
}

// Deletes todo from DB and markdown
func (m *markdownApplicationService) Delete(context.Context, *todo.UIDs) (errors *todo.Error, err error) {
	return
}

// Create UID method
// NewMongoDBInstance creator, List, Create and Update methods
// Parse, Add, Update and Delete methods for markdown
