package logger;

/*
 *  Scanner that finds all files whose names match a template.
 * 
 *  Copyright (C) 2016  The Winer Observatory
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *  For further information, please contact the Winer Observatory at
 *  http://www.winer.org. 
 */
 
import (
	"io/ioutil"
	"os"
	"regexp"
)

func FindMatchingDirEntries(dirname, match string) ([]os.FileInfo, error) {
	entries, err := ioutil.ReadDir(dirname)
	if err != nil {
		return entries, err
	}

	pregex, err := regexp.Compile(match)
	if err != nil {
		return nil, err
	}
	
	matching_entries := entries[:0]
	for _, entry := range entries {
		if pregex.MatchString(entry.Name()) {
			matching_entries = append(matching_entries, entry);
		}
	}
	
	return matching_entries, err
} 
