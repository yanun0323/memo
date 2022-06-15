## 在 VSCode 設置 Swift 開發環境

####安裝VSCode Extension

- <B>Swift</B>
    \- apple 開發者維護的Extension(Swift Server Work Group)
+ <B>Clang-Format</B>
    \- 編譯 swift 使用
- <B>apple-swift-format</B>
    \- swift 自動 format 功能
    ```
    安裝流程
    1. 安裝 Xcode
    2. 安裝 mint: $brew install mint
    3. 指定 xcode-selector 位置: 
        $sudo xcode-select -s /Applications/Xcode.app/Contents/Developer
    4. 安裝 apple 官方 format
        $mint install apple/swift-format@release/{VERSION}
        {VERSION}= 5.6 or others
    5. 找到 apple/swift-format 的存放路徑
        $mint which swift-format
    6. 到 VSCode preference 找到 'Apple-swift-format: Path'
       將 apple/swift-format 的路徑貼上去
    7. 完成！在任何 .swift 檔案案儲存就會自動format了
    ```


####source
- https://github.com/vknabel/vscode-apple-swift-format
- https://github.com/yonaskolb/Mint
- https://github.com/apple/swift-format/tree/release/5.7




