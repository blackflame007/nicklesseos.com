// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package gamePlateform

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Show(gameName string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex justify-center items-center flex-col mb-5 py-12\"><h1 class=\"text-3xl sm:text-5xl mb-5\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(cases.Title(language.English, cases.Compact).String(gameName))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/gamePlateform/show.templ`, Line: 11, Col: 103}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><iframe id=\"gameIframeId\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/games/%s/index.html", gameName))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/views/gamePlateform/show.templ`, Line: 12, Col: 79}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"864\" height=\"936\" class=\"max-w-screen-sm sm:max-w-full\"></iframe><script>\ndocument.addEventListener(\"DOMContentLoaded\", () => {\n    var iframe = document.getElementById(\"gameIframeId\");\n\n    iframe.onload = () => {\n        try {\n            var iframeWindow = iframe.contentWindow;\n\n            var request = iframeWindow.indexedDB.open(\"/userfs\", 21); // Adjust version as needed\n\n            request.onsuccess = function(event) {\n                var db = event.target.result;\n                var transaction = db.transaction([\"FILE_DATA\"], \"readonly\");\n                var objectStore = transaction.objectStore(\"FILE_DATA\");\n\n                var fileRequest = objectStore.get(\"/userfs/godot/app_userdata/flappyBirdClone/flappybird.json\");\n\n                fileRequest.onsuccess = function(event) {\n                    if (fileRequest.result && fileRequest.result.contents) {\n                        // Convert Int8Array to a string\n                        var contentsArray = Array.from(new Int8Array(fileRequest.result.contents));\n                        var jsonString = String.fromCharCode.apply(null, contentsArray);\n                        try {\n                            // Parse the JSON string\n                            var jsonData = JSON.parse(jsonString);\n                            console.log(\"High Score:\", jsonData.high_score);\n                        } catch (parseError) {\n                            console.error(\"Error parsing JSON data:\", parseError);\n                        }\n                    } else {\n                        console.log(\"No result found for the specified key.\");\n                    }\n                };\n\n                fileRequest.onerror = function(event) {\n                    console.error(\"Error reading file data:\", event);\n                };\n            };\n\n            request.onerror = function(event) {\n                console.error(\"Error opening IndexedDB:\", event);\n            };\n        } catch (error) {\n            console.error(\"Error accessing iframe content:\", error);\n        }\n    };\n});\n</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
