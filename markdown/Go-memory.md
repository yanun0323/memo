# Golang 的記憶體管理


<br>

## 基礎的 Stack & Heap 
#### Stack (堆疊)
- 暫時性記憶體空間，保存函數所產生的暫時性變數(區域變數)，一旦函式執行結束，這些記憶體就會自動被抹去
- 特性是 **LIFO** *( Last In First Out，後進先出 )*
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
> Box / Unbox = 效能損耗

> Box / Unbox 只會發生在 `Value To Reference` or `Reference To Value`
，若兩者皆為 `Reference` 就不會發生


## Go 的 Stack & Heap
根據[官方FAQ](https://golang.org/doc/faq#stack_or_heap)的說法，我們並不能單純靠看 Go 的代碼就判斷變數被分配到哪邊
> Go compilers will allocate variables that are local to a function in that function’s stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.
If a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic escape analysis recognizes some cases when such variables will not live past the return from the function and can reside on the stack.

但是 Go 提供了一個Tag  `-m` 來讓我們在 `run` `build` `test` 能觀察變數分配到記憶體的狀況

<br>
<br>
<br>

## 記憶體管理

記憶體管理包含三個部分 : `Mutator` `Allocator` `Collector`