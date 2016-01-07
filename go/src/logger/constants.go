package logger;

/*
 *  Constants used to log shadow band data.
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

var WHO_ARE_YOU = 'w'    // Request data source name
var MEASURE_LIGHT = 'm'  // Measure light distribution

var DEVICE_DIRECTORY = "/dev"    // Directory of I/O devices
var DATA_SOURCE_RE = "ttyUSB.*"  // Data source selector regex
