from datetime import datetime, timedelta, timezone
from app.core.config import configs 
from fastapi import HTTPException, status
from skyfield.api import load, EarthSatellite
from skyfield.toposlib import wgs84
from skyfield.positionlib import Geocentric
import re

from app.entities.sattellites import TLE
from app.schemas.sattelites_position import (
    SatellitesTimeRequest,
    SatellitesPositionResponce,
    SatellitesTimeResponce,
    RadarPositionGeograthRequest,
    VisionTime,
    SatelliteTime,
    SatellitePosition,
)

from app.repositories import S3Repository


class SatellitesPositions:
    def __init__(self, s3_repository: S3Repository):

        self.s3_repository =  s3_repository

        self.tle_file = ""
        self.TLE_array = []

        self.ts = load.timescale()
        self.time_step = timedelta(minutes=5)

        self.mask_visible = 0

    def load_sattelites_tle(self) -> list:
        file = self.s3_repository.get_tle(self.tle_file).tle_file
        lines = [line.strip() for line in file.split('\n') if line.strip()]
        TLE_array = []
        try: 

            i = 0
            while i < len(lines):

                if i + 2 >= len(lines):
                    break

                name_line = lines[i]
                line1 = lines[i + 1]
                line2 = lines[i + 2]

                if line1.startswith("1 ") and line2.startswith("2 "):
                    satellite_name, grouping = self.get_name(name_line)
                    
                    if(satellite_name == "" or grouping== ""):
                        i += 3
                        continue

                    print(f"{satellite_name} ({grouping})")
                    TLE_array.append({
                        "name": satellite_name,
                        "group": grouping,
                        "line1": line1,
                        "line2": line2
                    })
                    i += 3
                else:
                    i += 3

        except Exception as e:
            raise HTTPException(
                status_code=status.HTTP_409_CONFLICT,
                detail=f"Error deceptions: {e}",
                headers={"X-Error": "Custom header", "Error-type": "Parse TLE"},
            )
        return TLE_array

    def get_name(self,name_sat_group_string : str):

        name_sat_group = name_sat_group_string.strip().split(" ")

        if name_sat_group is None or len(name_sat_group) < 2:
            return ("", "")
        
        group = name_sat_group[0]
        name = ""

        full_name_satellite = " ".join(name_sat_group[1:])
        number = re.search(r'\((.*?)\)', full_name_satellite)
        
        if number:
            full_name_satellite = number.group(1)
            if "GALILEO" in full_name_satellite:
                group = "GALILEO"
                correct_name = full_name_satellite.split()
                if len(correct_name) == 2:
                    name = " ".join(full_name_satellite.split()[1:])
            else:
                name = re.sub(r'[^\d]', "", number.group(1))
                if group == "GALILEO":
                    name = ""

        if "COSMOS" in group:
            group = "GLONASS"
        elif "BEIDOU" in group:
            group = "BEIDOU"
        elif "GALILEO" in group:
            group = "GALILEO"
        elif "GPS" in group:
            group = "GALILEO"
        else:
            group = ""

        return (name, group)  
     

    def get_sattelites_positions(
        self, radar: RadarPositionGeograthRequest,
    ) -> SatellitesPositionResponce:

        satellite_positions = []

        if radar.tle_file != self.tle_file:
            self.tle_file = radar.tle_file
            self.TLE_array = self.load_sattelites_tle()

        
        dt = datetime.strptime(radar.inspection_time, configs.DATETIME_FORMAT)
        unix_time = int(dt.timestamp())

        for satellite in self.TLE_array:
            if (len(radar.satellites_name) == 0) or (satellite.name in radar.satellites_name):
                sattelite_props = self.get_sattelite_positions(
                    unix_time, satellite, radar
                )

                satellite_positions.append(
                    SatellitePosition(
                        Group = satellite.group,
                        Name = satellite.name,
                        Azimuth = sattelite_props["Azimuth"],
                        Elevation = sattelite_props["Elevation"],
                        Range = sattelite_props["Range"],
                    )
                )

        return SatellitesPositionResponce(
            Satellites = satellite_positions,
        )

    def get_sattelite_positions(
        self, current_time: datetime,  satellite: object, observer: RadarPositionGeograthRequest
    ) -> SatellitePosition:
        
    
        current_time = datetime.fromtimestamp(current_time, tz=timezone.utc)
        skyfield_time = self.ts.from_datetime(current_time)

        observer = wgs84.latlon(observer.radar_latitude, observer.radar_longitude, observer.radar_height)
    
        satellite = EarthSatellite(satellite.line1.strip(), satellite.line2.strip(), satellite.name.strip())
        position = satellite.at(skyfield_time)
        subpoint = wgs84.subpoint(position)

        difference = satellite - observer
        topocentric = difference.at(skyfield_time)
        
        elevation, azimuth, distance = topocentric.altaz()

        return {
            "Azimuth": round(azimuth.degrees, 2),
            "Range": round(distance.km, 2),
            "Elevation": round(elevation.degrees, 2),
            "Longitude": round(subpoint.longitude.degrees, 2),
            "Latitude": round(subpoint.latitude.degrees, 2),
            "Height": round(subpoint.elevation.km, 2),
        }
    

    def get_sattelites_times_vison(
        self, radar: SatellitesTimeRequest
    ) -> SatellitesTimeResponce:
        
        if radar.tle_file != self.tle_file:
            self.tle_file = radar.tle_file
            self.TLE_array = self.load_sattelites_tle()
        
        dt_start = datetime.strptime(radar.begin_time, configs.DATETIME_FORMAT)
        unix_time_start = int(dt_start.timestamp())
        begin_time = datetime.fromtimestamp(unix_time_start, tz=timezone.utc)

        dt_end = datetime.strptime(radar.end_time, configs.DATETIME_FORMAT)
        unix_time_end = int(dt_end.timestamp())
        end_time = datetime.fromtimestamp(unix_time_end, tz=timezone.utc)
        
        observer = wgs84.latlon(
            radar.radar_latitude,
            radar.radar_longitude,
            radar.radar_height
        )

        satellite_time = []
           

        for satellite in self.TLE_array:
            
            if not radar.satellites_name or radar.satellites_name.__contains__(satellite.name):
                satellite_data = EarthSatellite(
                    satellite.line1.strip(),
                    satellite.line2.strip(),
                    satellite.name
                )

                visibility_times = self.visibility_times(
                    begin_time, end_time, satellite_data, observer
                )
                
                if (len(visibility_times) > 0):
                    satellite_time.append(
                        SatelliteTime(
                            group = satellite.group,
                            name = satellite.name,
                            intervals=visibility_times,
                        )
                )


        return SatellitesTimeResponce(Satellites=satellite_time)


    def visibility_times(
        self,
        start_time: datetime,
        end_time: datetime,
        satellite: EarthSatellite,
        observer: Geocentric,
        min_elevation = 0,
    ) -> list[VisionTime]:
        
        current_time = start_time
        visibility_periods = []
        in_visibility = False
        period_start = None
        
        while current_time <= end_time:
            skyfield_time = self.ts.from_datetime(current_time)
            
            difference = satellite - observer
            topocentric = difference.at(skyfield_time)
            elevation = topocentric.altaz()[0].degrees
            
            
            if elevation >= min_elevation:
                if not in_visibility:
                    
                    period_start = current_time
                    in_visibility = True
            else:
                if in_visibility:
                    
                    visibility_periods.append(
                        VisionTime(
                            startDatetime=period_start.strftime(configs.DATETIME_FORMAT),
                            endDatetime=current_time.strftime(configs.DATETIME_FORMAT),
                        )
                    )
                    in_visibility = False
            
            current_time += self.time_step
        
        if in_visibility:
            visibility_periods.append(
                VisionTime(
                    startDatetime=period_start.strftime(configs.DATETIME_FORMAT),
                    endDatetime=end_time.strftime(configs.DATETIME_FORMAT),
                )
            )
            
        return visibility_periods
    