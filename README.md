# Password Generator

Easy Password Generator.

## Usage

`pg new [flags]`

Available flags:

- `--amb`       Include ambiguous symbols ({},(),`,~...)
- `-l, --len` `int`   Length of the password (default 16)
- `--low` Include lowercase letters  (default true)
- `--nums` Include numbers (0,1,2,3...) (default true)
- `--syms` Include symbols (*,&,$,#...) (default true)
- `--up` Include uppercase letters (default true)

## Example

```bash
pg new --len 32 --low=true --up=false --syms=true --amb=false

Your new password is: ^ig-1wn180c8sa$hcqx^kd4qbumaemp!
```

## Installation

`$ go install github.com/Phaseant/pg@latest`
