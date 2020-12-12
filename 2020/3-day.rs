use std::env;
use std::fs;

/* Instruction part 1 : You start on the open square (.) in the top-left corner and need to reach the bottom (below the bottom-most row on your map).
The toboggan can only follow a few specific slopes (you opted for a cheaper model that prefers rational numbers); start by counting all the trees you would encounter for the slope right 3, down 1:
From your starting position at the top-left, check the position that is right 3 and down 1. Then, check the position that is right 3 and down 1 from there, and so on until you go past the bottom of the map.

INFO : 
    - the same pattern repeats to the right many times
    - 

Exemple : The locations you'd check in the above example are marked here with O where there was an open square and X where there was a tree:
..##.........##.........##.........##.........##.........##.......  --->
#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..  --->
.#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.  --->
..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#  --->

*/
fn main() {
    let args: Vec<String> = env::args().collect();

    let filename = &args[1];

    let contents = fs::read_to_string(filename)
        .expect("Something went wrong reading the file");

    // set slope
    let move_h = 3;
    let move_v = 1;

    // set point
    let mut origin = 0;

    // var to be used when we will down more than once
    let mut skipped = 0;
   
    let mut number_tree = 0;
    let mut end_line = 0;
    
    /* Algo : 
        - skipped the first origin test
        - moving into the line by adding move_h for each iteration
        - checked of current origin : 
            - if origin < len line 
                normal check
            - else if origin > len line
                - add same pattern upto len of line > origin
                - normal check
    */ 
    for line in contents.lines() {
        end_line = line.len();
        if origin != 0 && origin < end_line {
            if line.chars().nth(origin).unwrap() == '#' {
                number_tree+=1;
            }
        }
        else if origin >= end_line {
            let mut str_buff = "".to_owned();
            while origin >= end_line {
                str_buff += line;
                end_line = str_buff.len();
            }
            if str_buff.chars().nth(origin).unwrap() == '#' {
                number_tree+=1;
            }
        } 
        origin += move_h;
    };
    println!("number tree : {}",number_tree)
}

