`flag.CommandLine.Var()`で、SetとStringのインターフェイスを持つ`Flag`型が追加される。
Flag型に追加されるときに、interfaceのString()が呼ばれて、DefValueに格納される。
Parseしたときにヘルプが必要であれば、このDefValueが表示されているので、単位がついたものが表示されている。
