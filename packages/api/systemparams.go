/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package api

import (
	"net/http"

	"github.com/IBAX-io/go-ibax/packages/consts"
	"github.com/IBAX-io/go-ibax/packages/model"
		logger.WithFields(log.Fields{"type": consts.DBError, "error": err}).Error("Getting all system parameters")
	}

	result := &ecosystemParamsResult{
		List: make([]paramResult, 0),
	}

	acceptNames := form.AcceptNames()
	for _, item := range list {
		if len(acceptNames) > 0 && !acceptNames[item.Name] {
			continue
		}
		result.List = append(result.List, paramResult{
			Name:       item.Name,
			Value:      item.Value,
			Conditions: item.Conditions,
		})
	}

	if len(result.List) == 0 {
		errorResponse(w, errParamNotFound.Errorf(form.Names), http.StatusBadRequest)
		return
	}

	jsonResponse(w, result)
}
