pub fn run() {
    day2::part1::run();
    day2::part2::run();
}

mod day2 {
    pub(crate) mod part1 {
        use crate::day2::day2::range_parser::{parse_ranges, Range};
        use crate::utils::utils::{read_file_from_env};

        // Part1
        // 38310256125 - correct
        pub fn run() {
            let lines = read_file_from_env(String::from("DAY2_INPUT"));
            for (_, line) in lines.iter().enumerate() {
                match parse_ranges(line) {
                    Ok(ranges) => {
                        let mut sum: u64 = 0;
                        for r in ranges {
                            let invalid_ids = process_range(r);
                            let sum_of: u64 = invalid_ids.iter().sum();
                            sum+= sum_of;
                        }
                        println!("invalid ids sum: {}", sum)
                    },
                    Err(e) => panic!("failed parsing ranges: {:?}", e),
                }
            }
        }
        fn process_range(input: Range) -> Vec<u64> {
            let mut res = Vec::new();
            for val in input.start..input.stop+1 {
                let val_str = val.to_string();
                let full_len = val_str.len();
                if full_len % 2 != 0 {
                    continue
                }
                let first_part = &val_str[0..full_len/2];
                let second_parth = &val_str[full_len/2..];
                if first_part == second_parth {
                    res.push(val);
                }
            }
            res
        }
    }
    pub(crate) mod part2 {
        use crate::day2::day2::range_parser::{parse_ranges, Range};
        use crate::utils::utils::read_file_from_env;

        // Part2
        // 58961152806 - correct
        pub fn run() {
            let lines = read_file_from_env(String::from("DAY2_INPUT"));
            for (_, line) in lines.iter().enumerate() {
                match parse_ranges(line) {
                    Ok(ranges) => {
                        let mut sum: u64 = 0;
                        for r in ranges {
                            let invalid_ids = process_range(r);
                            let sum_of: u64 = invalid_ids.iter().sum();
                            sum+= sum_of;
                        }
                        println!("invalid ids sum: {}", sum)
                    },
                    Err(e) => panic!("failed parsing ranges: {:?}", e),
                }
            }
        }
        fn process_range(input: Range) -> Vec<u64> {
            let mut res = Vec::new();
            for val in input.start..input.stop+1 {
                if process_id(val) {
                    res.push(val);
                }

            }
            res
        }
        fn process_id(val: u64) -> bool {
            let val_str = val.to_string();
            let full_len = val_str.len();
            let mut valid_group_sizes: Vec<usize> = Vec::new();
            for i in 1..full_len / 2 + 1 {
                if full_len % i == 0 {
                    valid_group_sizes.push(i);
                }
            }
            for (_, &i) in valid_group_sizes.iter().enumerate() {
                let mut previous_equal = false;
                let groups = full_len/i-1;
                for j in 0..groups {
                    let first_part_start = j*i;
                    let first_part_end = (j+1)*i;
                    let first_part = &val_str[first_part_start..first_part_end];
                    let second_part_start = (j+1)*i;
                    let second_part_end = second_part_start + (first_part_end-first_part_start);
                    let second_part = &val_str[second_part_start..second_part_end];
                    if j == 0 && first_part == second_part {
                        previous_equal = true;
                        continue
                    }
                    previous_equal = previous_equal && first_part == second_part;
                    if !previous_equal {
                        break
                    }
                }
                if previous_equal {
                    return true;
                }
            }
            false
        }
    }

    mod range_parser {
        #[derive(Debug)]
        pub struct Range {
            pub start: u64,
            pub stop: u64,
        }

        #[derive(Debug)]
        pub enum RangeParseError {
            InvalidFormat,
            ParseError(std::num::ParseIntError),
        }

        impl From<std::num::ParseIntError> for RangeParseError {
            fn from(err: std::num::ParseIntError) -> Self {
                RangeParseError::ParseError(err)
            }
        }

        pub fn parse_ranges(input: &str) -> Result<Vec<Range>, RangeParseError> {
            input
                .split(',')
                .map(|segment| {
                    let parts: Vec<&str> = segment.split('-').collect();
                    if parts.len() != 2 {
                        return Err(RangeParseError::InvalidFormat);
                    }
                    let start = parts[0].parse::<u64>()?;
                    let stop = parts[1].parse::<u64>()?;

                    Ok(Range { start, stop })
                })
                .collect()
        }
    }
}