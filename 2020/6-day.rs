use std::env;
use std::fs;

/* Instruction part 1 : 

INFO : 
    - 

Exemple : 

*/
fn main() {
    let args: Vec<String> = env::args().collect();

    let filename = &args[1];

    let contents = fs::read_to_string(filename)
        .expect("Something went wrong reading the file");

    let _sum_of_yes_ans = solve(&contents);
    
    // _sum_of_yes_ans = solve(&contents);
    println!("Part 1 : number of yes answer : {}",_sum_of_yes_ans);
}

fn solve(contents: &String) -> u64 {
    // init buffer
    let mut nb_soya : u64 = 0;

    let mut set_grp_ans = Vec::new();   
    
    /* Algo : 
        
    */ 
    for line in contents.lines() {
        if line.is_empty() {
            nb_soya += count_yes_answer(&set_grp_ans) as u64;
            set_grp_ans.clear();
        } else {
            set_grp_ans.push(line);
        }
    };
    nb_soya += count_yes_answer(&set_grp_ans) as u64;

    return nb_soya;
}

fn count_yes_answer(set: &Vec<&str>) -> u8 {
    // init buffer
    let mut nb_y = 0;

    let mut poll: Vec<u8> = Vec::new();

    init_poll(&mut poll);

    // fill poll 
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

    // count yes answer
    for q in poll.iter() {
        if q > &(0 as u8) {
            nb_y+=1
        }
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

// fn create_poll() -> HashMap<&'static char, u8> {
//     let alphabet = (b'a'..=b'z')           // Start as u8
//         .map(|c| c as char)            // Convert all to chars
//         .filter(|c| c.is_alphabetic()) // Filter only alphabetic chars
//         .collect::<Vec<_>>();          // Collect as Ve
//     let mut poll = HashMap::new();
    
//     for letter in alphabet.iter() {
//         poll.insert(letter, 0);
//     }

//     return poll;
// }
