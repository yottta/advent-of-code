pub mod utils {
    pub fn must_parse_to_i32(input: &str) -> i32 {
        match input.parse::<i32>() {
            Ok(number) => number,
            Err(e) => panic!("failed to convert {} to i32: {}", input, e),
        }
    }

    use std::env;
    use std::fs::File;
    use std::io::{BufRead, BufReader};
    use std::path::Path;

    pub fn read_file_from_env(env_key: String) -> Vec<String> {
        let mut lines: Vec<String> = Vec::new();
        let (raw_path, ok)= read_env(env_key);
        if !ok {
            return lines
        }
        let filepath= Path::new(&raw_path);
        let file = File::open(filepath);
        match file {
            Ok(f) => {
                let reader = BufReader::new(f);

                for line_result in reader.lines() {
                    match line_result {
                        Ok(line) => lines.push(line),
                        Err(e) => panic!(
                            "failed to read file line from {}: {}",
                            filepath.display(),
                            e
                        ),
                    }
                }
            }
            Err(e) => panic!(
                "failed to read file lines from {}: {}",
                filepath.display(),
                e
            ),
        }
        lines
    }

    fn read_env(key: String) -> (String, bool) {
        match env::var(&key) {
            Ok(val) => (val, true),
            Err(e) => {
                println!("no input defined for {}: {}", key, e);
                (String::new(), false)
            }
        }
    }
}
