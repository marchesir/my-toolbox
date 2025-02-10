use std::io;
use rand::Rng;
use std::cmp::Ordering;

fn main() {
    println!("Guess the number!");

    // Get random number between 1-100.
    let secret_number = rand::thread_rng().gen_range(1..=100);
    println!("The secret number is: {secret_number}");

    loop {
        println!("Please input your guess?");
        // By default variables are immutable,so 'mut' keyword is used to indicate mutability.
        // 'new' is an asociated function that is implemented on a type, in this case a growable UTF8 String. 
        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess) // Pass by reference where to read into from stdin.
            .expect("Failed to read line");  // Error management.

        // Convert inputted string u32 and,reuses inputted guess variable using shadowing.  
        // All eror will be filtered to a continue and force a new input.  
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };
        println!("You guessed: {guess}");

        // Use match powerfull pattern matching to compare 2 values. 
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}