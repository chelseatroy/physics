package physics

import "math"

const(
  g = 9.80
)

func Force(mass float64, acceleration float64, sigfig int) float64{
  return round(mass * acceleration, 0.5, sigfig)
}

func PendulumPeriod(length float64, sigfig int) float64{
  return round(2 * math.Pi * math.Sqrt(length/g), 0.5, sigfig)
}

func round(val float64, roundOn float64, places int ) (newVal float64) {
  var round float64
  pow := math.Pow(10, float64(places))
  digit := pow * val
  _, div := math.Modf(digit)
  _div := math.Copysign(div, val)
  _roundOn := math.Copysign(roundOn, val)
  if _div >= _roundOn {
    round = math.Ceil(digit)
  } else {
    round = math.Floor(digit)
  }
  newVal = round / pow
  return
}