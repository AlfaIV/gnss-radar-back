from pydantic import BaseModel


class TLE:
    def __init__(self, name):
        self._name = name
        self._line1 = None
        self._line2 = None

    @property
    def name(self):
        return self._name

    @name.setter
    def name(self, value):
        self._name = value

    @property
    def line1(self):
        return self._line1

    @line1.setter
    def line1(self, value):
        self._line1 = value

    @property
    def line2(self):
        return self._line2

    @line2.setter
    def line2(self, value):
        self._line2 = value


class RadarPosition(BaseModel):
    x_position: float
    y_position: float
    z_position: float
