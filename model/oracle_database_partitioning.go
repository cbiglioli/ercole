// Copyright (c) 2022 Sorint.lab S.p.A.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package model

// OracleDatabasePartitioning holds the informations about a partitioning.
type OracleDatabasePartitioning struct {
	Owner         string  `json:"owner" bson:"owner"`
	SegmentName   string  `json:"segmentName" bson:"segmentName"`
	PartitionName string  `json:"partitionName" bson:"partitionName"`
	SegmentType   string  `json:"segmentType" bson:"segmentType"`
	Mb            float64 `json:"mb" bson:"mb"`
}
