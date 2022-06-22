/* Coded by Faralaks https://github.com/Faralaks
This packege provides interface and implementation which allows u extract data from html message history
*/
package html_extractor

import (
	"errors"
	"os"
)

func Exists(p string) (bool, error) {
	if _, err := os.Stat(p); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
	}
}
