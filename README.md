wrs -- weighted random selection
================================

Usage:

1. Import the package

    import "github.com/sndb/wrs"

2. Create a new Chooser

    chr, err := NewChooser(
        Choice{5, "abc"},
        Choice{2, "xyz"},
        Choice{3, "zxc"},
    )
    if err != nil {
        log.Fatal(err)
    }

3. Pick the random element and assert its type

    chr.Pick().(string)
