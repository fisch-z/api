/*
 *  Copyright (C) 2021 7Cav.us
 *  This file is part of 7Cav-API <https://github.com/7cav/api>.
 *
 *  7Cav-API is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  7Cav-API is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with 7Cav-API. If not, see <http://www.gnu.org/licenses/>.
 */

package milpacs

type Roster struct {
	RosterID          uint64 `gorm:"primaryKey"`
	Title             string
	Description       string
	CreateDate        uint
	LastUpdateDate    uint
	DisplayOrder      uint
	Active            bool
	PositionCache     string
	FieldCache        string
	DefaultPositionId uint
	ExtraGroupIds     string

	Profiles []Profile
}

func (Roster) TableName() string {
	return "xf_nf_rosters"
}