wrs -- weighted random selection
================================

Usage:

1. Import the package

    import "github.com/sndb/wrs"

2. Create a Choices slice

    cs := Choices{{5, "abc"}, {2, "xyz}, {3, "zxc"}}

3. Choose the random element and assert its type

    cs.Choose().(string)
