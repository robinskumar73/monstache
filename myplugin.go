package main

import (
	"github.com/rwynn/monstache/monstachemap"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// a plugin to convert document values to uppercase
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
    doc := input.Document
		if payload, ok := doc["payload"]; ok {
			payload := payload.(map[string]interface{})
			// there's an element
			delete(payload, "cells");
      delete(payload, "descriptionInJSON");
      delete(payload, "descriptionInMD");
      
      if workspaceId, ok := payload["workspaceId"]; ok {
        switch workspaceId.(type){
          case string:
            break;
          case primitive.Undefined:
            delete(payload, "workspaceId");
          default:
            delete(payload, "workspaceId");
        }
      }
		}

		
    output = &monstachemap.MapperPluginOutput{Document: doc}
    return
}


func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error) {
  doc := input.Document
  BLOCK_MODEL_ID := "5UDGkdvYv"
  // Only keep document with model id of block.
  if payload, ok := doc["payload"]; ok {
    payload := payload.(map[string]interface{})
    if modelId, ok := payload["modelId"].(string); ok {
      keep =  modelId == BLOCK_MODEL_ID;
    }
  }
  return
}
