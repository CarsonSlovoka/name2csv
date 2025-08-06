# name2csv

將ttx對name表的輸出，改用csv來顯示

# INSTALL

```sh
go install github.com/CarsonSlovoka/name2csv@latest
```

```sh
git clone https://github.com/CarsonSlovoka/name2csv.git
cd name2csv
go install .
```

# USAGE

```lua
:%!name2csv     -- 請不要有其它的tag, 例如ttFont
'<,'>!name2csv
```

測試文本

```xml
<name>
  <namerecord nameID="0" platformID="1" platEncID="0" langID="0x0" unicode="True">
    Copyright © 2025 ... All rights reserved.
  </namerecord>
  <namerecord nameID="1" platformID="1" platEncID="0" langID="0x0" unicode="True">
    Qoo
  </namerecord>
  <namerecord nameID="2" platformID="1" platEncID="0" langID="0x0" unicode="True">
    Ultralight
  </namerecord>
  <namerecord nameID="5" platformID="1" platEncID="0" langID="0x0" unicode="True">
    Version 1.000
  </namerecord>
</name>
```

outout:

```csv
platformID,platEncID,langID,nameID,text
1,0,0x0,0,Copyright © 2025 ... All rights reserved.
1,0,0x0,1,Qoo
1,0,0x0,2,Ultralight
1,0,0x0,5,Version 1.000
```
