// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: aicoreops_menu.proto

package client

import (
	"context"

	"aicoreops_role/types"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateMenuRequest  = types.CreateMenuRequest
	CreateMenuResponse = types.CreateMenuResponse
	DeleteMenuRequest  = types.DeleteMenuRequest
	DeleteMenuResponse = types.DeleteMenuResponse
	GetMenuRequest     = types.GetMenuRequest
	GetMenuResponse    = types.GetMenuResponse
	ListMenusData      = types.ListMenusData
	ListMenusRequest   = types.ListMenusRequest
	ListMenusResponse  = types.ListMenusResponse
	Menu               = types.Menu
	UpdateMenuRequest  = types.UpdateMenuRequest
	UpdateMenuResponse = types.UpdateMenuResponse

	MenuService interface {
		// 创建菜单
		CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*CreateMenuResponse, error)
		// 获取菜单详情
		GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (*GetMenuResponse, error)
		// 更新菜单
		UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error)
		// 删除菜单
		DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...grpc.CallOption) (*DeleteMenuResponse, error)
		// 列出菜单
		ListMenus(ctx context.Context, in *ListMenusRequest, opts ...grpc.CallOption) (*ListMenusResponse, error)
	}

	defaultMenuService struct {
		cli zrpc.Client
	}
)

func NewMenuService(cli zrpc.Client) MenuService {
	return &defaultMenuService{
		cli: cli,
	}
}

// 创建菜单
func (m *defaultMenuService) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*CreateMenuResponse, error) {
	client := types.NewMenuServiceClient(m.cli.Conn())
	return client.CreateMenu(ctx, in, opts...)
}

// 获取菜单详情
func (m *defaultMenuService) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...grpc.CallOption) (*GetMenuResponse, error) {
	client := types.NewMenuServiceClient(m.cli.Conn())
	return client.GetMenu(ctx, in, opts...)
}

// 更新菜单
func (m *defaultMenuService) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*UpdateMenuResponse, error) {
	client := types.NewMenuServiceClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

// 删除菜单
func (m *defaultMenuService) DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...grpc.CallOption) (*DeleteMenuResponse, error) {
	client := types.NewMenuServiceClient(m.cli.Conn())
	return client.DeleteMenu(ctx, in, opts...)
}

// 列出菜单
func (m *defaultMenuService) ListMenus(ctx context.Context, in *ListMenusRequest, opts ...grpc.CallOption) (*ListMenusResponse, error) {
	client := types.NewMenuServiceClient(m.cli.Conn())
	return client.ListMenus(ctx, in, opts...)
}
