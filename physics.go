package physics

//calculates force in Newtons given a mass in kg and an acceleration in m/(s^2)
func Force(mass float64, acceleration float64) float64{
  return mass * acceleration
}