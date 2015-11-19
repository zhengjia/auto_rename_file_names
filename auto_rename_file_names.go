/*
The package auto_rename_file_names changes a file name to a unique value based on previously checked names.

Internally it keeps an slice of known file paths and if a name conflict is detected, an affix is added to the file name.

Example

	import "github.com/zhengjia/auto_rename_file_names"

	renamer := auto_rename_file_names.New()

	// you can repetitively call the method Get:
	// returns "/test1/1.txt"
	renamer.Get("/test1/1.txt")
	// returns "/test2/1.txt"
	renamer.Get("/test2/1.txt")
	// returns "/test1/1(1).txt"
	renamer.Get("/test1/1.txt")

	// Cleared previously checked file names and start from fresh
	renamer.Reset()
	// returns "/test1/1.txt"
	renamer.Get("/test1/1.txt")

*/
package auto_rename_file_names

import (
	"fmt"
	"path/filepath"
	"strings"
)

/*
You can use these consts to set the type and location of the affix.

By default it will be numeric and appended at the end of the file name

Example

	renamer.Type = StringSuffix
*/
const (
	NumericSuffix = iota
	NumericPrefix
	StringPrefix
	StringSuffix
)

/*
StringAffix: value for the affix.
For example, if it's set to "Copy", then the file name will be OriginalFileNameCopy
Only works when Type is StringPrefix or StringSuffix

Seperator: seperator between affixes
For example, if it's set to "_", then the file name will be OriginalFileName-Copy_Copy
Only works when Type is StringPrefix or StringSuffix

Connector: string inserted between original file name and the affix.
For example, if it's set to "-", then the file name will be OriginalFileName-Copy

NumericFormat: The template to generate the numeric affix. Default value is (%d)
For example, if it's set to "(%d)", then the file name will be OriginalFileName-(1)
Only works when Type is NumericPrefix or NumericSuffix
*/
type Renamer struct {
	Type          uint8
	StringAffix   string
	elems         []*elem
	Connector     string
	Seperator     string
	NumericFormat string
}

type elem struct {
	conflictTimes int64
	path          string
}

// Return a new object that can be called to get back unique file names
func New() (renamer *Renamer) {
	renamer = new(Renamer)
	renamer.NumericFormat = "(%d)"
	return
}

// Get back a unique file name
func (renamer *Renamer) Get(original_file_path string) (unique_file_path string) {
	var no_conflict bool
	original_file_path = strings.ToLower(original_file_path)
	unique_file_path, no_conflict = renamer.get_unique(original_file_path)
	if no_conflict {
		new_elem := &elem{
			path: original_file_path,
		}
		renamer.elems = append(renamer.elems, new_elem)
	}
	return
}

// Clear previously stored file names
func (renamer *Renamer) Reset() {
	renamer.elems = []*elem{}
}

func (renamer *Renamer) get_unique(original_file_path string) (unique_file_path string, no_conflict bool) {
	if len(renamer.elems) == 0 {
		unique_file_path = original_file_path
		no_conflict = true
	} else {
		for _, e := range renamer.elems {
			if strings.Compare(e.path, original_file_path) != 0 {
				unique_file_path = original_file_path
				continue
			} else {
				e.conflictTimes += 1

				var addition string
				extension := filepath.Ext(original_file_path)
				dir := filepath.Dir(original_file_path)
				file_name_without_extension := strings.TrimSuffix(filepath.Base(original_file_path), extension)
				path_without_extension := strings.TrimSuffix(original_file_path, extension)
				addition = e.getFileNameAddition(renamer)
				if renamer.Type == StringSuffix || renamer.Type == NumericSuffix {
					unique_file_path = path_without_extension + renamer.Connector + addition + extension
				} else if renamer.Type == StringPrefix || renamer.Type == NumericPrefix {
					unique_file_path = filepath.Join(dir, addition+renamer.Connector+file_name_without_extension+extension)
				}
				return
			}
		}
		no_conflict = true
	}
	return
}

func (e *elem) getFileNameAddition(renamer *Renamer) (addition string) {
	if renamer.Type == StringSuffix || renamer.Type == StringPrefix {
		ary := []string{}
		for i := 0; i < int(e.conflictTimes); i++ {
			ary = append(ary, renamer.StringAffix)
		}
		addition = strings.Join(ary, renamer.Seperator)
	} else if renamer.Type == NumericSuffix || renamer.Type == NumericPrefix {
		addition = fmt.Sprintf(renamer.NumericFormat, e.conflictTimes)
	}
	return
}
