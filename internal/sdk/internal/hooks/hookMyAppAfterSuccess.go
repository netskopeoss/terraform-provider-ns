package hooks

import (
	"net/http"
	"strings"
	"encoding/json"
    "io/ioutil"
	"fmt"
)

type MyAppResponse struct{}

var (
    //_ sdkInitHook       = (*MyAppResponse)(nil)
	//_ beforeRequestHook = (*MyAppResponse)(nil)
    _ afterSuccessHook  = (*MyAppResponse)(nil)
    //_ afterErrorHook    = (*MyAppResponse)(nil)
)

func (i *MyAppResponse) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	if hookCtx.OperationID == "createNPAPrivateApps" { 
        var responseMap map[string]interface{}

        // Read and unmarshal the response body
         body, err := ioutil.ReadAll(res.Body)
         if err != nil {
          return nil, fmt.Errorf("Error reading response body:", err)
         }

		// Unmarshal the raw response into a map
        if err := json.Unmarshal(body, &responseMap); err != nil {
            return nil, fmt.Errorf("failed to unmarshal response: %w", err)
        }

        // Modify the response as needed
        if val, ok := responseMap["app_name"]; ok { 
            if strVal, isString := val.(string); isString {
                cleanedValue := strings.Trim(strVal, "[]")
                responseMap["app_name"] = cleanedValue 
            }
        }

        // Marshal the modified response back to json.RawMessage
        modifiedResponse, err := json.Marshal(responseMap)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal modified response: %w", err)
        }
		return modifiedResponse, nil
    }
	return res, nil
}