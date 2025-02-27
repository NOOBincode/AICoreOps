/*
 * Copyright 2024 Bamboo
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * File: role.go
 */

package handler

import (
	"net/http"

	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/logic"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/svc"
	"github.com/GoSimplicity/AICoreOps/services/aicoreops_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type RoleHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRoleHandler(svcCtx *svc.ServiceContext) *RoleHandler {
	return &RoleHandler{
		svcCtx: svcCtx,
	}
}

// CreateRole 创建角色
func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var req types.CreateRoleRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.CreateRole(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// GetRole 获取角色详情
func (h *RoleHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	var req types.GetRoleRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.GetRole(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// UpdateRole 更新角色
func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	var req types.UpdateRoleRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.UpdateRole(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// DeleteRole 删除角色
func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	var req types.DeleteRoleRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.DeleteRole(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// ListRoles 获取角色列表
func (h *RoleHandler) ListRoles(w http.ResponseWriter, r *http.Request) {
	var req types.ListRolesRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.ListRoles(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// AssignPermissions 分配权限
func (h *RoleHandler) AssignPermissions(w http.ResponseWriter, r *http.Request) {
	var req types.AssignPermissionsRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.AssignPermissions(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// AssignRoleToUser 分配角色给用户
func (h *RoleHandler) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	var req types.AssignRoleToUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.AssignRoleToUser(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// RemoveUserPermissions 移除用户权限
func (h *RoleHandler) RemoveUserPermissions(w http.ResponseWriter, r *http.Request) {
	var req types.RemoveUserPermissionsRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.RemoveUserPermissions(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}

// RemoveRoleFromUser 移除用户角色
func (h *RoleHandler) RemoveRoleFromUser(w http.ResponseWriter, r *http.Request) {
	var req types.RemoveRoleFromUserRequest
	if err := httpx.Parse(r, &req); err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	l := logic.NewRoleLogic(r.Context(), h.svcCtx)
	resp, err := l.RemoveRoleFromUser(&req)
	if err != nil {
		httpx.OkJsonCtx(r.Context(), w, types.GeneralResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	resp.Code = http.StatusOK

	httpx.OkJsonCtx(r.Context(), w, resp)
}
