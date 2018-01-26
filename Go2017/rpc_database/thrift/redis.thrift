namespace go method

struct Args {
1: string key
2: string value
}

service Method{
	i32 Add(1: string key, 2: string value)
	string Get(1: string key)
}
