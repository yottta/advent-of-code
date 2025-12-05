pub fn run() {
    day1::part1::run();
    day1::part2::run();
}

mod day1 {
    pub(crate) mod part1 {
        use crate::day1::day1::normalize;
        use crate::utils::utils::{must_parse_to_i32, read_file_from_env};

        // 453 too low
        // 1059 - correct
        pub fn run() {
            let lines = read_file_from_env(String::from("DAY1_INPUT"));

            let mut current = 50;
            let mut no_of_zeros = 0;
            for (_, line) in lines.iter().enumerate() {
                current = go_to_next(line, current);
                if current == 0 {
                    no_of_zeros += 1;
                }
                // println!("prev {} act {} result {}", prev, line, current);
            }
            println!("password is {}", no_of_zeros);
        }

        fn go_to_next(input: &String, previous: i32) -> i32 {
            let dir = &input[0..1];
            let steps_str = &input[1..];
            let mut steps_no = must_parse_to_i32(steps_str);
            let mut no_over_99 = 0;
            if steps_no > 99 {
                no_over_99 = steps_no / 99;
                steps_no = steps_no % 99;
            }
            let delta = if dir == "R" { steps_no } else { steps_no * -1 };
            let mut temp = normalize(previous + delta);
            for _ in 0..no_over_99 {
                temp = normalize(previous + delta);
            }
            if dir == "R" {
                return normalize(temp - no_over_99);
            }
            normalize(temp + no_over_99)
        }
    }
    pub(crate) mod part2 {
        use crate::utils::utils::{must_parse_to_i32, read_file_from_env};

        // Part2
        // 6312 - too high
        // 6305 - right
        pub fn run() {
            let lines = read_file_from_env(String::from("DAY1_INPUT"));

            let mut current = 50;
            let mut no_of_zeros = 0;
            for (_, line) in lines.iter().enumerate() {
                let (new_current, clicks) = go_to_next_bruteforce(line, current);
                current = new_current;
                no_of_zeros += clicks;
            }
            println!("password is {}", no_of_zeros);
        }

        fn go_to_next_bruteforce(input: &String, previous: i32) -> (i32, i32) {
            let mut clicks = 0;
            let dir = &input[0..1];
            let steps_str = &input[1..];
            let steps_no = must_parse_to_i32(steps_str);
            if dir == "R" {
                let mut res = previous;
                for _ in 0..steps_no {
                    res += 1;
                    if res > 99 {
                        res = 0;
                        clicks += 1;
                    }
                }
                return (res, clicks);
            }
            let mut res = previous;
            for _ in 0..steps_no {
                res -= 1;
                if res == 0 {
                    clicks += 1;
                }
                if res < 0 {
                    res = 99
                }
            }
            (res, clicks)
        }
    }

    pub fn normalize(input: i32) -> i32 {
        if input > 99 {
            return input - 99 - 1;
        } else if input < 0 {
            return 99 + input + 1;
        }

        input
    }
}