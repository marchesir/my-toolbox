fn main() {
    // Data copied to heap with metadata on the stack.
    let s1 = String::from("hello");

    let mut s2 = s1; // make mutable so can be modified.
    // Heap data in s1 is now pointed to by s2, known as shallow copy.
    // Allocater will invalidate s2 to avoid double drop as both s1,s2 point to same heap data.
    // This can be changed to a deep copy whcih will have 2 seperated heap allocations using
    // clone() funmction, e.g. let s2 = s1.clone();
    
    // Wont compile as s1 is invalidated, s2 or clone() should be used.
    //println!("{}!", s1);

    // references, no ownership taken.
    change(&mut s2);
    let len = calculate_length(&s2); 
    println!("The length of '{}' is {}.", s2, len);

    // If we try to borrow 2 concurrent mut references the compiler will throe an error:
    //let r1 = &mut s2;
    //let r2 = &mut s2;
    //println!("{}, {}", r1, r2);
    // The reason of this restriction is to prvent data races.

    // We also cannot have a mutable reference while we have an immutable one to the same value:
    //let r1 = &s2; 
    //let r2 = &s2; 
    //let r3 = &mut s2;
    //println!("{}, {}, {}", r1, r2, r3);

}

fn calculate_length(s: &String) -> usize {
    s.len()
} // s goes out of scope but drop not called as no ownershiped taken as refernce used.


fn change(some_string: &mut String) { // refernece,no ownership taken but must be mutable to modify.
    some_string.push_str(", world!");
}

// The compliler will throw error if it detects dangling reference:
//fn dangle() -> &String { // dangle returns a reference to a String.
//    let s = String::from("hello");   // s is a new String on heap.
//    &s  // we return a reference to the String, s.
//}  // Danger: we try to return a refernece that has been deallocated, it wont compile.