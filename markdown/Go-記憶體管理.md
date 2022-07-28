# Golang 的記憶體管理

<br>

## 基礎的 Stack & Heap 
#### Stack (堆疊)
- 暫時性記憶體空間，保存函數所產生的暫時性變數(區域變數)，一旦函式執行結束，這些記憶體就會自動被抹去
- 特性是 **LIFO** *(Last In First Out，後進先出)*
- 不用擔心 **Memory Leak**

#### Heap (堆積)
- 保存全域變數 *(global variables)* 。預設情況下，所有全域變數都保存 heap memory

#### 變數的儲存
- `Value type` 變數與指標變數都會儲存在 **Stack**
- `Reference type` 變數(指標本身)也會儲存在 **Stack**，而指標指向的變數則儲存在 **Heap**

#### Box
- `Value type` -> `Reference type` 的過程
- Value 被放進 **Heap**，並在 **Stack** 產生一個指標變數指向 Value

#### Unbox
- `Reference type` -> `Value type` 的過程
- Reference 指向的變數(Heap)，複製一份到 **Stack** 
> Box/Unbox = 效能損耗

> Box/Unbox 只會發生在 `Value To Reference` or `Reference To Value`
，若兩者皆為 `Reference` 就不會發生


## Go 的 Stack & Heap
根據[官方FAQ](https://golang.org/doc/faq#stack_or_heap)的說法，我們並不能單純靠看 Go 的代碼就判斷變數被分配到哪邊
> Go compilers will allocate variables that are local to a function in that function’s stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.
If a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic escape analysis recognizes some cases when such variables will not live past the return from the function and can reside on the stack.

但是 Go 提供了一個Tag  `-m` 來讓我們在 `run` `build` `test` 能觀察變數分配到記憶體的狀況


## 記憶體管理

記憶體管理包含三個部分 : `Mutator` `Allocator` `Collector`

![Image](./asset/go-memory/1.png)

> `Mutator` 向 `Allocator` 申請記憶體
>
> `Allocator` 在 Heap 中初始化相對應的記憶體
>
> `Collector` 管理及回收記憶體

// TODO: Collector？？

## Allocator 設計原理

常見的 Allocator 有兩種類型
- Bump Allocator (Sequential Allocator)
- Free-List Allocator

### Bump Allocator
Bump Allocator 使用一種非常高效能的記憶體分配方式，他使用一個指針指向特定記憶體位置。
當程式申請記憶體時，他檢查指針後的剩餘記憶體空間，分配一段記憶體空間出去，接著移動指針到分配過的記憶體尾端。

![Image](./asset/go-memory/2.png)

Bump Allocator 雖然執行速度快、時間複雜度低，但若是分配過的記憶體被釋放了，這段空的記憶體並不會被重新利用。

![Image](./asset/go-memory/3.png)

> 紅色部分是被釋放的記憶體，Bump Allocator 並不會重新使用它

### Free-List Allocator
Free-List Allocator 在記憶體內部管理一個類似 `Link-List` 的結構，記錄所有空記憶體位址。
當程式申請記憶體，他會去遍歷這個 `Link-List`，找出足夠大的記憶體，分配並修改 `Link-List`

![Image](./asset/go-memory/4.png)

> 因為分配記憶體時要遍歷整個 `Link-List`，所以時間複雜度為 O(n)

Free-List Allocator 在設計時可以選擇不同的策略，常見的有以下四種
- **First-Fit**：
從頭開始遍歷，選擇第一個大於申請大小的記憶體
- **Next-Fit**：
從上次遍歷結束的位置開始遍歷，選擇第一個大於申請大小的記憶體
- **Best-Fit**：
遍歷整個 `Link-List`，選擇最合適大小的記憶體
- **Segregated-Fit**：
將記憶體分成多個 `Link-List`，每個`Link-List`內的記憶體大小都相同。申請記憶體時，先找到滿足條件的 `Link-List`，再開始遍歷

Go語言使用的記憶體分配策略類似第四種 `Segregated-Fit`

![Image](./asset/go-memory/5.png)

> 這邊將記憶體分成4、8、16、32 byte 的記憶體 `Link-List`
> 
> 當程式申請 8 byte 大小的記憶體，他會先找到 8 byte `Link-List`，從中分配記憶體