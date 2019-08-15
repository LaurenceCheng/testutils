module example.com/me/use

go 1.12

require example.com/me/lib v0.0.0

require example.com/me/testutils v0.0.0

replace example.com/me/lib => ../lib

replace example.com/me/testutils => ../testutils
