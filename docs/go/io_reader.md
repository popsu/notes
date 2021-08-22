# io.Reader

``` mermaid
%%{init: {'theme': 'base',  'themeVariables': { 'darkMode': true,
'primaryColor': '#2c698f', 'lineColor': '#fcd303',
'secondaryColor': '#00ff00', 'tertiaryColor': '#123456', 'primaryTextColor': '#e9ebd2'
}}}%%
graph LR
    B["[]byte"]
    R{io.Reader}
    str("string (text)")
    strFn("string (filename)")

    subgraph ReaderWriter [io.Writer<br>io.Reader]
        Fd[os.File]
        buf[bytes.Buffer]
    end

    subgraph Reader [io.Reader]
        br[bytes.Reader]
        sr[strings.Reader]
    end

    B -->|bytes.NewBuffer| buf
    B -->|bytes.NewReader| br
    B -->|"string"| str
    B -->|os.WriteFile| strFn
    str -->|"[]bytes"| B
    R -->|io.ReadAll| B
    strFn --> |os.ReadFile| B
    strFn --> |os.Open| Fd
    str --> |strings.NewReader| sr
```
