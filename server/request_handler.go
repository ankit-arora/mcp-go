// Code generated by `go generate`. DO NOT EDIT.
// source: server/internal/gen/request_handler.go.tmpl
package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ankit-arora/mcp-go/mcp"
)

// HandleMessage processes an incoming JSON-RPC message and returns an appropriate response
func (s *MCPServer) HandleMessage(
	ctx context.Context,
	message json.RawMessage,
) mcp.JSONRPCMessage {
	// Add server to context
	ctx = context.WithValue(ctx, serverKey{}, s)
	var err *requestError

	var baseMessage struct {
		JSONRPC string        `json:"jsonrpc"`
		Method  mcp.MCPMethod `json:"method"`
		ID      any           `json:"id,omitempty"`
	}

	if err := json.Unmarshal(message, &baseMessage); err != nil {
		return createErrorResponse(
			nil,
			mcp.PARSE_ERROR,
			"Failed to parse message",
		)
	}

	// Check for valid JSONRPC version
	if baseMessage.JSONRPC != mcp.JSONRPC_VERSION {
		return createErrorResponse(
			baseMessage.ID,
			mcp.INVALID_REQUEST,
			"Invalid JSON-RPC version",
		)
	}

	if baseMessage.ID == nil {
		var notification mcp.JSONRPCNotification
		if err := json.Unmarshal(message, &notification); err != nil {
			return createErrorResponse(
				nil,
				mcp.PARSE_ERROR,
				"Failed to parse notification",
			)
		}
		s.handleNotification(ctx, notification)
		return nil // Return nil for notifications
	}

	switch baseMessage.Method {
	case mcp.MethodInitialize:
		var request mcp.InitializeRequest
		var result *mcp.InitializeResult
		if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforeInitialize(ctx, baseMessage.ID, &request)
			result, err = s.handleInitialize(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterInitialize(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	case mcp.MethodPing:
		var request mcp.PingRequest
		var result *mcp.EmptyResult
		if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforePing(ctx, baseMessage.ID, &request)
			result, err = s.handlePing(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterPing(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	case mcp.MethodResourcesList:
		var request mcp.ListResourcesRequest
		var result *mcp.ListResourcesResult
		if s.capabilities.resources == nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.METHOD_NOT_FOUND,
				err:  fmt.Errorf("resources %w", ErrUnsupported),
			}
		} else if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforeListResources(ctx, baseMessage.ID, &request)
			result, err = s.handleListResources(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterListResources(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	case mcp.MethodResourcesTemplatesList:
		var request mcp.ListResourceTemplatesRequest
		var result *mcp.ListResourceTemplatesResult
		if s.capabilities.resources == nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.METHOD_NOT_FOUND,
				err:  fmt.Errorf("resources %w", ErrUnsupported),
			}
		} else if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforeListResourceTemplates(ctx, baseMessage.ID, &request)
			result, err = s.handleListResourceTemplates(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterListResourceTemplates(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	case mcp.MethodResourcesRead:
		var request mcp.ReadResourceRequest
		var result *mcp.ReadResourceResult
		if s.capabilities.resources == nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.METHOD_NOT_FOUND,
				err:  fmt.Errorf("resources %w", ErrUnsupported),
			}
		} else if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforeReadResource(ctx, baseMessage.ID, &request)
			result, err = s.handleReadResource(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterReadResource(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	case mcp.MethodPromptsList:
		var request mcp.ListPromptsRequest
		var result *mcp.ListPromptsResult
		if s.capabilities.prompts == nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.METHOD_NOT_FOUND,
				err:  fmt.Errorf("prompts %w", ErrUnsupported),
			}
		} else if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforeListPrompts(ctx, baseMessage.ID, &request)
			result, err = s.handleListPrompts(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterListPrompts(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	case mcp.MethodPromptsGet:
		var request mcp.GetPromptRequest
		var result *mcp.GetPromptResult
		if s.capabilities.prompts == nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.METHOD_NOT_FOUND,
				err:  fmt.Errorf("prompts %w", ErrUnsupported),
			}
		} else if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforeGetPrompt(ctx, baseMessage.ID, &request)
			result, err = s.handleGetPrompt(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterGetPrompt(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	case mcp.MethodToolsList:
		var request mcp.ListToolsRequest
		var result *mcp.ListToolsResult
		if s.capabilities.tools == nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.METHOD_NOT_FOUND,
				err:  fmt.Errorf("tools %w", ErrUnsupported),
			}
		} else if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforeListTools(ctx, baseMessage.ID, &request)
			result, err = s.handleListTools(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterListTools(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	case mcp.MethodToolsCall:
		var request mcp.CallToolRequest
		var result *mcp.CallToolResult
		if s.capabilities.tools == nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.METHOD_NOT_FOUND,
				err:  fmt.Errorf("tools %w", ErrUnsupported),
			}
		} else if unmarshalErr := json.Unmarshal(message, &request); unmarshalErr != nil {
			err = &requestError{
				id:   baseMessage.ID,
				code: mcp.INVALID_REQUEST,
				err:  &UnparseableMessageError{message: message, err: unmarshalErr, method: baseMessage.Method},
			}
		} else {
			s.hooks.beforeCallTool(ctx, baseMessage.ID, &request)
			result, err = s.handleToolCall(ctx, baseMessage.ID, request)
		}
		if err != nil {
			s.hooks.onError(ctx, baseMessage.ID, baseMessage.Method, &request, err)
			return err.ToJSONRPCError()
		}
		s.hooks.afterCallTool(ctx, baseMessage.ID, &request, result)
		return createResponse(baseMessage.ID, *result)
	default:
		return createErrorResponse(
			baseMessage.ID,
			mcp.METHOD_NOT_FOUND,
			fmt.Sprintf("Method %s not found", baseMessage.Method),
		)
	}
}
