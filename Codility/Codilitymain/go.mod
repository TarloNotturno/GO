module main

go 1.22.4
replace example.com/MaxCounters => ../MaxCounters
replace example.com/GenomicRangeQuery => ../GenomicRangeQuery
replace example.com/MissingInteger => ../MissingInteger

require example.com/MaxCounters v0.0.0-00010101000000-000000000000
require example.com/GenomicRangeQuery v0.0.0-00010101000000-000000000000
require example.com/MissingInteger v0.0.0-00010101000000-000000000000

