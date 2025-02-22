# พาย ไฟ ได - Pie Fire Dire

### 1. ติดตั้ง Dependencies
1. รันคำสั่งต่อไปนี้ใน terminal เพื่อดาวน์โหลด dependencies ที่จำเป็น:
```bash
go mod tidy
```

### 2. รันโปรแกรม
1. รันคำสั่ง `go run` เพื่อรันไฟล์ `main.go`
```bash
go run main.go
```


### 3. ได้ Output ออกมาแต่ละครั้งจะไม่เหมือนกันที่ ไป Fetch Data แต่ละครั้งจะถูก random ออกมาก

```plaintext
{
"beef":{
    "bresaola":30,
    "enim":40,
    "fatback":35,
    "jowl":25,
    "meatloaf":37,
    "pastrami":47,
    "pork":131,
    "t-bone":46
    }
}
```

### 4. Run Unit Test
```bash
go test -v 
```

### 5. ได้ Output จากการทำ Unit Test
```plaintext
=== RUN   TestCountBeef
--- PASS: TestCountBeef (0.00s)
PASS
ok      search-beef     0.475s
```
