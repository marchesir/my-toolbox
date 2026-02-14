use std::{any::type_name_of_val, io};

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

fn get_input(input: &mut String) {
    io::stdin()
        .read_line(input)
        .expect("Failed to read line");  
}

fn main() {
    let mut input = String::new();
   
    loop {
        println!("==>Please input the temperature and indciate the scale to use");
        println!("==>Example of valid inputs are 20C or 70F?");   
        // Clear string to avoid old values being appended.
        input.clear();
        get_input(&mut input);
        if input.trim().to_lowercase().ends_with('c') {
            // Convert to f32 so we can work in temperature unit.
            let temp: f32 = match &input[0..input.trim().len() -1].parse() {
               Ok(num) => *num,
               Err(_) => continue,
            };
            println!("\nTemperature Conversion");
            println!("======================");
            println!("Celsius:{}",input.trim().to_uppercase());
            println!("Fahrenheit:{}F",celsius_2_fahrenheit(temp));
            println!("======================");
        }
        else if input.trim().to_lowercase().ends_with('f' ) {
            // Convert to f32 so we can work in temperature unit.
            //input.trim().to_ascii_lowercase().pop()
            let temp: f32 = match &input[0..input.trim().len() -1].parse() {
               Ok(num) => *num,
               Err(_) => continue,
            };
            println!("\nTemperature Conversion");
            println!("======================");
            println!("Fahrenheit:{}",input.trim().to_uppercase());
            println!("Celsius:{}C",fahrenheit_2_celsius(temp));
            println!("======================");
        }
        else {
            continue;
        }
        break;
    }
}
