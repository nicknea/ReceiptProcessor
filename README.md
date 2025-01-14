Instructions
Manuever to RECEIPTPROCESSOR directory
go run .


Tests can be run with go test

Improvement Areas

Error Handling, currently just panics out. Handle bad requests like impromper UUIDs or malformed receipt data gracefully.

Store receipt instead of just points total. Could be useful for future extended functionality such as recalculating points with formula adjustments. 
