fn main() {
    // shadowing variables using the same name, cant be used with 'mut' variables.
    let x = 5;
    let x = x + 1;
    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}");
    }
    println!("The value of x is: {x}");  // as inner usage of x is now out of scope 6 is output.
}