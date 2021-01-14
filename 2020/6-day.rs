use std::env;
use std::fs;

/*  Instruction part 1 : A form asks a series of 26 yes-or-no questions marked a through z. 
All you need to do is identify the questions for which anyone in your group answers "yes".
For each of the people in their group, you write down the questions for which they answer "yes", one per line.

For each group, count the number of questions to which ANYONE answered "yes". What is the sum of those counts?

Exemple : 
    - abc

    - a
    - b
    - c 
    
    - ab
    - ac
    => 3 + 3 + 3 = 6

Instruction part 2 : Same as part 1 but you need to identify the questions to which EVERYONE answered "yes"!

For each group, count the number of questions to which everyone answered "yes". What is the sum of those counts?

Same Exemple : 
    - 3 + 0 + 1 = 4
*/
fn main() {
    let args: Vec<String> = env::args().collect();

    let filename = &args[1];

    let contents = fs::read_to_string(filename)
        .expect("Something went wrong reading the file");

    let _sum_of_yes_ans = solve(&contents, 1);
    println!("Part 1 : number of yes answer : {}",_sum_of_yes_ans);
    
    let _sum_of_yes_ans = solve(&contents, 2);
    println!("Part 2 : number of yes answer : {}",_sum_of_yes_ans);
}

fn solve(contents: &String, part: u8) -> u64 {
    // init buffer
    let mut nb_soya : u64 = 0;

    let mut set_grp_ans = Vec::new();   
    
    /* Algo : Add each answer of one group into a buffer vector.
    Once a blank line is read, process then clear the buffer.        
    */ 
    for line in contents.lines() {
        if line.is_empty() {
            nb_soya += count_yes_answer(&set_grp_ans, part) as u64;
            set_grp_ans.clear();
        } else {
            set_grp_ans.push(line);
        }
    };
    nb_soya += count_yes_answer(&set_grp_ans, part) as u64;

    return nb_soya;
}

fn count_yes_answer(set: &Vec<&str>, part: u8) -> u8 {
    // init buffer answer
    let mut nb_y = 0;

    // create a vector which represent the poll. Where Vec[0]..Vec[25] = question 'a'..'z'
    let mut poll: Vec<u8> = Vec::new();

    // init all element to 0. Maybe could be replaced by a direct assignation.
    init_poll(&mut poll);

    // fill vector with the given poll question
    fill_poll(set, &mut poll);

    // count all questions which ANYONE answered yes from a group
    if part == 1 {
        for q in poll.iter() {
            if q > &(0 as u8) {
                nb_y+=1
            }
        }
    }
    // count all questions which EVERYONE answered yes from a group 
    else if part == 2 {
        for q in poll.iter() {
            if q == &(set.len() as u8) {
                nb_y+=1
            }
        }
    }
    else {
        println!{"Error part must be 1 or 2"};
    }

    return nb_y;
}

fn init_poll(poll: &mut Vec<u8>){
    // let alphabet = String::from_utf8(
    //     (b'a'..=b'z').collect()
    // ).unwrap();     

    
    // let mut poll = HashMap::new();
    
    for _i in 0..26 {
        poll.push(0);
    }
}

/* Iterate over each elem of poll. 
If one elm is composed by more than one answer, iterate over it.
For each answer, convert the letter into ascii, substract 97 
and increment the value inside the dedicated memory inside the vector.
*/
fn fill_poll(set: &Vec<&str>, poll: &mut Vec<u8>){
    for elm in set.iter() {
        if elm.len() > 1 {
            for c in elm.chars() {
                // println!("{}",c as usize - 97);
                poll[c as usize - 97]+=1;
            }
        } else {
            // Possibly improvment here
            let buff: Vec<char> = elm.chars().collect();
            let c = buff[0];
            poll[c as usize - 97]+=1;
            // println!("{}", c as u8);
        }
    }
}
