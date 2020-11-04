package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "app.js",
		FileModTime: time.Unix(1604454102, 0),

		Content: string("const App = {\n    data() {\n        return {\n        counter: 0\n        }\n    },\n    template: \"<log-portal/>\"\n}\nconst app = Vue.createApp(App)\napp.component('log-portal', {\n    template: \"<div>ok</div>\"\n})\n \napp.mount('#app')\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1604452212, 0),

		Content: string("<html>\n    <head>\n        <script src=\"https://unpkg.com/vue@next\"></script>\n    </head>\n    <body>\n        <div id=\"app\">\n        </div>\n        <script src=\"app.js\" ></script>\n    </body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1604454102, 0),
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
		Time: time.Unix(1604454102, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"index.html": file3,
		},
	})
}
