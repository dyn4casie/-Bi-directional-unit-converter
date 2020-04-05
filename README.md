 build a bi-directional (converts both ways) unit converter.

Create a struct called Converter, It will contain methods that can convert various units in a bi-directional way, that is, if we attach a method to the Converter struct called FeetToCentimeter that converts every value in feet to centimeter, we must have a corresponding method that converts centimeter to feet called CentimeterToFeet. Every unit must be represented as a named type.

The following demonstrate the steps to create a FeetToCentimeter & CentimeterToFeet:

Create a struct type called Converter.

Create the named types for both units:

    type Feet float64    type Centimeter float64

Create FeetToCentimeter and CentimeterToFeet methods with both attached to the Converter struct.

Add a line in your main function that prints the result with sample data.

    func main() {      // ...      fmt.Println(CentimeterToFeet(13.50))      // ..    }

Hint 2: your method signatures should look like this func (cvr Converter) CentimeterToFeet(c Centimeter) Feet { // code }

Compulsory exercises
Minutes to Seconds
Seconds to Milliseconds
Centimeter to Feet
Feet to Centimeter
Celsius to Fahrenheit
Fahrenheit to Celsius
Radian to Degree (π radians is equal to 180°)
use the math package (just like we did with "fmt" package) from the standard library to access PI (math.Pi) value for better precision.
Degree to Radian
Kilogram to Pounds
Pounds to Kilogram
Optional
Feel free to add extra conversions