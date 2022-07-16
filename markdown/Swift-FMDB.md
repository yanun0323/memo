# FMDB - Swift SQLite 套件

在 Swift 專案中使用

1. Copy the relevant .m and .h files from the FMDB src folder into your project.

2. If prompted to create a "bridging header", you should do so. If not prompted and if you don't already have a bridging header, add one.

3. In your bridging header, add a line that says:
```swift
#import "FMDB.h"
```

4. Use the variations of executeQuery and executeUpdate with the sql and values parameters with try pattern, as shown below. These renditions of executeQuery and executeUpdate both throw errors in true Swift fashion.

> If you do the above, you can then write Swift code that uses FMDatabase. For example, as of Swift 3:
```swift
let fileURL = try! FileManager.default
    .url(for: .applicationSupportDirectory, in: .userDomainMask, appropriateFor: nil, create: true)
    .appendingPathComponent("test.sqlite")

let database = FMDatabase(url: fileURL)

guard database.open() else {
    print("Unable to open database")
    return
}

do {
    try database.executeUpdate("create table test(x text, y text, z text)", values: nil)
    try database.executeUpdate("insert into test (x, y, z) values (?, ?, ?)", values: ["a", "b", "c"])
    try database.executeUpdate("insert into test (x, y, z) values (?, ?, ?)", values: ["e", "f", "g"])

    let rs = try database.executeQuery("select x, y, z from test", values: nil)
    while rs.next() {
        if let x = rs.string(forColumn: "x"), let y = rs.string(forColumn: "y"), let z = rs.string(forColumn: "z") {
            print("x = \(x); y = \(y); z = \(z)")
        }
    }
} catch {
    print("failed: \(error.localizedDescription)")
}

database.close()
```
