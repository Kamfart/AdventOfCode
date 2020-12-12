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
    // let move_h = 1;
    // let move_v = 2;
    let mut _number_tree = 0;
    
    _number_tree = solve(&contents, 1, 1);
    _number_tree *= solve(&contents, 3, 1);
    _number_tree *= solve(&contents, 5, 1);
    _number_tree *= solve(&contents, 7, 1);
    _number_tree *= solve(&contents, 1, 2);
    println!("number tree : {}",_number_tree);
}

fn solve(contents: &String, move_h: usize, move_v: usize) -> u64 {
    // set point
    let mut origin = 0;

    // var to be used when we will down more than once
    let mut skipped = 0;
   
    let mut number_tree = 0;
    let mut _end_line = 0;
    
    /* Algo : 
        - skipped the first origin test
        - skipped the current iteration if move_v is superior to 0 (to simulate down)
            - skipped is decrementad at the end of each iteration,
            In fact, if skipped == 1, the instruction "if skipped > 0" will never be true.
        - moving into the line by adding move_h for each iteration
        - checked of current origin : 
            - if origin < len line 
                normal check
            - else if origin > len line
                - add same pattern upto len of line > origin
                - normal check
    */ 
    for line in contents.lines() {
        if skipped > 0 {
            skipped -= 1;
            continue;
        } else {
            skipped = move_v;
            _end_line = line.len();
            if origin != 0 && origin < _end_line {
                if line.chars().nth(origin).unwrap() == '#' {
                    number_tree+=1;
                }
            }
            else if origin >= _end_line {
                let mut str_buff = "".to_owned();
                while origin >= _end_line {
                    str_buff += line;
                    _end_line = str_buff.len();
                }
                if str_buff.chars().nth(origin).unwrap() == '#' {
                    number_tree+=1;
                }
            } 
            origin += move_h;
            skipped -= 1;
        }
    };

    return number_tree;
}

