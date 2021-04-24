package main

import (
	"github.com/rwynn/monstache/monstachemap"
)

// a plugin to convert document values to uppercase
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
    doc := input.Document
    // for k, v := range doc {
    //     switch v.(type) {
    //     case string:
    //         doc[k] = strings.ToUpper(v.(string))
    //     }
    // }
		if payload, ok := doc["payload"]; ok {
			payload := payload.(map[string]interface{})
			// there's an element
			delete(payload, "cells");
		}

		
    output = &monstachemap.MapperPluginOutput{Document: doc}
    return
}