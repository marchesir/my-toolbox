use std::io;

fn fahrenheit_2_celsius(f: f32) -> f32 {
  // C=(F−32)×5/9
  let c = (f - 32.0) * (5.0/9.0);   
  return c; 
}

fn celsius_2_fahrenheit(c: f32) -> f32 {
    // F=(Cx5/9)+32
    let f = (c * 9.0/5.0) + (32.0);
    return f; 
}

fn main() {
    let mut temp = String::new();

    loop {
        println!("Please input the temperature?");   
        
        io::stdin()
          .read_line(&mut temp)
          .expect("Failed to read line");  

        // convert to f32 so we can work in temperature unit.
        let temp: f32 = match temp.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        let f = fahrenheit_2_celsius(temp);
        println!("Your temperature in celsius is {}C",f);
        let c = celsius_2_fahrenheit(f);
        println!("Your temperature in fahrenheit is {}F",c);
        
        break;

    }
}
