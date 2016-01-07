package logger;

/*
 *  Tests that validate the directory scanner.
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
	"fmt"
	"strings"
	"testing"
)

var TEST_PATTERN = "tty.*"
var TEST_PATTERN_PREFIX = "tty"

func TestScanDirForStopStartDevices(t *testing.T) {
	start_stop_devices, err := FindMatchingDirEntries(
		DEVICE_DIRECTORY, TEST_PATTERN)
	if err != nil {
		t.Errorf("Error scanning /dev: %v.", err)
	}

	if len(start_stop_devices) == 0 {
		t.Error("No start stop devices found.")
	}
	
	for _, start_stop_device := range start_stop_devices {
		device_name := start_stop_device.Name()
		if !strings.HasPrefix(device_name, TEST_PATTERN_PREFIX) {
			t.Errorf("Device %s is not a start-stop device.", device_name)
		}
		fmt.Printf("%v -- OK.\n", start_stop_device);
	}
}
