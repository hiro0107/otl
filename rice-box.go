package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.js",
		FileModTime: time.Unix(1604467243, 0),

		Content: string("const App = {\n    template: \"<log-portal/>\"\n}\nconst app = Vue.createApp(App)\n\nconst repeat = async (time, func) => {\n    const f = async () => {\n        await func()\n        setTimeout(f, time)\n    }\n    await f()\n}\n\napp.component('log-portal', {\n    template: \"<div>ok</div>\",\n    setup(props) {\n        const logs = Vue.reactive([])\n        repeat(1000, async () => {\n            const response = await fetch(\"/logs\")\n            const json = await response.json()\n            console.log(json)\n        })\n        return {\n            logs\n        }\n    }\n})\n \napp.mount('#app')\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1604452212, 0),

		Content: string("<html>\n    <head>\n        <script src=\"https://unpkg.com/vue@next\"></script>\n    </head>\n    <body>\n        <div id=\"app\">\n        </div>\n        <script src=\"app.js\" ></script>\n    </body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1604467243, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "app.js"
			file3, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`http-files`, &embedded.EmbeddedBox{
		Name: `http-files`,
		Time: time.Unix(1604467243, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"index.html": file3,
		},
	})
}
