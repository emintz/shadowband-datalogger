package logger;

/*
 *  Data structure containing a single reading to be logged. One 
 *  instance will be allocated for each connected microcontroller.
 * 
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

/*
 * A complete set of data from a single microcontroller. Field contents
 * are:
 * 
 * Field              Contents
 * ---------------    --------------------------------------------------
 * instrument_id      an instrument identifier in the range [0 .. 511].
 *                    The ID will be burned into the microcontroller's
 *                    EEPROM, and will be unique for each copy.
 * data               10-bit light intensity values, one for each sensor
 *                    connected to the microcontroller.
 * duration_in_micros Data collection time in microseconds, accurate to
 *                    +/- 4 microseconds, per the microcontroller's 4
 *                    microsecond time resolution.
 */ 

type MeasurementConnection struct {
	instrument_id uint8        // Instrument ID, [0 .. 511] buried into the
	                           // microcontroller EEPROM
	reading_no int8            // A reading number in the range 
	                           // [0 .. 511], starting from 0. The 
	                           // serial number will be incremented
	                           // once for each read attempt, and will
	                           // roll over from 511 to 0. The idea is
	                           // to allign the following data on an
	                           // even byte boundary to improve
	                           // performance, and to have a way to
	                           // detect dropped data.
	sec int64                  // The time that the reading was
	                           // requested in UTC. See the Go Time
	                           // structure for details.
	nsec int32                 // Request time nanosecond offset. See
                               // the Go Time structure for details.
	data [8]uint16             // 10-bit light intensities 
	duration_in_micros uint32  // Data acquisition time in microseconds
}
